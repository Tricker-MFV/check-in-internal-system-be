package entity

// mapping for database

type Employee struct {
	ID       int    `db:"id" json:"id,omitempty"`
	Username string `db:"username" json:"username,omitempty"`
	Password string `db:"password" json:"password,omitempty"`
}
