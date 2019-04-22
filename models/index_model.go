package models

type IndexModel struct {
	DB *sql.DB
}

// Index
func (indexModel *IndexModel) Index() string { return "Hello, world." }
