package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryID  string `json:"category_id"`
}

func NewCourse(db *sql.DB) *Course {
	return &Course{db: db}
}

func (c *Course) CreateCourse(name, description, categoryID string) (Course, error) {
	ID := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)", ID, name, description, categoryID)
	if err != nil {
		return Course{}, err
	}
	return Course{db: c.db, ID: ID, Name: name, Description: description, CategoryID: categoryID}, nil
}

func (c *Course) FindAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT * FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		var curse Course
		if err := rows.Scan(&curse.ID, &curse.Name, &curse.Description, &curse.CategoryID); err != nil {
			return nil, err
		}
		courses = append(courses, curse)
	}
	return courses, nil
}

func (c *Course) FindByCategoryID(categoryID string) ([]Course, error) {
	rows, err := c.db.Query("SELECT * FROM courses WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		var curse Course
		if err := rows.Scan(&curse.ID, &curse.Name, &curse.Description, &curse.CategoryID); err != nil {
			return nil, err
		}
		courses = append(courses, curse)
	}
	return courses, nil
}
