package cmd

import (
	"bands/pkg/core"
	"fmt"

	au "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		token = viper.GetString("token")
		email = viper.GetString("email")
		file, _ = cmd.Flags().GetString("file")

		if file != "" {
			fmt.Println(``)
			fmt.Println(fmt.Sprintf("Checking %s...", file))
			fmt.Println(``)

			yaml, _, statusCode, err := core.ActionStatus(email, token, file)

			if err != nil {
				fmt.Println(au.Red(au.Bold(fmt.Sprintf("Error %d during up %s", statusCode, file))))
				return
			}

			message := yaml.Message
			url = yaml.Data.CheckoutUrl
			active := yaml.Data.Active

			if statusCode == 200 {
				if active == 1 {
					fmt.Println(au.Green(au.Bold("Your checkout page is up and running!")), au.Bold(url))
					fmt.Println(``)
					fmt.Println("You can also embed your checkout page into your own website:")
					fmt.Printf(`<iframe src="%s" frameborder="0" allowfullscreen style="width:100%%;height:100%%;"><iframe>`, url)
					fmt.Println(``)
					fmt.Println(``)
				}

				if active == 0 {
					fmt.Println(message)
					fmt.Println(``)
				}
			} else {
				fmt.Println(``)
				fmt.Println(au.Red(au.Bold(message)))
				fmt.Println(``)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	statusCmd.Flags().StringP("file", "f", "", "bands up -f <payments.yaml>")
}
