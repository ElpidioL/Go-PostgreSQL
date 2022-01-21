package CodePieces

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func UpdatingTable() {
	//storing the info to access the DB
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

	////////////Updating info to Table
	sqlStatement := `UPDATE Matches99
					SET match_id = $1, player_id = $2
					WHERE player_id = $3;`
	_, err = db.Exec(sqlStatement, "WeWillSeeItLater", "CakeBlack", "Cake99")
	if err != nil {
		panic(err)
	}
}
