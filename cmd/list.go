/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Uttkarsh-raj/library/lib"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the books",
	Long:  `List all the books in the library and there status.`,
	Run: func(cmd *cobra.Command, args []string) {
		books := &lib.Library{}
		if err := books.Load(libFile); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
		books.List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
