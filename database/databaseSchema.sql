CREATE DATABASE armada;
\c armada;

-- Connect to the database before creating tables!
-- Also, if you need to drop them, use DROP TABLE "Users"; (Have to wrap in quotes!)
-- The quotes are used to preserve case sensitivity.
-- In the prisma schema, make sure to lowercase all the model entries (first column)

CREATE TABLE users (
  uuid VARCHAR(255) PRIMARY KEY NOT NULL,
  username VARCHAR(255) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  firstName VARCHAR(255) NOT NULL,
  lastName VARCHAR(255) NOT NULL,
  isAdmin BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE cohort (
  id SERIAL PRIMARY KEY NOT NULL,
  name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE course (
  id SERIAL PRIMARY KEY NOT NULL,
  name VARCHAR(255) UNIQUE NOT NULL,
  cohortId INTEGER NOT NULL,
  FOREIGN KEY (cohortId) REFERENCES cohort(id),
  UNIQUE (name, cohortId)
);

CREATE TABLE workspace (
  uuid VARCHAR(255) PRIMARY KEY NOT NULL,
  desiredCount INTEGER NOT NULL,
  website VARCHAR(255) NOT NULL,
  userId VARCHAR(255) NOT NULL,
  courseId INTEGER NOT NULL,
  FOREIGN KEY (courseId) REFERENCES course(id) ON DELETE CASCADE,
  FOREIGN KEY (userId) REFERENCES users(uuid) ON DELETE CASCADE
);

CREATE TABLE user_cohort (
  userId VARCHAR(255) NOT NULL,
  cohortId INTEGER NOT NULL,
  FOREIGN KEY (userId) REFERENCES users(uuid) ON DELETE CASCADE,
  FOREIGN KEY (cohortId) REFERENCES cohort(id) ON DELETE CASCADE,
  PRIMARY KEY (userId, cohortId)
);

CREATE TABLE user_course (
  userId VARCHAR(255) NOT NULL,
  courseId INTEGER NOT NULL,
  FOREIGN KEY (userId) REFERENCES users(uuid) ON DELETE CASCADE,
  FOREIGN KEY (courseId) REFERENCES course(id) ON DELETE CASCADE,
  PRIMARY KEY (userId, courseId)
);

INSERT INTO users (uuid, username, email, firstname, lastname, isadmin) VALUES ('original_admin', 'armadaadmin', 'thefourofours@gmail.com', 'armada', 'admin', TRUE);