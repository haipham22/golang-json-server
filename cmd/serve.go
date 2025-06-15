package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"jsonserver/internal/serve"
)

// serveServerCmd represents the binance command
var serveServerCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

This is sample command.`,
	Run: func(cmd *cobra.Command, _ []string) {
		log := zap.S()

		port, err := cmd.Flags().GetInt64("port")
		if err != nil {
			log.Fatal("Invalid port config")
		}

		app, err := serve.InitServeApp(log)

	},
}

func init() {
	rootCmd.AddCommand(serveServerCmd)

	serveServerCmd.Flags().Int64("port", 8000, "API port listening")
}
