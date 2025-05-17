CREATE TABLE talaba_kurs (
    talaba_id INT REFERENCES talaba(id) ON DELETE CASCADE, 
    kurs_id INT REFERENCES kurs(id) ON DELETE CASCADE,      
    kurs_vaqti TIMESTAMP DEFAULT CURRENT_TIMESTAMP,       
    PRIMARY KEY (talaba_id, kurs_id)      
);
