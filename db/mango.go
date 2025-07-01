package db

import "fmt"

type mongoDB struct {
	hostname string
	port string
	user string
	password string
	database string
}

func (db mongoDB) Backup(t StrategieType) bool {
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

func (db mongoDB) Restore(t StrategieType) bool {
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

func (db mongoDB) Connection() bool {
	return true
}
