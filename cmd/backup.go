/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/abdellaouiamir/dbsave.git/db"
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
		for _, arg := range args {
			fmt.Println(arg)
		}
		dbmsType , _ := cmd.Flags().GetString("dbms")
		backupType , _ := cmd.Flags().GetString("type")
		host , _ := cmd.Flags().GetString("host")
		port , _ := cmd.Flags().GetString("port")
		user , _ := cmd.Flags().GetString("user")
		password , _ := cmd.Flags().GetString("password")
		dataBase , _ := cmd.Flags().GetString("database")
		allDb , _ := cmd.Flags().GetBool("all-databases")
		compress , _ := cmd.Flags().GetBool("compress")
		outputFile , _ := cmd.Flags().GetString("outputFile")
		fmt.Println(host, port, user, password, dataBase, allDb, outputFile)
		var newDB db.DbObject
		newDB.Hostname = host
		newDB.Port = port
		newDB.User = user
		newDB.Password = password
		newDB.DataBase = dataBase
		newDB.AllDataBase = allDb
		newDB.Compress = compress
		newDB.OutputFile = outputFile
		// select the dbms type
		var typeDbms db.DBMS
		switch dbmsType {
		case "postgresql":
			typeDbms = db.Postgresql
		case "mysql":
			typeDbms = db.Mysql
		case "mongo":
			typeDbms = db.Mongo
		default:
			fmt.Println("not supported type")
		}
		// create the dbms object
		newDB.Type = typeDbms
		dbms, err := db.NewDB(newDB) 
		if err != nil {
			fmt.Println(err)
		}
		// select the backup type
		var typeBackup db.StrategieType
		switch backupType {
			case "dump":
				typeBackup = db.Dump
			default:
				fmt.Println("bad choise")
		}
		// Run the backup function
		if err:= dbms.Backup(typeBackup);err != nil {
			fmt.Println("bad news")
		}
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)
	// Local Flags
	backupCmd.Flags().String("dbms", "", "the type of DBMS")
	backupCmd.Flags().String("type", "", "the type of resotre")
	backupCmd.Flags().String("host", "", "the host of the data base")
	backupCmd.Flags().String("port", "", "the data base port")
	backupCmd.Flags().String("user", "", "username for the data base")
	backupCmd.Flags().String("password", "", "password for the data base")
	backupCmd.Flags().String("database", "", "database name")
	backupCmd.Flags().Bool("all-databases", false, "all databases")
	backupCmd.Flags().Bool("compress", false, "use compressing")
	backupCmd.Flags().StringP("outputFile", "o", "", "the location of the output file")
	backupCmd.MarkFlagRequired("user")
}
