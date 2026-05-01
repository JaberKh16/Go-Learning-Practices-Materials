package data 

type SitemapIndex struct {
	Location []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string `xml:"keyword"`
	Locations string `xml:"loc"`
}


func WorkOnMap() map[string]NewsMap {
	// var s SitemapIndex
	var n News

	// Create a map to store news by keyword
	newsMap := make(map[string]NewsMap)
	
	// Populate the map with news data
	// for i, keyword := range n.Keywords {
	// 	if i < len(n.Locations) {
	// 		newsMap[keyword] = NewsMap{
	// 			Keyword:   keyword,
	// 			Locations: n.Locations[i],
	// 		}
	// 	}
	// }

	for idx, _ := range n.Locations {
		newsMap[n.Locations[idx]] = NewsMap{
			Keyword:   n.Keywords[idx],
			Locations: n.Locations[idx],
		}
	}
	
	return newsMap
}
