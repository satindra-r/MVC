CREATE TABLE `Orders` (
  `OrderId` int(11) NOT NULL,
  `UserId` int(11) DEFAULT NULL,
  `Price` decimal(10,2) DEFAULT NULL,
  `Paid` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`OrderId`),
  KEY `UserId` (`UserId`),
  CONSTRAINT `Orders_ibfk_1` FOREIGN KEY (`UserId`) REFERENCES `Users` (`UserId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;