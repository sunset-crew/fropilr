/*Package cmd decrypt this decrypts
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
    "fropilr/gpg"
    "fropilr/utils"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "errors"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "This decrypts encrypted files sent to the server.",
	Long: `Decrypts the file in the arguments`,
  Args: func(cmd *cobra.Command, args []string) error {
    if len(args) < 1 {
      return errors.New("requires at least a file to decrypt")
    }
    if viper.GetString("out") == "output.txt" {
      viper.Set("out", utils.FilenameWithoutExtension(args[0]))
    }
    return nil
  },
	Run: func(cmd *cobra.Command, args []string) {
    armorFlag, _ := cmd.Flags().GetBool("armor")
    if armorFlag == true {
      gpg.DecryptArmorFile(args[0])
      return
    }
    gpg.DecryptBinaryFile(args[0])
  },
}

func init() {
	rootCmd.AddCommand(decryptCmd)
  decryptCmd.Flags().String("out", "output.txt", "Output file")
  // decryptCmd.MarkFlagRequired("out")
  viper.BindPFlag("out", decryptCmd.Flags().Lookup("out"))
  viper.SetDefault("out", "output.txt")
}



