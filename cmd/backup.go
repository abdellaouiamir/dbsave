/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("backup called")
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
		outputFile , _ := cmd.Flags().GetString("outputFile")
		fmt.Println(host, port, user, password, dataBase, allDb, outputFile)
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	// Local Flags
	backupCmd.Flags().String("host", "localhost", "the host of the data base")
	backupCmd.Flags().String("port", "", "the data base port")
	backupCmd.Flags().String("user", "", "username for the data base")
	backupCmd.Flags().String("password", "", "password for the data base")
	backupCmd.Flags().String("database", "", "database name")
	backupCmd.Flags().Bool("all-databases", false, "all databases")
	backupCmd.Flags().StringP("outputFile", "o", "", "the location of the output file")
	backupCmd.MarkFlagRequired("user")
}
