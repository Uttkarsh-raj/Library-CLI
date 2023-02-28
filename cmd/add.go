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

const libFile = ".library.json"

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new book",
	Long:  `Add a new book to the existing library`,
	Run: func(cmd *cobra.Command, args []string) {
		books := &lib.Library{}
		if err := books.Load(libFile); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
		size := len(args)
		s := ""
		for idx := range args {
			if idx < size-1 {
				s = s + args[idx]
				s = s + " "
			}
		}
		title := s
		genre := args[size-1]

		err := books.Add(title, genre)
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
	rootCmd.AddCommand(addCmd)
}
