CREATE TABLE IF NOT EXISTS `taks` (
    `id` VARCHAR(255),
    `title` varchar(250),
    `description` varchar(250),
    `completed` boolean,
    PRIMARY KEY (`id`)
);