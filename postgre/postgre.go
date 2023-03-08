package postgre

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	// Initialize connection constants.
	HOST     = "localhost"
	PORT     = 5432
	DATABASE = "demopostgre"
	USER     = "postgres"
	PASSWORD = "kiper00"
)

var (
	db            *sql.DB
	err           error
	res           sql.Result
	sql_statement string
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func DemoPostgre() {
	// Initialize connection string.
	var connectionString string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DATABASE)

	// Initialize connection object.
	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database")
	defer db.Close()

	// Drop previous table of same name if one exists.
	_, err = db.Exec("DROP TABLE IF EXISTS inventory")
	checkError(err)
	fmt.Println("Finished dropping table (if existed)")

	// Create table.
	_, err = db.Exec("CREATE TABLE inventory (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);")
	checkError(err)
	fmt.Println("Finished creating table")

	// Insert some data into table.
	sql_statement = "INSERT INTO inventory (name, quantity) VALUES ($1, $2);"
	_, err = db.Exec(sql_statement, "banana", 150)
	checkError(err)
	_, err = db.Exec(sql_statement, "orange", 154)
	checkError(err)

	res, err = db.Exec(sql_statement, "apple", 100)
	fmt.Println(res)
	checkError(err)
	fmt.Println("Inserted 3 rows of data")

	// update
	// Modify some data in table.
	sql_statement = "UPDATE inventory SET quantity = $2 WHERE name = $1;"
	_, err = db.Exec(sql_statement, "banana", 200)
	checkError(err)
	fmt.Println("Updated 1 row of data")

	// delete
	// Delete some data from table.
	sql_statement := "DELETE FROM inventory WHERE name = $1;"
	_, err = db.Exec(sql_statement, "orange")
	checkError(err)
	fmt.Println("Deleted 1 row of data")
}
