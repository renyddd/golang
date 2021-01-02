// https://weread.qq.com/web/reader/f1e3207071eeeefaf1e138ak636320102206364d3f0ffdc
package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var Version bool
	var rootCmd = &cobra.Command{
		Use: "root [sub]",
		Short: "root command",
		Example: "xxx cccc ccccc",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Inside rootCmd run with args: %v\n", args)
			if Version {
				fmt.Println("Version: 1.0")
			}
		},
	}

	flags := rootCmd.Flags()
	flags.BoolVarP(&Version, "version", "v", false, "Print version info" +
		"rmation and quit")
	_ = rootCmd.Execute()
}