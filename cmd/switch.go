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

package cmd

import (
  "log"
  "fropilr/actions"
	"github.com/spf13/cobra"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch <new user>",
	Short: "Switches user profile to selected profile",
	Long: `Switching user profile to selected profile`,
	Run: func(cmd *cobra.Command, args []string) {
    if len(args) < 1 {
        log.Fatal("requires at least a user account to pick from (list)")
    }

    if actions.CheckForRestorableArchive(args[0]) == false {
        log.Fatal("Profile Does not exist")
    }

		actions.BackupActions()
		actions.RemoveActions(true)
		actions.RestoreActions(args[0])
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
}
