package CodePieces

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func CreateTable() {
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

	////////////Creat Table
	//table format that you want to use, you can use UNIQUE NOT NULL if you need ex:  player_id TEXT UNIQUE NOT NULL
	sqlStatement := `CREATE TABLE Matches99 (
					id SERIAL PRIMARY KEY,
					match_id TEXT,
					player_id TEXT,
			discord_id TEXT
	);`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		fmt.Println("erro")
	}
	///End

}
