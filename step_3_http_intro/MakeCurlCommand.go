package step3httpintro

import (
	"fmt"
	"strings"
)

func MakeCurlCommand(method, url, headers, body string) string {
    var parts []string
    
    parts = append(parts, "curl")
    
    if method != "GET" && method != "" {
        parts = append(parts, fmt.Sprintf("-X %s", method))
    }
    
    headerLines := strings.Split(headers, "\n")
    for _, header := range headerLines {
        header = strings.TrimSpace(header)
        if header != "" {
            parts = append(parts, fmt.Sprintf("-H '%s'", header))
        }
    }
    
    if body != "" {
        parts = append(parts, fmt.Sprintf("--data '%s'", body))
    }
    
    parts = append(parts, url)
    
    return strings.Join(parts, " ")
}
