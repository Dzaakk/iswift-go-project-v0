CREATE TABLE oauth_clients(
    `id` INT NOT NULL AUTO_INCREMENT,
    `client_id` VARCHAR(255) NOT NULL,
    `client_secret` VARCHAR(255) NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `redirect` VARCHAR(255) NOT NULL,
    `description` VARCHAR(255) NOT NULL,
    `scope` VARCHAR(255) NOT NULL,
    `created_by` INT NULL,
    `updated_by` INT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY(`id`)
    CONSTRAINT FK_oauth_clients_created_by FOREIGN KEY (`created_by`) REFERENCES admins(`id`) ON DELETE SET NULL,
    CONSTRAINT FK_oauth_clients_updated_by FOREIGN KEY (`updated_by`) REFERENCES admins(`id`) ON DELETE SET NULL
)