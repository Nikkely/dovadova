package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dovadova",
	Short: "dovadova is a scraping tool for shadoverse-evolve cardlist",
	Long: `dovadova get all card. please check robots.txt before running`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
