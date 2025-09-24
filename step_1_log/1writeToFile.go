package step1log

import (
	"os"
)

func WriteToLogFile(message string, fileName string) error{
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil{
		return err
	}
	file.WriteString(message)
	return nil
}