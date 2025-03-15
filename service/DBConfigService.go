package service

import (
	"bufio"
	"os"
	"strings"
)

// func scanFile(file *os.File) []string {
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		items := strings.Split(line, "=")
// 		return items
// 	}
// }

func GetDBConfig(file *os.File) []any {
	var info []any
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "=")
		info = append(info, items[1])
	}
	return info
}

// func GetDBName(file *os.File) string {
// 	scanFile(file)

// }

// func GetDBPort() {

// }

// func GetDBHostAddress() {

// }

// func GetDBUsername() {

// }

// func GetDBPassword() {

// }
