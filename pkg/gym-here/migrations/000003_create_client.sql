CREATE TABLE clients (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         age INT NOT NULL,
                         training_id INT NOT NULL,
                         created_at TIMESTAMPTZ DEFAULT NOW(),
                         updated_at TIMESTAMPTZ DEFAULT NOW(),
                         FOREIGN KEY (training_id) REFERENCES trainings(id)
);
