INSERT INTO ChefDB.Users (UserId,UserName,`Role`,PhoneNo,Address,Hash) VALUES
	 (1,'user1','User','1234567890','here','$2a$10$mP2tzk3PGz.42hSUgHlZ/uZ6LPt/gXP2DmJrTqcDWs720wVinodwa'),
	 (2,'chef','Chef','1234567890','','$2a$10$mP2tzk3PGz.42hSUgHlZ/uZ6LPt/gXP2DmJrTqcDWs720wVinodwa'),
	 (3,'admin','Admin','1234567890','qwertyui','$2a$10$xY.qRFDFNV7YE9A/6eTn9e2HZEH7JCFcwA2Dl1Ncuq5Em8rVfu./K'),
	 (4,'user2','User','1234567890','#2','$2a$10$mP2tzk3PGz.42hSUgHlZ/uZ6LPt/gXP2DmJrTqcDWs720wVinodwa');


INSERT INTO ChefDB.Sections (SectionId,SectionName,SectionOrder) VALUES
	 (1,'Appetizers',1),
	 (2,'Main Course',2),
	 (3,'Dessert',3);



INSERT INTO ChefDB.Items (ItemId,ItemName,SectionId,Price) VALUES
	 (1,'Spring Roll',1,120.00),
	 (2,'Naan',2,40.00),
	 (3,'Chocolate Ice Cream',3,60.00),
	 (4,'Momos',1,230.00),
	 (5,'Paneer',1,360.00),
	 (6,'Roti',2,150.00),
	 (7,'Kebab',1,123.00),
	 (8,'Kulcha',2,30.00),
	 (9,'Phulka',2,20.00),
	 (10,'Vannila Ice Cream',3,60.00),
	 (11,'Mango Ice Cream',3,60.00),
	 (12,'Mixed Ice Cream',3,10.00);

