package step2json

import "encoding/json"

func SerializeIntSlice(nums []int) ([]byte, error){
	res, err := json.Marshal(nums)
	return res, err

}