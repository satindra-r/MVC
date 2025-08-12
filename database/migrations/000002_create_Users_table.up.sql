CREATE TABLE `Users` (
                         `UserId` int(11) NOT NULL,
                         `UserName` varchar(255) DEFAULT NULL,
                         `Role` enum('User','Chef','Admin') DEFAULT NULL,
                         `PhoneNo` varchar(10) DEFAULT NULL,
                         `Address` varchar(511) DEFAULT NULL,
                         `Hash` char(60) DEFAULT NULL,
                         PRIMARY KEY (`UserId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;