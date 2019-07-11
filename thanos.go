package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func index_handler(w http.ResponseWriter, r *http.Request) {

	// Random array finder

	rand.Seed(time.Now().UnixNano())

	//Learnt that GoLang uses a determinsitic path to choose random so I had to import a time package to randomise it based on that, seems to work fairly well

	randomThanos = thanoslines[rand.Intn(len(thanoslines)-1)]

	fmt.Fprintf(w, randomThanos)
}

// All quotes

func all_handler(w http.ResponseWriter, r *http.Request) {

	allThanos := strings.Join(thanoslines, "\n")

	fmt.Fprintf(w, allThanos)
}

var thanoslines []string
var randomThanos string
var allThanos string

func main() {

	// Opens text file

	file, err := os.Open("thanosQuotes.txt")

	// If error in opening

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	// Splits the lines into an array after first reading

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// Apends the lines to the array

	for scanner.Scan() {
		thanoslines = append(thanoslines, scanner.Text())
	}

	// Closes text file

	file.Close()

	// Prints each line to console (for debug purposes)

	for _, eachline := range thanoslines {
		fmt.Println(eachline)
	}

	// Creates pages for the index of a random quote and all for all the quotes

	http.HandleFunc("/api/thanosapi/v1", index_handler)

	http.HandleFunc("/api/thanosapi/v1/all", all_handler)

	http.ListenAndServe(":8000", nil)
}
