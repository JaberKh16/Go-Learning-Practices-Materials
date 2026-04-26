package main

import (
	"fmt"
	"net/url"
)

func main ()  {
	siteUrl := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"

	parsedUrl, err := url.Parse(siteUrl)
	if err != nil {
		fmt.Println("can't parse the url", err)
		return
	}
	fmt.Printf("Type is: %T\n", parsedUrl)
	fmt.Println("Scheme is: ", parsedUrl.Scheme)
	fmt.Println("Scheme is: ", parsedUrl.Host)
	fmt.Println("Scheme is: ", parsedUrl.Path)
	fmt.Println("Scheme is: ", parsedUrl.RawQuery)
	fmt.Println("Scheme is: ", parsedUrl.RequestURI())

	// construct a URL string from a URL Object
	urlString := parsedUrl.String()
	fmt.Println("URL String: ", urlString)

}