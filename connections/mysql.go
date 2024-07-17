package connections

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var conn *sql.DB
var connStr string
var connProfile connectionProfile

type DbConnection interface {
	getDbConnection() *sql.DB
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type connectionProfile struct {
	user        string
	database    string
	host        string
	sslmode     string
	port        int
	maxPoolConn int
	password    string
}

type mysql struct {
	Connection DbConnection
}

func init() {
	readConfig()
	connProfile = connectionProfile{
		user:     viper.GetString("db.mysql.user"),
		database: viper.GetString("db.mysql.database"),
		host:     viper.GetString("db.mysql.host"),
		port:     viper.GetInt("db.mysql.port"),
		sslmode:  viper.GetString("db.mysql.sslmode"),
		password: viper.GetString("db.mysql.password"),
	}
	fmt.Printf("\nLoading database connection config %s at %s...\n",
		connProfile.database, connProfile.host)
	connStr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		connProfile.user,
		connProfile.password,
		connProfile.host,
		connProfile.port,
		connProfile.database)
}

func readConfig() {
	if configFromEnv := "dev"; configFromEnv != "" {
		if configPath := os.Getenv("CONFIG_PATH"); configPath != "" {
			fmt.Printf("\nLoading config from %s", configPath)
			viper.SetConfigFile(fmt.Sprintf("%s", configPath))
		} else {
			viper.SetConfigFile(fmt.Sprintf("./config/%s.toml", configFromEnv))
		}
	} else {
		panic("USR_ENV not defined.")
	}

	// Load config.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Print(err)
		panic(err)
	}
}

func (s *mysql) getDbConnection() *sql.DB {
	//func GetDbConnection() *sql.DB{
	if conn == nil {
		createConnection()
	}
	return conn
}

func createConnection() {
	con, err := sql.Open("mysql", connStr)
	fmt.Println(connStr)
	if err != nil {
		fmt.Println("Failed to connect to database")
		fmt.Println(err)
	} else {
		fmt.Println("Connected to database successfully")
	}
	conn = con
}

func New() DbConnection {
	return &mysql{}
}

func (p *mysql) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return p.getDbConnection().Query(query, args...)
}
