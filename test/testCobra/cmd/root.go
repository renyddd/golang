package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use: "mycmd fd",
	Short: "mycmd is just a test",
	Long:"xxxxxxxxx jfjj f asdf saf asf" +
		"fasdf adf asdfs d",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting mycmd")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

