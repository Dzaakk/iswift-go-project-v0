CREATE TABLE products(
    `id` INT NOT NULL AUTO_INCREMENT,
    `product_category_id` INT NOT NULL,
    `title` VARCHAR(255) NOT NULL,
    `price` INT NOT NULL,
    `description` TEXT NOT NULL,
    `image` VARCHAR NOT NULL,
    `video` VARCHAR NOT NULL,
    `created_by` INT NULL,
    `updated_by` INT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY(`id`)
    CONSTRAINT FK_products_created_by FOREIGN KEY (`created_by`) REFERENCES admins(`id`) ON DELETE SET NULL,
    CONSTRAINT FK_products_updated_by FOREIGN KEY (`updated_by`) REFERENCES admins(`id`) ON DELETE SET NULL
)