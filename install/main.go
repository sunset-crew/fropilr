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

package install

import (
  "os"
  "fmt"
  "bufio"
  "strings"
  "errors"
  "github.com/spf13/viper"
  "github.com/spf13/cobra"
  "fropilr/config"
  "fropilr/utils"
  "fropilr/gpg"
)



// SetPassphrase - sets the passphase for the current env
func SetPassphrase() {
      reader := bufio.NewReader(os.Stdin)
      fmt.Print("Passphrase: ")
      pass, _ := reader.ReadString('\n')
      fmt.Print("Enter Passphrase Again: ")
      pass2, _ := reader.ReadString('\n')
      if pass != pass2 {
          fmt.Printf("what the hell,\nThese don't match up\nText1: \"%s\", Text2: \"%s\"\n",strings.TrimSpace(pass), strings.TrimSpace(pass2))
          os.Exit(1)
      } else {
        fmt.Println("generating key")
        pass := strings.TrimSpace(pass)
        encpass, _ := gpg.EncryptCustomValue(strings.TrimSpace(pass))
        utils.Write2file(config.GetPassphraseFile(),encpass)
      }
}

// CheckAndValidate - Checks and Validates the Environment
func CheckAndValidate(cmd *cobra.Command) (int, error) {
    // can't find /.fropilr/live.yml
    home, _ := config.GetHomeDir()
    fexists, _ := utils.Fexists(config.GetConfigFile())
    if fexists == false {
      // check for .gnupg/
      _, err := os.Stat(fmt.Sprintf("%s/.gnupg/",home))
      if os.IsNotExist(err) {
        fmt.Println("Folder does not exist, setting up gpg")
        return 1, nil
      }
      //fmt.Println("gpg exists, please backup .gnupg and delete it")
      return 3, errors.New("gpg exists, please backup .gnupg and delete it")
    }
		return 2, errors.New(".fropilr/ exists, you should be able to use it")
}

//~ -      configDir := config.GetConfigDirectory()
//~ -      utils.CreateDirIfNotExist(configDir)
//~ -      utils.CreateFile(config.GetConfigFile())
//~ -      //reader := bufio.NewReader(os.Stdin)
//~ -      // Prompt and read



// ValidateAndCreate - Collects and Creates the New Key
func ValidateAndCreate(keyType string){
    reader := bufio.NewReader(os.Stdin)
    // Prompt and read
    fmt.Print("Name: ")
    name, _ := reader.ReadString('\n')
    fmt.Print("Email: ")
    email, _ := reader.ReadString('\n')
    // Trim whitespace and print
    fmt.Printf("%s if you can't think of one\n",utils.RandString(16))
    fmt.Print("Passphrase: ")
    pass, _ := reader.ReadString('\n')
    fmt.Print("Enter Passphrase Again: ")
    pass2, _ := reader.ReadString('\n')
    if pass != pass2 {
        fmt.Printf("what the hell,\nThese don't match up\nText1: \"%s\", Text2: \"%s\"\n",strings.TrimSpace(pass), strings.TrimSpace(pass2))
        os.Exit(1)
    }
    fmt.Println("generating key")

    pass = strings.TrimSpace(pass)
    name = strings.TrimSpace(name)
    email = strings.TrimSpace(email)

    configDir := config.GetConfigDirectory()
    utils.CreateDirIfNotExist(config.GetFropileDirectory())
    utils.CreateDirIfNotExist(config.GetBackupDirectory())
    utils.CreateDirIfNotExist(configDir)
    utils.CreateDirIfNotExist(config.GetDataDirectory())
    // utils.CreateFile(config.GetConfigFile())

    rsaKey := gpg.GenKey(name,email,keyType,[]byte(pass))
    utils.Write2file(config.GetPrivKey(), rsaKey)

    fmt.Println(gpg.ImportSecretkey(pass))
    pubkey, _ := gpg.ExportPubkey(name)
    utils.Write2file(config.GetPubKey(),pubkey)

    encpass, _ := gpg.EncryptCustomValue(strings.TrimSpace(pass))
    utils.Write2file(config.GetPassphraseFile(),encpass)

    fmt.Println("saving info")
    v := viper.New()
    v.AddConfigPath(configDir)
    v.SetConfigName("config")
    v.SetConfigType("yaml")
    v.Set("Name", name)
    v.Set("Email", email)
    // v.WriteConfig()
    err := v.SafeWriteConfig()
    if err != nil {
        fmt.Println(err)
    }
    // _ = os.Mkdir(config.GetBackupDirectory(), os.ModeDir)
    fmt.Println("Ok, everything is setup.")
}

// ValidateEnvironment Validates Environment
func ValidateEnvironment(){
    confDir := config.GetConfigDirectory()
    result, _ := utils.Fexists(confDir)
    if result == false {
      fmt.Println("valid")
    }
}
