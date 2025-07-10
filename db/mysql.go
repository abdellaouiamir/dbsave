package db

import (
	"errors"
	"fmt"
)

type mysqlDB struct {
	hostname string
	port string
	user string
	password string
	database string
	alldatabase bool
	compress bool
	outputFile string
}

func (db mysqlDB) Backup(t StrategieType) error {
	switch t{
	case Dump:
		fmt.Println("dump")
		return nil
	case Increment:
		fmt.Println("increment")
		return nil
	case File:
		fmt.Println("file")
		return nil
	}
	return errors.New("unsuported backup type")
}

func (db mysqlDB) Restore(t StrategieType) error {
	switch t{
	case Dump:
		fmt.Println("dump")
		return nil
	case Increment:
		fmt.Println("increment")
		return nil
	case File:
		fmt.Println("file")
		return nil
	default:
		return errors.New("unsuported restore type")
	}
}

func (db mysqlDB) Connection() bool {
	return true
}
