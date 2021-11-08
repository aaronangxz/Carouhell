CREATE DATABASE  IF NOT EXISTS `heroku_bdc39d4687a85d4` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `heroku_bdc39d4687a85d4`;
-- MySQL dump 10.13  Distrib 8.0.25, for macos11 (x86_64)
--
-- Host: us-cdbr-east-04.cleardb.com    Database: heroku_bdc39d4687a85d4
-- ------------------------------------------------------
-- Server version	5.6.50-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `acc_credentials_tab`
--

DROP TABLE IF EXISTS `acc_credentials_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `acc_credentials_tab` (
  `c_user_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_password` varchar(256) NOT NULL,
  PRIMARY KEY (`c_user_id`),
  CONSTRAINT `a_user_id, c_user_id` FOREIGN KEY (`c_user_id`) REFERENCES `acc_tab` (`a_user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `ValidateUserPassword` CHECK (length(user_password) >= 6)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `acc_credentials_tab`
--

LOCK TABLES `acc_credentials_tab` WRITE;
/*!40000 ALTER TABLE `acc_credentials_tab` DISABLE KEYS */;
/*!40000 ALTER TABLE `acc_credentials_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `acc_tab`
--

DROP TABLE IF EXISTS `acc_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `acc_tab` (
  `a_user_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(45) NOT NULL,
  `user_email` varchar(45) NOT NULL,
  `user_ctime` int(11) DEFAULT NULL,
  `user_status` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `user_last_login` int(11) DEFAULT NULL,
  PRIMARY KEY (`a_user_id`),
  CONSTRAINT `ValidateAccCredentials` CHECK (user_status BETWEEN 0 AND 3)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `acc_tab`
--

LOCK TABLES `acc_tab` WRITE;
/*!40000 ALTER TABLE `acc_tab` DISABLE KEYS */;
/*!40000 ALTER TABLE `acc_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary view structure for view `get_purchase_transactions_basic`
--

DROP TABLE IF EXISTS `get_purchase_transactions_basic`;
/*!50001 DROP VIEW IF EXISTS `get_purchase_transactions_basic`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `get_purchase_transactions_basic` AS SELECT 
 1 AS `lt_item_id`,
 1 AS `transaction_amount`,
 1 AS `transaction_type`,
 1 AS `transaction_ctime`,
 1 AS `wt_user_id`,
 1 AS `transaction_ref`,
 1 AS `lt_transaction_id`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `get_sales_transactions_basic`
--

DROP TABLE IF EXISTS `get_sales_transactions_basic`;
/*!50001 DROP VIEW IF EXISTS `get_sales_transactions_basic`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `get_sales_transactions_basic` AS SELECT 
 1 AS `lt_item_id`,
 1 AS `transaction_amount`,
 1 AS `transaction_type`,
 1 AS `transaction_ctime`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `get_wallet_transactions_basic`
--

DROP TABLE IF EXISTS `get_wallet_transactions_basic`;
/*!50001 DROP VIEW IF EXISTS `get_wallet_transactions_basic`*/;
SET @saved_cs_client     = @@character_set_client;
/*!50503 SET character_set_client = utf8mb4 */;
/*!50001 CREATE VIEW `get_wallet_transactions_basic` AS SELECT 
 1 AS `lt_item_id`,
 1 AS `transaction_amount`,
 1 AS `transaction_type`,
 1 AS `transaction_ctime`,
 1 AS `wt_user_id`*/;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `listing_reactions_tab`
--

