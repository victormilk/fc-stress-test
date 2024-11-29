/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/victormilk/fc-stress-test/request"
)

var rootCmd = &cobra.Command{
	Use:   "fc-stress-test",
	Short: "A simple stress test tool for making requests to a URL",
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")

		requestsService := request.NewRequestService()
		requestsService.MakeRequests(url, requests, concurrency)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "URL to make requests")
	rootCmd.Flags().IntP("requests", "r", 0, "Total requests to make")
	rootCmd.Flags().IntP("concurrency", "c", 0, "Concurrency of requests")
	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagRequired("requests")
	rootCmd.MarkFlagRequired("concurrency")
}
