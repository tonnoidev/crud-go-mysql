package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main() {
	fmt.Println("GO MySQL Tutorial")

	db, err := sql.Open("mysql", "root:p@ssw0rd2535@(127.0.0.1)/testdb")	
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("Successfully Connected MySQL database")

	// insert, err := db.Query("INSERT INTO Cars VALUES(9, 'Honda', 129932)")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()

	// fmt.Println("Inserted Succesfully")

	// update, err := db.Query("UPDATE Cars set price=12345 WHERE id=9")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer update.Close()

	// fmt.Println("Updated Successfully")

	delete, err := db.Query("DELETE FROM Cars WHERE id=9")
	if err != nil {
		panic(err.Error())
	}
	defer delete.Close()

	fmt.Println("Deleted Successfully")

}