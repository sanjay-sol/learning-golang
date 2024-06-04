//
// package file 
//
// import (
//     "bufio"
//     "fmt"
//     "log"
//     "os"
//     . "github.com/sanjay-sol/learning-go"
// )
//
// func ReadLines(filePath string) ([]string, error) {
//     file, err := os.Open(filePath)
//     if err != nil {
//         return nil, err
//     }
//     defer file.Close()
//
//     var lines []string
//     scanner := bufio.NewScanner(file)
//     for scanner.Scan() {
//         lines = append(lines, scanner.Text())
//     }
//     if err := scanner.Err(); err != nil {
//         return nil, err
//     }
//
//     return lines, nil
// }
//
//
//
// func compareFiles(file1Lines, file2Lines []string) {
//     maxLen := len(file1Lines)
//     if len(file2Lines) > maxLen {
//         maxLen = len(file2Lines)
//     }
//
//     for i := 0; i < maxLen; i++ {
//         var line1, line2 string
//         if i < len(file1Lines) {
//             line1 = file1Lines[i]
//         }
//         if i < len(file2Lines) {
//             line2 = file2Lines[i]
//         }
//
//         if line1 != line2 {
//             fmt.Printf("Line %d:\n", i+1)
//             if line1 != "" {
//                 fmt.Printf("- %s\n", line1)
//             }
//             if line2 != "" {
//                 fmt.Printf("+ %s\n", line2)
//             }
//         }
//     }
// }

package file

import "fmt"

func Hello() {
  fmt.Println("Hello from the client/file..")
}
