package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "CMD",
	Short: "just type your command",
	// 注视调 Run 函数可以阻止直接运行 rootCmd
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Println("your cmds is", args)
	//},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
