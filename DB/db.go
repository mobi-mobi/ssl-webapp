package DB

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func OpenDB() error {
	var err error
	DB, err = sql.Open("mysql", "dbuser:dbuserheslo@tcp(127.0.0.1:3306)/userdb")
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	fmt.Println("connected!")
	return nil
}

func CloseDB() error {
	err := DB.Close()
	if err != nil {
		return err
	}
	return nil
}

func QueryForUser(username string) {

}
