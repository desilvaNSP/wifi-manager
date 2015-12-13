
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";

--
-- Database: `portal`
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
