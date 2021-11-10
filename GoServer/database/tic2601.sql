CREATE DATABASE  IF NOT EXISTS `heroku_bdc39d4687a85d4` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `heroku_bdc39d4687a85d4`;
-- MySQL dump 10.13  Distrib 8.0.25, for macos11 (x86_64)
--
-- Host: us-cdbr-east-04.cleardb.com    Database: heroku_bdc39d4687a85d4
-- ------------------------------------------------------
-- Server version	5.6.50-log
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

INSERT INTO `` (`c_user_id`,`user_password`) VALUES (335,'$2a$14$/Cx4/.ElAB0QOpiqtfIf0.N/thWXTialtZbK1jJ/8LXa8wcYd.Xaq');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (215,'$2a$14$0c6tCHQ1ktbRWhiLlWoyM.8qmVvmHOiVoUnTDG.qnj2GsTHxFgH3a');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (525,'$2a$14$14XYiRvkx6FPcY3JfllUceTS3EuGdgVrXs7ILo9IwBgNnthzECsve');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (305,'$2a$14$4li6IsSHs9Xaa5PrUDVFuuoXa1N5TCwtKr02eb00kOn9dq/0EMsje');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (535,'$2a$14$64bcK4DBfF5ngMZp4f0ViucQNuAsWMR9MPXX2EgL7d4nJlufOviY.');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (485,'$2a$14$6HeZfYQtB16eQ.I7ECgAdeNID6FIi4wKfrf01.kswNwLiCbzX39Ia');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (145,'$2a$14$6W6zdd1SQ.SY1d/QbECr3ObOGpNIEwWp6r4S6Oar5BmUXrNe8X.dC');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (405,'$2a$14$7IwQCFlF8CqDmyfShDV3iO3gUrDlP0G4d2XVfnIkL6maL.j9maukO');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (185,'$2a$14$7T4GOUWIDJjzqKSIXle6f.4QaoZkEPhoONS.OeoXBTPksvsQ8msUq');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (345,'$2a$14$8kirvHHobTr7y7Dn7vNfq.vxgxwIBgEYnUICe3vVCxGmF8xzKFDEO');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (15,'$2a$14$9o.P8hx.h.a1vvlHXUGLwuWObGWFEm5n7QbmIcu4TmXZr4d2nwEAy');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (105,'$2a$14$aC88Z5hbTLRHKclh5jdGpunUMTWCyFb9HVJlLRX8zutW0mGffzBI6');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (65,'$2a$14$bF3GnXzFwm63ZuffDr/ZO.ImgcY4JVlW2d85NcHLigDGveTCrTyuu');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (5,'$2a$14$BxKow/n4RMJcCfDN2kmvp.4yxW8g5LLCmGqhZulkoyIW6puwd4VrS');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (465,'$2a$14$cHazrFF1UQXnPA0BgnA6GuLfELItoOwX9NHr4vGn4hTLj2YHkuEeG');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (85,'$2a$14$CoBCr5uOeAwsX4swzSNNEO8CfYy6yXsC6GGqSCAvZFAnAVFfGxAMm');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (585,'$2a$14$GXWeR86WoTwV5k05r0Azo.4B2gUlLIY2a5f5CTbtEQ9WrCyZrY/f2');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (315,'$2a$14$GYki3xGzSzSaG2Hx9qHDF.3G7WjiKIL68o2hLuUPe0vNK9gs3GV.O');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (155,'$2a$14$HmPpZXvIGnVKhrl1WB3ZdelNAF7137uQ3hJ8mrbpy0CEXJIgOElOW');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (555,'$2a$14$IMjCLIOUTRhA4M5KSZtwI.rbsuPl0NH1bzkmLbaxc//B/rZr60GMi');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (425,'$2a$14$IU11O.xUyYaTmtM8wx9D2eZ8UaRngYyAxi6uLWIP6.2CiVRr2teX2');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (265,'$2a$14$IxCcwbvs8kRtuHRtWWyrRePQHQ.ghfpeqlO3OalzAqaCURG7KOor.');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (505,'$2a$14$jKCcR6dwKR76W1yHo8JElOEb93qO9xHZSqQAF5hxQgoGCPSlkPsRS');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (295,'$2a$14$jRdRwfR57lkxRzZfMQXIn.DUvxZyAkD6OMJBGyL/gXvs9foGz1tV.');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (45,'$2a$14$KMDOWcxzp27VxplKUYDlMeTu2/.kZuH76qopfg7CAOckHkoNVJEfS');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (515,'$2a$14$LAcKKUnw.5ebSCNTJ4slJOOdZ5/yfrpBluie6vu0wk2nQf5yfdd5e');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (415,'$2a$14$ld.W.APz.fpEvdXdNEWVYOahYYafymt/Uxl/C5kZLObSE9n4qwEGS');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (225,'$2a$14$lIOmqEszlRGZ/WcvgR.iTOw0rhf7tsh1.xqCCaAQZ4t4SlkMgWhvu');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (385,'$2a$14$M8bhBdrCBTngUk9jbOSOoOmLFMBG1n9q8W0GPqgOoRCUM6exF4eQu');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (255,'$2a$14$MhjJvGISOHXR3/r.ftbkauvPRF5utjToDwPVWeNAAgP6DelP7BFVm');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (205,'$2a$14$mNHbfrk/JCnAULXSLY25ueFmbnj.A72hA4qz7DvD.4gAYT.AS085a');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (285,'$2a$14$mNyow38z62dPfQzydt.FYOLbJolXUW85Ss9vX0TROlZQzN9ecz7YO');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (395,'$2a$14$MZvpH0TzyMNxKqay6y2VCObywuLL7CqKLabNrV07isGT9F4suAX8a');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (245,'$2a$14$nI2hLcjtGe4jQzdjpoQRy.T8VEOfpYbNRwyeTKTvw9GMY1hVPBePa');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (135,'$2a$14$nVjGSU2.m8Ws6k7akgcEx.0JeywkU1aVUEZJQj7B2fDs53vz2FFmm');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (445,'$2a$14$oAZxI52So9BYPENxG6qd3ukQRqhCc0eTsj.bKVaJbMSvol8JduLka');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (325,'$2a$14$pBf87I0ipRvwsD.tkSh1q.SBJV5GbgfJSo2i.3/WNoLuC0aPrLfwe');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (545,'$2a$14$Pw80phbKdqvjaZzfglWYfO0YLnXozl0D4WQyhA1.ntoF8EGB8stFS');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (435,'$2a$14$PWlYsjNsKMNBS18313/fD.t3HzFyXyxp6vQ39/fWlNjzAqD7WcRZi');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (565,'$2a$14$qPrnqc03.rDdCoktcL2JNOCHphODc3WfIRGDXoFTZl9/yFwPM8912');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (75,'$2a$14$qVIiVTPVDTo5sXGIhiAwNOj/4I5y94AQk.ZNxq96.EScye6nL9bO.');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (125,'$2a$14$R7k9SY3NXIvOAES4J2oY4.mrIFActzB6PvlFVSbM.oxcfdZL5.TEa');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (95,'$2a$14$rhK97UInJRjfKNOE6BmE0O27Ua/xHPU5U.KZwJImlhY5qtFomq.ly');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (175,'$2a$14$SEHmc1D5mlBNT71L4N.aY.7D1IdhOl3XU8BkQmM37UMOpyII0mBR2');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (115,'$2a$14$Sf1x9Offte4R2.Bg1r0p/uk0E.ZGcdG5rcyjlxlzUcNEMENrLqDNS');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (195,'$2a$14$sgKKe9cZVZpmxzPmMEOnmewtYcqJcE4uNlcRxgc4AX9mCmPseWPEi');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (375,'$2a$14$ShJ5E6bCqsHuY90gialuVuIWfZW90x8oas0jwQNkxBRgGATz9kA0W');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (35,'$2a$14$sPlbhOVSGYSGp2OZ6Dkr6.ms4b/OFnfcRquVLnKOWEdjWaMgeba96');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (495,'$2a$14$TWiaLYbkpOxpZkb.5X/7Q.bIvDYFCacf1DTtvLnWpwgPRm9Jyxnym');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (365,'$2a$14$uRNmjfIrvUYObVpX6Tr.bOZDl3MGj1hEogCdXFaNuwjQwH4iG2FTy');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (455,'$2a$14$v1sKtkEpRQRzLW95eJ/Dx.7m2n4IqIHOfieTwkoJ2t4SDrOIfpmH6');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (275,'$2a$14$vlZyComVpDIhbQ.Zb8ZDJ.GEx.8fp9mN4G4fkmH39ObNX6oc6Tj2W');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (165,'$2a$14$w2ht2COCDuK7pX4yMmW67eDyv50.N4cRzEImcclNuaDeoEy4NmS4u');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (25,'$2a$14$wewqBYDb4ixBy/6kPvdkdeKSRj.Vwf9noN8bcP9yhm/rfQpABTjxy');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (235,'$2a$14$wKzesGUaI/keR6PiQvyIueDKGw0zz3hY6tsnuX/hK7twRyRZWaXsa');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (355,'$2a$14$xGliMjO1cpECAvto1jxase6o7939jTmaJ9l7.kZfxTVHCV5YPLp..');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (475,'$2a$14$yXW4iFnGxlmOl1rRvq0vvO2/YA8vO1Cx3S1VbHwZ8jt/SnSDKhfZe');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (55,'$2a$14$zMIXHaWEGErGhEtUHO1GCO/yuKLMnTAjofM6cLq/GSV8WZMwc8V4G');
INSERT INTO `` (`c_user_id`,`user_password`) VALUES (575,'$2a$14$ZwgIUzXQplImfvRCsW1rpeIQ.ZXjVdwqX1UXaZR7/Ijv3IU3r7ApK');


