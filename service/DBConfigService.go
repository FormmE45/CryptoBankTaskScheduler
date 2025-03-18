package service

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type DBconfig struct {
	Host     string
	Port     uint64
	User     string
	Password string
	Dbname   string
}

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

func InitDBConfig(fileConfigPath string) *DBconfig {
	//Read file that hold DB configuratio
	file := OpenFile(fileConfigPath)
	info := GetDBConfig(file)
	u, err := strconv.ParseUint(info[1].(string), 10, 64)
	Check(err)
	//Instantiate DBConfig
	dbConfig := DBconfig{info[0].(string), u, info[2].(string), info[3].(string), info[4].(string)}
	return &dbConfig
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
