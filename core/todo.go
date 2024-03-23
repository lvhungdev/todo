package core

type Todo struct {
	id   int    `db:"id"`
	name string `db:"name"`
}
