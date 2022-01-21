package CodePieces

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func DeleteTable() {
	//storing the info to acess the DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DbInfo.Host, DbInfo.Port, DbInfo.User, DbInfo.Password, DbInfo.Dbname)

	//starting db
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//don't forget to close it
	defer db.Close()
	//ping to check if its working
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	////////////Delete Table
	sqlStatement := `DROP TABLE IF EXISTS Matches99 RESTRICT`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		fmt.Println("erro")
	}
	///End

}
