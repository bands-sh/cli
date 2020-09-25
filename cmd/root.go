package cmd

import (
	"fmt"
	"log"
	"os"

	au "github.com/logrusorgru/aurora"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	debug       bool
	checkoutURL = "checkout.bands.sh"
	websiteURL  = "bands.sh"
	tagline     = "Payments as Code"
	email       string
	token       string
	setToken    string
	retToken    string
	retActive   bool
	file        string
	url         string
)

var rootCmd = &cobra.Command{
	Use:   "bands",
	Short: tagline,
	Long: `` +
		tagline + `.

Go from a yaml file to a fully working payment page. More at ` + websiteURL + `
	
Getting started:
  $ bands init --email <email>`,
	Run: func(cmd *cobra.Command, args []string) {

		cmd.Help()

		if email == "" {
			fmt.Println(au.Green(au.Bold("Try $ bands init --email <email>")))
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.bands.yaml")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "run in debug mode")
}

func initConfig() {
	home, err := homedir.Dir()
	configPath := fmt.Sprintf("%s/.bands.yaml", home)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".bands.yaml")
		viper.SetConfigType("yaml")
		viper.SetConfigFile(configPath)

		// Search config in home directory with name ".bands.yaml".
		if err := viper.ReadInConfig(); err != nil {
			_, file_err := os.OpenFile(configPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)

			if file_err != nil {
				log.Fatal(file_err)
			}

			fmt.Println("Global config file created in", home)
			viper.Set("email", "")
			viper.Set("token", "")

			if err := viper.WriteConfig(); err != nil {
				fmt.Println(err.Error())
			}
		}
	}

	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if viper.GetString("email") != "" {
			email = viper.GetString("email")
		}

		if viper.GetString("token") != "" {
			token = viper.GetString("token")
		}

		if debug == true {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
			fmt.Println(email + " " + token)
		}
	}
}
