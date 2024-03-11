package main

import (

	"fmt"
	"os"
	"strings"

)

func readaFile(filename string) string{

	data, err := os.ReadFile(filename)
	if err != nil {
		os.Exit(1)
	}
	return string(data)
}

func countChars(s string) int{
	return len(s)
}

func countWords(s string) int{

	words := strings.Fields(s)
	return len(words)

}

func countLines(s string) int{

	lines := strings.Split(s, "\n")
	return len(lines)
}


func main(){

	
	fmt.Println("Character Length:",countChars(readaFile("/etc/theHarvester/wordlists/dns-big.txt")))
	fmt.Println("Words Count:",countWords(readaFile("/etc/theHarvester/wordlists/dns-big.txt")))
	fmt.Println("Line Count:",countLines(readaFile("/etc/theHarvester/wordlists/dns-big.txt")))
	
}