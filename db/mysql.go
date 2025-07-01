package db

import "fmt"

type mysqlDB struct {
	hostname string
	port string
	user string
	password string
	database string
}

func (db mysqlDB) Backup(t StrategieType) bool {
	switch t{
	case Dump:
		fmt.Println("dump")
	case Increment:
		fmt.Println("increment")
	case File:
		fmt.Println("file")
	}
	return false
}

func (db mysqlDB) Restore(t StrategieType) bool {
	switch t{
	case Dump:
		fmt.Println("dump")
	case Increment:
		fmt.Println("increment")
	case File:
		fmt.Println("file")
	}
	return true
}

func (db mysqlDB) Connection() bool {
	return true
}
