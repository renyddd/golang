package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/klog"
)

func init() {
	rootCmd.PersistentFlags().StringP("all", "a", "", "say something.")
}

var rootCmd = &cobra.Command{
	Use:   "utils",
	Short: "utils is a collection.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		klog.Error(err)
	}
}
