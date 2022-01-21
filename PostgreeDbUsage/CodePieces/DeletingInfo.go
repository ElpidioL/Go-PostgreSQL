package CodePieces

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func DeletingInfo() {
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

	////////////Deleting info to Table
	sqlStatement := `
					DELETE FROM Matches99
					WHERE match_id = $1;`
	_, err = db.Exec(sqlStatement, "100")
	if err != nil {
		panic(err)
	}
	///End
}