DROP TABLE IF EXISTS `acc_tab`;
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

INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (5,'xuanze','e0649019@u.nus.edu',1636384442,1,1636554287);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (15,'TeslaSG','tes@la.com',1636385513,1,1636390617);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (25,'AppleSingapore','applesg@apple.com',1636385867,1,1636387252);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (35,'ASUSROG','rog@asus.com',1636389811,1,1636390037);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (45,'ChanelOfficial','sg@chanel.com',1636391329,1,1636392220);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (55,'HermanMillerSG','sg@hm.com',1636392859,1,1636392866);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (65,'Kerris','kerris@gmail.com',1636395011,1,1636479178);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (75,'brandyvealratfog','brandyvealratfog@gmail.com',1636545236,1,1636545236);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (85,'snowstormokranet','snowstormokranet@gmail.com',1636545248,1,1636545248);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (95,'flatspiritedaway','flatspiritedaway@gmail.com',1636545259,1,1636545259);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (105,'beforesunsetnix','beforesunsetnix@gmail.com',1636545268,1,1636545268);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (115,'oportunitypiefig','oportunitypiefig@gmail.com',1636545274,1,1636545274);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (125,'ostrichclimbing','ostrichclimbing@gmail.com',1636545282,1,1636545282);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (135,'beatblackeyepig','beatblackeyepig@gmail.com',1636545289,1,1636545289);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (145,'cometsbassorion','cometsbassorion@gmail.com',1636545297,1,1636545297);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (155,'capricorngandalf','capricorngandalf@gmail.com',1636545304,1,1636545304);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (165,'romanholidaynova','romanholidaynova@gmail.com',1636545311,1,1636545311);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (175,'marsnormanbates','marsnormanbates@gmail.com',1636545322,1,1636545322);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (185,'icethegraduate','icethegraduate@gmail.com',1636545330,1,1636545330);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (195,'tsunamispinach','tsunamispinach@gmail.com',1636545337,1,1636545337);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (205,'pumbaabootstoad','pumbaabootstoad@gmail.com',1636545344,1,1636545344);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (215,'applebrownbread','applebrownbread@gmail.com',1636545356,1,1636545356);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (225,'alienrockykale','alienrockykale@gmail.com',1636545363,1,1636545363);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (235,'vividmarsexpress','vividmarsexpress@gmail.com',1636545369,1,1636545369);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (245,'kumquattangerine','kumquattangerine@gmail.com',1636545376,1,1636545376);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (255,'doughnutshalobat','doughnutshalobat@gmail.com',1636545382,1,1636545382);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (265,'noteshighnoonnet','noteshighnoonnet@gmail.com',1636545414,1,1636545414);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (275,'earthpearpuckpie','earthpearpuckpie@gmail.com',1636545421,1,1636545421);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (285,'cookiesinsideout','cookiesinsideout@gmail.com',1636545427,1,1636545427);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (295,'pathsofgloryfox','pathsofgloryfox@gmail.com',1636545433,1,1636545433);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (305,'fogpulsarmoonnet','fogpulsarmoonnet@gmail.com',1636545440,1,1636545440);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (315,'droughtyipmandog','droughtyipmandog@gmail.com',1636545448,1,1636545448);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (325,'12angrymeniok-1','12angrymeniok-1@gmail.com',1636545456,1,1636545456);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (335,'thegreenmileveal-1','thegreenmileveal@gmail.com',1636545466,1,1636545466);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (345,'normanbatespie-1','normanbatespie@gmail.com',1636545473,1,1636545473);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (355,'strawberryantsea','strawberryantsea@gmail.com',1636545488,1,1636545488);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (365,'flatchickenleg','flatchickenleg@gmail.com',1636545500,1,1636545500);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (375,'jurassicparkred','jurassicparkred@gmail.com',1636545508,1,1636545508);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (385,'flyinginsectfig','flyinginsectfig@gmail.com',1636545516,1,1636545516);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (395,'thehustlerpiesun','thehustlerpiesun@gmail.com',1636545544,1,1636545544);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (405,'nixmilkapricot','nixmilkapricot@gmail.com',1636545550,1,1636545550);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (415,'yogamosscurling','yogamosscurling@gmail.com',1636545556,1,1636545556);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (425,'hydraduckboxing','hydraduckboxing@gmail.com',1636545563,1,1636545563);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (435,'crackersyogaice','crackersyogaice@gmail.com',1636545570,1,1636545570);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (445,'pieplanetarwine','pieplanetarwine@gmail.com',1636545576,1,1636545576);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (455,'marsexpressnix','marsexpressnix@gmail.com',1636545582,1,1636545582);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (465,'junowalrusbagel','junowalrusbagel@gmail.com',1636545655,1,1636545655);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (475,'wildflowerpotato','wildflowerpotato@gmail.com',1636545660,1,1636545660);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (485,'candywinterberry','candywinterberry@gmail.com',1636545667,1,1636545667);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (495,'birthdaycakeweb','birthdaycakeweb@gmail.com',1636545674,1,1636545674);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (505,'zebracranberry','zebracranberry@gmail.com',1636545681,1,1636545681);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (515,'lizardalivesnail','lizardalivesnail@gmail.com',1636545686,1,1636545686);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (525,'rataromaticaries','rataromaticaries@gmail.com',1636545693,1,1636545693);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (535,'squidcatplutobat','squidcatplutobat@gmail.com',1636545700,1,1636545700);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (545,'legolasclovernix','legolasclovernix@gmail.com',1636545705,1,1636545705);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (555,'chimpanzeearies','chimpanzeearies@gmail.com',1636545711,1,1636545711);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (565,'McDonaldsSG','mcd@mcdsg.com',1636546219,1,1636550068);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (575,'IKEAsg','sg@ikea.com',1636550139,1,1636550150);
INSERT INTO `` (`a_user_id`,`user_name`,`user_email`,`user_ctime`,`user_status`,`user_last_login`) VALUES (585,'SamsungSG','sg@samsung.com',1636551286,1,1636551292);


