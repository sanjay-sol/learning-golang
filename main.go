package main


import (
  "os"
  "log"
  "readLines"
  "compareFiles"
  "github.com/sanjay-sol/learning-go/file"
)

func main() {
    if len(os.Args) < 3 {
        log.Fatalf("Usage: %s <file1> <file2>\n", os.Args[0])
    }

    file1Path := os.Args[1]
    file2Path := os.Args[2]

    file1Lines, err := readLines(file1Path)
    if err != nil {
        log.Fatalf("Failed to read file %s: %s\n", file1Path, err)
    }

    file2Lines, err := readLines(file2Path)
    if err != nil {
        log.Fatalf("Failed to read file %s: %s\n", file2Path, err)
    }


