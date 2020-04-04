package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// define the database connection parameters
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "brainiac"
	dbname   = "go_test_db"
)

// User type defined
type User struct {
	ID        int    `json:"ID"`
	Age       int    `json:"Age"`
	FirstName string `json:"First Name"`
	LastName  string `json:"Last Name"`
	Email     string `json:"Email"`
}

func main() {
	// postgresInstanceInfo looks like "user:password@tcp(localhost:port)/dbname" when using "mysql"
	// postgresInstanceInfo looks like "host=%s port=%d user=%s password=%s dbname=% sslmode=disable" when using "postgres"
	postgresConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// connect to the database
	db, err := sql.Open("postgres", postgresConnStr) // this opens a connection and adds to the pool
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// connect to the database
	err = db.Ping() // this validates that the opened connection "db" is actually working
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully connected to the Postgres database")

	// insert data into the database
	sqlStatement := `INSERT INTO users (age, email, first_name, last_name) VALUES ($1, $2, $3, $4) RETURNING id` // inserts data and returns the primary key
	id := 0                                                                                                      // initialize to be an int
	err = db.QueryRow(sqlStatement, 85, "Nick.Cannon@marvel.com", "Nick", "Cannon").Scan(&id)                    // execute the sql, and placed the returned value inside `&id`
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("New record ID is: ", id) // print the row id that was automatically retrieved from the database

	// update data from database
	sqlStatement = `update users set first_name = $2, last_name = $3 where id = $1` // statement to update data with
	result, err := db.Exec(sqlStatement, 2, "Awokogbon", "Awokose")                 // execute the sql,by passing the required variable
	if err != nil {
		panic(err.Error())
	}
	affectedRowsCount, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(affectedRowsCount, "rows were updated") // print the number of rows that were updated

	// delete data from database
	sqlStatement = `delete from users where id = $1` // statement to delete data with
	result, err = db.Exec(sqlStatement, 1)           // execute the sql,by passing the required variable
	if err != nil {
		panic(err.Error())
	}
	affectedRowsCount, err = result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(affectedRowsCount, "rows were deleted") // print the number of rows that were deleted

	// query for specific data from database for a specific id
	sqlStatement = `select id, email from users where id = $1` // statement to delete data with
	var email string
	rowResult := db.QueryRow(sqlStatement, 4)       // execute the sql,by passing the required variable
	switch err = rowResult.Scan(&id, &email); err { // review differenct cases
	case sql.ErrNoRows: // sql.ErrNoRows is what Scan() returns when there is no row data to pass
		fmt.Println("No rows were returned")
	case nil: // this handles where there was data returned
		fmt.Println(id, email)
	// default case
	default:
		panic(err.Error()) // just throws a general error
	}

	// query for all data from database for a specific id
	sqlStatement = `select * from users where id = $1`
	var user User
	rowResult = db.QueryRow(sqlStatement, 3) // execute the sql,by passing the required variable
	err = rowResult.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email)
	switch err { // review differenct cases
	case sql.ErrNoRows: // sql.ErrNoRows is what Scan() returns when there is no row data to pass
		fmt.Println("No rows were returned")
	case nil: // this handles where there was data returned
		fmt.Println(user)
	// default case
	default:
		panic(err.Error())
	}

	// query for multiple records from database for multiple ids
	sqlStatement = `select * from users limit $1`  // get all colums from all records but limit the records to just $1
	multipleRows, err := db.Query(sqlStatement, 3) // execute the sql,by passing the required variable
	if err != nil {
		panic(err.Error())
	}
	defer multipleRows.Close() // needed in case this did not go well

	for multipleRows.Next() { // start iterating over the returned rows i.e.
		err = multipleRows.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user) // print the record extracted in each case
	}

	err = multipleRows.Err() // handle the error thrown when `multipleRos.Next() returns a false`
	if err != nil {
		panic(err.Error())
	}

}
