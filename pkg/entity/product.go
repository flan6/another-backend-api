package entity

type Product struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Value int    `db:"value"`
}
