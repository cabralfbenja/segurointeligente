-- db/init.sql
CREATE DATABASE IF NOT EXISTS segurointeligente;

USE segurointeligente;

CREATE TABLE IF NOT EXISTS insurance (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    insurance_type ENUM('car', 'phone', 'life') NOT NULL,
    time_from TIME NOT NULL,
    time_to TIME NOT NULL,
    value FLOAT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
