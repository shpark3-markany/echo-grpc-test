/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"local/fin/api"
	"local/fin/configs"
	"log"

	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Application run. Type is web",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("cmd.web.Run: %v", r)
			}
		}()
		api.OpenWeb(configs.WEB_PORT)
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}