CREATE 
    ALGORITHM = UNDEFINED 
    DEFINER = `b0bc6fadb8432d`@`%` 
    SQL SECURITY DEFINER
VIEW `get_purchase_transactions_basic` AS
    SELECT 
        `lt`.`lt_item_id` AS `lt_item_id`,
        `wt`.`transaction_amount` AS `transaction_amount`,
        `wt`.`transaction_type` AS `transaction_type`,
        `wt`.`transaction_ctime` AS `transaction_ctime`,
        `wt`.`wt_user_id` AS `wt_user_id`,
        `wt`.`transaction_ref` AS `transaction_ref`,
        `lt`.`lt_transaction_id` AS `lt_transaction_id`
    FROM
        (`wallet_transactions_tab` `wt`
        JOIN `listing_transactions_tab` `lt`)
    WHERE
        (`wt`.`transaction_ref` = `lt`.`lt_transaction_id`)

CREATE 
    ALGORITHM = UNDEFINED 
    DEFINER = `b0bc6fadb8432d`@`%` 
    SQL SECURITY DEFINER
VIEW `get_sales_transactions_basic` AS
    SELECT 
        `listing_transactions_tab`.`lt_item_id` AS `lt_item_id`,
        `listing_transactions_tab`.`transaction_amount` AS `transaction_amount`,
        2 AS `transaction_type`,
        `listing_transactions_tab`.`transaction_ctime` AS `transaction_ctime`
    FROM
        `listing_transactions_tab`

CREATE 
    ALGORITHM = UNDEFINED 
    DEFINER = `b0bc6fadb8432d`@`%` 
    SQL SECURITY DEFINER
