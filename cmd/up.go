package cmd

import (
	"bands/pkg/core"
	"fmt"

	au "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		token = viper.GetString("token")
		email = viper.GetString("email")
		file, _ = cmd.Flags().GetString("file")

		if file != "" {
			fmt.Println(``)
			fmt.Println(fmt.Sprintf("Deploying %s...", file))

			yaml, _, statusCode, err := core.ActionUp(email, token, file)

			if err != nil {
				fmt.Println(au.Red(au.Bold(fmt.Sprintf("Error %d during up %s", statusCode, file))))
				return
			}

			if statusCode == 200 {
				url = yaml.Data.CheckoutUrl
				viper.Set("url", url)
				viper.WriteConfig()
				fmt.Println(au.Green(au.Bold("Success!")))
				fmt.Println(``)
				fmt.Println("Your checkout page is live at", au.Bold(url))
				fmt.Println(``)
				fmt.Println("You can also embed your checkout page into your own website:")
				fmt.Printf(`<iframe src="%s" frameborder="0" allowfullscreen style="width:100%%;height:100%%;"><iframe>`, url)
				fmt.Println(``)
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
	rootCmd.AddCommand(upCmd)

	upCmd.Flags().StringP("file", "f", "", "bands up -f <payments.yaml>")
}
