-- migrate:up

INSERT INTO users(firstname, lastname, email, status)
VALUES ('John', 'Doe', 'john@example.com', 'ACTIVE'),
       ('Winnie', 'Michael', 'wmichael@example.com', 'ACTIVE'),
       ('James', 'Alphonso', 'alphonso@example.com', 'INVITED');

-- migrate:down
TRUNCATE TABLE users;