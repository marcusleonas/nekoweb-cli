/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"nekoweb-cli/internal/nekoweb"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get information about a nekoweb site",
	Long:  `Get information about a nekoweb site`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := nekoweb.GetAPIKey()
		client := nekoweb.NewClient(apiKey)

		domain, err := cmd.Flags().GetString("domain")
		if err != nil {
			fmt.Println("something went wrong", err)
			os.Exit(1)
		}

		allSites, err := cmd.Flags().GetBool("all")
		if err != nil {
			fmt.Println("something went wrong", err)
			os.Exit(1)
		}

		if allSites {
			sitesInfo, err := client.GetAllSitesInfo()
			if err != nil {
				fmt.Println("something went wrong", err)
				os.Exit(1)
			}
			for _, siteInfo := range *sitesInfo {
				printSiteInfo(&siteInfo)
			}
		} else {
			siteInfo, err := client.GetSiteInfo(domain)
			if err != nil {
				fmt.Println("something went wrong", err)
				os.Exit(1)
			}
			printSiteInfo(siteInfo)
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	infoCmd.Flags().StringP("domain", "d", "", "domain of the site you want info of. don't specify to get your 'main' site.")
	infoCmd.Flags().BoolP("all", "a", false, "list info for all your sites")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func printSiteInfo(siteInfo *nekoweb.SiteInfo) {
	fmt.Print(color.WhiteString("----- %v -----\n", siteInfo.Domain))

	formatted := color.GreenString(
		"Title: %s\nUpdates: %d\nFollowers: %d\nViews: %d\nCreated At: %d\nUpdated At: %d",
		siteInfo.Title,
		siteInfo.Updates,
		siteInfo.Followers,
		siteInfo.Views,
		siteInfo.CreatedAt,
		siteInfo.UpdatedAt,
	)

	fmt.Println(formatted)
}
