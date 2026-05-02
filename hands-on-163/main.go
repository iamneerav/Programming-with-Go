package main

func main() {

}

// TO DO

// User represents a user with an id and first name.
type User struct {
	ID    int
	First string
}

// MockDatastore is a temporary service that stores retrievable data.
type MockDatastore struct {
	Users map[int]User
}
