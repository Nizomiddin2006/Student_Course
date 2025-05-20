package course

import "fmt"

type Course struct {
    ID       int
    Name     string
    Faculty  string
}

func NewCourse(id int, name string, faculty string) *Course {
    return &Course{
        ID:      id,
        Name:    name,
        Faculty: faculty,
    }
}

func (c *Course) Print() {
    fmt.Printf("ID: %d, Name: %s, Faculty: %s\n", c.ID, c.Name, c.Faculty)
}
