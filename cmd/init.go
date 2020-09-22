package cmd

import (
	"bands/pkg/core"
	"fmt"

	au "github.com/logrusorgru/aurora"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	force       bool
	passedEmail string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		accResp, statusCode, err := core.AccountCreate(passedEmail, force, debug)

		if err != nil {
			fmt.Println(au.Red(au.Bold(fmt.Sprintf("Error %d during up %s", statusCode, file))))
			return
		}

		responseApiAccessToken := accResp.Data.ApiAccessToken
		responseApiActivationToken := accResp.Data.ApiActivationToken

		if responseApiAccessToken != "" {
			viper.Set("email", passedEmail)
			viper.Set("token", responseApiAccessToken)
			if err := viper.WriteConfig(); err != nil {
				fmt.Println(err.Error())
			}

			fmt.Println(``)
			fmt.Println(au.Green(au.Bold(fmt.Sprintf("Your Bands.sh API Account was created!"))))
			fmt.Println(au.Green(au.Bold(fmt.Sprintf("For your records, we also emailed you the below onboarding link."))))
			fmt.Println(``)
			fmt.Println("Please follow this link to complete onboarding your account setup:")
			fmt.Println(au.Bold(fmt.Sprintf("https://%s/onboard/start/?access_key=%s&activation_token=%s", checkoutURL, responseApiAccessToken, responseApiActivationToken)))
			fmt.Println(``)
		} else {
			fmt.Println(``)
			fmt.Println(au.Green(au.Bold(fmt.Sprintf("Your Bands.sh API Account is already active and token is set."))))
			fmt.Println(``)
			fmt.Println("If you need to reset your access token run:", au.Bold(fmt.Sprintf("$ bands init --email <email> --force")))
			fmt.Println(``)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.PersistentFlags().BoolVar(&force, "force", false, "force reset your account")
	rootCmd.PersistentFlags().StringVar(&passedEmail, "email", "", "Your email address.")
}
