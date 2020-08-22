package rpc

import (
	"fmt"
)

// College .
type College struct {
	database map[int]Student
}

// Student .
type Student struct {
	ID                  int
	FirstName, LastName string
}

// FullName .
func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

// Add .
func (c *College) Add(payload Student, reply *Student) error {
	if _, ok := c.database[payload.ID]; ok {
		return fmt.Errorf("Student with id %d already exists", payload.ID)
	}

	c.database[payload.ID] = payload
	*reply = payload
	return nil
}

// Get .
func (c *College) Get(payload int, reply *Student) error {
	result, ok := c.database[payload]
	if !ok {
		return fmt.Errorf("student with id %d does not exist", payload)
	}
	*reply = result
	return nil
}

// NewCollege .
func NewCollege() *College {
	return &College{database: make(map[int]Student)}
}
