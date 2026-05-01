package handlers

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

// root structure
type SiteMapIndex struct{
	XMLName xml.Name `xml:"sitemapindex"`
	Sitemaps []Sitemap `xml:"sitemap"`
}

// each sitemap entry
type Sitemap struct{
	XMLName xml.Name `xml:"sitemap"`
	Loc string `xml:"loc"`
	LastMod string `xml:"lastmod"`
}

func WorkingWithXMLData(){
	data :=[]byte (`<sitemapindex>
	<sitemap>
		<loc>https://...</loc>
		<lastmod>2026-05-01</lastmod>
	</sitemap>
	</sitemapindex>`)

	// Parse the XML
	var siteMapIndex SiteMapIndex
	err := xml.Unmarshal(data, &siteMapIndex)
	if err != nil {
		panic(err)
	}
	
	// Print the parsed data
	println("Parsed sitemaps:")
	for _, sitemap := range siteMapIndex.Sitemaps {
		fmt.Println("URL: ", sitemap.Loc)
		fmt.Println("Last Modified: ", sitemap.LastMod)
		fmt.Println("-----------------------")
	}
}

func PerformRequest() {
	response, err := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	
	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	
	// Parse the XML
	var siteMapIndex SiteMapIndex
	err = xml.Unmarshal(body, &siteMapIndex)
	if err != nil {
		panic(err)
	}
	
	// Print the parsed data
	println("Parsed sitemaps:")
	for _, sitemap := range siteMapIndex.Sitemaps {
		fmt.Println("URL: ", sitemap.Loc)
		fmt.Println("Last Modified: ", sitemap.LastMod)
		fmt.Println("-----------------------")
	}
}


func FetchAllSitemaps(sitemaps []Sitemap) {
	ch := make(chan string)

	for _, s := range sitemaps {
		go func(url string) {
			resp, _ := http.Get(url)
			defer resp.Body.Close()
			ch <- url
		}(s.Loc)
	}

	for range sitemaps {
		fmt.Println("Fetched:", <-ch)
	}
}