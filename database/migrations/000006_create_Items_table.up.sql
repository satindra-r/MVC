CREATE TABLE `Items` (
  `ItemId` int(11) NOT NULL,
  `ItemName` varchar(255) DEFAULT NULL,
  `SectionId` int(11) DEFAULT NULL,
  `Price` decimal(10,2) DEFAULT NULL,
  PRIMARY KEY (`ItemId`),
  KEY `SectionId` (`SectionId`),
  CONSTRAINT `Items_ibfk_1` FOREIGN KEY (`SectionId`) REFERENCES `Sections` (`SectionId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;