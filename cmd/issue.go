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

// issueCmd represents the issue command
var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Issue a book",
	Long:  `Issue a book from the library.`,
	Run: func(cmd *cobra.Command, args []string) {
		books := &lib.Library{}
		if err := books.Load(libFile); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
		size := len(args)
		s := ""
		for idx := range args {
			if idx < size {
				s = s + args[idx]
				s = s + " "
			}
		}
		title := s

		err := books.Issue(title)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = books.Store(libFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(issueCmd)
}
