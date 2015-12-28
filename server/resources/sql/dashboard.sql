
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";

--
-- Database: `dashboard`
--
CREATE DATABASE IF NOT EXISTS `dashboard` DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;
USE `dashboard`;

-- --------------------------------------------------------
--
-- Table structures for dashboard
--
CREATE TABLE IF NOT EXISTS `tenants` (
  `tenantid` INT NOT NULL AUTO_INCREMENT,
  `domain` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `createdon` TIMESTAMP,
  PRIMARY KEY(`tenantid`)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `users` (
  `userid` BIGINT NOT NULL AUTO_INCREMENT,
  `tenantid` INT,
  `username` varchar(255) DEFAULT NULL,
  `password` VARCHAR(255) DEFAULT NULL,
  `email` VARCHAR(255) DEFAULT NULL,
  `status` VARCHAR(255) DEFAULT NULL,
  `lastupdatedtime` TIMESTAMP,
  PRIMARY KEY(`userid`),
  FOREIGN KEY(tenantid) REFERENCES tenants(tenantid) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `permissions` (
  `permissionid` BIGINT NOT NULL AUTO_INCREMENT,
  `tenantid` INT,
  `name`  VARCHAR(255) DEFAULT NULL,
  `action` VARCHAR(255) DEFAULT NULL,
  PRIMARY KEY(`permissionid`),
  FOREIGN KEY(tenantid) REFERENCES tenants(tenantid) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `userpermissions` (
  `permissionid` BIGINT,
  `userid` BIGINT,
  PRIMARY KEY(`permissionid`, `userid`),
  FOREIGN KEY(userid) REFERENCES users(userid) ON DELETE CASCADE,
  FOREIGN KEY(permissionid) REFERENCES permissions(permissionid) ON DELETE CASCADE
) ENGINE=InnoDB;

INSERT INTO tenants (domain, email, status)
VALUES ('wislabs.com','admin@wislabs.com','active');

INSERT INTO users (tenantid, username, password, email, status)
VALUES (1,'admin','$2a$10$FesfnIBKqhH2MuF1hmss0umXNrrx28AW1E4re9OCAwib3cIOKBz3C', 'admin@wsilabs.com', 'active');

INSERT INTO permissions (tenantid, name, action)
VALUES (1, 'wifi_location', 'read'),
       (1, 'wifi_location', 'write'),
       (1, 'wifi_location', 'execute')
       ;

INSERT INTO userpermissions (userid, permissionid)
VALUES (1, 1),
       (1, 2),
       (1, 3)
       ;

CREATE TABLE IF NOT EXISTS `aplocations` (
  `locationid` BIGINT NOT NULL AUTO_INCREMENT,
  `ssid`  VARCHAR(255) NOT NULL,
  `mac` VARCHAR(255) DEFAULT NULL,
  `longitude` FLOAT ,
  `latitude` FLOAT ,
  `groupname` varchar(255) NOT NULL,
  PRIMARY KEY(`locationid`, `ssid`, `mac`)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `apgroups` (
  `groupid` BIGINT NOT NULL AUTO_INCREMENT,
  `locationid` BIGINT,
  `groupname` VARCHAR(255) DEFAULT NULL,
  PRIMARY KEY (`groupid`),
  FOREIGN KEY(`locationid`) REFERENCES aplocations(locationid) ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `useragentinfo` (
  `date` datetime DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL ,
  `locationid` BIGINT,
  `device` varchar(200) DEFAULT NULL,
  `browser` varchar(200) DEFAULT NULL,
  `os` varchar(200) DEFAULT NULL,
  `ua` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`username`, `ua`, `locationid`),
  FOREIGN KEY(locationid) REFERENCES aplocations(locationid) ON DELETE CASCADE
)ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `ipexclusion` (
  `locationid` BIGINT,
  `excludeip` varchar(255) DEFAULT NULL,
  PRIMARY KEY(`locationid`,`excludeip`),
  FOREIGN KEY(locationid) REFERENCES aplocations(locationid) ON DELETE CASCADE
) ENGINE=InnoDB;

--
-- Adding sample dataset
--

INSERT INTO aplocations (locationid, ssid, mac, groupname)
VALUES
  (1, 'Free_Darebin_Wi-Fi', 'f0:b0:52:37:ed:d0', 'Preston'),
  (2, 'Free_Darebin_Wi-Fi', 'd4:68:4d:03:83:20', 'Preston');

INSERT INTO apgroups(groupid, locationid, groupname)
VALUES
  (1, 2, 'ISL'),
  (2, 2, 'ISL');

INSERT INTO useragentinfo (date, username, locationid, device, browser, os, ua)
VALUES
  (NOW() - INTERVAL 3 MONTH, 'anu123',1,'Android','Chrome', 'Linux', 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.125 Safari/537.36'),
  (NOW() - INTERVAL 3 DAY, 'samee',2,'Apple','Chrome', 'IOS', 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.125 Safari/537.36');