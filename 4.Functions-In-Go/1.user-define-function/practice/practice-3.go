package main

var ErrorNotFound = "not found"

func findItem(items []string, target string) (string, error) {
	for _, item := range items {
		if item == target {
			return item, nil
		}
	}
	return "", ErrorNotFound
}

func main() {
	items := []string{"apple", "banana", "orange"}
	item, err := findItem(items, "banana")
	if err != nil {
		println(err)
		return
	}
	println(item)
}
