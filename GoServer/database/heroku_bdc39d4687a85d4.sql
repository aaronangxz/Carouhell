-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema heroku_bdc39d4687a85d4
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `heroku_bdc39d4687a85d4` ;

-- -----------------------------------------------------
-- Schema heroku_bdc39d4687a85d4
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `heroku_bdc39d4687a85d4` DEFAULT CHARACTER SET utf8 ;
USE `heroku_bdc39d4687a85d4` ;

-- -----------------------------------------------------
-- Table `heroku_bdc39d4687a85d4`.`acc_tab`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `heroku_bdc39d4687a85d4`.`acc_tab` ;

CREATE TABLE IF NOT EXISTS `heroku_bdc39d4687a85d4`.`acc_tab` (
  `a_user_id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_name` VARCHAR(45) NOT NULL,
  `user_email` VARCHAR(80) NOT NULL,
  `user_ctime` INT(11) NULL,
  `user_status` TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0 - Online\n1 - Offline\n',
  `user_image` BLOB NULL DEFAULT NULL,
  `user_last_login` INT(11) NULL DEFAULT NULL,
  `user_rating` DOUBLE(2,1) UNSIGNED NULL DEFAULT 0,
  PRIMARY KEY (`a_user_id`),
  UNIQUE INDEX `user_email_UNIQUE` (`user_email` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `heroku_bdc39d4687a85d4`.`notification_tab`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `heroku_bdc39d4687a85d4`.`notification_tab` ;

CREATE TABLE IF NOT EXISTS `heroku_bdc39d4687a85d4`.`notification_tab` (
  `n_notification_id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `n_user_id` INT(10) UNSIGNED NOT NULL,
  `notification_text` VARCHAR(256) NOT NULL,
  `notification_url` VARCHAR(256) NOT NULL,
  `notification_time` INT(11) NOT NULL,
  `notification_ui_info` BLOB NULL DEFAULT NULL,
  PRIMARY KEY (`n_notification_id`, `n_user_id`),
  INDEX `a_user_id, n_user_id_idx` (`n_user_id` ASC) VISIBLE,
  CONSTRAINT `a_user_id, n_user_id`
    FOREIGN KEY (`n_user_id`)
    REFERENCES `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `heroku_bdc39d4687a85d4`.`listing_tab`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `heroku_bdc39d4687a85d4`.`listing_tab` ;

CREATE TABLE IF NOT EXISTS `heroku_bdc39d4687a85d4`.`listing_tab` (
  `l_item_id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `item_name` VARCHAR(256) NOT NULL,
  `item_price` INT(10) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT 'Datatype - DOUBLE(10, 2)',
  `item_quantity` SMALLINT(3) UNSIGNED NOT NULL DEFAULT 1,
  `item_purchasedquantity` SMALLINT(3) UNSIGNED NULL DEFAULT 0,
  `item_description` VARCHAR(256) NOT NULL,
  `item_shippinginfo` TINYINT(2) NULL DEFAULT 0,
  `item_location` VARCHAR(20) NOT NULL,
  `item_status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT 'ITEM_STATUS_ALL = 0\nITEM_STATUS_NORMAL = 1\nITEM_STATUS_SOLDOUT = 2\nITEM_STATUS_DELETED = 3',
  `item_category` TINYINT(2) NOT NULL COMMENT '1 - Fruits & Vegetables\n2 - Meat & Seafood\n3 - Dairy & Chilled\n4 - Breakfast\n5 - Beer, Wine & Spirits\n6 - Food Pantry\n7 - Kids & Baby\n8 - Healthcare & Medicines\n9 - Skin Care\n10 - Personal Care\n11 - Apparels\n12 - Mobile & Gadgets\n13 - Home Appliances\n14 - Sports & Outdoors\n15 - Automotive\n16 - Jewellery & Accessories\n17 - Home & Living\n18 - Video Games\n19 - Books\n20 - Pets\n21 - Travel\n22 - Coupons & Services\n23 - Beverages\n',
  `item_image` BLOB NULL DEFAULT NULL,
  `l_seller_id` INT(10) UNSIGNED NOT NULL,
  `listing_ctime` INT(11) UNSIGNED NOT NULL DEFAULT 0,
  `listing_mtime` INT(11) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`l_item_id`),
  INDEX `user_id_idx` (`l_seller_id` ASC) VISIBLE,
  CONSTRAINT `a_user_id,l_seller_id`
    FOREIGN KEY (`l_seller_id`)
    REFERENCES `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `heroku_bdc39d4687a85d4`.`wallet_tab`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `heroku_bdc39d4687a85d4`.`wallet_tab` ;

CREATE TABLE IF NOT EXISTS `heroku_bdc39d4687a85d4`.`wallet_tab` (
  `w_user_id` INT(10) UNSIGNED NOT NULL,
  `wallet_balance` INT(10) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT 'Datatype - DOUBLE(10, 2)',
  `wallet_status` TINYINT(1) NULL,
  `last_topup` INT(11) NULL,
  `last_used` INT(11) NULL COMMENT '\n',
  PRIMARY KEY (`w_user_id`),
  INDEX `a_user_id, w_user_id_idx` (`w_user_id` ASC) VISIBLE,
  CONSTRAINT `a_user_id, w_user_id`
    FOREIGN KEY (`w_user_id`)
    REFERENCES `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `heroku_bdc39d4687a85d4`.`wallet_transaction`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `heroku_bdc39d4687a85d4`.`wallet_transaction` ;

CREATE TABLE IF NOT EXISTS `heroku_bdc39d4687a85d4`.`wallet_transaction` (
  `wt_transaction_id` INT(10) UNSIGNED NULL,
  `wt_user_id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `transaction_details` VARCHAR(256) NULL DEFAULT NULL,
  `transaction_ctime` INT(11) NULL,
  `transaction_mtime` INT(11) NULL,
  `transaction_amount` INT(10) UNSIGNED NOT NULL DEFAULT 0.00 COMMENT 'Datatype - DOUBLE(10, 2)',
  `transaction_status` TINYINT(1) NULL DEFAULT 2 COMMENT '1 - success\n2 - failed\n3 - cancelled\n',
  PRIMARY KEY (`wt_user_id`),
  UNIQUE INDEX `transaction_id_UNIQUE` (`wt_transaction_id` ASC) VISIBLE,
  CONSTRAINT `a_user_id, wt_user_id`
    FOREIGN KEY (`wt_user_id`)
    REFERENCES `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `heroku_bdc39d4687a85d4`.`listing_transactions_tab`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `heroku_bdc39d4687a85d4`.`listing_transactions_tab` ;

CREATE TABLE IF NOT EXISTS `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (
  `lt_transaction_id` INT(10) UNSIGNED NULL AUTO_INCREMENT,
  `lt_item_id` INT(10) UNSIGNED NOT NULL,
  `lt_user_id` INT(10) UNSIGNED NOT NULL,
  `transaction_ctime` INT(11) NOT NULL,
  `transaction_mtime` INT(11) NOT NULL,
  `transaction_price` INT(11) NULL DEFAULT NULL COMMENT 'Datatype - DOUBLE(10, 2)',
  `transaction_status` TINYINT(1) NULL DEFAULT 0 COMMENT 'Successful = 0\nFailed = 1\nCancelled = 2',
  PRIMARY KEY (`lt_item_id`, `lt_user_id`),
  UNIQUE INDEX `transaction_id_UNIQUE` (`lt_transaction_id` ASC) VISIBLE,
  INDEX `a_user_id,lt_user_id_idx` (`lt_user_id` ASC) VISIBLE,
  CONSTRAINT `l_item_id,lt_item_id`
    FOREIGN KEY (`lt_item_id`)
    REFERENCES `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `a_user_id,lt_user_id`
    FOREIGN KEY (`lt_user_id`)
    REFERENCES `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `heroku_bdc39d4687a85d4`.`acc_credentials_tab`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `heroku_bdc39d4687a85d4`.`acc_credentials_tab` ;

CREATE TABLE IF NOT EXISTS `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (
  `c_user_id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_password` VARCHAR(256) NOT NULL,
  `user_security_question` TINYINT(1) UNSIGNED NOT NULL,
  `user_security_answer` VARCHAR(256) NOT NULL,
  PRIMARY KEY (`c_user_id`),
  CONSTRAINT `a_user_id, c_user_id`
    FOREIGN KEY (`c_user_id`)
    REFERENCES `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `heroku_bdc39d4687a85d4`.`user_review_tab`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `heroku_bdc39d4687a85d4`.`user_review_tab` ;

CREATE TABLE IF NOT EXISTS `heroku_bdc39d4687a85d4`.`user_review_tab` (
  `review_id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `rv_user_id` INT(10) UNSIGNED NOT NULL,
  `rv_seller_id` INT(10) UNSIGNED NOT NULL,
  `ratings` TINYINT(1) NOT NULL DEFAULT -1,
  `review_text` VARCHAR(256) NULL DEFAULT NULL,
  `ctime` INT(11) UNSIGNED NOT NULL,
  PRIMARY KEY (`review_id`),
  INDEX `a_user_id, rv_user_id_idx` (`rv_user_id` ASC) INVISIBLE,
  INDEX `l_seller_id, rv_seller_id_idx` (`rv_seller_id` ASC) VISIBLE,
  CONSTRAINT `a_user_id, rv_user_id`
    FOREIGN KEY (`rv_user_id`)
    REFERENCES `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `l_seller_id, rv_seller_id`
    FOREIGN KEY (`rv_seller_id`)
    REFERENCES `heroku_bdc39d4687a85d4`.`listing_tab` (`l_seller_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `heroku_bdc39d4687a85d4`.`listing_reactions_tab`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `heroku_bdc39d4687a85d4`.`listing_reactions_tab` ;

CREATE TABLE IF NOT EXISTS `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (
  `reactions_id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `rt_user_id` INT(10) UNSIGNED NOT NULL,
  `rt_item_id` INT(10) UNSIGNED NOT NULL,
  `reaction_type` TINYINT(1) UNSIGNED NOT NULL DEFAULT 0,
  `comment` VARCHAR(256) NULL DEFAULT NULL,
  `ctime` INT(11) UNSIGNED NOT NULL,
  PRIMARY KEY (`reactions_id`),
  INDEX `a_user_id, rt_user_id_idx` (`rt_user_id` ASC) VISIBLE,
  INDEX `l_item_id, rt_item_id_idx` (`rt_item_id` ASC) VISIBLE,
  CONSTRAINT `a_user_id, rt_user_id`
    FOREIGN KEY (`rt_user_id`)
    REFERENCES `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE,
  CONSTRAINT `l_item_id, rt_item_id`
    FOREIGN KEY (`rt_item_id`)
    REFERENCES `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Data for table `heroku_bdc39d4687a85d4`.`acc_tab`
-- -----------------------------------------------------
START TRANSACTION;
USE `heroku_bdc39d4687a85d4`;
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (1, 'Gus Amaral', 'gus_amaral@gmail.com', 1313192021, 0, NULL, 1463790569, 5);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (2, 'Jeanne Madrigal', 'jeanne_m@gmail.com', 1361681057, 0, NULL, 1466947661, 5);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (3, 'Cristin Allums', 'cristin_allums@gmail.com', 1363790569, 0, NULL, 1467812688, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (4, 'Marlene Mctaggart', 'marlene_mctaggart@gmail.com', 1366947661, 1, NULL, 1463790569, 3);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (5, 'Larisa Komar', 'larisa_komar@gmail.com', 1367812688, 0, NULL, 1466947661, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (6, 'Jackie Hayton', 'jackie_hayton@gmail.com', 1343421873, 1, NULL, 1458247651, 4);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (7, 'Shella Belote', 'shella_belote@gmail.com', 1358247651, 0, NULL, 1453535532, 4);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (8, 'Antonetta Concepcion', 'antonetta_concepcion@gmail.com', 1353535532, 0, NULL, 1464286205, 5);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (9, 'Winnie Malia', 'winnie_malia@gmail.com', 1343421873, 1, NULL, 1442625424, 5);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (10, 'Roman Bush', 'roman_bush@gmail.com', 1358247651, 0, NULL, 1443421873, 5);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (11, 'Toya Guerette', 'toya_guerette@gmail.com', 1353535532, 0, NULL, 1458247651, 4);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (12, 'Pedro Routt', 'pedro_routt@gmail.com', 1364286205, 1, NULL, 1414744275, 4);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (13, 'Arnette Mcmurray', 'arnette_mcmurray@gmail.com', 1342625424, 1, NULL, 1456406392, 4);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (14, 'Arletta Winburn', 'arletta_winburn@gmail.com', 1313744275, 0, NULL, 1474924465, 4);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (15, 'Stephen Ebert', 'stephen_ebert@gmail.com', 1356406392, 1, NULL, 1414744275, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (16, 'Shantay Coster', 'shantay_coster@gmail.com', 1374924465, 0, NULL, 1458247651, 3);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (17, 'Noelle Kehrer', 'noelle_kehrer@gmail.com', 1364926820, 0, NULL, 1453535532, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (18, 'Julienne Abston', 'julienne_abston@gmail.com', 1389345471, 1, NULL, 1464286205, 4);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (19, 'Helaine Tilson', 'helaine_tilson@gmail.com', 1344975285, 0, NULL, 1420871670, 3);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (20, 'Penni Printup', 'penni_printup@gmail.com', 1326928868, 0, NULL, 1467896775, 4);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (21, 'Katherina Marasco', 'katherina_marasco@gmail.com', 1320871670, 0, NULL, 1448562440, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (22, 'Pennie Mcginley', 'pennie_mcginley@hotmail.com', 1314072452, 1, NULL, 1485033404, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (23, 'Faustina Knouse', 'fautina_knouse@hotmail.com', 1342248046, 1, NULL, 1448562440, 4);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (24, 'Selma Allsop', 'selma_allsop@hotmail.com', 1397358109, 0, NULL, 1497358109, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (25, 'Audrea Carraway', 'audrea_carraway@hotmail.com', 1367896775, 0, NULL, 1467896775, 4);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (26, 'Nadia Collington', 'nadia_collington@hotmail.com', 1348562440, 0, NULL, 1485033404, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (27, 'Joette Baily', 'joette_baily@hotmail.com', 1385033404, 1, NULL, 1485033404, 4);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (28, 'Maryetta Poppell', 'maryetta_poppell@hotmail.com', 1339107621, 0, NULL, 1439107621, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (29, 'Teodoro Ceniceros', 'teodoro_ceniceros@hotmail.com', 1337834963, 0, NULL, 1437834963, 5);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (30, 'Ilda Medders', 'ilda_medders@hotmail.com', 1344428381, 1, NULL, 1444428381, 4);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (31, 'Judith Police', 'judith_police@hotmail.com', 1362454583, 1, NULL, 1462454583, 5);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (32, 'Charley Vanfleet', 'charley_vanfleet@hotmail.com', 1372241246, 0, NULL, 1453535532, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (33, 'Cheryl Parrilla', 'cheryl_parrila@hotmail.com', 1340525753, 0, NULL, 1443421873, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (34, 'Eveline Weathersby', 'eveline_weathersby@hotmail.com', 1324984974, 0, NULL, 1458247651, 4);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (35, 'Roselee Monfort', 'roselee_monfort@hotmail.com', 1351851856, 1, NULL, 1456406392, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (36, 'Gerry Maudlin', 'gerry_maudlin@hotmail.com', 1372561716, 0, NULL, 1474924465, 4);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (37, 'Kiesha Carone', 'kiesha_carone@yahoo.com', 1383204801, 1, NULL, 1464926820, 5);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (38, 'Audra Ewert', 'audra_ewert@yahoo.com', 1357557670, 0, NULL, 1489345471, 5);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (39, 'Anja Ranieri', 'anja_ranieri@yahoo.com', 1375419828, 0, NULL, 1475419828, 3);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (40, 'Letitia Delatorre', 'letitia_delatorre@yahoo.com', 1346232138, 1, NULL, 1446232148, 3);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (41, 'Lyndia Ali', 'lyndia_ali@yahoo.com', 1352091561, 0, NULL, 1452091561, 4);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (42, 'Elisa Meeker', 'elisa_meeker@yahoo.com', 1332828484, 1, NULL, 1432828484, 4);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (43, 'Keva Sauseda', 'keva_sauseda@yahoo.com', 1318931454, 1, NULL, 1418931454, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (44, 'Alycia Lamontagne', 'alycia_lamontagne@yahoo.com', 1345714626, 0, NULL, 1445714626, 5);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (45, 'Miguelina Mast', 'miguelina_mast@yahoo.com', 1355644622, 0, NULL, 1455644622, 5);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (46, 'Orville Mccaskey', 'orville_mccaskey@yahoo.com', 1335608987, 0, NULL, 1420871670, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (47, 'Mabel Ascencio', 'mabel_ascencio@yahoo.com', 1341052651, 0, NULL, 1414072452, 3);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (48, 'Jenise Radtke', 'jenise_radtke@yahoo.com', 1368728807, 0, NULL, 1442248046, 3);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (49, 'Reyna Haymond', 'reyna_haymond@yahoo.com', 1332375922, 1, NULL, 1497358109, 3);
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_ctime`, `user_status`, `user_image`, `user_last_login`, `user_rating`) VALUES (50, 'Masako Leitzel', 'masako_leitzel@yahoo.com', 1356267140, 0, NULL, 1444428381, 5);

COMMIT;


-- -----------------------------------------------------
-- Data for table `heroku_bdc39d4687a85d4`.`notification_tab`
-- -----------------------------------------------------
START TRANSACTION;
USE `heroku_bdc39d4687a85d4`;
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (1, 1, 'SALE', 'https://my_shop_db/sales', 1555496468, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (2, 2, 'PROMOTION ', 'https://my_shop_db/sales', 1577574753, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (3, 3, 'ITEM IS SHIPPED ', 'https://my_shop_db/shipping', 1593731470, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (4, 4, 'PAYMENT COMPLETED ', 'https://my_shop_db/payment', 1540805529, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (5, 5, 'SOMETHING IN CART', 'https://my_shop_db', 1551411071, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (6, 6, 'PROMOTION ', 'https://my_shop_db/sales', 1550033439, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (7, 7, 'SALE', 'https://my_shop_db/sales', 1578971413, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (8, 8, 'SALE', 'https://my_shop_db/sales', 1575023286, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (9, 9, 'SALE', 'https://my_shop_db/sales', 1551036678, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (10, 10, 'PROMOTION ', 'https://my_shop_db/sales', 1449776464, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (11, 11, 'ITEM IS SHIPPED ', 'https://my_shop_db/shipping', 1496135797, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (12, 12, 'PROMOTION ', 'https://my_shop_db/sales', 1500947796, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (13, 13, 'PROMOTION ', 'https://my_shop_db/sales', 1469676602, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (14, 14, 'SALE', 'https://my_shop_db/sales', 1468373981, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (15, 15, 'PROMOTION ', 'https://my_shop_db/sales', 1489078911, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (16, 16, 'PROMOTION ', 'https://my_shop_db/sales', 1480523887, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (17, 17, 'ITEM IS SHIPPED ', 'https://my_shop_db/shipping', 1438691072, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (18, 18, 'PROMOTION ', 'https://my_shop_db/sales', 1469968837, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (19, 19, 'PROMOTION ', 'https://my_shop_db/sales', 1493938666, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (20, 20, 'PROMOTION ', 'https://my_shop_db/sales', 1602285084, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (21, 21, 'ITEM IS SHIPPED ', 'https://my_shop_db/shipping', 1540126699, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (22, 22, 'PAYMENT COMPLETED ', 'https://my_shop_db/payment', 1564427132, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (23, 23, 'PROMOTION ', 'https://my_shop_db/sales', 1564594822, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (24, 24, 'PROMOTION ', 'https://my_shop_db/sales', 1576165335, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (25, 25, 'ITEM IS SHIPPED ', 'https://my_shop_db/shipping', 1599705120, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (26, 26, 'PAYMENT COMPLETED ', 'https://my_shop_db/payment', 1541936015, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (27, 27, 'PROMOTION ', 'https://my_shop_db/sales', 1546043842, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (28, 28, 'ITEM IS SHIPPED ', 'https://my_shop_db/shipping', 1572604299, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (29, 29, 'PAYMENT COMPLETED ', 'https://my_shop_db/payment', 1537296559, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (30, 30, 'SALE', 'https://my_shop_db/sales', 1602886700, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (31, 31, 'SALE', 'https://my_shop_db/sales', 1587965821, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (32, 32, 'PROMOTION ', 'https://my_shop_db/sales', 1570320134, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (33, 33, 'ITEM IS SHIPPED ', 'https://my_shop_db/shipping', 1575851364, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (34, 34, 'PAYMENT COMPLETED ', 'https://my_shop_db/payment', 1541839997, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (35, 35, 'SOMETHING IN CART', 'https://my_shop_db', 1561406577, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (36, 36, 'SALE', 'https://my_shop_db/sales', 1457061757, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (37, 37, 'PROMOTION ', 'https://my_shop_db/sales', 1463178849, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (38, 38, 'SALE', 'https://my_shop_db/sales', 1445686446, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (39, 39, 'SALE', 'https://my_shop_db/sales', 1476999264, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (40, 40, 'PROMOTION ', 'https://my_shop_db/sales', 1467682136, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (41, 41, 'PROMOTION ', 'https://my_shop_db/sales', 1444556248, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (42, 42, 'PROMOTION ', 'https://my_shop_db/sales', 1478654727, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (43, 43, 'ITEM IS SHIPPED ', 'https://my_shop_db/shipping', 1463064852, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (44, 44, 'PAYMENT COMPLETED ', 'https://my_shop_db/payment', 1463824321, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (45, 45, 'SOMETHING IN CART', 'https://my_shop_db', 1448104036, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (46, 46, 'SALE', 'https://my_shop_db/sales', 1467758404, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (47, 47, 'PROMOTION ', 'https://my_shop_db/sales', 1442108552, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (48, 48, 'ITEM IS SHIPPED ', 'https://my_shop_db/shipping', 1468003088, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (49, 49, 'PAYMENT COMPLETED ', 'https://my_shop_db/payment', 1470154497, NULL);
INSERT INTO `heroku_bdc39d4687a85d4`.`notification_tab` (`n_notification_id`, `n_user_id`, `notification_text`, `notification_url`, `notification_time`, `notification_ui_info`) VALUES (50, 50, 'SOMETHING IN CART', 'https://my_shop_db', 1437584342, NULL);

COMMIT;


-- -----------------------------------------------------
-- Data for table `heroku_bdc39d4687a85d4`.`listing_tab`
-- -----------------------------------------------------
START TRANSACTION;
USE `heroku_bdc39d4687a85d4`;
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (1, 'Apple', 2, 8, 6, 'Their dense flesh is creamy yellow and crisp, offering a mildly sweet flavor.', 0, 'Malaysia', 1, 1, NULL, 1, 1427609346, 1427954946);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (2, 'Orange', 3, 5, 3, 'Delicious and juicy orange fruit, good for juicing and containing an impressive list of essential nutrients, vitamins, minerals for normal growth and development and overall well-being.', 2, 'Taiwan', 0, 1, NULL, 2, 1427782146, 1428127746);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (3, 'Pear', 3, 8, 6, 'Packham pears have a wide-bottomed shape and a smooth green skin that ripens to yellow.', 1, 'Australia', 1, 1, NULL, 3, 1427954946, 1428300546);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (4, 'Blueberry', 4, 7, 5, 'Plump, juicy, and sweet, with vibrant colours ranging from deep purple-blue to blue-black and highlighted by a silvery sheen called a bloom, blueberries are one of nature\'s great treasures.', 2, 'Japan', 0, 1, NULL, 4, 1428127746, 1428473346);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (5, 'Watermelon', 6, 1, 1, 'The large round fruit has a hard green rind, a watery red pulp and small brown seeds.', 0, 'Taiwan', 1, 1, NULL, 5, 1428300546, 1428646146);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (6, 'Chicken ', 10, 4, 2, 'The chickens not kept in cages but are instead raised in a modernised, temperature-controlled and environmentally friendly barn which gives the chickens ample space to roam with access to food and water.', 0, 'South Korea', 0, 2, NULL, 6, 1428473346, 1428818946);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (7, 'Pork ', 9, 1, 1, 'Pork Loin Boneless trimmed & ready to used, no added artificial coloring & flavours. it is portion to approx 125g each and consist 4 pcs Individually freeze, approx packed 500-550g.', 2, 'Phillippines', 0, 2, NULL, 7, 1428646146, 1428991746);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (8, 'Beef ', 15, 10, 8, 'Black Angus 150days Grain fed Beef Tenderloin with 100% natural, no added preservative, coloring & additives.', 0, 'Phillippines', 2, 2, NULL, 8, 1428818946, 1429164546);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (9, 'Fish ', 12, 9, 7, 'Daily processed fresh Batang fillet', 0, 'South Korea', 1, 2, NULL, 9, 1428991746, 1429337346);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (10, 'Duck ', 9, 2, 0, ' Processed locally to ensure low carbon footprint, all products are thermal-vacuum sealed for superior product quality.', 0, 'Colombia', 2, 2, NULL, 10, 1429164546, 1429510146);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (11, 'Cheese', 12, 7, 5, 'Cheese slices make from FreshMilk and high quality butter cream which give a rich cheesy flavour. ', 1, 'Singapore', 2, 3, NULL, 11, 1429337346, 1429682946);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (12, 'Yoghurt ', 6, 2, 0, 'Made with only milk, cream and cultures. ', 2, 'Hong Kong', 1, 3, NULL, 12, 1429510146, 1429855746);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (13, 'Milk ', 3, 4, 2, 'Made from 100% fresh milk.Pasteurised and homogenised.', 0, 'Hong Kong', 3, 3, NULL, 13, 1429682946, 1430028546);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (14, 'Whipping Cream', 9, 7, 5, 'It is a full bodied cream of 35,1% fat which makes it ideal for whipping.', 0, 'Phillippines', 1, 3, NULL, 14, 1429855746, 1430201346);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (15, 'Butter', 5, 4, 2, 'It has been gracing kitchens with the richest aroma and flavor, perfect for adding life to your favorite recipes. No. 1 in Singapore', 0, 'Hong Kong', 3, 3, NULL, 15, 1430028546, 1430374146);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (16, 'Egg', 3, 9, 7, '15 medium sized eggs, perfect for everyday use. Ours eggs are sourced from Malaysian farms of the highest standards.', 2, 'Japan', 1, 4, NULL, 16, 1430201346, 1430546946);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (17, 'Bread ', 3, 5, 3, 'Baked from high protein flour and enriched with vitamins and minerals. Especially popular with bigger families, it\'s high in vitamins B1, B2, B3, Calcium and Iron, and has no trans fat.', 0, 'Japan', 2, 4, NULL, 17, 1430374146, 1430719746);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (18, 'Water', 2, 7, 5, 'Natural mineral water collected from the Dewa Sanzan mountains', 2, 'Japan', 2, 4, NULL, 18, 1430546946, 1430892546);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (19, 'cereal ', 6, 5, 3, 'Post Selects Banana Nut Crunch Cereal is a delicious mix of real bananas baked into multi grain clusters, multi grain flakes, and specially selected walnuts. It\'s naturally flavored and provides 4 g of fiber and 40 g of whole grains per serving.', 1, 'South Korea', 0, 4, NULL, 19, 1430719746, 1431065346);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (20, 'orange juice', 3, 9, 7, 'Concentrated Low Calorie Orange and Mango Soft Drink with Sweeteners. Robinson Fruit Creations - At Robinsons we travel the world to seek out the very best flavours. ', 1, 'Hong Kong', 3, 4, NULL, 20, 1430892546, 1431238146);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (21, 'Sparkling juice', 2, 6, 4, 'Add some fun to your life with this 100-Percent non-alcoholic sparkling apple juice. Made from the best ingredients, best enjoyed when served chill for get-togethers.', 2, 'Phillippines', 3, 5, NULL, 21, 1431065346, 1431410946);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (22, 'red wine ', 325, 3, 1, 'Fresh red fruits, together with toasted notes highlighting coffee and chocolate. ', 0, 'Japan', 3, 5, NULL, 22, 1431238146, 1431583746);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (23, 'white wine ', 625, 6, 4, 'Irresistibly, refreshingly good. Bright aromatic herbal notes of this wine works well with some feta or goats cheese. ', 0, 'Maldives', 0, 5, NULL, 23, 1431410946, 1431756546);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (24, 'tiger beer', 60, 5, 3, '\nBrewed fresh in Singapore, for Singapore. Tropical lagered since 1932 for a full bodied yet refreshing taste.', 0, 'China', 1, 5, NULL, 24, 1431583746, 1431929346);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (25, 'champagne ', 123, 6, 4, 'Made from a blend of 80-percent Pinot Noir and 20-percent Chardonnay, this Champagne evokes the typical character of our vineyard by hillsides and the power of the Pinots Noirs.', 0, 'China', 2, 5, NULL, 25, 1431756546, 1432102146);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (26, 'croissant ', 3, 6, 4, 'Delicious butter croissants free from artificial colors and flavorings.', 0, 'United States', 2, 6, NULL, 26, 1431929346, 1432274946);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (27, 'danish ', 2, 1, 1, 'The crown shape Danish pastry with butter.', 1, 'France', 2, 6, NULL, 27, 1432102146, 1432447746);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (28, 'puff pastry ', 3, 9, 7, 'Pastry is so mouth wateringly delicious, light & flaky.', 2, 'France', 0, 6, NULL, 28, 1432274946, 1432620546);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (29, 'brownie ', 1, 7, 5, '\nIndividually packed ready-to-eat brownies. Vegetarian-Friendly with no alcoholic ingredients', 1, 'United States', 3, 6, NULL, 29, 1432447746, 1432793346);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (30, 'apple strudel ', 10, 6, 4, 'Classic Apple Pastry, layer of buttery puff pastry filled with the caramelised apple & custard lattice.', 2, 'United States', 2, 6, NULL, 30, 1432620546, 1432966146);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (31, 'wet wipes ', 62, 2, 0, 'Anti-Bacterial Wet Tissue effectively kills 99.99-percent of bacteria, keeping your baby safe.', 0, 'Hong Kong', 3, 7, NULL, 31, 1432793346, 1433138946);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (32, 'diapers ', 40, 7, 5, '\nOrganic cotton is contained in the surface sheets of the diapers, which makes babies comfortable as if in mom\'s hands! It has a soft touch to the skin, so no need to worry about delicate babies\' skin. 128 pcs (4 packs).', 2, 'France', 0, 7, NULL, 32, 1432966146, 1433311746);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (33, 'pacifier ', 15, 2, 0, 'An anatomically shaped mouth shield and motifs to make your baby smile.', 0, 'Isle of Man', 0, 7, NULL, 33, 1433138946, 1433484546);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (34, 'milk bottle ', 58, 3, 1, 'Super-sensitive, easi-vent valve eliminates excessive air flow', 2, 'China', 2, 7, NULL, 34, 1433311746, 1433657346);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (35, 'milk powder', 55, 6, 4, 'support cognitive functioning, strong bones and teeth ', 1, 'Indonesia', 0, 7, NULL, 35, 1433484546, 1433830146);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (36, 'protein shake ', 42, 10, 8, 'contains clean plant protein, as well as probiotics and digestive enzymes for easier absorption and utilisation', 1, 'China', 2, 8, NULL, 36, 1433657346, 1434002946);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (37, 'vitamin C', 23, 1, 1, 'protects cells from oxidative damage and increases iron absorption', 2, 'Singapore', 1, 8, NULL, 37, 1433830146, 1434175746);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (38, 'collagen drink ', 45, 7, 5, 'minimize the signs of ageing such as fine lines, wrinkles and dryness', 0, 'Italy', 3, 8, NULL, 38, 1434002946, 1434348546);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (39, 'blood pressure monitor ', 200, 1, 1, 'Accurate Blood Pressure and Heart Rate Monitoring', 0, 'Indonesia', 1, 8, NULL, 39, 1434175746, 1434521346);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (40, 'Oximeter', 15, 3, 1, 'a noninvasive method for monitoring a person\'s oxygen saturation', 0, 'Indonesia', 0, 8, NULL, 40, 1434348546, 1434694146);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (41, 'moisturizer', 16, 5, 3, 'This extra-rich body lotion formula, enriched with Certified Shea Butter, provides deep hydration', 2, 'Indonesia', 1, 9, NULL, 41, 1434521346, 1434866946);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (42, 'sun block ', 22, 1, 1, 'It is water resistant, sweat resistant, resists rub-off, non-comedogenic, dermatologist tested, oil free and PABA free', 0, 'Singapore', 1, 9, NULL, 42, 1434694146, 1435039746);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (43, 'toner', 19, 2, 0, 'gently removes impurities and minimizes appearance of acne and large pores', 0, 'Sri Lanka', 0, 9, NULL, 43, 1434866946, 1435212546);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (44, 'bb cream ', 25, 9, 7, 'Contains Hyaluronic Acid Complex, Natural Green Tea extract, prevents pollutants from touching skin directly', 0, 'China', 0, 9, NULL, 44, 1435039746, 1435385346);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (45, 'cc cream ', 21, 9, 7, 'Formulated with smart colour-match capsule, colour pigments encapsulated by powder burst out and mix during application that blend into skin colour', 0, 'Indonesia', 1, 9, NULL, 45, 1435212546, 1435558146);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (46, 'shampoo ', 15, 3, 1, 'Deeply cleanses hair roots, removing excess oil and dirt ', 0, 'China', 1, 10, NULL, 46, 1435385346, 1435730946);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (47, 'conditioner ', 8, 3, 1, 'Leaves hair feeling silky and smooth ', 2, 'Australia', 0, 10, NULL, 47, 1435558146, 1435903746);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (48, 'face wash ', 20, 8, 6, 'encourage the healing and detoxification of the skin', 2, 'Cambodia', 1, 10, NULL, 48, 1435730946, 1436076546);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (49, 'body wash ', 16, 2, 0, 'Creamy, bubbly and gentle on the skin', 0, 'China', 0, 10, NULL, 49, 1435903746, 1436249346);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_ctime`, `listing_mtime`) VALUES (50, 'hand soap ', 5, 7, 5, 'Rich-lathering and skin-loving cleanser ', 1, 'South Korea', 0, 10, NULL, 50, 1436076546, 1436422146);

COMMIT;


-- -----------------------------------------------------
-- Data for table `heroku_bdc39d4687a85d4`.`wallet_tab`
-- -----------------------------------------------------
START TRANSACTION;
USE `heroku_bdc39d4687a85d4`;
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1, 194, 2, 1353413607, 1502111537);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (2, 187, 0, 1363845694, 1493868046);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (3, 50, 1, 1395346315, 1467729411);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (4, 65, 2, 1395680122, 1437884362);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (5, 76, 0, 1378396589, 1497370444);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (6, 170, 0, 1369640673, 1459902832);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (7, 75, 0, 1385647904, 1476696831);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (8, 31, 1, 1359756657, 1500211567);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (9, 37, 2, 1355355602, 1493165913);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (10, 43, 2, 1338469780, 1474858048);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (11, 197, 2, 1387496378, 1481706083);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (12, 39, 1, 1339873150, 1477352589);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (13, 92, 2, 1374270735, 1472499732);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (14, 146, 2, 1338922388, 1496030633);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (15, 39, 1, 1365875221, 1466841380);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (16, 96, 2, 1390079873, 1446101680);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (17, 56, 0, 1340243545, 1439437083);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (18, 60, 2, 1389213951, 1486527833);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (19, 183, 0, 1372411264, 1449710191);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (20, 113, 2, 1357898998, 1475399811);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (21, 173, 1, 1339507346, 1470039273);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (22, 179, 2, 1338502551, 1465110100);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (23, 166, 1, 1354470691, 1440130257);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (24, 181, 2, 1393262462, 1444448314);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (25, 186, 2, 1341320159, 1476042717);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (26, 65, 1, 1346710308, 1441540940);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (27, 45, 0, 1355770084, 1465413220);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (28, 201, 2, 1336839810, 1476740355);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (29, 170, 1, 1368899449, 1465066604);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (30, 104, 2, 1364168131, 1475877566);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (31, 41, 2, 1341853741, 1489453644);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (32, 66, 2, 1363957775, 1439589374);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (33, 191, 1, 1369472881, 1479924202);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (34, 209, 0, 1388992562, 1448518136);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (35, 156, 0, 1341604263, 1473672965);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (36, 120, 0, 1346399688, 1439127306);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (37, 107, 0, 1391893807, 1435859547);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (38, 174, 2, 1396711850, 1452705216);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (39, 202, 1, 1335814372, 1462768505);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (40, 177, 2, 1388389870, 1444307362);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (41, 122, 1, 1374584612, 1494312784);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (42, 205, 0, 1336313317, 1463791703);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (43, 16, 0, 1362213031, 1444027958);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (44, 154, 1, 1371498875, 1471654257);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (45, 115, 1, 1347814848, 1449224149);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (46, 15, 0, 1368486447, 1438073486);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (47, 146, 2, 1372439255, 1435799504);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (48, 197, 1, 1341556954, 1478630859);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (49, 47, 1, 1400344926, 1466246931);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_tab` (`w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (50, 107, 0, 1398865821, 1435683953);

COMMIT;


-- -----------------------------------------------------
-- Data for table `heroku_bdc39d4687a85d4`.`wallet_transaction`
-- -----------------------------------------------------
START TRANSACTION;
USE `heroku_bdc39d4687a85d4`;
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (1, 1, 'SINGLE ITEM', 1494784934, 1535981853, 83, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (2, 2, 'SINGLE ITEM', 1449899570, 1535779348, 97, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (3, 3, 'SINGLE ITEM', 1474708730, 1536232801, 92, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (4, 4, 'SINGLE ITEM', 1480991474, 1535444828, 11, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (5, 5, 'MULTIPLE ITEMS', 1499593663, 1535441812, 103, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (6, 6, 'MULTIPLE ITEMS', 1477800415, 1535418025, 77, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (7, 7, 'SINGLE ITEM', 1451787412, 1536326906, 30, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (8, 8, 'SINGLE ITEM', 1470294395, 1535291393, 27, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (9, 9, 'MULTIPLE ITEMS', 1492812625, 1535550906, 97, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (10, 10, 'MULTIPLE ITEMS', 1481713513, 1536362467, 108, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (11, 11, 'MULTIPLE ITEMS', 1444520329, 1535730556, 41, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (12, 12, 'SINGLE ITEM', 1471184882, 1535526555, 80, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (13, 13, 'SINGLE ITEM', 1442840927, 1535911062, 19, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (14, 14, 'SINGLE ITEM', 1442423898, 1535717462, 79, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (15, 15, 'SINGLE ITEM', 1464702360, 1535433482, 41, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (16, 16, 'SINGLE ITEM', 1493256098, 1536253482, 98, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (17, 17, 'MULTIPLE ITEMS', 1471742172, 1536098875, 56, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (18, 18, 'MULTIPLE ITEMS', 1465835161, 1536097489, 65, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (19, 19, 'MULTIPLE ITEMS', 1447216345, 1536146045, 35, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (20, 20, 'MULTIPLE ITEMS', 1487044721, 1536036537, 92, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (21, 21, 'MULTIPLE ITEMS', 1502378273, 1535461825, 51, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (22, 22, 'MULTIPLE ITEMS', 1451078646, 1535264083, 50, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (23, 23, 'MULTIPLE ITEMS', 1466992394, 1536089692, 13, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (24, 24, 'SINGLE ITEM', 1440273689, 1536201496, 33, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (25, 25, 'MULTIPLE ITEMS', 1477958401, 1535848605, 79, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (26, 26, 'MULTIPLE ITEMS', 1444716024, 1536151258, 87, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (27, 27, 'SINGLE ITEM', 1488718948, 1535969201, 34, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (28, 28, 'SINGLE ITEM', 1491822428, 1535298310, 21, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (29, 29, 'MULTIPLE ITEMS', 1440158002, 1535500901, 13, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (30, 30, 'SINGLE ITEM', 1454213956, 1536313176, 25, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (31, 31, 'SINGLE ITEM', 1495192998, 1535650706, 108, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (32, 32, 'SINGLE ITEM', 1454955222, 1535819532, 50, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (33, 33, 'MULTIPLE ITEMS', 1499307751, 1535520891, 35, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (34, 34, 'MULTIPLE ITEMS', 1464082282, 1535822961, 95, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (35, 35, 'SINGLE ITEM', 1495585080, 1536201447, 132, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (36, 36, 'SINGLE ITEM', 1502402782, 1536082297, 79, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (37, 37, 'MULTIPLE ITEMS', 1496660805, 1536032312, 58, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (38, 38, 'MULTIPLE ITEMS', 1438481399, 1535789452, 103, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (39, 39, 'MULTIPLE ITEMS', 1470521022, 1535626653, 88, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (40, 40, 'MULTIPLE ITEMS', 1502603502, 1536355673, 79, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (41, 41, 'MULTIPLE ITEMS', 1452648390, 1535780454, 34, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (42, 42, 'SINGLE ITEM', 1475489311, 1535347064, 80, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (43, 43, 'MULTIPLE ITEMS', 1461887245, 1535841129, 33, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (44, 44, 'SINGLE ITEM', 1471626738, 1535268086, 100, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (45, 45, 'MULTIPLE ITEMS', 1496499534, 1536147018, 14, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (46, 46, 'SINGLE ITEM', 1492344891, 1536151150, 21, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (47, 47, 'MULTIPLE ITEMS', 1490489834, 1536310957, 85, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (48, 48, 'SINGLE ITEM', 1447933733, 1535853889, 46, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (49, 49, 'SINGLE ITEM', 1474109263, 1535574505, 132, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`wallet_transaction` (`wt_transaction_id`, `wt_user_id`, `transaction_details`, `transaction_ctime`, `transaction_mtime`, `transaction_amount`, `transaction_status`) VALUES (50, 50, 'SINGLE ITEM', 1460720468, 1535694590, 44, 0);

COMMIT;


-- -----------------------------------------------------
-- Data for table `heroku_bdc39d4687a85d4`.`listing_transactions_tab`
-- -----------------------------------------------------
START TRANSACTION;
USE `heroku_bdc39d4687a85d4`;
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (1, 1, 1, 1436482377, 1436655177, 422, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (2, 2, 2, 1436655177, 1436827977, 720, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (3, 3, 3, 1436827977, 1437000777, 933, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (4, 4, 4, 1437000777, 1437173577, 845, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (5, 5, 5, 1437173577, 1437346377, 483, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (6, 6, 6, 1437346377, 1437519177, 190, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (7, 7, 7, 1437519177, 1437691977, 386, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (8, 8, 8, 1437691977, 1437864777, 778, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (9, 9, 9, 1437864777, 1438037577, 923, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (10, 10, 10, 1438037577, 1438210377, 706, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (11, 11, 11, 1438210377, 1438383177, 854, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (12, 12, 12, 1438383177, 1438555977, 798, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (13, 13, 13, 1438555977, 1438728777, 298, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (14, 14, 14, 1438728777, 1438901577, 49, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (15, 15, 15, 1438901577, 1439074377, 349, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (16, 16, 16, 1439074377, 1439247177, 590, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (17, 17, 17, 1439247177, 1439419977, 496, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (18, 18, 18, 1439419977, 1439592777, 401, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (19, 19, 19, 1439592777, 1439765577, 453, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (20, 20, 20, 1439765577, 1439938377, 567, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (21, 21, 21, 1439938377, 1440111177, 322, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (22, 22, 22, 1440111177, 1440283977, 504, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (23, 23, 23, 1440283977, 1440456777, 839, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (24, 24, 24, 1440456777, 1440629577, 363, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (25, 25, 25, 1440629577, 1440802377, 241, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (26, 26, 26, 1440802377, 1440975177, 514, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (27, 27, 27, 1440975177, 1441147977, 644, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (28, 28, 28, 1441147977, 1441320777, 772, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (29, 29, 29, 1441320777, 1441493577, 591, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (30, 30, 30, 1441493577, 1441666377, 293, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (31, 31, 31, 1441666377, 1441839177, 152, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (32, 32, 32, 1441839177, 1442011977, 198, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (33, 33, 33, 1442011977, 1442184777, 835, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (34, 34, 34, 1442184777, 1442357577, 18, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (35, 35, 35, 1442357577, 1442530377, 119, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (36, 36, 36, 1442530377, 1442703177, 489, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (37, 37, 37, 1442703177, 1442875977, 180, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (38, 38, 38, 1442875977, 1443048777, 847, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (39, 39, 39, 1443048777, 1443221577, 859, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (40, 40, 40, 1443221577, 1443394377, 167, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (41, 41, 41, 1443394377, 1443567177, 941, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (42, 42, 42, 1443567177, 1443739977, 874, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (43, 43, 43, 1443739977, 1443912777, 900, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (44, 44, 44, 1443912777, 1444085577, 180, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (45, 45, 45, 1444085577, 1444258377, 625, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (46, 46, 46, 1444258377, 1444431177, 757, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (47, 47, 47, 1444431177, 1444603977, 624, 2);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (48, 48, 48, 1444603977, 1444776777, 724, 1);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (49, 49, 49, 1444776777, 1444949577, 559, 0);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_transactions_tab` (`lt_transaction_id`, `lt_item_id`, `lt_user_id`, `transaction_ctime`, `transaction_mtime`, `transaction_price`, `transaction_status`) VALUES (50, 50, 50, 1444781777, 1445449577, 218, 2);

COMMIT;


-- -----------------------------------------------------
-- Data for table `heroku_bdc39d4687a85d4`.`acc_credentials_tab`
-- -----------------------------------------------------
START TRANSACTION;
USE `heroku_bdc39d4687a85d4`;
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (1, 'zPhpJB1ZVV', 1, 'you');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (2, 'QtreljWqjC', 1, 'me');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (3, 'TlmT6ePmoP', 2, '??????');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (4, 'i0D6n8iJUT', 3, 'bleh');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (5, '4TuBvqDj0o', 2, '???');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (6, 'lAMS1VMzfE', 4, '??');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (7, 'bzjlf8t85p', 4, 'confused');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (8, 'sfb!@f..', 1, 'him');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (9, 'iisdfs24350', 1, 'her');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (10, 'sfdg522rsd2', 1, 'them');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (11, 'asd1f2fqwe4f21', 4, 'random');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (12, 'adsf212adf54', 4, 'random');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (13, '1=09-=-gj4\'', 4, 'random');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (14, 'dfg894;\'sv', 4, 'hi');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (15, 'sb\';rger', 1, 'meh');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (16, 'f', 3, 'bye');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (17, 'vvcv', 1, 'confused');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (18, '8rPlpER8Jm', 4, 'bleh');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (19, 'FosbRg4YHM', 3, 'bleh');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (20, 'FYmL06mcDW', 4, 'confused');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (21, 'w0rLebQkf3', 1, 'bleh');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (22, '5GaTFGFHtA', 2, 'what do you mean?');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (23, 'CJ8s0lr1G2', 4, 'booo');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (24, 'UaImvi4yQG', 2, 'what do you mean?');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (25, 'iKcp78yZf0', 4, 'booo');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (26, 'ZCJr78GRev', 1, '!?!?!?');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (27, 'uhnuXTamND', 4, 'what do you mean?');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (28, 'hbPXwi7CAD', 2, 'nani');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (29, 'r7ezyEyZGY', 1, 'what do you mean?');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (30, 'oG5c1w9MfQ', 4, 'nani');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (31, 'DRIXsZxfPo', 1, '!???!?!?!?!?!');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (32, 'c6SrIN0gHF', 2, 'booo');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (33, 'ULITfMVjfI', 1, 'haiz');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (34, 'Bq7SKLKXUZ', 4, 'hey');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (35, 'TPonn0eKrH', 2, 'see');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (36, 'xZsn31Mpw2', 4, 'me');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (37, 'M9aiCn93Aq', 1, 'yo');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (38, 'RcolriDy7q', 1, 'yo');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (39, 'pNikMM69wn', 3, 'yo');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (40, 'RE9SbXgIrn', 3, 'watch me');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (41, 'nLAcGag4ld', 4, 'watch me');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (42, 'VreWX3GHlk', 4, 'watch');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (43, 'tV4gF2fUju', 2, 'cmon');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (44, 'DGDeOkP99E', 1, 'you');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (45, 'H0a1reavyM', 1, 'can');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (46, 'yiezLLLIlT', 2, 'do');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (47, 'tj5J3pyWdl', 3, 'it');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (48, 'NbMJO3FPJs', 3, ' cmon');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (49, 'TqxuFlFWYh', 3, 'go');
INSERT INTO `heroku_bdc39d4687a85d4`.`acc_credentials_tab` (`c_user_id`, `user_password`, `user_security_question`, `user_security_answer`) VALUES (50, 'SIjOFnF9Td', 1, 'go');

COMMIT;


-- -----------------------------------------------------
-- Data for table `heroku_bdc39d4687a85d4`.`user_review_tab`
-- -----------------------------------------------------
START TRANSACTION;
USE `heroku_bdc39d4687a85d4`;
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (1, 1, 1, 5, 'This', 1463790569);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (2, 2, 2, 5, 'maybe', 1466947661);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (3, 3, 3, 2, 'i', 1467812688);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (4, 4, 4, 3, 'can', 1463790569);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (5, 5, 5, 2, 'only', 1466947661);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (6, 6, 6, 4, 'so', 1458247651);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (7, 7, 7, 4, '...', 1453535532);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (8, 8, 8, 5, '..', 1464286205);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (9, 9, 9, 5, 'maybe', 1442625424);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (10, 10, 10, 5, '...', 1443421873);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (11, 11, 11, 4, '....', 1458247651);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (12, 12, 12, 4, '.....', 1414744275);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (13, 13, 13, 4, '......', 1456406392);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (14, 14, 14, 4, '.....', 1474924465);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (15, 15, 15, 1, '....', 1414744275);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (16, 16, 16, 3, 'how long', 1458247651);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (17, 17, 17, 1, 'will', 1453535532);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (18, 18, 18, 4, 'is', 1464286205);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (19, 19, 19, 3, '@$#%&!', 1420871670);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (20, 20, 20, 4, 'like', 1467896775);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (21, 21, 21, 1, 'taking', 1448562440);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (22, 22, 22, 2, 'complain', 1485033404);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (23, 23, 23, 4, 'here', 1448562440);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (24, 24, 24, 2, '!', 1497358109);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (25, 25, 25, 4, '!!!', 1467896775);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (26, 26, 26, 1, 'dear', 1485033404);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (27, 27, 27, 4, '<3333', 1485033404);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (28, 28, 28, 2, 't', 1439107621);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (29, 29, 29, 5, 'i', 1437834963);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (30, 30, 30, 4, '6', 1444428381);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (31, 31, 31, 5, '0', 1462454583);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (32, 32, 32, 2, 'misery', 1453535532);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (33, 33, 33, 1, '&&&&&', 1443421873);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (34, 34, 34, 4, '&&', 1458247651);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (35, 35, 35, 2, '&', 1456406392);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (36, 36, 36, 4, 'i', 1474924465);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (37, 37, 37, 5, 'maybe', 1464926820);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (38, 38, 38, 5, 'i', 1489345471);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (39, 39, 39, 3, 'can', 1475419828);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (40, 40, 40, 3, 'c', 1446232148);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (41, 41, 41, 4, '2', 1452091561);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (42, 42, 42, 4, 'dear', 1432828484);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (43, 43, 43, 2, 'this', 1418931454);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (44, 44, 44, 5, '1', 1445714626);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (45, 45, 45, 5, '----', 1455644622);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (46, 46, 46, 2, '--', 1420871670);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (47, 47, 47, 3, 'lord', 1414072452);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (48, 48, 48, 3, 'save', 1442248046);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (49, 49, 49, 3, 'me', 1497358109);
INSERT INTO `heroku_bdc39d4687a85d4`.`user_review_tab` (`review_id`, `rv_user_id`, `rv_seller_id`, `ratings`, `review_text`, `ctime`) VALUES (50, 50, 50, 5, 'end', 1444428381);

COMMIT;


-- -----------------------------------------------------
-- Data for table `heroku_bdc39d4687a85d4`.`listing_reactions_tab`
-- -----------------------------------------------------
START TRANSACTION;
USE `heroku_bdc39d4687a85d4`;
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (1, 1, 1, 1, 'dislike', 1563790569);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (2, 2, 2, 0, 'like', 1566947661);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (3, 3, 3, 0, 'like', 1567812688);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (4, 4, 4, 1, 'dislike', 1563790569);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (5, 5, 5, 0, 'like', 1566947661);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (6, 6, 6, 1, 'dislike', 1558247651);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (7, 7, 7, 0, 'like', 1553535532);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (8, 8, 8, 1, 'dislike', 1564286205);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (9, 9, 9, 0, 'like', 1542625424);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (10, 10, 10, 0, 'like', 1543421873);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (11, 11, 11, 1, 'dislike', 1558247651);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (12, 12, 12, 1, 'dislike', 1515744275);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (13, 13, 13, 0, 'like', 1556406392);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (14, 14, 14, 0, 'like', 1574924465);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (15, 15, 15, 0, 'like', 1515744275);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (16, 16, 16, 1, 'dislike', 1558247651);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (17, 17, 17, 0, 'like', 1553535532);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (18, 18, 18, 0, 'like', 1564286205);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (19, 19, 19, 0, 'like', 1520871670);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (20, 20, 20, 1, 'dislike', 1567896775);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (21, 21, 21, 0, 'like', 1548562440);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (22, 22, 22, 0, 'like', 1585033404);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (23, 23, 23, 0, 'like', 1548562440);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (24, 24, 24, 1, 'dislike', 1597358109);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (25, 25, 25, 1, 'dislike', 1567896775);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (26, 26, 26, 1, 'dislike', 1585033404);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (27, 27, 27, 1, 'dislike', 1585033404);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (28, 28, 28, 0, 'like', 1539107621);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (29, 29, 29, 1, 'dislike', 1537834963);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (30, 30, 30, 1, 'dislike', 1544428381);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (31, 31, 31, 0, 'like', 1562454583);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (32, 32, 32, 0, 'like', 1553535532);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (33, 33, 33, 0, 'like', 1543421873);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (34, 34, 34, 0, 'like', 1558247651);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (35, 35, 35, 0, 'like', 1556406392);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (36, 36, 36, 0, 'like', 1574924465);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (37, 37, 37, 1, 'dislike', 1564926820);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (38, 38, 38, 1, 'dislike', 1589345471);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (39, 39, 39, 1, 'dislike', 1575419828);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (40, 40, 40, 0, 'like', 1546232158);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (41, 41, 41, 0, 'like', 1552091561);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (42, 42, 42, 1, 'dislike', 1532828484);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (43, 43, 43, 1, 'dislike', 1518931554);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (44, 44, 44, 1, 'dislike', 1545715626);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (45, 45, 45, 1, 'dislike', 1555644622);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (46, 46, 46, 0, 'like', 1520871670);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (47, 47, 47, 0, 'like', 1515072452);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (48, 48, 48, 0, 'like', 1542248046);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (49, 49, 49, 1, 'dislike', 1597358109);
INSERT INTO `heroku_bdc39d4687a85d4`.`listing_reactions_tab` (`reactions_id`, `rt_user_id`, `rt_item_id`, `reaction_type`, `comment`, `ctime`) VALUES (50, 50, 50, 0, 'like', 1544428381);

COMMIT;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
