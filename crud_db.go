package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to the database
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Create the Coffee_menu table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Coffee_menu (no INT, Speciality VARCHAR(30), Prices INT)")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Table Coffee_menu created")

	// Insert a record into the table
	_, err = db.Exec("INSERT INTO Coffee_menu (no, Speciality, Prices) VALUES (1, 'Latte', 249), (2, 'espresso', 199), (3, 'black tea', 149)")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Record inserted into the table")

	// Display records from the table
	rows, err := db.Query("SELECT * FROM Coffee_menu")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var no int
		var speciality string
		var prices int
		err = rows.Scan(&no, &speciality, &prices)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("No: %d, Speciality: %s, Prices: %d\n", no, speciality, prices)
	}

	// Update a record in the table
	_, err = db.Exec("UPDATE Coffee_menu SET Speciality = 'Black Mamba' WHERE no = 1")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Record updated successfully")

	// Delete a record from the table
	_, err = db.Exec("DELETE FROM Coffee_menu WHERE no = 1")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Record deleted successfully")
}
