package main

/*

func routeHandler() {
	r := mux.NewRouter()

	// r.HandleFunc("/api/books", getbooks).Methods("GET")
	// r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	// r.HandleFunc("/api/addBook/", createBooks).Methods("POST")
	// r.HandleFunc("/api/book/{id}", upateBook).Methods("PUT")
	// r.HandleFunc("/api/book/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/api/servers", getServers).Methods("GET")

	log.Fatal(http.ListenAndServe(portName, r))
}

func main() {

	// MOCK Data
	books = append(books, Book{ID: "1", Title: "Demo "})
	books = append(books, Book{ID: "2", Title: "Demo 2"})

	servers = append(servers, Servers{Name: "USA", Ip: "192.168.1.1"})
	servers = append(servers, Servers{Name: "France", Ip: "192.168.8.8"})
	// Route Hanlders/ Endpoints
	routeHandler()
}

func getbooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parms := mux.Vars(r) // get parms
	for _, item := range books {
		if item.ID == parms["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000)) // Mock ID
	books = append(books, book)
	json.NewEncoder(w).Encode(books)

}

func upateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}*/
