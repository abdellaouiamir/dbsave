package db

import (
	"errors"
	"fmt"
)

type mongoDB struct {
	hostname string
	port string
	user string
	password string
	database string
	alldatabase bool
	compress bool
	outputFile string
}

func (db mongoDB) Backup(t StrategieType) error {
	switch t{
	case Dump:
		fmt.Println("dump")
	case Increment:
		fmt.Println("increment")
	case File:
		fmt.Println("file")
	}
	return errors.New("not supported backup type")
}

func (db mongoDB) Restore(t StrategieType) error {
	switch t{
	case Dump:
		fmt.Println("dump")
	case Increment:
		fmt.Println("increment")
	case File:
		fmt.Println("file")
	}
	return errors.New("not supported restore type")
}

func (db mongoDB) Connection() bool {
	return true
}
