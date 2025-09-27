package step3httpintro


func BuildHTTPRequest(method, url, host, headers, body string) string{
    request := method + " " + url + " HTTP/1.1\r\n"
    
    request += "Host: " + host + "\r\n"
    
    if headers != "" {
        request += headers
    }
    
    request += "\r\n"

    if body != "" {
        request += body
    }
    
    return request
}