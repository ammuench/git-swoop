// git-swoop - a quick little cli util to *swoop* to another branch, pull down the latest from remote, and then return to where you started
// Copyright (C) 2025 Alex Muench
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime/debug"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("ERROR: Need a singular branch name to swoop from")
		os.Exit(126)
	}

	swoopBranch := os.Args[1]

	if swoopBranch == "-version" || swoopBranch == "--version" || swoopBranch == "-v" {
		if info, ok := debug.ReadBuildInfo(); ok {
			swoopVersionParsed := swoopVersion
			if swoopVersionParsed == "" {
				swoopVersionParsed = info.Main.Version
			}
			fmt.Printf("\ngit-swoop %s, built with %s\n", swoopVersionParsed, info.GoVersion)
			fmt.Println("")
			fmt.Println("git-swoop Copyright (C) 2025 Alex Muench")
			fmt.Println("This program comes with ABSOLUTELY NO WARRANTY")
			fmt.Println("This is free software, and you are welcome to redistribute it")
			fmt.Println("under certain conditions; check the LICENSE.md file at `https://github.com/ammuench/git-swoop`")
		}
		return
	}

	if swoopBranch == "-h" || swoopBranch == "-help" || swoopBranch == "--help" {
		fmt.Println("")
		fmt.Println("usage: git-swoop <target-branch-name>")
		fmt.Println("   git-swoop will try to checkout your target branch, pull down the latest from remote,")
		fmt.Println("   and then return to where you started")
		fmt.Println("")
		fmt.Println("flags:")
		fmt.Println("   --help (alias: -h, -help): prints this message")
		fmt.Println("   --version (alias: -v, -version): prints the version and basic package info")
		fmt.Println("")

		return
	}

	_, err := exec.Command("git", "status").Output()
	if err != nil {
		fmt.Println("ERROR: Not in a git repository")
		return
	}

	currBranchBytes, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	currBranch := string(currBranchBytes[0:(len(currBranchBytes) - 1)])
	if err != nil {
		fmt.Println("ERROR: Unable to determine current working branch")
		return
	}

	checkoutCmd := exec.Command("git", "checkout", swoopBranch)
	checkoutOutput, err := checkoutCmd.CombinedOutput()
	if err != nil {
		fmt.Printf("\n%v\n", string(checkoutOutput))
		fmt.Printf("\nERROR: Unable to checkout swoop branch `%s`\n", swoopBranch)
		fmt.Printf("\nStill on origninal branch `%s`\n", currBranch)
		return
	}

	pullCmd := exec.Command("git", "pull")
	pullOutput, err := pullCmd.CombinedOutput()
	if err != nil {
		fmt.Printf("\n%v\n", string(pullOutput))
		fmt.Printf("\nERROR: Unable to pull from swoop branch `%s`\n", swoopBranch)
		reCheckoutErr := reCheckout(currBranch)
		if reCheckoutErr != nil {
			fmt.Printf("\nStill on swoop branch `%s`\n", swoopBranch)
		} else {
			fmt.Printf("\nReturned to original branch `%s`\n", currBranch)
		}
		return
	} else {
		fmt.Printf("\n%v\n", string(pullOutput))
	}

	err = reCheckout(currBranch)
	if err != nil {
		fmt.Printf("\nStill on swoop branch `%s`\n", swoopBranch)
		return
	}

	fmt.Printf("\nSuccessfully swooped from branch `%s` and returned to branch `%s`\n", swoopBranch, currBranch)
}

func reCheckout(currBranch string) error {
	reCheckoutCmd := exec.Command("git", "checkout", currBranch)
	reCheckoutOutput, err := reCheckoutCmd.CombinedOutput()
	if err != nil {
		fmt.Printf("\n%v\n", string(reCheckoutOutput))
		fmt.Printf("\nERROR: Unable to return to original branch `%s`\n", currBranch)
	}

	return err
}
