package CodePieces

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

func AddInfoTable() {
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

	////////////Add info to Table
	//I saw that using "database/sql" is already the way to prevent SQL Injection, since I'm novice with go and postgre I'm not sure
	//but since a random dude on internet told me that this prevent it, I'll trust without trying.

	for i := 99; i <= 101; i++ { // ignore the loop, its just so i can have examples later.
		sqlStatement := `
					INSERT INTO Matches99 (match_id, player_id, discord_id)
					VALUES ($1, $2, $3)`
		_, err = db.Exec(sqlStatement, strconv.Itoa(i), "Cake"+strconv.Itoa(i), "black#9999")

		//simplified version db.Exec(sqlStatement, "10", "Cake", "black#9999")
		if err != nil {
			panic(err)
		}
	}
	///End
}
