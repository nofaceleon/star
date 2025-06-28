
CREATE TABLE `users` (
                         `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                         `username` varchar(50) COLLATE utf8mb4_bin NOT NULL,
                         `password` char(32) COLLATE utf8mb4_bin NOT NULL,
                         `email` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
                         `created_at` datetime DEFAULT NULL,
                         `updated_at` datetime DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `users_username_uindex` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin



CREATE TABLE words (
                       id INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
                       uid INT UNSIGNED NOT NULL,
                       word VARCHAR ( 255 ) NOT NULL,
                       definition TEXT,
                       example_sentence TEXT,
                       chinese_translation VARCHAR ( 255 ),
                       pronunciation VARCHAR ( 255 ),
                       proficiency_level SMALLINT UNSIGNED,
                       created_at DATETIME,
                       updated_at DATETIME
);

ALTER TABLE words ADD UNIQUE (uid, word);