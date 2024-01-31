/*
Package cmd setup this sets up the environment
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
	"fropilr/install"
	"fropilr/utils"
	"github.com/spf13/cobra"
	"log"
)

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Sets up the gnupg environment",
	Long:  `This installs everything`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var val int
		if val, err = install.CheckAndValidate(cmd); val == 1 {
			keyType, _ := cmd.Flags().GetString("key-type")
			if utils.FindIn([]string{"rsa", "x25519"}, keyType) == false {
				log.Fatal("That key type doesn't exist")
			}
			install.ValidateAndCreate(keyType)
			log.Println("Installed")
		}

		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	setupCmd.Flags().StringP("key-type", "k", "rsa", "Encryption key type [rsa,x25519]")
}
