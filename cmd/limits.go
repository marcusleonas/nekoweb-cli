/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"nekoweb-cli/internal/nekoweb"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
)

// limitsCmd represents the limits command
var limitsCmd = &cobra.Command{
	Use:   "limits",
	Short: "Get your current rate limit status",
	Long:  `Get your current rate limit status`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := nekoweb.GetAPIKey()
		client := nekoweb.NewClient(apiKey)

		res, err := client.GetLimits()
		if err != nil {
			fmt.Println("something went wrong", err)
			os.Exit(1)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

		fmt.Fprintf(w, "Category\tLimit\tRemaining\tReset\n") // header
		fmt.Fprintf(w, "--------\t-----\t---------\t-----\n") // optional underline

		fmt.Fprintf(
			w,
			"General\t%d\t%d\t%s\n",
			res.General.Limit, res.General.Remaining,
			formatReset(res.General.Reset),
		)

		fmt.Fprintf(
			w,
			"Big Uploads (over 100MB)\t%d\t%d\t%s\n",
			res.BigUploads.Limit, res.BigUploads.Remaining,
			formatReset(res.BigUploads.Reset),
		)

		fmt.Fprintf(
			w,
			"ZIPs\t%d\t%d\t%s\n",
			res.Zip.Limit, res.Zip.Remaining,
			formatReset(res.Zip.Reset),
		)

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(limitsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// limitsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// limitsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func formatReset(reset int64) string {
	if reset <= 0 {
		return "n/a"
	}
	return time.UnixMilli(reset).Format(time.RFC1123)
}
