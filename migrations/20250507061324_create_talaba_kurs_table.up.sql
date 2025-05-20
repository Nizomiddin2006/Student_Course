CREATE TABLE student_course (
    student_id INT REFERENCES student(id) ON DELETE CASCADE, 
    course_id INT REFERENCES course(id) ON DELETE CASCADE,      
    course_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,       
    PRIMARY KEY (student_id, course_id)      
);
