package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) CreateCategory(name, description string) (Category, error) {
	ID := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", ID, name, description)
	if err != nil {
		return Category{}, err
	}
	return Category{db: c.db, ID: ID, Name: name, Description: description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := []Category{}
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

// fazer um medotodo que retorne a categoria com base no id do curso passado
func (c *Category) FindByCourseID(courseID string) (Category, error) {
	var category Category
	if err := c.db.QueryRow("SELECT c.* FROM categories c INNER JOIN courses co ON c.id = co.category_id WHERE co.id = $1", courseID).Scan(&category.ID, &category.Name, &category.Description); err != nil {
		return Category{}, err
	}
	return category, nil
}

func (c *Category) FindByID(ID string) (Category, error) {
	var category Category
	if err := c.db.QueryRow("SELECT * FROM categories WHERE id = $1", ID).Scan(&category.ID, &category.Name, &category.Description); err != nil {
		return Category{}, err
	}
	return category, nil
}
