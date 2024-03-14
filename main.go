// Package main contains the main function for executing the code.
package main

import (
    "fmt"
    "os"
    "strings"
)

// readFile reads the content of a file specified by the input string.
// If an error occurs, it prints the error.
// Returns the content of the file as a string.
func readFile(s string) string{
    res, err := os.ReadFile(s)
    if err != nil {
        fmt.Println(err)
    }
    return string(res)
}

// countChars counts the number of characters in a string and returns the count.
func countChars(s string) int{
    return len(s)
}

// countWords splits the string into words and returns the count of words.
func countWords(s string) int{
    splitchar:=strings.Fields(s)
    return len(splitchar)
}

// lineCount counts the number of lines in a string separated by newline characters.
func lineCount(s string) int{
    res := strings.Split(s, "\n")
    return len(res)
}

// wordFrequency calculates the frequency of each word in a string and returns a map.
func wordFrequency(s string) map[string]int{
    res:=strings.Fields(s)
    freq:= make(map[string]int)

    for _, word:=range res{
        freq[word]++
    }
    return freq
}

// The main function prompts the user to enter a file path, reads the file, and displays character count,
// word count, line count, and word frequency in the file.
func main() {
    var file string
    fmt.Println("Enter file full path:");fmt.Scan(&file)

    fileContent:=readFile(file)
    fmt.Println("Character Count:", countChars(fileContent))
    fmt.Println("Word Count:", countWords(fileContent))
    fmt.Println("Line Count:", lineCount(fileContent))
    fmt.Println("Word Frequency:")

    for key, value := range wordFrequency(fileContent){
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }
}