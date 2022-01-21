package CodePieces

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ReadingSingle() {
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

	////////////Reading single info from Table
	//remember that we created this value on the update
	sqlStatement := `SELECT match_id FROM matches99 WHERE player_id='CakeBlack';`
	var playerID string

	row := db.QueryRow(sqlStatement)
	err = row.Scan(&playerID)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	case nil:
		fmt.Println("\n" + playerID + "\n")
	default:
		panic(err)
	}
	///End
}
