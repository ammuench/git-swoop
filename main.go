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
	"strings"
)

// Exit codes
const (
	ExitSuccess = 0
	ExitFailure = 1
)

// Flag constants
const (
	FlagVersion      = "--version"
	FlagVersionShort = "-v"
	FlagVersionAlt   = "-version"
	FlagHelp         = "--help"
	FlagHelpShort    = "-h"
	FlagHelpAlt      = "-help"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("ERROR: Need a singular branch name to swoop from\n")
		printHelpInfo()
		os.Exit(ExitFailure)
	}

	swoopBranch := os.Args[1]

	// Handle flags
	if handleFlags(swoopBranch) {
		return
	}

	// Verify we're in a git repository
	if err := verifyGitRepo(); err != nil {
		fmt.Println("ERROR: Not in a git repository")
		os.Exit(ExitFailure)
	}

	// Get current branch
	currBranch, err := getCurrentBranch()
	if err != nil {
		fmt.Println("ERROR: Unable to determine current working branch:", err)
		os.Exit(ExitFailure)
	}

	// Checkout target branch
	checkoutOutput, err := checkoutBranch(swoopBranch)
	if err != nil {
		fmt.Printf("\n%v\n", string(checkoutOutput))
		fmt.Printf("\nERROR: Unable to checkout swoop branch `%s`\n", swoopBranch)
		fmt.Printf("\nStill on original branch `%s`\n", currBranch)
		os.Exit(ExitFailure)
	}

	// Pull latest changes
	pullCmd := exec.Command("git", "pull")
	pullOutput, err := pullCmd.CombinedOutput()
	if err != nil {
		fmt.Printf("\n%v\n", string(pullOutput))
		fmt.Printf("\nERROR: Unable to pull from swoop branch `%s`\n", swoopBranch)

		// Try to go back to original branch
		reCheckoutOutput, reCheckoutErr := checkoutBranch(currBranch)
		if reCheckoutErr != nil {
			fmt.Printf("\n%v\n", string(reCheckoutOutput))
			fmt.Printf("\nERROR: Unable to return to original branch `%s`\n", currBranch)
			fmt.Printf("\nStill on swoop branch `%s`\n", swoopBranch)
		} else {
			fmt.Printf("\nReturned to original branch `%s`\n", currBranch)
		}
		os.Exit(ExitFailure)
	} else {
		fmt.Printf("\n%v\n", string(pullOutput))
	}

	// Return to original branch
	_, err = checkoutBranch(currBranch)
	if err != nil {
		fmt.Printf("\nStill on swoop branch `%s`\n", swoopBranch)
		os.Exit(ExitFailure)
	}

	fmt.Printf("\nSuccessfully swooped from branch `%s` and returned to branch `%s`\n", swoopBranch, currBranch)
}

// processes command line flags and returns true if a flag was handled
func handleFlags(flag string) bool {
	switch flag {
	case FlagVersion, FlagVersionShort, FlagVersionAlt:
		printVersionInfo()
		os.Exit(ExitSuccess)
		return true
	case FlagHelp, FlagHelpShort, FlagHelpAlt:
		printHelpInfo()
		os.Exit(ExitSuccess)
		return true
	default:
		return false
	}
}

func printVersionInfo() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		fmt.Println("ERROR: Unable to read build information")
		return
	}

	swoopVersionParsed := swoopVersion
	if swoopVersionParsed == "" {
		swoopVersionParsed = info.Main.Version
	}

	goVersionParsed := goVersion
	if goVersionParsed == "" {
		goVersionParsed = info.GoVersion
	}

	fmt.Printf("\ngit-swoop %s, built with %s\n", swoopVersionParsed, goVersionParsed)
	fmt.Println("")
	fmt.Println("git-swoop Copyright (C) 2025 Alex Muench")
	fmt.Println("This program comes with ABSOLUTELY NO WARRANTY")
	fmt.Println("This is free software, and you are welcome to redistribute it")
	fmt.Println("under certain conditions; check the LICENSE.md file at `https://github.com/ammuench/git-swoop`")
}

func printHelpInfo() {
	fmt.Println("")
	fmt.Println("usage: git-swoop <target-branch-name>")
	fmt.Println("   git-swoop will try to checkout your target branch, pull down the latest from remote,")
	fmt.Println("   and then return to where you started")
	fmt.Println("")
	fmt.Println("flags:")
	fmt.Println("   --help (alias: -h, -help): prints this message")
	fmt.Println("   --version (alias: -v, -version): prints the version and basic package info")
	fmt.Println("")
}

// verifyGitRepo checks if the current directory is a git repository by checking `git status`
func verifyGitRepo() error {
	_, err := exec.Command("git", "status").Output()
	return err
}

func getCurrentBranch() (string, error) {
	currBranchBytes, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(currBranchBytes)), nil
}

func checkoutBranch(branch string) ([]byte, error) {
	checkoutCmd := exec.Command("git", "checkout", branch)
	return checkoutCmd.CombinedOutput()
}
