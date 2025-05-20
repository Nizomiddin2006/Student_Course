package studentcourse

import (
    "fmt"
    "golang-template/internal/student"
    "golang-template/internal/course"
)

func StudentCourse(s *student.Student, c *course.Course) {
    fmt.Printf("%s student enrolled in the %s course.\n", s.Name, c.Name)
}
