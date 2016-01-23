-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.1.10-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             9.1.0.4867
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- Dumping database structure for allowance
CREATE DATABASE IF NOT EXISTS `allowance` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `allowance`;


-- Dumping structure for table allowance.buckets
CREATE TABLE IF NOT EXISTS `buckets` (
  `BucketID` int(11) NOT NULL AUTO_INCREMENT,
  `KidID` int(11) NOT NULL,
  `Name` varchar(50) NOT NULL,
  `DefaultAllocation` int(11) NOT NULL DEFAULT '0',
  `CurrentTotal` decimal(10,2) NOT NULL DEFAULT '0.00',
  `CreateDate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `UpdateDate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`BucketID`),
  KEY `FK__kids` (`KidID`),
  CONSTRAINT `FK__kids` FOREIGN KEY (`KidID`) REFERENCES `kids` (`KidID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Data exporting was unselected.


-- Dumping structure for table allowance.family
CREATE TABLE IF NOT EXISTS `family` (
  `FamilyID` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(50) NOT NULL,
  `CreateDate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `UpdateDate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`FamilyID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Data exporting was unselected.


-- Dumping structure for table allowance.kids
CREATE TABLE IF NOT EXISTS `kids` (
  `KidID` int(11) NOT NULL AUTO_INCREMENT,
  `FamilyID` int(11) NOT NULL,
  `Name` varchar(50) NOT NULL,
  `Email` varchar(128) DEFAULT '',
  `CreateDate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `UpdateDate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`KidID`),
  KEY `FK_kids_family` (`FamilyID`),
  CONSTRAINT `FK_kids_family` FOREIGN KEY (`FamilyID`) REFERENCES `family` (`FamilyID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Data exporting was unselected.


-- Dumping structure for table allowance.kidtokens
CREATE TABLE IF NOT EXISTS `kidtokens` (
  `Token` char(8) NOT NULL,
  `KidID` int(11) NOT NULL,
  `CreateDate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`Token`),
  KEY `FK_kidtokens_kids` (`KidID`),
  CONSTRAINT `FK_kidtokens_kids` FOREIGN KEY (`KidID`) REFERENCES `kids` (`KidID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Data exporting was unselected.


-- Dumping structure for table allowance.parents
CREATE TABLE IF NOT EXISTS `parents` (
  `ParentID` int(11) NOT NULL AUTO_INCREMENT,
  `FamilyID` int(11) NOT NULL,
  `Name` varchar(50) NOT NULL,
  `Email` varchar(128) NOT NULL,
  `PasswordHash` binary(60) NOT NULL,
  `CreateDate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `UpdateDate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ParentID`),
  UNIQUE KEY `Idx_email` (`Email`),
  KEY `FK_parents_family` (`FamilyID`),
  CONSTRAINT `FK_parents_family` FOREIGN KEY (`FamilyID`) REFERENCES `family` (`FamilyID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Data exporting was unselected.


-- Dumping structure for table allowance.transactions
CREATE TABLE IF NOT EXISTS `transactions` (
  `TransactionID` int(11) NOT NULL AUTO_INCREMENT,
  `BucketID` int(11) NOT NULL,
  `CreateParentID` int(11) NOT NULL,
  `Amount` decimal(10,2) NOT NULL,
  `Note` varchar(256) NOT NULL,
  `UUID` char(36) NOT NULL,
  `CreateDate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `UpdateDate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`TransactionID`),
  KEY `FK__buckets` (`BucketID`),
  KEY `FK__parents` (`CreateParentID`),
  CONSTRAINT `FK__buckets` FOREIGN KEY (`BucketID`) REFERENCES `buckets` (`BucketID`),
  CONSTRAINT `FK__parents` FOREIGN KEY (`CreateParentID`) REFERENCES `parents` (`ParentID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Data exporting was unselected.


-- Dumping structure for trigger allowance.transactions_after_insert
SET @OLDTMP_SQL_MODE=@@SQL_MODE, SQL_MODE='STRICT_TRANS_TABLES,NO_ENGINE_SUBSTITUTION';
DELIMITER //
CREATE TRIGGER `transactions_after_insert` AFTER INSERT ON `transactions` FOR EACH ROW BEGIN
  UPDATE buckets SET CurrentTotal = CurrentTotal+NEW.Amount 
  WHERE buckets.BucketID = NEW.BucketID;
END//
DELIMITER ;
SET SQL_MODE=@OLDTMP_SQL_MODE;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