VIEW `get_wallet_transactions_basic` AS
    SELECT 
        NULL AS `lt_item_id`,
        `wallet_transactions_tab`.`transaction_amount` AS `transaction_amount`,
        `wallet_transactions_tab`.`transaction_type` AS `transaction_type`,
        `wallet_transactions_tab`.`transaction_ctime` AS `transaction_ctime`,
        `wallet_transactions_tab`.`wt_user_id` AS `wt_user_id`
    FROM
        `wallet_transactions_tab`
    WHERE
        (`wallet_transactions_tab`.`transaction_type` = 0)

DROP TABLE IF EXISTS `listing_reactions_tab`;
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

INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (5,5,5,1,'TESLA stonks when?',1636386892);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (15,5,25,1,'Can include COE? Fast deal',1636386928);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (25,5,25,0,NULL,1636386942);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (55,5,95,0,NULL,1636389294);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (65,5,105,1,'i\'m interested',1636390653);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (75,5,135,1,'hello, any nego?',1636390732);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (105,5,145,0,NULL,1636456567);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (115,5,125,1,'hello',1636459735);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (125,65,205,0,NULL,1636479475);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (135,65,165,0,NULL,1636479479);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (145,65,205,1,'I\'m interested! How can we deal?',1636479517);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (155,555,215,1,'Yummy yummy :p',1636546734);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (165,545,215,1,'Best food ever NGL',1636546749);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (175,535,215,1,'How much to upsize??',1636546787);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (185,525,215,1,'Can I get one without the patty?',1636546818);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (195,515,215,1,'I love the lettuce',1636546837);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (205,555,215,0,NULL,1636546865);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (215,545,215,0,NULL,1636546868);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (225,535,215,0,NULL,1636546871);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (235,525,215,0,NULL,1636546874);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (245,515,215,0,NULL,1636546876);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (255,505,215,0,NULL,1636546879);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (265,495,215,0,NULL,1636546882);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (275,485,215,0,NULL,1636546887);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (285,475,215,0,NULL,1636546890);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (295,465,215,0,NULL,1636546893);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (305,455,215,0,NULL,1636546896);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (315,445,215,0,NULL,1636546899);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (325,435,215,0,NULL,1636546902);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (335,425,215,0,NULL,1636546905);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (345,415,215,0,NULL,1636546907);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (355,415,235,0,NULL,1636547141);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (365,405,235,0,NULL,1636547143);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (375,395,235,0,NULL,1636547147);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (385,385,235,0,NULL,1636547149);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (395,375,235,0,NULL,1636547152);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (405,365,235,0,NULL,1636547154);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (415,355,235,0,NULL,1636547157);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (425,345,235,0,NULL,1636547159);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (435,345,225,0,NULL,1636547173);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (445,335,225,0,NULL,1636547176);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (455,325,225,0,NULL,1636547178);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (465,315,225,0,NULL,1636547181);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (475,305,225,0,NULL,1636547183);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (485,405,225,1,'Why mine is not crispy?',1636547225);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (495,395,225,1,'top up $0.50 to get one more piece can?',1636547253);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (505,355,205,1,'god damn so expensive?',1636547306);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (515,345,195,1,'1k fast deal',1636547333);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (525,335,185,1,'So pretty!!',1636547355);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (535,335,145,1,'Omg..I hope my bf sees this',1636547386);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (545,5,225,0,NULL,1636547412);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (555,5,215,0,NULL,1636547414);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (565,5,235,0,NULL,1636547417);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (575,5,165,0,NULL,1636547434);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (585,5,185,0,NULL,1636547522);
INSERT INTO `` (`reactions_id`,`rt_user_id`,`rt_item_id`,`reaction_type`,`comment`,`ctime`) VALUES (595,5,265,0,NULL,1636549608);

--
-- Table structure for table `listing_tab`
--

