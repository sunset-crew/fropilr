/*Package cmd encrypt this encrypts
Copyright © 2020 Joe Siwiak <joe@unherd.info>

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
package cmd

import (
  "errors"
  "fropilr/utils"
  "fropilr/config"
  "fropilr/gpg"
  "fmt"
  "github.com/ProtonMail/gopenpgp/v2/helper"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
)

var thefile string

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Using a pubkey, this encrypts the message",
	Long: `Using a public key, this encrypts a message to be
sent across the waves`,
  Args: func(cmd *cobra.Command, args []string) error {
    if len(args) < 1 {
      return errors.New("Encrypt requires <file>")
    }
    f, err := utils.Fexists(args[0])

    if err != nil {
      panic(err)
    }
    if f == false {
      return errors.New("cleartext file missing")
    }
    thefile = args[0]
    return nil
  },
	Run: func(cmd *cobra.Command, args []string) {
    pubkey := utils.ReadFile(viper.GetString("pubkey"))
    data := utils.ReadFile(args[0])
    if viper.GetString("outputfile") == "encrypted.gpg" {
        viper.Set("outputfile", fmt.Sprintf("%s.gpg",args[0]))
    }
    armorFlag, _ := cmd.Flags().GetBool("armor")
    if armorFlag == true {
        armor,err:= helper.EncryptMessageArmored(pubkey,data)
        if err != nil {
          panic(err)
        }
        utils.Write2file(viper.GetString("outputfile"),armor)
        return
    }
    binary,err:= gpg.EncryptMessageBinary(pubkey,data)
    if err != nil {
        panic(err)
    }
    utils.WriteBinary2file(viper.GetString("outputfile"),binary)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

  defaultName := fmt.Sprintf("%s.gpg","encrypted")
  encryptCmd.Flags().BoolP("armor", "a", false, "outputs armor instead of binary files")
  encryptCmd.Flags().StringP("pubkey", "p", config.GetPubKey(), "Custom Public Key (uses the one in .fropilr/pubkey")
  encryptCmd.Flags().StringP("outputfile", "f", defaultName, "Custom Output File (default: encrypted.gpg")
	viper.BindPFlag("pubkey", encryptCmd.Flags().Lookup("pubkey"))
	viper.BindPFlag("outputfile", encryptCmd.Flags().Lookup("outputfile"))
	viper.SetDefault("pubkey", config.GetPubKey())
	viper.SetDefault("outputfile", defaultName)
}
