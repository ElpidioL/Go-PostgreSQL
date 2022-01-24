# GoPostgre
 An example usage of PostgreSQL with GO, very simple since the objective is that others can read and learn.
You need to install PostgreSQL first, of course
First, i recommend you creating a .env file to keep your info

![image](https://user-images.githubusercontent.com/96087622/150814356-c7de63cc-f754-4966-b1a8-0092d946a12d.png)

You need to change your password and the DbName for it to work.
You can call yours .env variables with the import ("os" "github.com/joho/godotenv")

![image](https://user-images.githubusercontent.com/96087622/150814691-7887ceae-f5c6-4324-bb20-827b18c5f759.png)

I'm also checking if the port is a INT in this case.

///////////////////////////////////////////////////////////////////////////////////////////////////////////////

Then you need to initialize  the PostgreeDB

You need to import ("database/sql"
                    _ "github.com/lib/pq")

![image](https://user-images.githubusercontent.com/96087622/150815247-2c6ba0ef-e1b1-41b4-b551-fd9fc04ee4db.png)

The you can start working with sql statements

![image](https://user-images.githubusercontent.com/96087622/150815955-59c718d7-fa14-438c-a328-bfe2a5bc3abe.png)


OBS: when needed you need to close with defer dbInstance.Close()
it need to be done when you first initialize the DB and when you user Query

You will be using db.exec() for delete, create, update values
db.QueryRow for reading single values and d.Query for multiple values.

I recommend you go deeper into this topic if you are making a serious application.
