package step2json

import "encoding/json"

func DeserializeStringMap(data string) (map[string]string, error){
	res := make(map[string]string)

	err := json.Unmarshal([]byte(data), &res)

	return res, err
}