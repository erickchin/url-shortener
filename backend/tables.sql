CREATE TABLE urls(
 id serial PRIMARY KEY,
 url_code VARCHAR(50) UNIQUE NOT NULL,
 short_url VARCHAR(2083) NOT NULL,
 original_url VARCHAR(2083) NOT NULL,
 created_on TIMESTAMP NOT NULL
);

CREATE TABLE logs(
 id serial PRIMARY KEY,
 url_code VARCHAR(50) NOT NULL,
 ip_address VARCHAR(45) NOT NULL,
 accessed_on TIMESTAMP NOT NULL
); 