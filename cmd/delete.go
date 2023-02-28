package cmd

import (
	"fmt"
	"os"

	"github.com/Uttkarsh-raj/library/lib"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "del",
	Short: "Remove a book",
	Long:  `Remove any particular book from the library.`,
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
		err := books.Delete(title)
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
	rootCmd.AddCommand(deleteCmd)
}
