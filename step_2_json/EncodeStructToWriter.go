package step2json

import (
	"encoding/json"
	"io"
)



func EncodeStudentsToWriter(w io.Writer, students []Student) error{
	err := json.NewEncoder(w).Encode(students)

	return err
}