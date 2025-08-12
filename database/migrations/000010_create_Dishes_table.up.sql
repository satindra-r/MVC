CREATE TABLE `Dishes` (
  `DishId` int(11) NOT NULL,
  `ItemId` int(11) DEFAULT NULL,
  `OrderId` int(11) DEFAULT NULL,
  `DishCount` int(11) DEFAULT NULL,
  `SplInstructions` text DEFAULT NULL,
  `Prepared` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`DishId`),
  KEY `OrderId` (`OrderId`),
  KEY `ItemId` (`ItemId`),
  CONSTRAINT `Dishes_ibfk_1` FOREIGN KEY (`OrderId`) REFERENCES `Orders` (`OrderId`),
  CONSTRAINT `Dishes_ibfk_2` FOREIGN KEY (`ItemId`) REFERENCES `Items` (`ItemId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;