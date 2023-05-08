CREATE TABLE users(
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
    `code_verified` VARCHAR(255) NOT NULL,
    `email_verified_at` TIMESTAMP NULL,
    `created_by` INT NULL,
    `updated_by` INT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY(`id`)
    UNIQUE KEY users_email_unique(`email`),
    INDEX idx_users_email(`email`),
    INDEX idx_users_created_by(`created_by`),
    INDEX idx_users_updated_by(`updated_by`),
    CONSTRAINT FK_users_created_by FOREIGN KEY (`created_by`) REFERENCES admins(`id`) ON DELETE SET NULL, 
    CONSTRAINT FK_users_updated_by FOREIGN KEY (`updated_by`) REFERENCES admins(`id`) ON DELETE SET NULL
)