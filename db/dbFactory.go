package db

import "fmt"

// Strategie Supported
type StrategieType int
const (
	Dump StrategieType = iota
	File
	Increment
)

type DB interface {
	Backup(t StrategieType) error
	Restore(t StrategieType) error
	Connection() bool
}
type DbObject struct {
	Hostname string
	Port string
	User string
	Password string
	DataBase string
	AllDataBase bool
	OutputFile string
	Compress bool
	Type DBMS
}
func (d DbObject) Error() string {
	return fmt.Sprintf("not supported DBMS type %s", d.Type)
}
// DBMS Supported types
type DBMS string
const (
	Postgresql DBMS = "postgresql"
	Mysql DBMS = "mysql"
	Mongo DBMS = "mongo"
)

func NewDB(dbO DbObject) (DB, error) {
	var d DB = nil
	switch dbO.Type {
	case Postgresql:
		d = postgresqlDB{
			hostname: dbO.Hostname,
			port: dbO.Port,
			user: dbO.User,
			password: dbO.Password,
			database: dbO.DataBase,
			alldatabase: dbO.AllDataBase,
			compress: dbO.Compress,
			outputFile: dbO.OutputFile,
		}
	case Mysql:
		d = mysqlDB{
			hostname:    dbO.Hostname,
			port:        dbO.Port,
			user:        dbO.User,
			password:    dbO.Password,
			database:    dbO.DataBase,
			alldatabase: dbO.AllDataBase,
			compress:    dbO.Compress,
			outputFile:  dbO.OutputFile,
		}
	case Mongo:
		d = mongoDB{
			hostname: dbO.Hostname,
			port: dbO.Port,
			user: dbO.User,
			password: dbO.Password,
			database: dbO.DataBase,
			alldatabase: dbO.AllDataBase,
			compress:    dbO.Compress,
			outputFile:  dbO.OutputFile,
		}
		default:
		return d, dbO
	}
	return d, nil
}
