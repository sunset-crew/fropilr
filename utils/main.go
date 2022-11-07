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

package utils

import (
  "os"
  "os/exec"
  "fmt"
  "bytes"
  "io/ioutil"
  "strings"
  "path"
)

// CreateFile this creates a single file kinda like the unix touch command
func CreateFile(filename string) {
    tFile, err := os.Create(filename)
    Errormsg(err)
    //some actions happen here
    tFile.Truncate(0)
    tFile.Seek(0,0)
    tFile.Sync()
    tFile.Close()
}


// Write2file Writes To files
func Write2file(filename string, input string) {
    f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
    if err != nil {
        fmt.Println(err)
        return
    }
    // Great test for bytes
    //d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
    _, err = f.WriteString(input)
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    //fmt.Println(n2, "bytes written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}

// WriteBinary2file Writes To files
func WriteBinary2file(filename string, input []byte) {
    f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
    if err != nil {
        fmt.Println(err)
        return
    }
    // Great test for bytes
    //d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
    _, err = f.Write(input)
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    //fmt.Println(n2, "bytes written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}


// CreateDirIfNotExist Creates directory even if it doesn't exist
func CreateDirIfNotExist(dir string) {
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        err = os.MkdirAll(dir, 0755)
        if err != nil {
            panic(err)
        }
    }
}

// Fexists checks if the file exists
func Fexists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}

// Errormsg display the error message
func Errormsg(err error){
    if err != nil {
        panic(err)
    }
    return
}

// DeleteFile deletes the file at the path string
func DeleteFile(path string) {
    // delete file
    var err = os.Remove(path)
    if err != nil {
        return
    }
    fmt.Println("File Deleted")
}

// ReadFile this reads the files
func ReadFile(path string) string {
    data, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println("File reading error", err)
        return ""
    }
    return string(data)
}

// ReadBinaryFile - this reads the binary files
func ReadBinaryFile(path string) []byte {
    data, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println("File reading error", err)
        return nil
    }
    return data
}

// FilenameWithoutExtension returns a string with the filename only
func FilenameWithoutExtension(fn string) string {
      return strings.TrimSuffix(fn, path.Ext(fn))
}

// RunCmdGrabOutput this grabs and processes the output of stdout.
func RunCmdGrabOutput(args []string) ([]byte, error) {
    c := exec.Command(args[0], args[1:]...)
    var stdout bytes.Buffer
    c.Stdout = &stdout
    c.Stderr = os.Stderr
    err := c.Run()
    if err != nil {
        return nil, err
    }
    return stdout.Bytes(),nil
}

// RunInteractive this runs shell commands that are interactive
func RunInteractive(args []string){
    c := exec.Command(args[0], args[1:]...)
    c.Stdin = os.Stdin
    c.Stdout = os.Stdout
    c.Stderr = os.Stderr
    err := c.Run()
    if err != nil {
        //fmt.Printf("%s\n", args)
        fmt.Printf("cmd.Run() failed with %s\n", err)
        os.Exit(1)
    }
}

// StandardizeSpaces - this sets a standard amt of spaces
func StandardizeSpaces(s string) string {
    return strings.Join(strings.Fields(s), " ")
}

// FindIn - string in a slice
func FindIn(slice []string, val string) (bool) {
    for _, item := range slice {
        if item == val {
            return true
        }
    }
    return false
}
