CREATE TABLE IF NOT EXISTS products
(
    `id`          INT UNSIGNED   NOT NULL AUTO_INCREMENT,
    `name`        VARCHAR(255)   NOT NULL,
    `description` TEXT           NOT NULL,
    `image`       VARCHAR(255),
    `price`       DECIMAL(10, 2) NOT NULL,
    `quantity`    INT UNSIGNED   NOT NULL,
    `createdAt`   TIMESTAMP      NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id)
);