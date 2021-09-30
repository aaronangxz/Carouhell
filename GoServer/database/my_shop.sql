-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema my_shop
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `my_shop` ;

-- -----------------------------------------------------
-- Schema my_shop
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `my_shop` DEFAULT CHARACTER SET utf8 ;
USE `my_shop` ;

-- -----------------------------------------------------
-- Table `my_shop`.`acc_tab`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `my_shop`.`acc_tab` ;

CREATE TABLE IF NOT EXISTS `my_shop`.`acc_tab` (
  `a_user_id` INT(10) UNSIGNED NOT NULL,
  `user_name` VARCHAR(45) NOT NULL,
  `user_email` VARCHAR(80) NOT NULL,
  `user_creationdate` INT(11) NOT NULL,
  `user_status` TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0 - Offline\n1 - Online',
  `user_type` VARCHAR(6) NOT NULL COMMENT '\'Buyer || Seller\'',
  `user_image` BLOB NULL DEFAULT NULL,
  `user_lastlogin` INT(11) NULL DEFAULT NULL,
  `user_rating` INT(2) UNSIGNED NULL DEFAULT 0,
  PRIMARY KEY (`a_user_id`),
  UNIQUE INDEX `user_email_UNIQUE` (`user_email` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `my_shop`.`notification_tab`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `my_shop`.`notification_tab` ;

CREATE TABLE IF NOT EXISTS `my_shop`.`notification_tab` (
  `n_notification_id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `n_user_id` INT(10) UNSIGNED NOT NULL,
  `notification_text` VARCHAR(256) NOT NULL,
  `notification_url` VARCHAR(256) NOT NULL,
  `notification_time` DATETIME NOT NULL,
  `notification_ui_info` BLOB NOT NULL,
  PRIMARY KEY (`n_notification_id`, `n_user_id`),
  INDEX `a_user_id, n_user_id_idx` (`n_user_id` ASC) VISIBLE,
  CONSTRAINT `a_user_id, n_user_id`
    FOREIGN KEY (`n_user_id`)
    REFERENCES `my_shop`.`acc_tab` (`a_user_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `my_shop`.`listing_tab`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `my_shop`.`listing_tab` ;

CREATE TABLE IF NOT EXISTS `my_shop`.`listing_tab` (
  `l_item_id` INT(10) UNSIGNED NOT NULL,
  `item_name` VARCHAR(256) NOT NULL,
  `item_price` DOUBLE(10,2) UNSIGNED NOT NULL DEFAULT 0.00,
  `item_quantity` SMALLINT(3) UNSIGNED NOT NULL DEFAULT 1,
  `item_purchasedquantity` SMALLINT(3) UNSIGNED NULL DEFAULT 0,
  `item_description` VARCHAR(256) NOT NULL,
  `item_shippinginfo` TINYINT(1) NULL DEFAULT 0,
  `item_paymentinfo` TINYINT(1) NULL DEFAULT 0,
  `item_location` VARCHAR(20) NOT NULL,
  `item_status` VARCHAR(15) NOT NULL DEFAULT 0 COMMENT '\'Available | Unavailable',
  `item_category` VARCHAR(30) NOT NULL,
  `item_image` BLOB NULL DEFAULT NULL,
  `l_seller_id` INT(10) UNSIGNED NOT NULL,
  `listing_date` INT(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`l_item_id`),
  INDEX `user_id_idx` (`l_seller_id` ASC) VISIBLE,
  CONSTRAINT `a_user_id,l_seller_id`
    FOREIGN KEY (`l_seller_id`)
    REFERENCES `my_shop`.`acc_tab` (`a_user_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `my_shop`.`wallet_tab`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `my_shop`.`wallet_tab` ;

CREATE TABLE IF NOT EXISTS `my_shop`.`wallet_tab` (
  `w_wallet_id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `w_user_id` INT(10) UNSIGNED NOT NULL,
  `wallet_balance` DOUBLE(10,2) UNSIGNED NOT NULL DEFAULT 0.00,
  `wallet_status` VARCHAR(10) NULL,
  `last_topup` INT(11) NULL,
  `last_used` INT(11) NULL,
  PRIMARY KEY (`w_wallet_id`, `w_user_id`),
  INDEX `a_user_id, w_user_id_idx` (`w_user_id` ASC) VISIBLE,
  CONSTRAINT `a_user_id, w_user_id`
    FOREIGN KEY (`w_user_id`)
    REFERENCES `my_shop`.`acc_tab` (`a_user_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `my_shop`.`wallet_transaction`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `my_shop`.`wallet_transaction` ;

CREATE TABLE IF NOT EXISTS `my_shop`.`wallet_transaction` (
  `wt_transaction_id` INT(10) UNSIGNED NULL AUTO_INCREMENT,
  `wt_wallet_id` INT(10) UNSIGNED NOT NULL,
  `transaction_details` VARCHAR(256) NULL DEFAULT NULL,
  `transaction_value` DOUBLE(10,2) UNSIGNED NOT NULL DEFAULT 0.00,
  `transaction_status` VARCHAR(10) NULL DEFAULT 'Failed' COMMENT '\n',
  PRIMARY KEY (`wt_wallet_id`),
  UNIQUE INDEX `transaction_id_UNIQUE` (`wt_transaction_id` ASC) VISIBLE,
  CONSTRAINT `a_user_id, wt_user_id`
    FOREIGN KEY (`wt_wallet_id`)
    REFERENCES `my_shop`.`acc_tab` (`a_user_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `my_shop`.`listing_transactions_tab`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `my_shop`.`listing_transactions_tab` ;

CREATE TABLE IF NOT EXISTS `my_shop`.`listing_transactions_tab` (
  `lt_transaction_id` INT(10) UNSIGNED NULL AUTO_INCREMENT,
  `lt_item_id` INT(10) UNSIGNED NOT NULL,
  `lt_user_id` INT(10) UNSIGNED NOT NULL,
  `transaction_time` INT(11) NULL,
  `transaction_price` DOUBLE(10,2) NULL DEFAULT NULL COMMENT 'transaction_price = item_price',
  `transaction_status` VARCHAR(10) NULL DEFAULT 'Failed',
  PRIMARY KEY (`lt_item_id`, `lt_user_id`),
  UNIQUE INDEX `transaction_id_UNIQUE` (`lt_transaction_id` ASC) VISIBLE,
  INDEX `a_user_id,lt_user_id_idx` (`lt_user_id` ASC) VISIBLE,
  CONSTRAINT `l_item_id,lt_item_id`
    FOREIGN KEY (`lt_item_id`)
    REFERENCES `my_shop`.`listing_tab` (`l_item_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `a_user_id,lt_user_id`
    FOREIGN KEY (`lt_user_id`)
    REFERENCES `my_shop`.`acc_tab` (`a_user_id`)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Data for table `my_shop`.`acc_tab`
-- -----------------------------------------------------
START TRANSACTION;
USE `my_shop`;
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000002, 'Jeanne Madrigal', 'jeanne_m@gmail.com', 2019-12-03, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000003, 'Cristin Allums', 'cristin_allums@gmail.com', 2019-12-03, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000004, 'Marlene Mctaggart', 'marlene_mctaggart@gmail.com', 2020-10-22, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000005, 'Larisa Komar', 'larisa_komar@gmail.com', 2019-12-03, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000006, 'Jackie Hayton', 'jackie_hayton@gmail.com', 2019-12-03, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000007, 'Shella Belote', 'shella_belote@gmail.com', 2020-07-17, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000008, 'Antonetta Concepcion', 'antonetta_concepcion@gmail.com', 2019-12-05, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000009, 'Winnie Malia', 'winnie_malia@gmail.com', 2019-12-15, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000010, 'Roman Bush', 'roman_bush@gmail.com', 2019-08-03, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000011, 'Toya Guerette', 'toya_guerette@gmail.com', 2020-08-18, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000012, 'Pedro Routt', 'pedro_routt@gmail.com', 2021-06-12, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000013, 'Arnette Mcmurray', 'arnette_mcmurray@gmail.com', 2020-10-22, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000014, 'Arletta Winburn', 'arletta_winburn@gmail.com', 2021-01-12, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000015, 'Stephen Ebert', 'stephen_ebert@gmail.com', 2021-03-12, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000016, 'Shantay Coster', 'shantay_coster@gmail.com', 2020-09-25, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000017, 'Noelle Kehrer', 'noelle_kehrer@gmail.com', 2020-10-22, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000018, 'Julienne Abston', 'julienne_abston@gmail.com', 2021-07-18, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000019, 'Helaine Tilson', 'helaine_tilson@gmail.com', 2019-11-06, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000020, 'Penni Printup', 'penni_printup@gmail.com', 2020-04-12, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000021, 'Katherina Marasco', 'katherina_marasco@gmail.com', 2020-01-01, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000022, 'Pennie Mcginley', 'pennie_mcginley@hotmail.com', 2020-10-16, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000023, 'Faustina Knouse', 'fautina_knouse@hotmail.com', 2020-04-05, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000024, 'Selma Allsop', 'selma_allsop@hotmail.com', 2020-03-09, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000025, 'Audrea Carraway', 'audrea_carraway@hotmail.com', 2020-01-30, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000026, 'Nadia Collington', 'nadia_collington@hotmail.com', 2021-02-03, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000027, 'Joette Baily', 'joette_baily@hotmail.com', 2020-10-22, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000028, 'Maryetta Poppell', 'maryetta_poppell@hotmail.com', 2019-11-25, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000029, 'Teodoro Ceniceros', 'teodoro_ceniceros@hotmail.com', 2020-07-07, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000030, 'Ilda Medders', 'ilda_medders@hotmail.com', 2020-09-11, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000031, 'Judith Police', 'judith_police@hotmail.com', 2020-05-30, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000032, 'Charley Vanfleet', 'charley_vanfleet@hotmail.com', 2019-08-09, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000033, 'Cheryl Parrilla', 'cheryl_parrila@hotmail.com', 2021-03-01, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000034, 'Eveline Weathersby', 'eveline_weathersby@hotmail.com', 2020-11-11, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000035, 'Roselee Monfort', 'roselee_monfort@hotmail.com', 2020-02-26, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000036, 'Gerry Maudlin', 'gerry_maudlin@hotmail.com', 2020-03-03, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000037, 'Kiesha Carone', 'kiesha_carone@yahoo.com', 2020-04-05, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000038, 'Audra Ewert', 'audra_ewert@yahoo.com', 2020-04-05, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000039, 'Anja Ranieri', 'anja_ranieri@yahoo.com', 2020-05-30, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000040, 'Letitia Delatorre', 'letitia_delatorre@yahoo.com', 2020-02-26, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000041, 'Lyndia Ali', 'lyndia_ali@yahoo.com', 2020-07-07, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000042, 'Elisa Meeker', 'elisa_meeker@yahoo.com', 2020-10-22, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000043, 'Keva Sauseda', 'keva_sauseda@yahoo.com', 2021-03-01, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000044, 'Alycia Lamontagne', 'alycia_lamontagne@yahoo.com', 2019-10-22, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000045, 'Miguelina Mast', 'miguelina_mast@yahoo.com', 2019-11-22, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000046, 'Orville Mccaskey', 'orville_mccaskey@yahoo.com', 2019-11-29, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000047, 'Mabel Ascencio', 'mabel_ascencio@yahoo.com', 2020-12-31, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000048, 'Jenise Radtke', 'jenise_radtke@yahoo.com', 2020-12-01, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000049, 'Reyna Haymond', 'reyna_haymond@yahoo.com', 2020-01-19, 0, 'Seller', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000050, 'Masako Leitzel', 'masako_leitzel@yahoo.com', 2021-05-01, 0, 'Buyer', NULL, NULL, NULL);
INSERT INTO `my_shop`.`acc_tab` (`a_user_id`, `user_name`, `user_email`, `user_creationdate`, `user_status`, `user_type`, `user_image`, `user_lastlogin`, `user_rating`) VALUES (1000000001, 'Gus Amaral', 'gus_amaral@gmail.com', 2020-10-22, 0, 'Buyer', NULL, NULL, NULL);

COMMIT;


-- -----------------------------------------------------
-- Data for table `my_shop`.`listing_tab`
-- -----------------------------------------------------
START TRANSACTION;
USE `my_shop`;
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000001, 'Apple', 2.5, 8, NULL, 'Their dense flesh is creamy yellow and crisp, offering a mildly sweet flavor.', NULL, NULL, 'Malaysia', 'Unavailable', 'Fruits & Vegetables', NULL, 1000000002, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000002, 'Orange', 3, 5, NULL, 'Delicious and juicy orange fruit, good for juicing and containing an impressive list of essential nutrients, vitamins, minerals for normal growth and development and overall well-being.', NULL, NULL, 'Taiwan', 'Available', 'Fruits & Vegetables', NULL, 1000000004, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000003, 'Pear', 3.5, 8, NULL, 'Packham pears have a wide-bottomed shape and a smooth green skin that ripens to yellow.', NULL, NULL, 'Australia', 'Available', 'Fruits & Vegetables', NULL, 1000000005, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000004, 'Blueberry', 4.2, 7, NULL, 'Plump, juicy, and sweet, with vibrant colours ranging from deep purple-blue to blue-black and highlighted by a silvery sheen called a bloom, blueberries are one of nature\'s great treasures.', NULL, NULL, 'Japan', 'Unavailable', 'Fruits & Vegetables', NULL, 1000000007, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000005, 'Watermelon', 6.8, 1, NULL, 'The large round fruit has a hard green rind, a watery red pulp and small brown seeds.', NULL, NULL, 'Taiwan', 'Available', 'Fruits & Vegetables', NULL, 1000000009, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000006, 'Chicken ', 10.9, 4, NULL, 'The chickens not kept in cages but are instead raised in a modernised, temperature-controlled and environmentally friendly barn which gives the chickens ample space to roam with access to food and water.', NULL, NULL, 'South Korea', 'Available', 'Meat & Seafood', NULL, 1000000010, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000007, 'Pork ', 9.8, 1, NULL, 'Pork Loin Boneless trimmed & ready to used, no added artificial coloring & flavours. it is portion to approx 125g each and consist 4 pcs Individually freeze, approx packed 500-550g.', NULL, NULL, 'Phillippines', 'Unavailable', 'Meat & Seafood', NULL, 1000000012, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000008, 'Beef ', 15.5, 10, NULL, 'Black Angus 150days Grain fed Beef Tenderloin with 100% natural, no added preservative, coloring & additives.', NULL, NULL, 'Phillippines', 'Available', 'Meat & Seafood', NULL, 1000000013, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000009, 'Fish ', 12.2, 9, NULL, 'Daily processed fresh Batang fillet', NULL, NULL, 'South Korea', 'Available', 'Meat & Seafood', NULL, 1000000014, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000010, 'Duck ', 9.5, 2, NULL, ' Processed locally to ensure low carbon footprint, all products are thermal-vacuum sealed for superior product quality.', NULL, NULL, 'Colombia', 'Available', 'Meat & Seafood', NULL, 1000000016, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000011, 'Cheese', 12.8, 7, NULL, 'Cheese slices make from FreshMilk and high quality butter cream which give a rich cheesy flavour. ', NULL, NULL, 'Singapore', 'Unavailable', 'Dairy & Chilled', NULL, 1000000017, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000012, 'Yoghurt ', 6.8, 2, NULL, 'Made with only milk, cream and cultures. ', NULL, NULL, 'Hong Kong', 'Available', 'Dairy & Chilled', NULL, 1000000019, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000013, 'Milk ', 3.6, 4, NULL, 'Made from 100% fresh milk.Pasteurised and homogenised.', NULL, NULL, 'Hong Kong', 'Available', 'Dairy & Chilled', NULL, 1000000031, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000014, 'Whipping Cream', 9.9, 7, NULL, 'It is a full bodied cream of 35,1% fat which makes it ideal for whipping.', NULL, NULL, 'Phillippines', 'Unavailable', 'Dairy & Chilled', NULL, 1000000033, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000015, 'Butter', 5.6, 4, NULL, 'It has been gracing kitchens with the richest aroma and flavor, perfect for adding life to your favorite recipes. No. 1 in Singapore', NULL, NULL, 'Hong Kong', 'Available', 'Dairy & Chilled', NULL, 1000000034, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000016, 'Egg', 3.7, 9, NULL, '15 medium sized eggs, perfect for everyday use. Ours eggs are sourced from Malaysian farms of the highest standards.', NULL, NULL, 'Japan', 'Unavailable', 'Breakfast', NULL, 1000000035, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000017, 'Bread ', 3.5, 5, NULL, 'Baked from high protein flour and enriched with vitamins and minerals. Especially popular with bigger families, it\'s high in vitamins B1, B2, B3, Calcium and Iron, and has no trans fat.', NULL, NULL, 'Japan', 'Available', 'Breakfast', NULL, 1000000037, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000018, 'Water', 2.1, 7, NULL, 'Natural mineral water collected from the Dewa Sanzan mountains', NULL, NULL, 'Japan', 'Unavailable', 'Breakfast', NULL, 1000000038, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000019, 'cereal ', 6.5, 5, NULL, 'Post Selects Banana Nut Crunch Cereal is a delicious mix of real bananas baked into multi grain clusters, multi grain flakes, and specially selected walnuts. It\'s naturally flavored and provides 4 g of fiber and 40 g of whole grains per serving.', NULL, NULL, 'South Korea', 'Available', 'Breakfast', NULL, 1000000040, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000020, 'orange juice', 3.7, 9, NULL, 'Concentrated Low Calorie Orange and Mango Soft Drink with Sweeteners. Robinson Fruit Creations - At Robinsons we travel the world to seek out the very best flavours. ', NULL, NULL, 'Hong Kong', 'Available', 'Breakfast', NULL, 1000000043, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000021, 'Sparkling juice', 2.5, 6, NULL, 'Add some fun to your life with this 100-Percent non-alcoholic sparkling apple juice. Made from the best ingredients, best enjoyed when served chill for get-togethers.', NULL, NULL, 'Phillippines', 'Unavailable', 'Beer, Wine & Spirits', NULL, 1000000045, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000022, 'red wine ', 325, 3, NULL, 'Fresh red fruits, together with toasted notes highlighting coffee and chocolate. ', NULL, NULL, 'Japan', 'Unavailable', 'Beer, Wine & Spirits', NULL, 1000000048, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000023, 'white wine ', 625, 6, NULL, 'Irresistibly, refreshingly good. Bright aromatic herbal notes of this wine works well with some feta or goats cheese. ', NULL, NULL, 'Maldives', 'Unavailable', 'Beer, Wine & Spirits', NULL, 1000000049, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000024, 'tiger beer', 60, 5, NULL, '\nBrewed fresh in Singapore, for Singapore. Tropical lagered since 1932 for a full bodied yet refreshing taste.', NULL, NULL, 'China', 'Available', 'Beer, Wine & Spirits', NULL, 1000000033, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000025, 'champagne ', 123, 6, NULL, 'Made from a blend of 80-percent Pinot Noir and 20-percent Chardonnay, this Champagne evokes the typical character of our vineyard by hillsides and the power of the Pinots Noirs.', NULL, NULL, 'China', 'Unavailable', 'Beer, Wine & Spirits', NULL, 1000000034, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000026, 'croissant ', 3, 6, NULL, 'Delicious butter croissants free from artificial colors and flavorings.', NULL, NULL, 'United States', 'Available', 'Food Pantry', NULL, 1000000035, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000027, 'danish ', 2.5, 1, NULL, 'The crown shape Danish pastry with butter.', NULL, NULL, 'France', 'Unavailable', 'Food Pantry', NULL, 1000000037, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000028, 'puff pastry ', 3.6, 9, NULL, 'Pastry is so mouth wateringly delicious, light & flaky.', NULL, NULL, 'France', 'Unavailable', 'Food Pantry', NULL, 1000000033, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000029, 'brownie ', 1.8, 7, NULL, '\nIndividually packed ready-to-eat brownies. Vegetarian-Friendly with no alcoholic ingredients', NULL, NULL, 'United States', 'Available', 'Food Pantry', NULL, 1000000012, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000030, 'apple strudel ', 10, 6, NULL, 'Classic Apple Pastry, layer of buttery puff pastry filled with the caramelised apple & custard lattice.', NULL, NULL, 'United States', 'Unavailable', 'Food Pantry', NULL, 1000000013, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000031, 'wet wipes ', 62, 2, NULL, 'Anti-Bacterial Wet Tissue effectively kills 99.99-percent of bacteria, keeping your baby safe.', NULL, NULL, 'Hong Kong', 'Unavailable', 'Mum & Baby', NULL, 1000000014, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000032, 'diapers ', 40, 7, NULL, '\nOrganic cotton is contained in the surface sheets of the diapers, which makes babies comfortable as if in mom\'s hands! It has a soft touch to the skin, so no need to worry about delicate babies\' skin. 128 pcs (4 packs).', NULL, NULL, 'France', 'Available', 'Mum & Baby', NULL, 1000000016, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000033, 'pacifier ', 15, 2, NULL, 'An anatomically shaped mouth shield and motifs to make your baby smile.', NULL, NULL, 'Isle of Man', 'Unavailable', 'Mum & Baby', NULL, 1000000017, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000034, 'milk bottle ', 58, 3, NULL, 'Super-sensitive, easi-vent valve eliminates excessive air flow', NULL, NULL, 'China', 'Unavailable', 'Mum & Baby', NULL, 1000000043, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000035, 'milk powder', 55, 6, NULL, 'support cognitive functioning, strong bones and teeth ', NULL, NULL, 'Indonesia', 'Available', 'Mum & Baby', NULL, 1000000045, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000036, 'protein shake ', 42, 10, NULL, 'contains clean plant protein, as well as probiotics and digestive enzymes for easier absorption and utilisation', NULL, NULL, 'China', 'Available', 'Health', NULL, 1000000048, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000037, 'vitamin C', 23, 1, NULL, 'protects cells from oxidative damage and increases iron absorption', NULL, NULL, 'Singapore', 'Unavailable', 'Health', NULL, 1000000016, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000038, 'collagen drink ', 45, 7, NULL, 'minimize the signs of ageing such as fine lines, wrinkles and dryness', NULL, NULL, 'Italy', 'Available', 'Health', NULL, 1000000007, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000039, 'blood pressure monitor ', 200, 1, NULL, 'Accurate Blood Pressure and Heart Rate Monitoring', NULL, NULL, 'Indonesia', 'Unavailable', 'Health', NULL, 1000000009, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000040, 'Oximeter', 15, 3, NULL, 'a noninvasive method for monitoring a person\'s oxygen saturation', NULL, NULL, 'Indonesia', 'Unavailable', 'Health', NULL, 1000000010, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000041, 'moisturizer', 16, 5, NULL, 'This extra-rich body lotion formula, enriched with Certified Shea Butter, provides deep hydration', NULL, NULL, 'Indonesia', 'Unavailable', 'Skin Care', NULL, 1000000012, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000042, 'sun block ', 22, 1, NULL, 'It is water resistant, sweat resistant, resists rub-off, non-comedogenic, dermatologist tested, oil free and PABA free', NULL, NULL, 'Singapore', 'Unavailable', 'Skin Care', NULL, 1000000013, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000043, 'toner', 19, 2, NULL, 'gently removes impurities and minimizes appearance of acne and large pores', NULL, NULL, 'Sri Lanka', 'Available', 'Skin Care', NULL, 1000000002, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000044, 'bb cream ', 25, 9, NULL, 'Contains Hyaluronic Acid Complex, Natural Green Tea extract, prevents pollutants from touching skin directly', NULL, NULL, 'China', 'Unavailable', 'Skin Care', NULL, 1000000004, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000045, 'cc cream ', 21, 9, NULL, 'Formulated with smart colour-match capsule, colour pigments encapsulated by powder burst out and mix during application that blend into skin colour', NULL, NULL, 'Indonesia', 'Available', 'Skin Care', NULL, 1000000005, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000046, 'shampoo ', 15, 3, NULL, 'Deeply cleanses hair roots, removing excess oil and dirt ', NULL, NULL, 'China', 'Unavailable', 'Personal Care', NULL, 1000000031, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000047, 'conditioner ', 8, 3, NULL, 'Leaves hair feeling silky and smooth ', NULL, NULL, 'Australia', 'Available', 'Personal Care', NULL, 1000000033, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000048, 'face wash ', 20, 8, NULL, 'encourage the healing and detoxification of the skin', NULL, NULL, 'Cambodia', 'Unavailable', 'Personal Care', NULL, 1000000034, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000049, 'body wash ', 16, 2, NULL, 'Creamy, bubbly and gentle on the skin', NULL, NULL, 'China', 'Unavailable', 'Personal Care', NULL, 1000000017, 0);
INSERT INTO `my_shop`.`listing_tab` (`l_item_id`, `item_name`, `item_price`, `item_quantity`, `item_purchasedquantity`, `item_description`, `item_shippinginfo`, `item_paymentinfo`, `item_location`, `item_status`, `item_category`, `item_image`, `l_seller_id`, `listing_date`) VALUES (2000000050, 'hand soap ', 5, 7, NULL, 'Rich-lathering and skin-loving cleanser ', NULL, NULL, 'South Korea', 'Available', 'Personal Care', NULL, 1000000017, 0);

COMMIT;


-- -----------------------------------------------------
-- Data for table `my_shop`.`wallet_tab`
-- -----------------------------------------------------
START TRANSACTION;
USE `my_shop`;
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000001, 4000000001, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000002, 4000000002, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000003, 4000000003, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000004, 4000000004, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000005, 4000000005, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000006, 4000000006, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000007, 4000000007, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000008, 4000000008, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000009, 4000000009, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000010, 4000000010, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000011, 4000000011, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000012, 4000000012, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000013, 4000000013, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000014, 4000000014, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000015, 4000000015, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000016, 4000000016, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000017, 4000000017, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000018, 4000000018, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000019, 4000000019, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000020, 4000000020, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000021, 4000000021, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000022, 4000000022, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000023, 4000000023, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000024, 4000000024, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000025, 4000000025, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000026, 4000000026, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000027, 4000000027, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000028, 4000000028, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000029, 4000000029, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000030, 4000000030, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000031, 4000000031, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000032, 4000000032, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000033, 4000000033, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000034, 4000000034, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000035, 4000000035, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000036, 4000000036, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000037, 4000000037, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000038, 4000000038, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000039, 4000000039, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000040, 4000000040, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000041, 4000000041, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000042, 4000000042, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000043, 4000000043, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000044, 4000000044, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000045, 4000000045, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000046, 4000000046, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000047, 4000000047, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000048, 4000000048, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000049, 4000000049, 0.00, NULL, NULL, NULL);
INSERT INTO `my_shop`.`wallet_tab` (`w_wallet_id`, `w_user_id`, `wallet_balance`, `wallet_status`, `last_topup`, `last_used`) VALUES (1000000050, 4000000050, 0.00, NULL, NULL, NULL);

COMMIT;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
