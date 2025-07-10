package db

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

type postgresqlDB struct {
	hostname string
	port string
	user string
	password string
	database string
	alldatabase bool
	compress bool
	outputFile string
}
func (pg postgresqlDB) Error() string {
	return fmt.Sprintf("See the error below")
}

func (db postgresqlDB) Backup(t StrategieType) error {
	switch t{
	case Dump:
		return postgreSQLDumpBackup(db)
	case Increment:
		fmt.Println("increment")
	case File:
		fmt.Println("file")
	}
	return errors.New("bad choise")
}

func (db postgresqlDB) Restore(t StrategieType) error {
	switch t{
	case Dump:
		return postgreSQLDumpRestore(db)
	case Increment:
		fmt.Println("increment")
	case File:
		fmt.Println("file")
	}
	return errors.New("bad resoter")
}

func (db postgresqlDB) Connection() bool {
	return true
}

// Implementation

func postgreSQLDumpBackup(d postgresqlDB) error {
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
	if d.database != "" && !d.alldatabase {
		options = append(options, "-d", d.database)
	}
	if d.compress == true {
		options = append(options, "-Fc")
	}
	options = append(options, "-f", d.outputFile)
	var out, errout bytes.Buffer
	var cmd *exec.Cmd
	if d.alldatabase == true {
		cmd = exec.Command("pg_dumpall", options...)
	} else {
		cmd = exec.Command("pg_dump", options...)
	}
	cmd.Stdout = &out
	cmd.Stderr = &errout
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s:=\n%s", err, errout.String())
		return d
	}
	fmt.Println(out.String())
	return nil
}

func postgreSQLDumpRestore (d postgresqlDB) error {
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
	if d.database != "" && !d.alldatabase {
		options = append(options, "-d", d.database)
	}
	options = append(options, "-f", "test.sql")
	var out, errout bytes.Buffer
	var cmd *exec.Cmd
	if d.compress == true {
		cmd = exec.Command("pg_restore", options...)
	} else {
		cmd = exec.Command("pg_restore", options...)
	}
	cmd.Stdout = &out
	cmd.Stderr = &errout
	err := cmd.Run()
	if err != nil {
		fmt.Printf("%s:=\n%s", err, errout.String())
		return d
	}
	fmt.Println(out.String())
	return nil
}
