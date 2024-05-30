-- Create Database
CREATE DATABASE IF NOT EXISTS TaskDB;
-- Move on TaskDB and create table
USE TaskDB;
CREATE TABLE IF NOT EXISTS tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status INT NOT NULL
);

-- Insert some test data
INSERT INTO tasks (name, status) VALUES ('Task 1', 0);
INSERT INTO tasks (name, status) VALUES ('Task 2', 1);
INSERT INTO tasks (name, status) VALUES ('Task 3', 0);
INSERT INTO tasks (name, status) VALUES ('Task 4', 1);
INSERT INTO tasks (name, status) VALUES ('Task 5', 0);
INSERT INTO tasks (name, status) VALUES ('Task 6', 1);
INSERT INTO tasks (name, status) VALUES ('Task 7', 0);
INSERT INTO tasks (name, status) VALUES ('Task 8', 1);
INSERT INTO tasks (name, status) VALUES ('Task 9', 0);
INSERT INTO tasks (name, status) VALUES ('Task 10', 1);