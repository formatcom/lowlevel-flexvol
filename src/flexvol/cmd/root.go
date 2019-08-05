package cmd

import (
	"fmt"
	"errors"
	"os"
	"path/filepath"
	"github.com/spf13/cobra"
	"github.com/sirupsen/logrus"
)

var (
	APP string
	//The verbose flag value
	verbose string
)

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {

		// invoke help
		cmd.HelpFunc()(cmd, []string{})
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		level, err := logrus.ParseLevel(verbose)
		if err != nil {
			return errors.New(fmt.Sprintf(
				"not a valid Level: %v\n", verbose))
		}
		logrus.SetLevel(level)

		file, err := os.OpenFile(
				fmt.Sprintf("/var/log/%v.log", APP),
				os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			logrus.Info("Failed to log to file, using default stderr")
		}else {
			logrus.SetOutput(file)
		}

		return nil
	},

}

func init() {
	rootCmd.PersistentFlags().StringVarP(&verbose, "verbosity", "v",
		logrus.WarnLevel.String(),
		"Log level (debug, info, warn, error, fatal, panic")
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
