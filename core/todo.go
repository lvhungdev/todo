package core

type Todo struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
