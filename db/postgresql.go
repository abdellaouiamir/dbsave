package db

import (
	"bytes"
	"fmt"
	"os/exec"
)

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
		postgreSQLDump(db)
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
		fmt.Println("restore dump")
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

func postgreSQLDump(d postgresqlDB) {
	var options []string 
	if d.hostname != ""{
		options = append(options, "-h", d.hostname)
	}
	if d.port != ""{
		options = append(options, "-p", d.port)
	}
	if d.user != ""{
		options = append(options, "-U", d.user)
	}
	if d.password != ""{
		options = append(options, "-P", d.password)
	}
	if d.database != ""{
		options = append(options, "-db", d.database)
	}
	var out, errout bytes.Buffer
	cmd := exec.Command("echo", options...)
	cmd.Stdout = &out
	cmd.Stderr = &errout
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s:= %s", err, errout.String())
	}
	fmt.Println(out.String())
}
