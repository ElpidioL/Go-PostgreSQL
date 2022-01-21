package FullCode

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

//This struct holds my DB table format
type MatchUser struct {
	ID        int
	MatchID   string
	PlayerID  string
	DiscordID string
}

//This struc holds the Postgree acess format
type DbAcess struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

var DbInfo DbAcess

func init() {
	//getting env values to hide my sensitive info
	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("Error loading .env file")
		return
	}
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		//error if the provided PORT in .env its not an INT  "10"
		fmt.Println("Port error, verify if its a number", err)
		os.Exit(0)
	}
	DbInfo.Host = os.Getenv("HOST")
	DbInfo.Port = port
	DbInfo.User = os.Getenv("USER")
	DbInfo.Password = os.Getenv("PASSWORD")
	DbInfo.Dbname = os.Getenv("DB_NAME")
}

func ExampleFull() {
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

	////////////////////////////////////////Using the DB/////////////////////////////////////
	//THIS IS A EXAMPLE CODE, ITS TO BE USED AS REFERENCE AND IT WILL DROP/CREATE THE SAME TABLE OVER AND OVER.
	//there is difference between db.Exec and db.Query, but I'm not a pro at this.
	//usually db.Exec for CUD, Create, Update, Delete and Query for Ready multiple values i guess.
	//you don't need to close db.Exec, but you need to read all values and Close db.Query
	//db.QueryRow() for single values and db.Query for multiple values.
	//I recommend you looking further on this.

	////////////Delete Table
	sqlStatement := `DROP TABLE IF EXISTS Matches99 RESTRICT`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		fmt.Println("erro")
	}
	///End

	////////////Creat Table
	//table format that you want to use, you can use UNIQUE NOT NULL if you need ex:  player_id TEXT UNIQUE NOT NULL
	sqlStatement = `CREATE TABLE Matches99 (
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

	////////////Add info to Table
	//I saw that using "database/sql" is already the way to prevent SQL Injection, since I'm novice with go and postgre I'm not sure
	//but since a random dude on internet told me that this prevent it, I'll trust without trying.

	for i := 99; i <= 101; i++ { // ignore the loop, its just so i can have examples later.
		sqlStatement = `
					INSERT INTO Matches99 (match_id, player_id, discord_id)
					VALUES ($1, $2, $3)`
		_, err = db.Exec(sqlStatement, strconv.Itoa(i), "Cake"+strconv.Itoa(i), "black#9999")

		//simplified version db.Exec(sqlStatement, "10", "Cake", "black#9999")
		if err != nil {
			panic(err)
		}
	}
	///End

	////////////Updating info to Table
	sqlStatement = `UPDATE Matches99
					SET match_id = $1, player_id = $2
					WHERE player_id = $3;`
	_, err = db.Exec(sqlStatement, "WeWillSeeItLater", "CakeBlack", "Cake99")
	if err != nil {
		panic(err)
	}

	////////////Deleting info to Table
	sqlStatement = `
					DELETE FROM Matches99
					WHERE match_id = $1;`
	_, err = db.Exec(sqlStatement, "100")
	if err != nil {
		panic(err)
	}
	///End

	////////////Reading single info from Table
	//remember that we created this value on the update
	sqlStatement = `SELECT match_id FROM matches99 WHERE player_id='CakeBlack';`
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
