// LML - LiveMediaList
// Service for choice LiveMedia
// Roadmap
// 1 - Pattern
// 2 - Read start data from file and write to DB
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func init() {
	log.Println("Start Init func: ")
	readStartDb()
}

func main() {
	log.Println("Start Main func: ")
}

func readStartDb() {
	file, err := os.Open("./DB/StartDB.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			fmt.Println("----------------------------")
		}
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
