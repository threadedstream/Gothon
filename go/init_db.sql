CREATE TABLE IF NOT EXISTS statistics(
    id SERIAL PRIMARY KEY,
    date DATE NOT NULL DEFAULT CURRENT_DATE,
    views INT,
    clicks INT,
    cost   FLOAT,
);

