/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// listRepoCmd represents the listRepo command
var listRepoCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listRepo called")
		// Prepare the command
		listRepoCmd := exec.Command("gh", "repo", "list")

		// Run the command and capture the output
		output, err := listRepoCmd.CombinedOutput()
		if err != nil {
			log.Fatalf("Error running 'gh repo list': %v", err)
		}
		// Convert the output into a string
		repoList := string(output)
		// Print the repositories
		if len(repoList) == 0 {
			fmt.Println("No repositories found.")
		} else {
			// Optionally, format the output
			fmt.Println("Repositories:")
			// If the output has multiple lines (one repo per line), you could do extra formatting
			repoLines := strings.Split(repoList, "\n")
			for _, line := range repoLines {
				if line != "" {
					fmt.Println("  -", line)
				}
			}
		}
	},
}

func init() {
	repoCmd.AddCommand(listRepoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listRepoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listRepoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
