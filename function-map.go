package main

func MapFunctions() {
	// book := make(map[string]string)
	book := map[string]string{
		"title":     "Belajar Golang",
		"author":    "Mustofa Adny",
		"publisher": "Penerbit Golang",
	}

	delete(book, "publisher")
}