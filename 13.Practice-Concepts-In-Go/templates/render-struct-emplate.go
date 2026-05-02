package templates

type Data struct {
	Name string
	Desc string
	Fields []Field
}


type Field struct {
	Name string
	Type string
}


func setupData() Data {
	Data := []Data{
		{Name: "User", Desc: "User description", Fields: []Field{{Name: "Name", Type: "string"}, {Name: "Age", Type: "int"}}},
		{Name: "Product", Desc: "Product description", Fields: []Field{{Name: "Name", Type: "string"}, {Name: "Price", Type: "float64"}}},
		{Name: "Order", Desc: "Order description", Fields: []Field{{Name: "Name", Type: "string"}, {Name: "Price", Type: "float64"}}},
		{Name: "Category", Desc: "Category description", Fields: []Field{{Name: "Name", Type: "string"}}},
		{Name: "Tag", Desc: "Tag description", Fields: []Field{{Name: "Name", Type: "string"}}},
		{Name: "Review", Desc: "Review description", Fields: []Field{{Name: "Name", Type: "string"}, {Name: "Rating", Type: "int"}}},
		{Name: "Address", Desc: "Address description", Fields: []Field{{Name: "Name", Type: "string"}}},
		{Name: "Payment", Desc: "Payment description", Fields: []Field{{Name: "Name", Type: "string"}}},
		{Name: "Shipping", Desc: "Shipping description", Fields: []Field{{Name: "Name", Type: "string"}}},
		{Name: "Discount", Desc: "Discount description", Fields: []Field{{Name: "Name", Type: "string"}}},
		{Name: "Coupon", Desc: "Coupon description", Fields: []Field{{Name: "Name", Type: "string"}}},
		{Name: "Wishlist", Desc: "Wishlist description", Fields: []Field{{Name: "Name", Type: "string"}}},
	}
	return Data
} 

func workWithTheSetupData() {
	data := setupData()
	fmt.Println(data)

	tmpl, err := template.ParseFiles("./files/struct.tpl")
	if err != nil {
		log.Fatalf("coutn not parse struct template: %v\n", err)
	}	

	// process the data
	var processed bytes.Buffer
	err = tmpl.Execute(&processed, data)
	if err != nil {
		log.Fatalf("coutn not execute struct template: %v\n", err)
	}

	// formate the output
	// fmt.Println(processed.String())
	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		log.Fatalf("coutn not format struct template: %v\n", err)
	}
	
	fmt.Println(string(formatted))
	
}
