// Copyright © 2016 Wei-Ting Kuo <waitingkuo0527@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/waitingkuo/deck/utils"
	"os"
	"strings"
)

// package_installCmd represents the package_install command
var packageInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// consider moving this to global init or utils
		utils.MkdirDeckLocalBinPath()

		if len(args) == 0 {
			os.Exit(0) // FIXME define better exit method, maybe change the code number
		}

		// assume it has bin name now
		// FIXME need to refactor
		a := strings.Split(args[0], "/")
		packageName := a[0]
		binName := a[1]
		fmt.Println(packageName, binName)
		if err := utils.InstallBin(packageName, binName); err != nil {
			panic(err)
		}

	},
}

func init() {
	packageCmd.AddCommand(packageInstallCmd)
}
