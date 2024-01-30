package cmd

import (
	"os"

	"github.com/mateusmatinato/goexpert-stress-test/internal/stresstest"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goexpert-stress-test",
	Short: "Execute Stress Test on a given URL with concurrency",
	Long: `Execute Stress Test on a given URL with concurrency. Receives a URL, concurrency and number of 
requests as parameters. Returns the total number of requests, duration and HTTP status code group.`,
	Run: func(cmd *cobra.Command, args []string) {
		url := cmd.Flag("url").Value.String()
		concurrency, _ := cmd.Flags().GetInt("concurrency")
		requests, _ := cmd.Flags().GetInt("requests")

		params := stresstest.Params{
			URL:         url,
			Concurrency: concurrency,
			Requests:    requests,
		}
		params.PrintStart()
		uc := stresstest.NewUseCase()
		resp := uc.Execute(params)
		resp.PrintResult()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "URL to be tested")
	rootCmd.Flags().IntP("concurrency", "c", 10, "Number of concurrent requests")
	rootCmd.Flags().IntP("requests", "r", 10, "Number of requests")

	_ = rootCmd.MarkFlagRequired("url")
}
