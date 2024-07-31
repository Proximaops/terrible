/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "terrible",
	Short: "Ansible in terraform with ease",
	Long:  `Terrible is a tool that allows you to use ansible playbooks in hosts deployed with terraform on many hosting providers with ease in AWX.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			`terrible is tool to execute terraform scripts and run ansible playbooks on the deployed hosts

Find more information at: https://github.com/Proximaops/terrible
Usage:
  terrible [file] [flags]

Use "terrible --help" for more information.`)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
