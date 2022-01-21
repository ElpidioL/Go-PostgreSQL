package CodePieces

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ReadingMultiple() {
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

	////////////Reading mutiple values from Table
	// here we are using the struct defined on the start.
	var matches []MatchUser // going to use it to see all information

	rows, err := db.Query("SELECT * FROM matches99;")

	if err != nil {
		panic(err)
	}
	defer rows.Close() //Remember to close for Query

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var mtc MatchUser
		if err := rows.Scan(&mtc.ID, &mtc.MatchID, &mtc.PlayerID,
			&mtc.DiscordID); err != nil {
			fmt.Println("Something went wrong")
		}
		matches = append(matches, mtc)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("DEU ERRO")
	}
	fmt.Println(matches)
	///End
}
