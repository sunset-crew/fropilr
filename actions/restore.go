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

package actions

import (
  "os"
	"log"
  "fmt"
  "fropilr/tar"
  "fropilr/config"
)

// RestoreActions
func RestoreActions(archiveName string){
    if _, err := os.Stat(config.GetConfigDirectory()); !os.IsNotExist(err) {
        log.Fatal(".fropilr already exists")
    }
    home, _ := config.GetHomeDir()
    backupArchive := fmt.Sprintf("%s/%s.tar.gz",config.GetBackupDirectory(),archiveName)
    tar.Unftar(home,backupArchive)
    fmt.Printf("%s Restore\n",backupArchive)
}

// CheckForRestorableArchive
func CheckForRestorableArchive(archiveName string) bool {
    if _, err := os.Stat(fmt.Sprintf("%s/%s.tar.gz",config.GetBackupDirectory(),archiveName)); !os.IsNotExist(err) {
        return true
    }
    return false
}
