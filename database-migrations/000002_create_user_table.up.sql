CREATE TABLE IF NOT EXISTS service_users(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    username VARCHAR(50) NOT NULL,
    password_hash VARCHAR(100) NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50)
);

INSERT INTO service_users(created_at, updated_at, username, password_hash, first_name, last_name)
VALUES(now(), now(), 'molsbee', '$2a$10$lBPmVTA40xA68utxTHqp8es/hPpgoNe27QH.Qs1CGHra1IzL.rn0y', 'William', 'Molsbee');