CREATE TABLE kurs (
    id SERIAL PRIMARY KEY,          
    fanlar VARCHAR(100) NOT NULL,     
    bosqich VARCHAR(50) NOT NULL, 
    oquvchi_soni INT DEFAULT 0,    
    boshlangan_vaqti TIMESTAMP NOT NULL, 
    tugash_vaqti TIMESTAMP NOT NULL
    
);