DROP TABLE IF EXISTS `listing_reactions_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `listing_reactions_tab` (
  `reactions_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `rt_user_id` int(10) unsigned NOT NULL,
  `rt_item_id` int(10) unsigned NOT NULL,
  `reaction_type` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `comment` varchar(256) DEFAULT NULL,
  `ctime` int(11) unsigned NOT NULL,
  PRIMARY KEY (`reactions_id`),
  KEY `a_user_id, rt_user_id_idx` (`rt_user_id`),
  KEY `l_item_id, rt_item_id_idx` (`rt_item_id`),
  CONSTRAINT `a_user_id, rt_user_id` FOREIGN KEY (`rt_user_id`) REFERENCES `acc_tab` (`a_user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `l_item_id, rt_item_id` FOREIGN KEY (`rt_item_id`) REFERENCES `listing_tab` (`l_item_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `ValidateListingReactionType` CHECK (reaction_type BETWEEN 0 AND 1)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `listing_reactions_tab`
--

LOCK TABLES `listing_reactions_tab` WRITE;
/*!40000 ALTER TABLE `listing_reactions_tab` DISABLE KEYS */;
/*!40000 ALTER TABLE `listing_reactions_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `listing_tab`
--

DROP TABLE IF EXISTS `listing_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `listing_tab` (
  `l_item_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `item_name` varchar(50) NOT NULL,
  `item_price` int(10) unsigned NOT NULL DEFAULT '0',
  `item_quantity` smallint(3) unsigned NOT NULL DEFAULT '1',
  `item_purchased_quantity` smallint(3) unsigned DEFAULT '0',
  `item_description` varchar(256) NOT NULL,
  `item_location` tinyint(2) NOT NULL,
  `item_status` tinyint(1) NOT NULL DEFAULT '0',
  `item_category` tinyint(2) NOT NULL,
  `l_seller_id` int(10) unsigned NOT NULL,
  `listing_ctime` int(11) unsigned NOT NULL DEFAULT '0',
  `listing_mtime` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`l_item_id`),
  KEY `a_user_id, l_seller_id` (`l_seller_id`),
  FULLTEXT KEY `item_name` (`item_name`),
  FULLTEXT KEY `item_description` (`item_description`),
  FULLTEXT KEY `search_term` (`item_name`,`item_description`),
  CONSTRAINT `a_user_id, l_seller_id` FOREIGN KEY (`l_seller_id`) REFERENCES `acc_tab` (`a_user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `ValidateItemLocation` CHECK (item_location BETWEEN 0 AND 44),
  CONSTRAINT `ValidateItemStatus` CHECK (item_status BETWEEN 0 AND 3),
  CONSTRAINT `ValidateItemCategory` CHECK (item_category BETWEEN 0 AND 17)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `listing_tab`
--

LOCK TABLES `listing_tab` WRITE;
/*!40000 ALTER TABLE `listing_tab` DISABLE KEYS */;
/*!40000 ALTER TABLE `listing_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `listing_transactions_tab`
--

DROP TABLE IF EXISTS `listing_transactions_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `listing_transactions_tab` (
  `lt_transaction_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `lt_item_id` int(10) unsigned NOT NULL,
  `lt_user_id` int(10) unsigned NOT NULL,
  `transaction_ctime` int(11) NOT NULL,
  `transaction_quantity` smallint(3) NOT NULL,
  `transaction_amount` int(11) DEFAULT NULL,
  PRIMARY KEY (`lt_transaction_id`),
  UNIQUE KEY `transaction_id_UNIQUE` (`lt_transaction_id`),
  KEY `a_user_id, lt_user_id` (`lt_user_id`),
  KEY `l_item_id, lt_item_id` (`lt_item_id`),
  CONSTRAINT `a_user_id, lt_user_id` FOREIGN KEY (`lt_user_id`) REFERENCES `acc_tab` (`a_user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `l_item_id, lt_item_id` FOREIGN KEY (`lt_item_id`) REFERENCES `listing_tab` (`l_item_id`) ON DELETE NO ACTION ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `listing_transactions_tab`
--

LOCK TABLES `listing_transactions_tab` WRITE;
/*!40000 ALTER TABLE `listing_transactions_tab` DISABLE KEYS */;
/*!40000 ALTER TABLE `listing_transactions_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_review_tab`
--

DROP TABLE IF EXISTS `user_review_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_review_tab` (
  `review_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `rv_user_id` int(10) unsigned NOT NULL,
  `rv_seller_id` int(10) unsigned NOT NULL,
  `ratings` tinyint(1) NOT NULL DEFAULT '0',
  `review_text` varchar(256) DEFAULT NULL,
  `ctime` int(11) unsigned NOT NULL,
  PRIMARY KEY (`review_id`),
  KEY `a_user_id, rv_user_id_idx` (`rv_user_id`),
  KEY `l_seller_id, rv_seller_id_idx` (`rv_seller_id`),
  CONSTRAINT `a_user_id, rv_seller_id` FOREIGN KEY (`rv_seller_id`) REFERENCES `acc_tab` (`a_user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `a_user_id, rv_user_id` FOREIGN KEY (`rv_user_id`) REFERENCES `acc_tab` (`a_user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `ValidateUserReviewRatings` CHECK (ratings BETWEEN 1 AND 5)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_review_tab`
--

LOCK TABLES `user_review_tab` WRITE;
/*!40000 ALTER TABLE `user_review_tab` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_review_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wallet_tab`
--

DROP TABLE IF EXISTS `wallet_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `wallet_tab` (
  `w_user_id` int(10) unsigned NOT NULL,
  `wallet_balance` int(10) unsigned NOT NULL DEFAULT '0',
  `wallet_status` tinyint(1) DEFAULT '0',
  `last_top_up` int(11) DEFAULT NULL,
  `last_used` int(11) DEFAULT NULL,
  PRIMARY KEY (`w_user_id`),
  KEY `a_user_id, wallet_id_idx` (`w_user_id`),
  CONSTRAINT `a_user_id, w_user_id` FOREIGN KEY (`w_user_id`) REFERENCES `acc_tab` (`a_user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `ValidateWallet` CHECK (wallet_status BETWEEN 0 AND 2)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wallet_tab`
--

LOCK TABLES `wallet_tab` WRITE;
/*!40000 ALTER TABLE `wallet_tab` DISABLE KEYS */;
/*!40000 ALTER TABLE `wallet_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `wallet_transactions_tab`
--

DROP TABLE IF EXISTS `wallet_transactions_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `wallet_transactions_tab` (
  `wt_transaction_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `wt_user_id` int(10) unsigned NOT NULL,
  `transaction_ctime` int(11) DEFAULT NULL,
  `transaction_amount` int(10) unsigned NOT NULL DEFAULT '0',
  `transaction_type` tinyint(1) DEFAULT '0',
  `transaction_ref` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`wt_transaction_id`),
  UNIQUE KEY `transaction_id_UNIQUE` (`wt_transaction_id`),
  KEY `a_user_id, wt_user_id` (`wt_user_id`),
  KEY `transaction_ref, lt_transaction_id` (`transaction_ref`),
  CONSTRAINT `a_user_id, wt_user_id` FOREIGN KEY (`wt_user_id`) REFERENCES `acc_tab` (`a_user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `transaction_ref, lt_transaction_id` FOREIGN KEY (`transaction_ref`) REFERENCES `listing_transactions_tab` (`lt_transaction_id`) ON DELETE NO ACTION ON UPDATE CASCADE,
  CONSTRAINT `ValidateWalletTransactionType` CHECK (transaction_type BETWEEN 0 AND 2)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `wallet_transactions_tab`
--

LOCK TABLES `wallet_transactions_tab` WRITE;
/*!40000 ALTER TABLE `wallet_transactions_tab` DISABLE KEYS */;
/*!40000 ALTER TABLE `wallet_transactions_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'heroku_bdc39d4687a85d4'
--
/*!50003 DROP PROCEDURE IF EXISTS `add_review` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`b0bc6fadb8432d`@`%` PROCEDURE `add_review`(IN userid INT(10), IN sellerid INT(10), IN rating INT(1), 
IN reviewtext VARCHAR(255), OUT status INT(1), OUT newrating DOUBLE(1,1))
BEGIN
DECLARE _rollback BOOL DEFAULT 0;
DECLARE CONTINUE HANDLER FOR SQLEXCEPTION SET _rollback = 1;
START TRANSACTION;
INSERT INTO user_review_tab (rv_user_id, rv_seller_id, ratings, review_text,ctime) 
SELECT userid,sellerid,rating,reviewtext,unix_timestamp() FROM dual
WHERE NOT EXISTS (SELECT * FROM user_review_tab WHERE rv_user_id = userid AND rv_seller_id = sellerid);
SELECT Round(( Sum(ratings) / Count(ratings) ), 1) AS new_rating
FROM   user_review_tab
WHERE  rv_seller_id = sellerid;
IF _rollback THEN
		set newrating = 0.0;
		set status = -1;
        ROLLBACK;
    ELSE
		set newrating = new_rating;
		set status = 0;
        COMMIT;
    END IF;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `create_user` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`b0bc6fadb8432d`@`%` PROCEDURE `create_user`(IN username VARCHAR(255), IN useremail VARCHAR(255), IN userpassword VARCHAR(255), OUT status INT(1))
BEGIN
DECLARE _rollback BOOL DEFAULT 0;

/*
DECLARE exit handler for sqlexception
BEGIN
SET _rollback = 1;
GET DIAGNOSTICS CONDITION 1
@p1 = RETURNED_SQLSTATE, @p2 = MESSAGE_TEXT;
SELECT @p1 as RETURNED_SQLSTATE  , @p2 as MESSAGE_TEXT;
ROLLBACK;
END;

DECLARE exit handler for sqlwarning
BEGIN
SET _rollback = 1;
GET DIAGNOSTICS CONDITION 1
@p1 = RETURNED_SQLSTATE, @p2 = MESSAGE_TEXT;
SELECT @p1 as RETURNED_SQLSTATE  , @p2 as MESSAGE_TEXT;
ROLLBACK;
END;
*/

DECLARE CONTINUE HANDLER FOR SQLEXCEPTION SET _rollback = 1;
START TRANSACTION;
INSERT INTO acc_tab (user_name, user_email, user_ctime, user_status, user_last_login) 
SELECT username,useremail,unix_timestamp(),1,unix_timestamp() FROM dual
WHERE NOT EXISTS (SELECT * FROM acc_tab WHERE user_name = username OR user_email = useremail);
SELECT LAST_INSERT_ID() INTO @userid;
INSERT INTO acc_credentials_tab  
SELECT @userid, userpassword FROM dual
WHERE NOT EXISTS (SELECT * FROM acc_credentials_tab WHERE c_user_id = @userid);
INSERT INTO wallet_tab (w_user_id)  
SELECT @userid FROM dual
WHERE NOT EXISTS (SELECT * FROM wallet_tab WHERE w_user_id = @userid);
IF _rollback THEN
    SET status = -1;
    Rollback;
  ELSE
    SET status = 0;
	COMMIT;
END If;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 DROP PROCEDURE IF EXISTS `get_user_transactions` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`b0bc6fadb8432d`@`%` PROCEDURE `get_user_transactions`(IN userid INT(10))
BEGIN
START TRANSACTION;
SELECT transaction_history.lt_item_id AS item_id, 
transaction_history.transaction_amount, 
transaction_history.transaction_type, 
transaction_history.transaction_ctime, 
item_info.item_name FROM(
SELECT lt_item_id , transaction_amount, transaction_type, transaction_ctime FROM(
SELECT lt_item_id, transaction_amount, transaction_type, transaction_ctime FROM get_sales_transactions_basic 
WHERE lt_item_id IN
	(SELECT l_item_id FROM listing_tab 
	WHERE l_seller_id = userid)
UNION ALL
SELECT lt_item_id, transaction_amount, transaction_type, transaction_ctime FROM get_purchase_transactions_basic
WHERE wt_user_id = userid
UNION ALL
SELECT lt_item_id, transaction_amount, transaction_type, transaction_ctime FROM get_wallet_transactions_basic
WHERE wt_user_id = userid
) AS transactions) AS transaction_history
LEFT JOIN
(
	SELECT l_item_id, item_name FROM listing_tab
) AS item_info ON transaction_history.lt_item_id = item_info.l_item_id ORDER BY transaction_ctime DESC;
END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Final view structure for view `get_purchase_transactions_basic`
--

/*!50001 DROP VIEW IF EXISTS `get_purchase_transactions_basic`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`b0bc6fadb8432d`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `get_purchase_transactions_basic` AS select `lt`.`lt_item_id` AS `lt_item_id`,`wt`.`transaction_amount` AS `transaction_amount`,`wt`.`transaction_type` AS `transaction_type`,`wt`.`transaction_ctime` AS `transaction_ctime`,`wt`.`wt_user_id` AS `wt_user_id`,`wt`.`transaction_ref` AS `transaction_ref`,`lt`.`lt_transaction_id` AS `lt_transaction_id` from (`wallet_transactions_tab` `wt` join `listing_transactions_tab` `lt`) where (`wt`.`transaction_ref` = `lt`.`lt_transaction_id`) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `get_sales_transactions_basic`
--

/*!50001 DROP VIEW IF EXISTS `get_sales_transactions_basic`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`b0bc6fadb8432d`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `get_sales_transactions_basic` AS select `listing_transactions_tab`.`lt_item_id` AS `lt_item_id`,`listing_transactions_tab`.`transaction_amount` AS `transaction_amount`,2 AS `transaction_type`,`listing_transactions_tab`.`transaction_ctime` AS `transaction_ctime` from `listing_transactions_tab` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `get_wallet_transactions_basic`
--

/*!50001 DROP VIEW IF EXISTS `get_wallet_transactions_basic`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8mb4 */;
/*!50001 SET character_set_results     = utf8mb4 */;
/*!50001 SET collation_connection      = utf8mb4_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`b0bc6fadb8432d`@`%` SQL SECURITY DEFINER */
/*!50001 VIEW `get_wallet_transactions_basic` AS select NULL AS `lt_item_id`,`wallet_transactions_tab`.`transaction_amount` AS `transaction_amount`,`wallet_transactions_tab`.`transaction_type` AS `transaction_type`,`wallet_transactions_tab`.`transaction_ctime` AS `transaction_ctime`,`wallet_transactions_tab`.`wt_user_id` AS `wt_user_id` from `wallet_transactions_tab` where (`wallet_transactions_tab`.`transaction_type` = 0) */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-11-08 23:00:01
