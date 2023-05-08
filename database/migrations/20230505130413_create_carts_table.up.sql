CREATE TABLE carts(
    `id` INT NOT NULL AUTO_INCREMENT,
    `product_id` INT NOT NULL,
    `user_id` INT NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    PRIMARY KEY(`id`),
    CONSTRAINT FK_class_rooms_product_id FOREIGN KEY (`product_id`) REFERENCES products(`id`) ON DELETE SET NULL, 
    CONSTRAINT FK_class_rooms_user_id FOREIGN KEY (`user_id`) REFERENCES users(`id`) ON DELETE SET NULL, 
)