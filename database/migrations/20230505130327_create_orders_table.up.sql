CREATE TABLE orders (
    `id` INT NOT NULL AUTO_INCREMENT,
    `discount_id` INT NOT NULL,
    `user_id` INT NOT NULL,
    `checkout_link` VARCHAR(255) NOT NULL,
    `price` INT NOT NULL,
    `total_price` INT NOT NULL,
    `external_id` VARCHAR(255) NOT NULL,
    `status` VARCHAR(255) NOT NULL,
    `created_by` INT NULL,
    `updated_by` INT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY(`id`),
    CONSTRAINT FK_orders_created_by FOREIGN KEY (`created_by`) REFERENCES admins(`id`) ON DELETE SET NULL, 
    CONSTRAINT FK_orders_updated_by FOREIGN KEY (`updated_by`) REFERENCES admins(`id`) ON DELETE SET NULL,
    CONSTRAINT FK_orders_user_id FOREIGN KEY (`user_id`) REFERENCES users(`id`) ON DELETE SET NULL 
)