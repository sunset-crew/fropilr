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
	"fmt"
	"fropilr/config"
	"fropilr/tar"
	"fropilr/web"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Saved Profiles",
	Long: `List the Saved Profiles
This way you can view all the profiles that
have been saved.`,
	Run: func(cmd *cobra.Command, args []string) {
		const padding = 3
		w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.AlignRight)
		if _, err := os.Stat(config.GetBackupDirectory()); !os.IsNotExist(err) {
			// path/to/whatever exists
			remote, _ := cmd.Flags().GetBool("remote")
			if remote == true {
				fmt.Println("Get list from remote")
				web.ListRemoteProfiles()
				return
			}
			for _, s := range tar.ListProfiles(config.GetBackupDirectory()) {
				fmt.Fprintf(w, "\t%s\n", s)
			}
			w.Flush()
			return
		}
		fmt.Println("no profile backup directory")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("remote", "r", false, "Help message for toggle")
}
