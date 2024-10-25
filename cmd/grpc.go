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

var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Application run. Type is grpc",
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("cmd.grpc.Run: %v", r)
			}
		}()
		api.OpenRpc(configs.GRPC_PORT)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
}
