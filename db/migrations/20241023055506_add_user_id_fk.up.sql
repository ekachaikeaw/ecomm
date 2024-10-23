ALTER TABLE `orders`
    ADD COLUMN `user_id` int NOT NOT NULL,
    ADD CONSTRAINT `user_id_fk` FOREIGN KEY (`user_id`)
        REFERECES `users` (`id`);