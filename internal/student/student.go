package student

import "fmt"

type Student struct {
    ID       int
    Name     string
    Age      int
    Faculty  string
}

func NewStudent(id int, name string, age int, faculty string) *Student {
    return &Student{
        ID:      id,
        Name:    name,
        Age:     age,
        Faculty: faculty,
    }
}

func (s *Student) Print() {
    fmt.Printf("ID: %d, Name: %s, Age: %d, Faculty: %s\n", s.ID, s.Name, s.Age, s.Faculty)
}
