package main

import (
	"database/sql"
	"fmt"
	"log"

	// blank import for MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

const driverName = "mysql"

type Hello struct {
	ID int32
	Text string
}

func main()  {
	user := "root"
	password := "dojo-training"
	host := "127.0.0.1"
	port := "3306"
	database := "training"

	c, err := sql.Open(driverName,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database))
	if err != nil {
		log.Fatal(err)
	}

	row := c.QueryRow("SELECT * FROM hello;")

	h := Hello{}
	if err := row.Scan(&h.ID, &h.Text); err != nil {
		log.Fatal(err)
	}

	fmt.Println(h.Text)
}
