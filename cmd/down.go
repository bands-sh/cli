package cmd

import (
	"bands/pkg/core"
	"fmt"

	au "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Deactive the checkout page",
	Long: `This command will deactivate the checkout page.
It will also archive all objects created from the yaml input.

Nothing will get permanently deleted, and this checkout 
page can be brought back up with "bands up" command.`,
	Run: func(cmd *cobra.Command, args []string) {
		token = viper.GetString("token")
		email = viper.GetString("email")
		file, _ = cmd.Flags().GetString("file")

		if file != "" {
			fmt.Println(``)
			fmt.Println(fmt.Sprintf("Deactivating %s...", file))

			yaml, _, statusCode, err := core.ActionDown(email, token, file)

			if err != nil {
				fmt.Println(au.Red(au.Bold(fmt.Sprintf("Error %d during up %s", statusCode, file))))
				return
			}

			if statusCode == 200 {
				url = yaml.Data.CheckoutUrl
				viper.WriteConfig()
				fmt.Println(``)
				fmt.Println(au.Green(au.Bold("Success!")), "Checkout page", au.Bold(url), "has been deactivated.")
				fmt.Println(``)
			} else {
				fmt.Println(``)
				fmt.Println(au.Red(au.Bold(yaml.Message)))
				fmt.Println(``)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(downCmd)
	downCmd.Flags().StringP("file", "f", "", "bands up -f <payments.yaml>")
	downCmd.MarkFlagRequired("file")
}
