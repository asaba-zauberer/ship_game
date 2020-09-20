use ca_hack;

SET NAMES utf8mb4;

INSERT INTO `collection_item` (`id`,`name`, `rarity`) VALUES ("1","スキン01", 1);
INSERT INTO `collection_item` (`id`,`name`, `rarity`) VALUES ("2","スキン02", 1);
INSERT INTO `collection_item` (`id`,`name`, `rarity`) VALUES ("3","スキン03", 1);
INSERT INTO `collection_item` (`id`,`name`, `rarity`) VALUES ("4","スキン04", 1);
INSERT INTO `collection_item` (`id`,`name`, `rarity`) VALUES ("5","スキン05", 2);
INSERT INTO `collection_item` (`id`,`name`, `rarity`) VALUES ("6","スキン06", 3);
INSERT INTO `collection_item` (`id`,`name`, `rarity`) VALUES ("7","スキン07", 2);

INSERT INTO `gacha_probability` (`collection_item_id`,`ratio`) VALUES ("1",5);
INSERT INTO `gacha_probability` (`collection_item_id`,`ratio`) VALUES ("2",5);
INSERT INTO `gacha_probability` (`collection_item_id`,`ratio`) VALUES ("3",5);
INSERT INTO `gacha_probability` (`collection_item_id`,`ratio`) VALUES ("4",5);
INSERT INTO `gacha_probability` (`collection_item_id`,`ratio`) VALUES ("5",3);
INSERT INTO `gacha_probability` (`collection_item_id`,`ratio`) VALUES ("6",3);
INSERT INTO `gacha_probability` (`collection_item_id`,`ratio`) VALUES ("7",1);
