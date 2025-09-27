package step3httpintro

func BuildHTTPResponse(statusLine, headers, body string) string{
	response := statusLine + "\r\n"
    
    if headers != "" {
        response += headers
    }
    
    response += "\r\n"

    if body != "" {
        response += body
    }
    
    return response
}