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
  "fmt"
	"log"
  "fropilr/tar"
  "fropilr/config"
)

func BackupActions(){
    backupDir := config.GetBackupDirectory()
    if _, err := os.Stat(backupDir); os.IsNotExist(err) {
      // path/to/whatever does not exist
      err := os.Mkdir(backupDir, 0755)
      if err != nil {
        log.Fatal(err)
      }
    }
    home, _ := config.GetHomeDir()
    backupName := fmt.Sprintf("%s/%s.tar.gz",backupDir,config.GetBackupName())
    if _, err := os.Stat(backupName); ! os.IsNotExist(err) {
        // move file first
        log.Printf("%s removed old",backupName)
        err = os.Remove(backupName)
        if err != nil {
            log.Fatal(err)
        }
    }
    tar.BackupProfile(home,backupName)
    log.Printf("%s backed up",backupName)
}
