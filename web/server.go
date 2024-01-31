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

package web

import (
	"encoding/json"
	"fmt"
	"fropilr/config"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	// "path/filepath"
)

type ListEntry struct {
	Name     string `json:name`
	Basename string `json:basename`
}

type UploadFileInfo struct {
	Name   string            `json:"name"`
	Email  string            `json:"email"`
	Errors map[string]string `json:"errors"`
}

var rxEmail = regexp.MustCompile(".+@.+\\..+")

func (uploadFileInfo *UploadFileInfo) Validate() bool {
	uploadFileInfo.Errors = make(map[string]string)

	match := rxEmail.Match([]byte(uploadFileInfo.Email))
	if match == false {
		uploadFileInfo.Errors["Email"] = "Please enter a valid email address"
	}

	if strings.TrimSpace(uploadFileInfo.Name) == "" {
		uploadFileInfo.Errors["Content"] = "Please enter a message"
	}

	return len(uploadFileInfo.Errors) == 0
}

// https://gist.github.com/mattetti/5914158
// https://www.golanglearn.com/json-post-payload-example-in-golang/
// https://zetcode.com/golang/http-server/
// https://compositecode.blog/2020/05/04/simple-http-server-starter-in-15-minutes/
// curl -F "myFile=@archive.tar.gz"  -F "name=joee" -F "email=kf4jasgmail.com" http://localhost:9999/upload
func uploadFile(w http.ResponseWriter, r *http.Request) {

	//err := json.NewDecoder(r.Body).Decode(&postFields)
	//if err != nil {
	//    fmt.Println("Error Retrieving the File")
	//    fmt.Println(err)
	//    return
	//}

	fmt.Println("File Upload Endpoint Hit")

	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	// Here, the uploaded file is specified to be in range of max size 10 MB to 20 MB.
	r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file

	postFields := &UploadFileInfo{
		Name:  r.FormValue("name"),
		Email: r.FormValue("email"),
	}

	if postFields.Validate() == false {
		outPut, err := json.Marshal(postFields)
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(outPut)
		return
	}
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)

	// fmt.Printf("MIME Header: %+v\n", handler.Header)

	name := r.FormValue("name")
	email := r.FormValue("email")

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Email: %s\n", email)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	//~ tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	//~ if err != nil {
	//~ fmt.Println(err)
	//~ }
	//~ defer tempFile.Close()
	// read all of the contents of our uploaded file into a
	// byte array

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	// write this byte array to our temporary file
	//tempFile.Write(fileBytes)

	backup_file := fmt.Sprintf("%s/%s.tar.gz", config.GetBackupDirectory(), config.TransformNameEmail(name, email))
	err = ioutil.WriteFile(backup_file, fileBytes, 0644)
	// return that we have successfully uploaded our file!
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Fprintf(w, "Successfully Uploaded File\n")
	outPut, err := json.Marshal(postFields)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(outPut)
}

// curl -F "name=joeeeee" -F "email=kf4jas@gmail.com" http://localhost:9999/download --output archive.tar.gz

func downloadFile(w http.ResponseWriter, r *http.Request) {
	postFields := &UploadFileInfo{
		Name:  r.FormValue("name"),
		Email: r.FormValue("email"),
	}
	if postFields.Validate() == false {
		outPut, err := json.Marshal(postFields)
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(outPut)
		return
	}
	backup_file := fmt.Sprintf("%s/%s.tar.gz", config.GetBackupDirectory(), config.TransformNameEmail(postFields.Name, postFields.Email))
	http.ServeFile(w, r, backup_file)
}

func listBackups(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Successfully Uploaded File\n")
	var entries []ListEntry
	files, err := ioutil.ReadDir(config.GetBackupDirectory())
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		basename := f.Name()
		name := strings.TrimSuffix(basename, ".tar.gz")
		le := ListEntry{Basename: basename, Name: name}
		entries = append(entries, le)
	}
	outPut, err := json.Marshal(entries)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(outPut)
}

func setupRoutes() {
	http.HandleFunc("/list", listBackups)
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/download", downloadFile)
	http.ListenAndServe(":9999", nil)
}

func Server() {
	setupRoutes()
}
