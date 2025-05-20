CREATE TABLE course (
    id SERIAL PRIMARY KEY,          
    subjects VARCHAR(100) NOT NULL,     
    level VARCHAR(50) NOT NULL, 
    student_count INT DEFAULT 0,    
    start_time TIMESTAMP NOT NULL, 
    end_time TIMESTAMP NOT NULL
);
