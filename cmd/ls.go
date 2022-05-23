/*
Copyright © 2022 Aditya Agrawal adiag1200@gmail.com

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

	"github.com/crossphoton/drive-encrypt/src"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list files with path",
	Run: func(cmd *cobra.Command, args []string) {

		if b, _ := cmd.Flags().GetBool("all"); b {
			files, err := src.ListAllFiles(getPassword(cmd))
			if err != nil {
				log.Fatal(err)
			}

			printFiles(files)
			return
		}

		if len(args) == 0 {
			files, err := src.ListPath("", getPassword(cmd))
			if err != nil {
				log.Fatal(err)
			}

			printFiles(files)
		} else {
			for _, arg := range args {
				files, err := src.ListPath(arg, getPassword(cmd))
				if err != nil {
					log.Fatal(err)
				}

				printFiles(files)

			}
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().BoolP("all", "a", false, "Print all files (Path independent)")
}
