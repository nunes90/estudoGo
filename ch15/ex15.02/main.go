// ex15.02 - holding prime numbers in a database
package main

import (
	"database/sql"
	"fmt"
	"math/big"

	_ "github.com/lib/pq"
)

func main() {
	var number int64
	var prop string
	var primeSum int64
	var newNumber int64

	db, err := sql.Open("postgres", "user=postgres password=Start!123 host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the DB was successfully initialized!")
	}

	allTheNumbers := "SELECT * FROM Number"
	numbers, err := db.Prepare(allTheNumbers)
	if err != nil {
		panic(err)
	}
	primeSum = 0
	result, err := numbers.Query()
	fmt.Println("The list of prime numbers:")
	for result.Next() {
		err = result.Scan(&number, &prop)
		if err != nil {
			panic(err)
		}
		if big.NewInt(number).ProbablyPrime(0) {
			primeSum += number
			fmt.Print(" ", number)
		}
	}
	err = numbers.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("\nThe total sum of primes numbers in this range is:", primeSum)

	remove := "DELETE FROM Number WHERE Property=$1"
	removeResult, err := db.Exec(remove, "Even")
	if err != nil {
		panic(err)
	}
	modifiedRecords, err := removeResult.RowsAffected()
	fmt.Println("The number of rows removed: ", modifiedRecords)
	fmt.Println("Updateing numbers...")

	update := "UPDATE Number SET Number=$1 WHERE Number=$2 AND Property=$3"
	allTheNumbers = "SELECT * FROM Number"
	numbers, err = db.Prepare(allTheNumbers)
	if err != nil {
		panic(err)
	}
	result, err = numbers.Query()
	for result.Next() {
		err = result.Scan(&number, &prop)
		if err != nil {
			panic(err)
		}
		newNumber = number + primeSum
		_, err = db.Exec(update, newNumber, number, prop)
		if err != nil {
			panic(err)
		}
	}
	numbers.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("The execution is now complete...")
	db.Close()
}
