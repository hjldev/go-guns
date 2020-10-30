package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-guns/cmd/admin"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "go-guns",
	Short:        "go-guns",
	SilenceUsage: true,
	Long:         `go-guns`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {

		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", `欢迎使用go-guns系统`)
	},
}

func init() {
	// 所有启动命令在此添加
	rootCmd.AddCommand(admin.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
