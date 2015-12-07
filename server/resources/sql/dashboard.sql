
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

DROP TABLE IF EXISTS `users`;

CREATE TABLE IF NOT EXISTS `users` (
  `tenantid` int(10) DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` VARCHAR(255) DEFAULT NULL,
  `email` VARCHAR(255) DEFAULT NULL,
  `status` VARCHAR(255) DEFAULT NULL,
  `lastupdatedtime` TIMESTAMP,
  PRIMARY KEY(`username`, `tenantid`)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `roles` (
  `name` varchar(255) DEFAULT NULL,
  `tenantid` int(10) DEFAULT 0,
  PRIMARY KEY(`name`, `tenantid`)
) ENGINE=InnoDB;


CREATE TABLE IF NOT EXISTS `permissions` (
  `tenantid` int(10) DEFAULT 0,
  `resourceid` int(10) DEFAULT 0,
  `action` VARCHAR(255) DEFAULT NULL,
  PRIMARY KEY(`tenantid`,`resourceid`)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `usersroles` (
  `username` varchar(255) DEFAULT NULL,
  `password` VARCHAR(255) DEFAULT NULL,
  `email` VARCHAR(255) DEFAULT NULL,
  `activate` BIT DEFAULT 0,
  PRIMARY KEY(`username`)
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS `userpermissions` (
  `username` varchar(255) DEFAULT NULL,
  `tenantid` int(10) DEFAULT 0,
  `resourceid` int(10) DEFAULT 0,
  `action` VARCHAR(255) DEFAULT NULL,
  PRIMARY KEY(`username`)
) ENGINE=InnoDB;

INSERT INTO users (tenantid, username, password, email, status)
VALUES (1,'admin','$2a$10$FesfnIBKqhH2MuF1hmss0umXNrrx28AW1E4re9OCAwib3cIOKBz3C', 'admin@wsilabs.com', 'active');

INSERT INTO roles (name,tenantid)
VALUES ('admin',1),
       ('user',1);

-- --------------------------------------------------------
--
-- INSERT int(11)O radcheck (username,attribute,op,value) VALUES ('test','CLeartext-Password',':=','test')
--