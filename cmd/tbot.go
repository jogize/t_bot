/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	//Teletoken bot
	TeleToken = os.Getenv("TELE_TOKEN")
)

// tbotCmd represents the tbot command
var tbotCmd = &cobra.Command{
	Use:   "tbot",
	Short: "A brief description of your command",
	Aliases: []string{"start"},
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("tbot %s started", appVersion)
		b, err := telebot.NewBot(telebot.Settings{
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env variable. %s ", err)
			return
		}

		b.Handle(telebot.OnText, func(c telebot.Context) error {
			log.Print(c.Message().Payload, c.Text())
			payload := c.Message().Payload

			switch payload {
			case "hello": 
				err = c.Send(fmt.Sprintf("Hello I'm tbot %s!", appVersion))
			}

			return err
		})

		// b.Handle("/start", func(c telebot.Context) error {
		// 	return c.Send("Hello world!")
		// })

		b.Start()
	},
}

func init() {
	rootCmd.AddCommand(tbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
