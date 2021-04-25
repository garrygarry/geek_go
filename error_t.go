package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql",
		"root:zaq12wsx@tcp(127.0.0.1:3306)/mydevops")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var name string
	err = db.QueryRow("select username from t_user where id = ?", 4).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err)
			// there were no rows, but otherwise no error occurred
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println(name)
}