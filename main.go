package main

import (

	"fmt"
	"os"
	"strings"

)

/* func createFile(filename string, s string) (error){

	data := []byte(s)
	err := os.WriteFile(filename, data, 0666)
	if err != nil {
		return err
	}
	log.Println(data)
	return nil
} */

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

func countWords(s string) []string{

	words := strings.Fields(s)
	return words
	// return len(words)

}

func countLines(s string) int{

	lines := strings.Split(s, "\n")
	return len(lines)
}

func countFrequency(s []string) map[string]int{

	m := make(map[string]int)
	for _, value := range s{
		count, ok := m[value]
		if !ok {
			count = 1
		} else {
			count++
		}
		m[value] = 1

	}
	return m
}

func main(){

	
	fmt.Println("Character Length:",countChars(readaFile("/etc/X11/rgb.txt")))
	fmt.Println("Words Count:",countWords(readaFile("/etc/X11/rgb.txt")))
	fmt.Println("Line Count:",countLines(readaFile("/etc/X11/rgb.txt")))
	fmt.Println("Word Frequency:",countFrequency(countWords(readaFile("/etc/X11/rgb.txt"))))

	
}