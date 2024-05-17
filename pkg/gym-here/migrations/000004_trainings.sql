CREATE TABLE trainings (
                           id SERIAL PRIMARY KEY,
                           name VARCHAR(255) NOT NULL UNIQUE,
                           workout_time VARCHAR(255),
                           workout_days VARCHAR(255),
                           instructor_id INT NOT NULL,
                           created_at TIMESTAMPTZ DEFAULT NOW(),
                           updated_at TIMESTAMPTZ DEFAULT NOW(),
                           FOREIGN KEY (instructor_id) REFERENCES instructors(id)
);
