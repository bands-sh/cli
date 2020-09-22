package cmd

import (
	"bands/pkg/core"
	"fmt"

	au "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		token = viper.GetString("token")
		email = viper.GetString("email")
		file, _ = cmd.Flags().GetString("file")

		if file != "" {
			fmt.Println(``)
			fmt.Println(fmt.Sprintf("Bringing it down %s...", file))

			yaml, _, statusCode, err := core.ActionDown(email, token, file)

			if err != nil {
				fmt.Println(au.Red(au.Bold(fmt.Sprintf("Error %d during up %s", statusCode, file))))
				return
			}

			if statusCode == 200 {
				viper.WriteConfig()
				fmt.Println(au.Green(au.Bold("Success!")))
				fmt.Println(``)
				fmt.Println("Your checkout page was brought down.")
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
}
