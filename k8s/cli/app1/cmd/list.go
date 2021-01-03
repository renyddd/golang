package cmd

//
//import (
//	"fmt"
//	"github.com/spf13/cobra"
//	"log"
//	"os"
//)
//
//type listCmd struct {
//	cmd *cobra.Command
//}
//
//func NewListCmd() *listCmd {
//	return &listCmd{
//		cmd: &cobra.Command{
//			Use: "[path]",
//			Run: func(cmd *cobra.Command, args []string) {
//				file, err := os.Getwd()
//				if err != nil {
//					log.Fatalln(err)
//				}
//				fmt.Println("you list:", file)
//			},
//		},
//	}
//}
//
//func (l *listCmd) Execute() {
//	if err := l.cmd.Execute(); err != nil {
//		log.Fatalln(err)
//	}
//}
