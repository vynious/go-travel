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
                      title VARCHAR(255),
                      country VARCHAR(50),
                      start_date DATE,
                      end_date DATE
);

CREATE TABLE connections (
                             party_a VARCHAR(255),
                             party_b VARCHAR(255),
                             PRIMARY KEY (party_a, party_b),
                             FOREIGN KEY (party_a) REFERENCES users(id) ON DELETE CASCADE,
                             FOREIGN KEY (party_b) REFERENCES users(id) ON DELETE CASCADE,
                             connected_date TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE user_trip (
                           trip_id BIGSERIAL,
                           user_id VARCHAR(255),
                           PRIMARY KEY (trip_id, user_id),
                           FOREIGN KEY (trip_id) REFERENCES trip(id) ON DELETE CASCADE,
                           FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE travel_entry (
                       id BIGSERIAL PRIMARY KEY,
                       user_id VARCHAR(255),
                       trip_id BIGSERIAL,
                       location VARCHAR(255),
                       description TEXT,
                       FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
                       FOREIGN KEY (trip_id) REFERENCES trip(id) ON DELETE CASCADE
);

CREATE TABLE media (
                       id BIGSERIAL PRIMARY KEY,
                       entry_id BIGSERIAL,
                       url VARCHAR(255),
                       FOREIGN KEY (entry_id) REFERENCES travel_entry(id) ON DELETE CASCADE
);

CREATE TABLE comment (
                         id BIGSERIAL PRIMARY KEY,
                         entry_id BIGSERIAL,
                         user_id VARCHAR(255),
                         content TEXT,
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


