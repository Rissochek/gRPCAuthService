package utils

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func OpenFile(filename *string) (*os.File){
	file, err := os.OpenFile(*filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil{
		log.Fatalf("Failed")
	}
	return file
}

func CreateUser(username *string, password *string, file *os.File){
	result_string := fmt.Sprintf("%s %s\n", *username, *password)
	_, err := file.WriteString(result_string)
	if err != nil{
		log.Fatalf("Failed to write string in file: %v", err)
	}
}

func CloseFile(file *os.File){
	err := file.Close()
	if err != nil{
		log.Fatalf("Failed to close file: %v", err)
	}
}

func SearchInFile(username *string, password *string, file *os.File) error{
	user_info := fmt.Sprintf("%s %s", *username, *password)
	log.Println(user_info)
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		log.Println("jasodhwifuop")
        if strings.Contains(line, user_info) {
            fmt.Println("Found:", line)
			return nil
        }
	}
	return errors.New("usesss not found")
}