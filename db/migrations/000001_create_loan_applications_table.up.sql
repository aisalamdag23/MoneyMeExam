CREATE TABLE IF NOT EXISTS loan_applications(
   id serial PRIMARY KEY,
   amount_required VARCHAR (255) NOT NULL,
   term VARCHAR (50) NOT NULL,
   first_name VARCHAR (300) NOT NULL,
   last_name VARCHAR (300) NOT NULL,
   date_of_birth DATE NOT NULL,
   mobile VARCHAR(50) NOT NULL,
   email VARCHAR(100) NOT NULL,
   created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at timestamp NULL DEFAULT NULL,
   deleted_at timestamp NULL DEFAULT NULL
);