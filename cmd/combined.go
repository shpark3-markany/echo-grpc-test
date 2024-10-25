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

var combinedCmd = &cobra.Command{
	Use:   "combined",
	Short: "Application run. Type is combined",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("cmd.conbined.Run: %v", r)
			}
		}()
		go api.OpenRpc(configs.GRPC_PORT)
		api.OpenWeb(configs.WEB_PORT)
	},
}

func init() {
	rootCmd.AddCommand(combinedCmd)
}
