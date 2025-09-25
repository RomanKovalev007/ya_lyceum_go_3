package step2json

import (
	"encoding/json"
	"io"
)

type Student struct {
    Name  string `json:"name"`
    Grade int    `json:"grade"`
}

func DecodeStudentFromReader(r io.Reader) (Student, error){
	var student Student

	err := json.NewDecoder(r).Decode(&student)

	return student, err
}