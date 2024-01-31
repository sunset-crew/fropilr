/*
Copyright © 2021 Joe Siwiak <joe@unherd.info>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package tar

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// BackupProfile backs up the full profile
func BackupProfile(home string, backupLocation string) {
	// var out []string
	var buf bytes.Buffer
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	os.Chdir(home)

	// alphabetical order matters
	out := WalkList(".fropilr/config")
	data := WalkList(".fropilr/data")
	gnupg := WalkList(".gnupg")
	out = append(out, data...)
	out = append(out, gnupg...)
	out = append(out, ".muttrc")
	fmt.Println(out)
	CompressList(out, &buf)
	os.Chdir(path)

	fileToWrite, err := os.OpenFile(backupLocation, os.O_CREATE|os.O_RDWR, os.FileMode(0644))
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(fileToWrite, &buf); err != nil {
		panic(err)
	}
	if err := fileToWrite.Close(); err != nil {
		log.Fatal(err)
	}
}

// WalkList Walks a directory and makes a list of files
func WalkList(directory string) []string {
	var out []string
	err := filepath.Walk(directory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			out = append(out, path)
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return out
}

// CompressList compresses a list of files
func CompressList(fileList []string, buf io.Writer) error {
	zr := gzip.NewWriter(buf)
	tw := tar.NewWriter(zr)
	for _, file := range fileList {
		fi, err := os.Stat(file)
		if err != nil {
			return err
		}

		// mode := fi.Mode()
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}

		// must provide real name
		// (see https://golang.org/src/archive/tar/common.go?#L626)
		header.Name = filepath.ToSlash(file)

		// write header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		// if not a dir, write file content
		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			if _, err := io.Copy(tw, data); err != nil {
				return err
			}
		}
	}
	// produce tar
	if err := tw.Close(); err != nil {
		return err
	}
	// produce gzip
	if err := zr.Close(); err != nil {
		return err
	}
	return nil
}

// validRelPath check for path traversal and correct forward slashes
func validRelPath(p string) bool {
	if p == "" || strings.Contains(p, `\`) || strings.HasPrefix(p, "/") || strings.Contains(p, "../") {
		return false
	}
	return true
}

// Decompress the io stream
func Decompress(src io.Reader, dst string) error {
	// ungzip
	zr, err := gzip.NewReader(src)
	if err != nil {
		return err
	}
	// untar
	tr := tar.NewReader(zr)

	// uncompress each element
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return err
		}
		target := header.Name

		// validate name against path traversal
		if !validRelPath(header.Name) {
			return fmt.Errorf("tar contained invalid name error %q", target)
		}

		// add dst + re-format slashes according to system
		target = filepath.Join(dst, header.Name)
		// if no join is needed, replace with ToSlash:
		// target = filepath.ToSlash(header.Name)

		// check the type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it (with 0755 permission)
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}
		// if it's a file create it (with same permission)
		case tar.TypeReg:
			fileToWrite, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			// copy over contents
			if _, err := io.Copy(fileToWrite, tr); err != nil {
				return err
			}
			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			fileToWrite.Close()
		}
	}

	//
	return nil
}

// Unftar  restoreDir and archive
func Unftar(restoreDir string, archive string) {
	// untar write
	b, e := ioutil.ReadFile(archive)
	if e != nil {
		panic(e)
	}
	r := bytes.NewReader(b)
	if err := Decompress(r, restoreDir); err != nil {
		// probably delete uncompressHere?
		log.Fatal(err)
	}
}

// ListProfiles - lists the saved profile
func ListProfiles(backupDirectory string) []string {
	files, err := ioutil.ReadDir(backupDirectory)
	if err != nil {
		log.Fatal(err)
	}
	out := []string{}
	for _, f := range files {
		basename := f.Name()
		name := strings.TrimSuffix(basename, ".tar.gz")
		out = append(out, name)
	}
	return out
}
