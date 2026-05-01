package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func showCity(w http.ResponseWriter, r *http.Request) {
	city := City{}
	queryParams := "%" + r.URL.Path[1:] + "%"
	rows, err := database.Query("SELECT * FROM cities WHERE name LIKE ?", queryParams)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	for rows.Next() {
		err := rows.Scan(&city.ID, &city.Name, &city.CountryCode, &city.District, &city.Population)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%s (%s), Population: %d \n", city.Name, city.CountryCode, city.Population)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func showCityHTML(w http.ResponseWriter, r *http.Request) {
	city := City{}
	queryParam := "%" + r.URL.Path[1:] + "%"
	rows, err := database.Query("SELECT Name, CountryCode, Population
	FROM city WHERE Name LIKE ?", queryParam)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	html := "<html><head><title>City
	Search</title></head><body><h1>Search for" + queryParam + "</h1><table
	border='1'><tr><th>City</th><th>Country
	Code</th><th>Population</th></tr>"

	for rows.Next() {
		err := rows.Scan(&city.Name, &city.CountryCode,&city.Population)
		if err != nil {
			log.Fatal(err)
		}
		html +=
		fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%d</td></tr>", city.Name,
		city.CountryCode, city.Population)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	} else {
		html += "</table></body></html>"
		fmt.Fprintln(w, html)
	}
}
}


var database *sql.DB

func init() {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}

	database = db
	log.Println("Database connected")
	defer db.Close()
}
