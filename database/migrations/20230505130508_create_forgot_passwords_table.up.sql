CREATE TABLE forgot_passwords(
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `valid` BOOLEAN NOT NULL,
    `expired_at` TIMESTAMP,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY(`id`),
    CONSTRAINT FK_forgot_passwords_user_id FOREIGN KEY (`user_id`) REFERENCES users(`id`) ON DELETE SET NULL
)