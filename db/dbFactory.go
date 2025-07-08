package db

// Strategie Supported
type StrategieType int
const (
	Dump StrategieType = iota
	File
	Increment
)

type DB interface {
	Backup(t StrategieType) bool
	Restore(t StrategieType) bool
	Connection() bool
}
type DbObject struct {
	Hostname string
	Port string
	User string
	Password string
	DataBase string
}
// DBMS Supported types
type DBMS string
const (
	Postgresql DBMS = "postgresql"
	Mysql DBMS = "mysql"
	Mongo DBMS = "mongo"
)

func NewDB(dbType DBMS, dbO DbObject) DB {
	var d DB = nil
	switch dbType {
	case Postgresql:
		d = postgresqlDB{
			hostname: dbO.Hostname,
			port: dbO.Port,
			user: dbO.User,
			password: dbO.Password,
			database: dbO.DataBase,
		}
	case Mysql:
		d = mysqlDB{
			hostname: dbO.Hostname,
			port: dbO.Port,
			user: dbO.User,
			password: dbO.Password,
			database: dbO.DataBase,
		}
	case Mongo:
		d = mongoDB{
			hostname: dbO.Hostname,
			port: dbO.Port,
			user: dbO.User,
			password: dbO.Password,
			database: dbO.DataBase,
		}
	}
	return d
}
