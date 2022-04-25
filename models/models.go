package models

// User schema of the user table
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Age      int64  `json:"age"`
}

// CREATE TABLE users (
//     userid SERIAL PRIMARY KEY,
//     name TEXT,
//     age INT,
//     location TEXT
// );
