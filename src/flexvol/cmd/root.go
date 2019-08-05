package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
)

var (
	APP string
)

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {

		// invoke help
		cmd.HelpFunc()(cmd, []string{})
	},
}

func Execute(version string) {
	APP = filepath.Clean(os.Args[0])

	rootCmd.Use = APP
	rootCmd.Short = fmt.Sprintf("%v kubelet\n", APP)
	rootCmd.Version = version

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
