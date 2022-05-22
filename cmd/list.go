package cmd

import (
	"log"

	"github.com/Nikkely/dovadova/fetcher"
	"github.com/spf13/cobra"
)

const (
	noHeadlessFlg = "no-headless"
	outputFlg     = "output"
)

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().Bool(noHeadlessFlg, false, "run chrome (no headless)")
	listCmd.PersistentFlags().StringP(outputFlg, "o", "output/cardlist.json", "run chrome (no headless)")
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "fetch cardlist",
	Long:  `You can get cardlist. Dovadova fetch with headless chrome.`,
	Run: func(cmd *cobra.Command, args []string) {
		noHeadless, err := cmd.Flags().GetBool(noHeadlessFlg)
		if err != nil {
			log.Fatalln(err.Error())
		}
		output, err := cmd.Flags().GetString(outputFlg)
		if err != nil {
			log.Fatalln(err.Error())
		}

		if err := fetcher.FetchCardList(output, !noHeadless); err != nil {
			log.Fatalln(err.Error())
		}
	},
}
