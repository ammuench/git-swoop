package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 2 {
		log.Printf("ERROR: Need a singular branch name to swoop from")
		os.Exit(126)
	}

	swoopBranch := os.Args[1]

	currBranch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		log.Panic("Error! Not in a git repository")
	}

	fmt.Printf("\nYou are on branch %s\n", currBranch)
	fmt.Printf("\nYou want to swoop on branch %s\n", swoopBranch)

	checkoutCmd := exec.Command("git", "checkout", swoopBranch)
	err = checkoutCmd.Run()
	if err != nil {
		fmt.Printf("\nERROR: Unable to checkout swoop branch `%s`", swoopBranch)
		os.Exit(126)
	}

	pullCmd := exec.Command("git", "pull", swoopBranch)
	err = pullCmd.Run()
	if err != nil {
		fmt.Printf("\nERROR: Unable to pull from swoop branch `%s`", swoopBranch)
    reCheckoutErr := reCheckout(string(currBranch))
    if reCheckoutErr != nil {

    fmt.Printf("\nStill on swoop branch `%s`", swoopBranch)
    } else {

    fmt.Printf("\nReturned to original branch `%s`", string(currBranch))
    }
		os.Exit(126)
	}

  err = reCheckout(string(currBranch))
  if err != nil {
    fmt.Printf("\nStill on swoop branch `%s`", swoopBranch)
    os.Exit(126);
  }

	fmt.Printf("\nSuccessfully swooped from branch `%s` and returned to branch `%s`", swoopBranch, string(currBranch))
}

func reCheckout(currBranch string) (error){
	reCheckoutCmd := exec.Command("git", "checkout", string(currBranch))
	err := reCheckoutCmd.Run()
	if err != nil {
		fmt.Printf("\nERROR: Unable to return to original branch `%s`", string(currBranch))
	}

	return err;
}
