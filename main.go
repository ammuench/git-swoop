package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("ERROR: Need a singular branch name to swoop from")
		os.Exit(126)
	}

	swoopBranch := os.Args[1]

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
