package db

import "fmt"

type postgresqlDB struct {
	hostname string
	port string
	user string
	password string
	database string
}

func (db postgresqlDB) Backup(t StrategieType) bool {
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

func (db postgresqlDB) Restore(t StrategieType) bool {
	switch t{
	case Dump:
		postgreSQLDump()
	case Increment:
		fmt.Println("increment")
	case File:
		fmt.Println("file")
	}
	return true
}

func (db postgresqlDB) Connection() bool {
	return true
}

// Implementation

func postgreSQLDump() {
	fmt.Println("dump")
}
