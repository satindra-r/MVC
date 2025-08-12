CREATE TABLE `Sections` (
  `SectionId` int(11) NOT NULL,
  `SectionName` varchar(255) DEFAULT NULL,
  `SectionOrder` int(11) DEFAULT NULL,
  PRIMARY KEY (`SectionId`),
  UNIQUE KEY `SectionOrder` (`SectionOrder`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;