CREATE TABLE instructors (
                             id SERIAL PRIMARY KEY,
                             name VARCHAR(255) NOT NULL,
                             profile_sport VARCHAR(255) NOT NULL,
                             qualification VARCHAR(255),
                             work_experience INT,
                             created_at TIMESTAMPTZ DEFAULT NOW(),
                             updated_at TIMESTAMPTZ DEFAULT NOW()
);
