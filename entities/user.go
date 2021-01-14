package entities

// User is this guest-books admin user
type User struct {
	ID       int
	Username string
	Password string
	State    int
}