DROP TABLE IF EXISTS `listing_tab`;
CREATE TABLE `listing_tab` (
  `l_item_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `item_name` varchar(50) NOT NULL,
  `item_price` int(10) unsigned NOT NULL DEFAULT '0',
  `item_quantity` int(5) unsigned NOT NULL DEFAULT '1',
  `item_stock` int(5) unsigned NOT NULL,
  `item_description` varchar(1000) NOT NULL,
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

INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (5,'$TSLA$',150000,1,1,'Tesla, Inc. is an American electric vehicle and clean energy company based in Palo Alto, California, United States. Tesla designs and manufactures electric cars, battery energy storage from home to grid-scale, solar panels and solar roof tiles, and related products and services.',11,1,15,5,1636384939,1636389571);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (25,'Tesla Model S Plaid',49900000,5,5,'With the longest range and quickest acceleration of any electric vehicle in production, Model S Plaid is the highest performing sedan ever built. All Model S powertrains, with updated battery architecture, are capable of back-to-back, consistent 1/4 mile r',0,1,7,15,1636385795,1636385795);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (35,'iPhone 13 Pro Max 1TB',262900,4,5,'Oh. So. Pro. A dramatically more powerful camera system. A display so responsive, every interaction feels new again. The world’s fastest smartphone chip. Exceptional durability. And a huge leap in battery life. Let’s Pro. Super Retina XDR display with ProM',19,1,2,25,1636386106,1636386106);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (45,'Apple Watch Series 7',64900,97,99,'The aluminium case is lightweight and made from 100 per cent recycled aerospace-grade alloy.The Sport Band is made from a durable yet surprisingly soft high-performance fluoroelastomer, with an innovative pin-and-tuck closure.',19,1,2,25,1636386440,1636386440);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (55,'Apple iPad Pro 2TB',351900,99,99,'Faster performance and graphics. The 8‑core CPU of the Apple M1 chip delivers up to 50 per cent faster performance. And the M1 chip has an 8‑core GPU in a class of its own, providing up to 40 per cent faster graphics performance to iPad Pro. So you can bui',19,1,2,25,1636386800,1636386800);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (65,'Apple AirPods Pro 3',26900,999,999,'AirPods are lightweight and offer a contoured design. They sit at just the right angle for comfort and to better direct audio to your ear. The stem is 33 per cent shorter than AirPods (2nd generation) and includes a force sensor to easily control music and calls.',19,1,2,25,1636388115,1636388115);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (75,'MacBook Pro 16 M1 MAX',524900,999,999,'Supercharged for pros. Buy now. With the blazing-fast M1 Pro or M1 Max chip. Up to 21 hr battery life. 1080p FaceTime HD camera. Liquid Retina XDR display. Touch ID. HDMI port.',19,1,2,25,1636388261,1636388261);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (85,'Pro Display XDR',1037700,42,45,'32-inch Retina 6K. Astonishing colour accuracy. Super-wide viewing angle. And Extreme Dynamic Range.',19,1,2,25,1636388373,1636388373);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (95,'Tesla Model X Plaid',52200000,4,5,'With the most power and quickest acceleration of any SUV, Model X Plaid is the highest performing SUV ever built. All Model X powertrains, with updated battery architecture, can deliver instant torque at any speed. Model X platforms unite powertrain and battery technologies for unrivaled performance, range and efficiency. New module and pack thermal architecture allows faster charging and gives you more power and endurance in all conditions.',11,1,7,15,1636388647,1636388749);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (105,'ROG Strix RTX™ 3090',220000,2,2,'NVIDIA Ampere Streaming Multiprocessors: The building blocks for the world’s fastest, most efficient GPU, the all-new Ampere SM brings 2X the FP32 throughput and improved power efficiency. 2nd Generation RT Cores: Experience 2X the throughput of 1st gen RT Cores, plus concurrent RT and shading for a whole new level of ray tracing performance. 3rd Generation Tensor Cores: Get up to 2X the throughput with structural sparsity and advanced AI algorithms such as DLSS. Now with support for up to 8K resolution, these cores deliver a massive boost in game performance and all-new AI capabilities.',22,1,1,35,1636390001,1636390058);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (115,'ROG Swift PG65UQ',499900,10,10,'ROG Swift PG65 Big Format Gaming Display with NVIDIA® G-SYNC™- 65” 4K HDR( 3840 X 2160), 120Hz+, Ultra-low latency, NVIDIA SHIELD™',22,1,1,35,1636390198,1636390198);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (125,'ROG RAMPAGE VI EXTREME OMEGA',148800,10,10,'Intel X299 EATX gaming motherboard LGA 2066 for Intel Core X-Series processors, with ROG DIMM.2, DDR4 4266MHz , onboard 802.11ac Wi-Fi, 10Gbps LAN, USB 3.1 Gen 2, SATA, Quad M.2 and Aura Sync RGB lighting',22,1,1,35,1636390396,1636390396);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (135,'ROG G35CG',599900,10,10,'Compete at a new level with a tournament-ready gaming desktop built to deliver an exhilarating esports experience with Windows 10 Home. Up to an NVIDIA® GeForce RTX™ 3090 graphics card amps up your frames per second to make fast-paced gaming silky smooth. A powerful 11th Gen Intel® Core™ i9 CPU accelerates content creation and can even be overclocked with an AI-enhanced system that makes overclocking easier for everyone. The GT35\'s high end performance is bolstered by a cooling-focused, multi-chambered chassis design and liquid cooling system. Aura Sync accents and Keystone II technology deepen the customization options, allowing it to truly become the perfect machine for every gamer.',22,1,1,35,1636390556,1636390556);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (145,'Classic Handbag Lambskin Black',1319000,8,10,'The CHANEL iconic bag, a symbol of elegance and luxury. Discover The CHANEL Classic Handbag on the official website or in your CHANEL boutique. Icons of CHANEL.',22,1,4,45,1636391577,1636391577);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (155,'Classic Handbag Lambskin White',1319000,10,10,'The CHANEL iconic bag, a symbol of elegance and luxury. Discover The CHANEL Classic Handbag on the official website or in your CHANEL boutique. Icons of CHANEL.',22,1,4,45,1636391700,1636391700);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (165,'Small Vanity Case Beige & Black',888800,10,10,'The CHANEL iconic bag, a symbol of elegance and luxury. Discover The CHANEL Classic Handbag on the official website or in your CHANEL boutique. Icons of CHANEL. Beech Wood, Calfskin & Gold-Tone Metal. Beige & Black',22,1,4,45,1636392196,1636392272);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (175,'Chanel N°5',115000,10,10,'In 1921, Gabrielle Chanel turned the codes of perfumery upside down and created N°5. A radical creation that revolutionised the traditions of its era. A design piece turned icon whose bottle and label are immediately recognizable.\nN°5 celebrates 100 years of fame in 2021. To mark the occasion and in honor of the holiday season, CHANEL is offering a calendar unlike any other, inspired by the signature silhouette of the N°5 perfume bottle. This collector\'s item is comprised of 27 boxes numbered 5 to 31, each of which contains a full-size fragrance or makeup product, miniature, or other surprise marked with Gabrielle Chanel\'s lucky number. A piece to treasure for years to come.',22,1,8,45,1636392381,1636392410);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (185,'Chanel ROUGE ALLURE',5600,97,99,'LIMITED EDITION - N°5 HOLIDAY 2021 COLLECTION LUMINOUS INTENSE LIP COLOUR',22,1,8,45,1636392513,1636392513);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (195,'Aeron - Graphite',240000,30,30,'Aeron Remastered is the perfect marriage of performance and design.Our best-selling office chair still defines expectations for ergonomic comfort more than 20 years after its debut. Comes in Three Sizes Small, Medium, or Large, Aeron’s got your back',21,1,6,55,1636393131,1636393131);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (205,'Embody Chair - Sync Black',252900,30,30,'You feel Embody’s Pixelated Support™ the moment you sit down— a sense that you are floating, yet perfectly balanced. The seat distributes your weight evenly while supporting your body’s micro-movements. The narrow backrest allows you to move freely and naturally as it automatically adjusts to support a full range of seated postures.By reducing seated pressure and encouraging freedom of movement, Embody allows blood and oxygen to flow more freely, which helps keep you focused. Form doesn’t just follow function with Embody. Function is on full display.',21,1,6,55,1636393301,1636393301);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (215,'McSpicy',700,94,100,'If you’re one of those people who like your chicken big on spiciness, this is the sandwich for you. A thick, juicy cutlet of chicken thigh and drum sits fiery hot on a bed of crunchy lettuce between toasted sesame seed buns – shiok!',34,1,0,565,1636546379,1636546379);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (225,'Chicken McCrispy® (6pc)',840,100,100,'Even more crispy, juicy, tender chunks of chicken for you to sink your teeth into. Share it with friends and family, or simply have the entire bucket to yourself—who could blame you?',34,1,0,565,1636546482,1636546482);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (235,'Buttermilk Crispy Chicken',695,100,100,'Crispy whole-muscle chicken thigh flavoured with buttermilk packed in a glazed burger bun. Served with white cheddar cheese, romaine lettuce, black pepper mayo; and topped with grilled pineapple rings and crisp purple cabbage.In short, perfection in every bite!',34,1,0,565,1636546573,1636546573);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (245,'Evian Natural Mineral Water',350,100,100,'Meet evian+, a sparkling mineral enhanced drink made with the natural spring water you know and love, but enhanced with magnesium and zinc',31,1,0,505,1636547738,1636547738);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (255,'BE@RBRICK Basquiat 1000%',120000,1,1,'Dimensions: 1000% height 70cm',3,1,12,475,1636547976,1636547976);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (265,'LEGO Starwars Millenium Falcon',119000,2,2,'The ultimate LEGO® Star Wars Millennium Falcon has landed! With 7,500 elements, Han Solo\'s super-fast Corellian ship is crammed with the coolest details and fun features. Marvel at the intricate hull detailing, sensor dish, upper and lower quad laser cannons, 7 landing legs, lowering boarding ramp and hidden blaster cannon. Seat up to 4 minifigures in the cockpit with detachable canopy and remove the outer panels to reveal the highly detailed interior. The main hold has seating for the minifigures, a Dejarik holographic game, engineering station with a turning seat for a minifigure, and a doorway build with passageway decoration. Make your way to the rear compartment to monitor the hyperdrive from the console. There\'s also a hidden floor compartment, 2 escape pod hatches and an access ladder to the gunnery station seat and a detachable hull panel with a fully rotating quad laser cannon. Display the Millennium Falcon alongside your favorite sets for the perfect collector\'s item… and whe',7,1,12,475,1636548167,1636548167);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (275,'Off-White Galaxy Brush Zipper Hoodie',60000,10,10,'Off-White',10,1,3,395,1636548395,1636548761);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (285,'Nike Air Jordan 1 Retro High OG Pro',20000,2,2,'Cushioning, Slip-resistant, Wearable, Wrapping, Supportive',15,1,3,385,1636548892,1636548892);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (295,'Superman Dog costume',2000,20,20,'So cuteeeeeeee',20,1,5,365,1636549078,1636549078);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (305,'Sushi Cat costume',2000,20,20,'So cuteeeeeeee',20,1,5,365,1636549191,1636549191);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (315,'Solid Wood Dining Table',1999999,2,2,'Natural wood obtained from the forests, having a consistent composition of material throughout its body without hollow gaps. Strong, durable, and resistant.',2,1,6,355,1636550080,1636550080);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (325,'IKEA x ROG MATCHSPEL',25900,20,20,'MATCHSPEL gaming chair helps you play at the top of your game. The whole body enjoys nice support and you can adjust the height of the chair, neck and armrests to sit really comfy when the game begins.',38,1,6,575,1636550317,1636550317);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (335,'VINTER 2021',12900,45,45,'The perfect choice for those who want a really tall and full-bodied Christmas tree. This tree is guaranteed needle-free and takes up little storage space. Just put it in its spot and start decorating.',38,1,6,575,1636550686,1636550739);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (345,'Level 1000 Collector\'s Aegis of Champions',999999,1,1,'All Battle Pass owners who reach Battle Level 1000 are invited to receive The International 2016 Collector\'s Aegis, an exclusive 1/5th-scale, bronze-plated alloy replica of the famed champion\'s prize.',25,1,9,305,1636550924,1636550924);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (355,'Crysis 3',4990,10,10,'Crysis 3 is a 2013 first-person shooter video game developed by Crytek and published in 2013 by Electronic Arts. It is the third and final main game of the Crysis series, a sequel to the 2011 video game Crysis 2. The multiplayer portion of the game was developed by Crytek UK.',8,1,9,305,1636551062,1636551062);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (365,'Family Hub™ Multi Door, 550L',450000,10,10,'Manage your groceries efficiently with View Inside app',8,1,10,585,1636551497,1636551547);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (375,'EcoBubble™, 9Kg, Washer Dryer',115000,30,30,'Intuitive Display with AI Control. Hygiene Steam. EcoBubble™',8,1,10,585,1636551671,1636551671);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (385,'QN900A Neo QLED 8K 85\'',2824000,5,5,'The powerful evolution of Neo QLED 8K comes with a backlight dimming technology that controls our proprietary Quantum Mini LEDs with perfect precision. Witness unimaginable details expressed in both the darkest black to the purest white with x1.5 more lighting zones than normal Quantum Matrix Technology.',8,1,10,585,1636551802,1636551802);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (395,'TI11 Tickets',88888,5,5,'Leaked frotn row tickets',5,1,13,255,1636552080,1636552080);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (405,'五月天 JUST ROCK IT Tickets',66600,5,5,'Mayday Just Rock It!!! BLUE in Singapore is proactively being rescheduled to 03 December 2022 (Saturday) in view of the latest authorities advisory.',5,1,13,245,1636552374,1636552374);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (415,'SUPREME RAWLINGS CHROME MAPLE WOOD BAS',50000,1,1,'100% Authentic',15,1,14,235,1636552554,1636552554);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (425,'SUPREME Brick',39900,1,1,'100% Authentic',15,1,10,235,1636552705,1636552745);
INSERT INTO `` (`l_item_id`,`item_name`,`item_price`,`item_quantity`,`item_stock`,`item_description`,`item_location`,`item_status`,`item_category`,`l_seller_id`,`listing_ctime`,`listing_mtime`) VALUES (435,'Penthouse for sale',2000000000,1,1,'In the living area, the high ceilings and the view of the city were enhanced by to the use of communicating spaces between the dining area and the living room. Elements from the recent Milano collectioncreate an atmosphere of modern luxury, interacting with warm green, beige and brown colours. In the sleeping area the atmosphere is more intimate, the furniture from the Milano collection are proposed in light beige tones in contrast with the darker colours of the walls. The Milano bed upholsterd in fabric becomes the centre piece of the bedroom, a comfortable piece from which one can admire the fascinating surrounding landscape.',22,1,16,205,1636552972,1636552972);

--
-- Table structure for table `listing_transactions_tab`
--

DROP TABLE IF EXISTS `listing_transactions_tab`;
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

INSERT INTO `` (`lt_transaction_id`,`lt_item_id`,`lt_user_id`,`transaction_ctime`,`transaction_quantity`,`transaction_amount`) VALUES (15,45,5,1636386853,1,64900);
INSERT INTO `` (`lt_transaction_id`,`lt_item_id`,`lt_user_id`,`transaction_ctime`,`transaction_quantity`,`transaction_amount`) VALUES (25,95,5,1636389658,1,52200000);
INSERT INTO `` (`lt_transaction_id`,`lt_item_id`,`lt_user_id`,`transaction_ctime`,`transaction_quantity`,`transaction_amount`) VALUES (35,145,5,1636391811,1,1319000);
INSERT INTO `` (`lt_transaction_id`,`lt_item_id`,`lt_user_id`,`transaction_ctime`,`transaction_quantity`,`transaction_amount`) VALUES (45,145,5,1636391829,1,1319000);
INSERT INTO `` (`lt_transaction_id`,`lt_item_id`,`lt_user_id`,`transaction_ctime`,`transaction_quantity`,`transaction_amount`) VALUES (55,185,65,1636395048,1,5600);
INSERT INTO `` (`lt_transaction_id`,`lt_item_id`,`lt_user_id`,`transaction_ctime`,`transaction_quantity`,`transaction_amount`) VALUES (65,185,65,1636395087,1,5600);
INSERT INTO `` (`lt_transaction_id`,`lt_item_id`,`lt_user_id`,`transaction_ctime`,`transaction_quantity`,`transaction_amount`) VALUES (75,45,5,1636459773,1,64900);
INSERT INTO `` (`lt_transaction_id`,`lt_item_id`,`lt_user_id`,`transaction_ctime`,`transaction_quantity`,`transaction_amount`) VALUES (85,35,285,1636546156,1,262900);
INSERT INTO `` (`lt_transaction_id`,`lt_item_id`,`lt_user_id`,`transaction_ctime`,`transaction_quantity`,`transaction_amount`) VALUES (95,215,555,1636546669,1,700);
INSERT INTO `` (`lt_transaction_id`,`lt_item_id`,`lt_user_id`,`transaction_ctime`,`transaction_quantity`,`transaction_amount`) VALUES (105,215,545,1636546676,1,700);
INSERT INTO `` (`lt_transaction_id`,`lt_item_id`,`lt_user_id`,`transaction_ctime`,`transaction_quantity`,`transaction_amount`) VALUES (115,215,535,1636546679,1,700);
INSERT INTO `` (`lt_transaction_id`,`lt_item_id`,`lt_user_id`,`transaction_ctime`,`transaction_quantity`,`transaction_amount`) VALUES (125,215,525,1636546681,1,700);
INSERT INTO `` (`lt_transaction_id`,`lt_item_id`,`lt_user_id`,`transaction_ctime`,`transaction_quantity`,`transaction_amount`) VALUES (135,215,515,1636546685,1,700);
INSERT INTO `` (`lt_transaction_id`,`lt_item_id`,`lt_user_id`,`transaction_ctime`,`transaction_quantity`,`transaction_amount`) VALUES (145,215,505,1636546688,1,700);
INSERT INTO `` (`lt_transaction_id`,`lt_item_id`,`lt_user_id`,`transaction_ctime`,`transaction_quantity`,`transaction_amount`) VALUES (155,185,5,1636547496,1,5600);

--
-- Table structure for table `user_review_tab`
--

DROP TABLE IF EXISTS `user_review_tab`;
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

INSERT INTO `` (`review_id`,`rv_user_id`,`rv_seller_id`,`ratings`,`review_text`,`ctime`) VALUES (5,5,25,5,'Apple fan boy here. 5/5',1636386990);
INSERT INTO `` (`review_id`,`rv_user_id`,`rv_seller_id`,`ratings`,`review_text`,`ctime`) VALUES (15,5,15,5,'Elon stonksss',1636389671);
INSERT INTO `` (`review_id`,`rv_user_id`,`rv_seller_id`,`ratings`,`review_text`,`ctime`) VALUES (25,65,45,5,'highly recommended',1636395058);
INSERT INTO `` (`review_id`,`rv_user_id`,`rv_seller_id`,`ratings`,`review_text`,`ctime`) VALUES (35,555,565,5,'Best food ever',1636546964);
INSERT INTO `` (`review_id`,`rv_user_id`,`rv_seller_id`,`ratings`,`review_text`,`ctime`) VALUES (45,405,565,5,'5/5 for sure!! yummy',1636546982);
INSERT INTO `` (`review_id`,`rv_user_id`,`rv_seller_id`,`ratings`,`review_text`,`ctime`) VALUES (55,395,565,2,'Unhealthy AF',1636546998);
INSERT INTO `` (`review_id`,`rv_user_id`,`rv_seller_id`,`ratings`,`review_text`,`ctime`) VALUES (65,355,565,2,'Mummy say cannot eat fast food',1636547018);
INSERT INTO `` (`review_id`,`rv_user_id`,`rv_seller_id`,`ratings`,`review_text`,`ctime`) VALUES (75,305,565,5,'Noice noice',1636547042);
INSERT INTO `` (`review_id`,`rv_user_id`,`rv_seller_id`,`ratings`,`review_text`,`ctime`) VALUES (85,375,565,5,'BEST',1636547053);
INSERT INTO `` (`review_id`,`rv_user_id`,`rv_seller_id`,`ratings`,`review_text`,`ctime`) VALUES (95,295,565,5,':):)',1636547066);
INSERT INTO `` (`review_id`,`rv_user_id`,`rv_seller_id`,`ratings`,`review_text`,`ctime`) VALUES (105,5,45,5,'expensive but noice',1636547507);

--
-- Table structure for table `wallet_tab`
--

DROP TABLE IF EXISTS `wallet_tab`;
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

INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (5,45526600,0,1636389645,1636547496);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (15,500,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (25,327800,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (35,0,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (45,5600,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (55,0,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (65,1100,0,1636480657,1636395087);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (75,0,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (85,0,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (95,6000,0,1636546055,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (105,0,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (115,0,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (125,5000,0,1636545757,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (135,0,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (145,6000,0,1636546048,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (155,0,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (165,0,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (175,0,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (185,0,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (195,0,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (205,1000,0,1636546062,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (215,6000,0,1636546031,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (225,0,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (235,1000,0,1636545792,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (245,0,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (255,0,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (265,0,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (275,0,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (285,238100,0,1636546137,1636546156);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (295,0,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (305,70000,0,1636546023,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (315,35000,0,1636545987,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (325,3000,0,1636545950,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (335,35000,0,1636546001,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (345,7000,0,1636545921,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (355,3000,0,1636545946,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (365,35000,0,1636545981,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (375,2000,0,1636545974,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (385,7000,0,1636545918,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (395,10000,0,1636545957,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (405,5000,0,1636545774,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (415,1000,0,1636545806,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (425,5000,0,1636545907,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (435,5000,0,1636545910,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (445,5500,0,1636545814,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (455,5000,0,1636545889,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (465,5000,0,1636545903,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (475,11900,0,1636545822,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (485,10000,0,1636545881,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (495,10000,0,1636545868,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (505,6200,0,1636545833,1636546688);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (515,9300,0,1636545783,1636546685);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (525,8300,0,1636545860,1636546681);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (535,6200,0,1636545828,1636546679);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (545,3300,0,1636545842,1636546676);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (555,4300,0,1636545767,1636546669);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (565,4200,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (575,0,0,NULL,NULL);
INSERT INTO `` (`w_user_id`,`wallet_balance`,`wallet_status`,`last_top_up`,`last_used`) VALUES (585,0,0,NULL,NULL);

--
-- Table structure for table `wallet_transactions_tab`
--

DROP TABLE IF EXISTS `wallet_transactions_tab`;
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

INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (5,5,1636385385,500000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (15,5,1636386853,64900,1,15);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (25,5,1636389645,100000000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (35,5,1636389658,52200000,1,25);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (45,5,1636391811,1319000,1,35);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (55,5,1636391829,1319000,1,45);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (65,65,1636395036,10000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (75,65,1636395048,5600,1,55);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (85,65,1636395080,1200,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (95,65,1636395087,5600,1,65);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (105,5,1636459773,64900,1,75);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (115,65,1636480656,1100,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (125,125,1636545757,5000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (135,475,1636545763,5000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (145,555,1636545767,5000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (155,405,1636545774,5000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (165,515,1636545783,10000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (175,235,1636545792,1000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (185,415,1636545806,1000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (195,445,1636545814,5500,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (205,475,1636545822,6900,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (215,535,1636545828,6900,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (225,505,1636545833,6900,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (235,545,1636545842,4000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (245,525,1636545860,9000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (255,495,1636545868,10000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (265,485,1636545881,10000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (275,455,1636545889,5000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (285,465,1636545903,5000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (295,425,1636545907,5000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (305,435,1636545910,5000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (315,385,1636545918,7000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (325,345,1636545921,7000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (335,355,1636545946,3000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (345,325,1636545950,3000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (355,395,1636545957,10000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (365,375,1636545974,2000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (375,365,1636545981,35000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (385,315,1636545987,35000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (395,335,1636546001,35000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (405,305,1636546023,70000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (415,215,1636546031,6000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (425,145,1636546048,6000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (435,95,1636546055,6000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (445,205,1636546062,1000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (455,285,1636546068,1000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (465,285,1636546137,500000,0,NULL);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (475,285,1636546156,262900,1,85);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (485,555,1636546669,700,1,95);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (495,545,1636546676,700,1,105);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (505,535,1636546679,700,1,115);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (515,525,1636546681,700,1,125);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (525,515,1636546685,700,1,135);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (535,505,1636546688,700,1,145);
INSERT INTO `` (`wt_transaction_id`,`wt_user_id`,`transaction_ctime`,`transaction_amount`,`transaction_type`,`transaction_ref`) VALUES (545,5,1636547496,5600,1,155);

--
-- Dumping routines for database 'heroku_bdc39d4687a85d4'
--
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

DELIMITER ;;
CREATE DEFINER=`b0bc6fadb8432d`@`%` PROCEDURE `create_user`(IN username VARCHAR(255), IN useremail VARCHAR(255), IN userpassword VARCHAR(255), OUT status INT(1))
BEGIN
DECLARE _rollback BOOL DEFAULT 0;

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