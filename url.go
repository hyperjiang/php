package php

import (
	"encoding/base64"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Base64Encode encodes data with MIME base64
func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// Base64Decode decodes data encoded with MIME base64.
// Returns error if data is in invalid format.
func Base64Decode(data string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(data)

	return string(b), err
}

// GetHeaders fetches all the headers sent by the server in response to an HTTP request
func GetHeaders(url string) (http.Header, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	resp.Header.Set("status", resp.Status)

	return resp.Header, nil
}

// GetMetaTags extracts all meta tag content attributes from http content and returns an array
func GetMetaTags(url string) (map[string]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ParseDocument(resp.Body), nil
}

// ParseDocument parses a document for meta data
func ParseDocument(doc io.Reader) map[string]string {
	data := make(map[string]string)
	z := html.NewTokenizer(doc)

	var titleFound bool
	for {
		tt := z.Next()
		t := z.Token()
		switch tt {
		case html.ErrorToken:
			return data
		case html.EndTagToken:
			if t.Data == "head" {
				return data
			}
		case html.StartTagToken, html.SelfClosingTagToken:
			if t.Data == "title" {
				titleFound = true
			}
			if t.Data == "meta" {
				var property, content string
				for _, attr := range t.Attr {
					switch attr.Key {
					case "property", "name":
						property = strings.ToLower(attr.Val)
					case "content":
						content = attr.Val
					}
				}

				if property != "" {
					data[strings.TrimSpace(property)] = content
				}
			}
		case html.TextToken:
			if titleFound {
				data["title"] = t.Data
				titleFound = false
			}
		}
	}

	return data
}
