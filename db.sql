CREATE DATABASE `golang` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

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
                                          check_in_time DATETIME,
                                          latitude_check_in  DOUBLE,
                                        longitude_check_in DOUBLE,
                                          check_out_time DATETIME,
                                          latitude_check_out  DOUBLE,
                                          longitude_check_out DOUBLE,
                                          total_hours FLOAT,
                                          over_time FLOAT,
                                          bonus DECIMAL(10,2),
                                          salary DECIMAL(10,2),
                                          location VARCHAR(255),
                                          FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

-- Tạo bảng WeeklyPoints
CREATE TABLE WeeklyPoints (
                              id SERIAL PRIMARY KEY,
                              user_id varchar(255),
                              week_start_date DATE,
                              week_end_date DATE,
                              points INT,
                              FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

-- Tạo bảng PrizePool
CREATE TABLE PrizePool (
                           id SERIAL PRIMARY KEY,
                           name VARCHAR(255),
                           description TEXT,
                           value DECIMAL(10, 2), -- Giá trị có thể điều chỉnh tùy theo yêu cầu
                           active BOOLEAN
);

drop table if exists `Attendance`;