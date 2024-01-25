CREATE TABLE `order` (
  `id` varchar(128) NOT NULL,
  `price` float DEFAULT NULL,
  `tax` float DEFAULT NULL,
  `final_price` float DEFAULT NULL,
  PRIMARY KEY (`id`)
) 