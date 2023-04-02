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
    "fmt"
	"io"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"os"
    "net/url"
    "time"
)


func DownloadFromServer(name string, email string) error {
	// Get the data
	resp, err := http.PostForm("http://localhost:9999/download", url.Values{"name": {name},"email": {email}})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create("/home/joee/arch.tar.gz")
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// https://github.com/mikicaivosevic/golang-json-client
func getResponseData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}
	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	return contents
}

func ListRemoteProfiles() {
    var entries []ListEntry
    var netClient = &http.Client{
      Timeout: time.Second * 10,
    }
    resp, err := netClient.Get("http://localhost:9999/list")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    err = json.Unmarshal(body, &entries)
    if err != nil {
      fmt.Println("error:", err)
    }
    // fmt.Println("Response status:", resp.Status)
    for _, entry := range entries {
          fmt.Printf("  %s\n",entry.Name)
    }
}
//~ func GetArchive(name,email string){
    //~ response, err := http.PostForm("https://example.com/api", url.Values{"id": {"8"}})
    //~ if err != nil {
        //~ // postForm error happens
    //~ } else {
        //~ defer response.Body.Close()
        //~ body, err := ioutil.ReadAll(response.Body)

        //~ if err != nil {
            //~ // read response error
        //~ } else {
            //~ // now handle the response
        //~ }
    //~ }
//~ }

//~ func Post(url string, jsonData string) string {
	//~ var jsonStr = []byte(jsonData)

	//~ req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//~ req.Header.Set("Content-Type", "application/json")

	//~ client := &http.Client{}
	//~ resp, err := client.Do(req)
	//~ if err != nil {
		//~ panic(err)
	//~ }
	//~ defer resp.Body.Close()

	//~ fmt.Println("response Status:", resp.Status)
	//~ fmt.Println("response Headers:", resp.Header)
	//~ body, _ := ioutil.ReadAll(resp.Body)
	//~ return string(body)
//~ }
