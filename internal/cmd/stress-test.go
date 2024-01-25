package cmd

import (
	stress_test "github.com/gabrielsc1998/go-stress-test/internal/app"
	"github.com/spf13/cobra"
)

func test() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")
		stressTest := stress_test.New(url, requests, concurrency)
		stressTest.Run()
		return nil
	}
}

func StressTestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "test <url> <requests> <concurrency>",
		Example: `test --url=http://localhost:8080 --requests=100 --concurrency=10`,
		RunE:    test(),
	}
	var requiredFlags = []string{"url", "requests", "concurrency"}
	for _, flag := range requiredFlags {
		cmd.MarkPersistentFlagRequired(flag)
	}
	cmd.PersistentFlags().String("url", "", "Specify the URL")
	cmd.PersistentFlags().Int("requests", 0, "Specify the number of requests")
	cmd.PersistentFlags().Int("concurrency", 0, "Specify the number of concurrency")
	return cmd
}
