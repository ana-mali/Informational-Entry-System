package models

import "time"

type Note struct {
	ID        int
	Text      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
type Task struct {
	ID int
	Name string
	CreatedAt time.Time 
	Priority *string 
	DueDate *time.Time 
	UpdatedAt *time.Time
}

type List struct{
	ID int 
	Name string
	Items []Item 
	CreatedAt time.Time 
	UpdatedAt *time.Time
}

type Item struct{
	ID int
	Text string
	Check bool 
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func (n Note) GetID() int {
	return n.ID
}
func (t Task) GetID() int {
	return t.ID
}
func (l List) GetID() int {
	return l.ID
}
type Identifiable interface {
	GetID() int
}