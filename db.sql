CREATE TABLE IF NOT EXISTS `Users` (
                                       `id` int NOT NULL AUTO_INCREMENT,
                                       `user_id` varchar(255),
                                       `name` varchar(255) DEFAULT NULL,
                                       `pass_word` varchar(255) DEFAULT NULL,
                                       `email` varchar(255) DEFAULT NULL,
                                       `role` enum('admin','employee') DEFAULT NULL,
                                       `created_at` datetime(3) DEFAULT NULL,
                                       `updated_at` datetime(3) DEFAULT NULL,
                                       `token` varchar(255),
                                       PRIMARY KEY (`id`),
                                       UNIQUE KEY `user_id_UNIQUE` (`user_id`),
                                       UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE IF NOT EXISTS Attendance (
                                          id INT AUTO_INCREMENT PRIMARY KEY,
                                          user_id varchar(255),
                                          department VARCHAR(255),
                                          position VARCHAR(255),
                                          check_in_time DATETIME,
                                          check_out_time DATETIME,
                                          total_hours FLOAT,
                                          over_time FLOAT,
                                          bonus DECIMAL(10,2),
                                          salary DECIMAL(10,2),
                                          location VARCHAR(255),
                                          FOREIGN KEY (user_id) REFERENCES Users(user_id)
);