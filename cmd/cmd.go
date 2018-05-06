package cmd

import (
	"github.com/ckeyer/logrus"
	"github.com/funxdata/landlady/global"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "landlady",
		Short: "",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logrus.SetFormatter(&logrus.JSONFormatter{})
			if global.IsDebug() {
				logrus.SetLevel(logrus.DebugLevel)
			}
			logrus.Debugf("debug ?: %v", global.IsDebug())
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&global.Debug, "debug", "D", false, "debug")
}

func Execute() {
	rootCmd.Execute()
}
