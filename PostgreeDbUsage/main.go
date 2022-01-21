package main

import (
	SmallCode "db/training/CodePieces"
	//FullCode "db/training/FullCode"
)

//For this code work you need a .env File and define some sensitive variables there like
//HOST = "localhost"
//PORT = "5432" << THE DEFAULT PORT, IF YOU DIDN'T CHANGE, ITS THIS.
//USER = "postgres"  << YOUR USER, PROBABLY ITS POSTGRES
//PASSWORD = "YOURPASSWORD"
//DB_NAME = "YOURDBNAME"

//you can read your .env file importing ("github.com/joho/godotenv" and "os")
//the usage is / X = os.Getenv("HOST") /

func main() {
	//Here you can some functions to interact with postgree, you can read the script in CodePieces to understand what is
	//happening in small pieces.
	SmallCode.DeleteTable()
	SmallCode.CreateTable()
	SmallCode.AddInfoTable()
	SmallCode.UpdatingTable()
	SmallCode.DeletingInfo()
	SmallCode.ReadingSingle()
	SmallCode.ReadingMultiple()

	//Alternatively, you can go to Full Code and see everything in the same script
	//there you will get a better understanding since i implement my structs in the same Code
	//but since there is a lot of code, can be harder to understand.

	//FullCode.ExampleFull()

	//both the small pieces and the full code do the same, i did this split to make easier to understand.
}
