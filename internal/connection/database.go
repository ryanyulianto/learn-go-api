package connection

import (
	"belajar-go-api/internal/config"
	"database/sql"
	"log"
	"strings"

	"fmt"

	_ "github.com/lib/pq"
)

func GetDatabase(conf config.Database) *sql.DB {
	var dsnParts []string
	dsnParts = append(dsnParts, fmt.Sprintf("host=%s", conf.Host))
	dsnParts = append(dsnParts, fmt.Sprintf("port=%s", conf.Port))
	dsnParts = append(dsnParts, fmt.Sprintf("dbname=%s", conf.Name))
	dsnParts = append(dsnParts, "sslmode=disable")
	dsnParts = append(dsnParts, fmt.Sprintf("TimeZone=%s", conf.Tz))

	if conf.Username != "" {
		dsnParts = append(dsnParts, fmt.Sprintf("user=%s", conf.Username))
		if conf.Password != "" {
			dsnParts = append(dsnParts, fmt.Sprintf("password=%s", conf.Password))
		}
	}
	dsn := strings.Join(dsnParts, " ")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
	}
	ping := db.Ping()
	if ping != nil {
		log.Fatal("Error pinging database: ", ping.Error())
	}
	return db
}
