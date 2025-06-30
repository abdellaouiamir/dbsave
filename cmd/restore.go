/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("restore called")
		fmt.Println("compress", compress, "verbose", verbose, "queit", quiet, "version", version)
		for _, arg := range args {
			fmt.Println(arg)
		}
		host , _ := cmd.Flags().GetString("host")
		port , _ := cmd.Flags().GetString("port")
		user , _ := cmd.Flags().GetString("user")
		password , _ := cmd.Flags().GetString("password")
		dataBase , _ := cmd.Flags().GetString("database")
		allDb , _ := cmd.Flags().GetBool("all-databases")
		fmt.Println(host, port, user, password, dataBase, allDb)
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
	// Local Flags
	restoreCmd.Flags().String("host", "localhost", "the host of the data base")
	restoreCmd.Flags().String("port", "", "the data base port")
	restoreCmd.Flags().String("user", "", "username for the data base")
	restoreCmd.Flags().String("password", "", "password for the data base")
	restoreCmd.Flags().String("database", "", "database name")
	restoreCmd.Flags().Bool("all-databases", false, "all databases")
	restoreCmd.MarkFlagRequired("user")
}
