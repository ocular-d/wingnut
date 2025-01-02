/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("info called")
		fmt.Printf("OS: %s\n", runtime.GOOS)
		fmt.Printf("Arch: %s\n", runtime.GOARCH)
		fmt.Printf("Go Version: %s\n\n", runtime.Version())
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
