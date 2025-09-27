package step3httpintro

import (
	"strconv"
	"strings"
)

func ParseHTTPStatus(statusLine string) (code int, reason string){
	parts := strings.Split(statusLine, " ")
	code, err := strconv.Atoi(parts[1])
	if err != nil{
		return 0, ""
	}

	reason = ""
	for _,p := range parts[2:]{
		reason += p + " "
	} 

	reason = strings.TrimRight(reason, " ")

	return code, reason
}