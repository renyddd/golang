package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list [dir]",
	Short: "just imitate ls",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			os.Exit(1)
		}
		// args come after LIST command
		f, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
		}

		fs, err := f.Readdir(-1)
		for _, v := range fs {
			fmt.Println(v.Name())
		}
	},
}
