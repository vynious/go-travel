-- +goose Up
CREATE TABLE users (
                       id VARCHAR(255) PRIMARY KEY,
                       name VARCHAR(100) NOT NULL,
                       username VARCHAR(50) NOT NULL UNIQUE,
                       email VARCHAR(100) NOT NULL UNIQUE,
                       profile_picture VARCHAR(255),
                       creation_date TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE trip (
                      id BIGSERIAL PRIMARY KEY,
                      title VARCHAR(255) NOT NULL,
                      country VARCHAR(50) NOT NULL,
                      start_date DATE NOT NULL,
                      end_date DATE NOT NULL
);

CREATE TABLE connections (
                             party_a VARCHAR(255) NOT NULL,
                             party_b VARCHAR(255) NOT NULL,
                             PRIMARY KEY (party_a, party_b),
                             FOREIGN KEY (party_a) REFERENCES users(id) ON DELETE CASCADE,
                             FOREIGN KEY (party_b) REFERENCES users(id) ON DELETE CASCADE,
                             connected_date TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE user_p (
                           trip_id BIGSERIAL NOT NULL,
                           user_id VARCHAR(255) NOT NULL,
                           PRIMARY KEY (trip_id, user_id),
                           FOREIGN KEY (trip_id) REFERENCES trip(id) ON DELETE CASCADE,
                           FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE travel_entry (
                       id BIGSERIAL PRIMARY KEY,
                       user_id VARCHAR(255) NOT NULL,
                       trip_id BIGSERIAL NOT NULL,
                       location VARCHAR(255) NOT NULL,
                       description TEXT NOT NULL,
                       FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
                       FOREIGN KEY (trip_id) REFERENCES trip(id) ON DELETE CASCADE
);

CREATE TABLE media (
                       entry_id BIGSERIAL NOT NULL,
                       key VARCHAR(255) NOT NULL,
                       PRIMARY KEY (entry_id, key),
                       FOREIGN KEY (entry_id) REFERENCES travel_entry(id) ON DELETE CASCADE
);

CREATE TABLE comment (
                         id BIGSERIAL PRIMARY KEY,
                         entry_id BIGSERIAL NOT NULL,
                         user_id VARCHAR(255) NOT NULL,
                         content TEXT NOT NULL,
                         commented_on DATE NOT NULL DEFAULT NOW(),
                         FOREIGN KEY (entry_id) REFERENCES travel_entry(id) ON DELETE CASCADE,
                         FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE IF EXISTS comment;
DROP TABLE IF EXISTS media;
DROP TABLE IF EXISTS travel_entry;
DROP TABLE IF EXISTS user_trip;
DROP TABLE IF EXISTS connections;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS trip;


