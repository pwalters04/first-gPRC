CREATE TABLE ToDo (
    id SERIAL PRIMARY KEY,
    title VARCHAR(40) DEFAULT NULL,
    description VARCHAR(1024) DEFAULT NULL,
    reminder TIMESTAMP NULL DEFAULT  NULL,
    PRIMARY KEY(id)
);
