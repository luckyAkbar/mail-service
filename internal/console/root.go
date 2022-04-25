package console

import (
	"os"

	"mail-service/internal/config"

	"github.com/spf13/cobra"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	log "github.com/sirupsen/logrus"
)

var RootCmd = &cobra.Command{
	Use: "mail-service",
}

// Execute: :nodoc:
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func init() {
	setupLogger()
}

func setupLogger() {
	formatter := runtime.Formatter{
		ChildFormatter: &log.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		},
		Line: true,
		File: true,
	}

	log.SetFormatter(&formatter)
	log.SetOutput(os.Stdout)

	logLevel, err := log.ParseLevel(config.LogLevel())

	if err != nil {
		logLevel = log.DebugLevel
	}

	log.SetLevel(logLevel)
}
