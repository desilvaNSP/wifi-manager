SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";

--
-- Database: `dashboard`
--
CREATE DATABASE IF NOT EXISTS `dashboard`
  DEFAULT CHARACTER SET latin1
  COLLATE latin1_swedish_ci;
USE `dashboard`;

-- --------------------------------------------------------
--
-- Table structures for dashboard
--
CREATE TABLE IF NOT EXISTS `tenants` (
  `tenantid`  INT NOT NULL AUTO_INCREMENT,
  `domain`    VARCHAR(255) DEFAULT NULL,
  `status`    VARCHAR(255) DEFAULT NULL,
  `createdon` TIMESTAMP,
  PRIMARY KEY (`tenantid`)
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `users` (
  `userid`          BIGINT NOT NULL AUTO_INCREMENT,
  `tenantid`        INT,
  `username`        VARCHAR(255)    DEFAULT NULL,
  `password`        VARCHAR(255)    DEFAULT NULL,
  `email`           VARCHAR(255)    DEFAULT NULL,
  `status`          VARCHAR(255)    DEFAULT NULL,
  `lastupdatedtime` TIMESTAMP,
  PRIMARY KEY (`userid`),
  FOREIGN KEY (tenantid) REFERENCES tenants (tenantid)
    ON DELETE CASCADE
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `permissions` (
  `permissionid` BIGINT NOT NULL AUTO_INCREMENT,
  `tenantid`     INT,
  `name`         VARCHAR(255)    DEFAULT NULL,
  `action`       VARCHAR(255)    DEFAULT NULL,
  PRIMARY KEY (`permissionid`),
  FOREIGN KEY (tenantid) REFERENCES tenants (tenantid)
    ON DELETE CASCADE
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `userpermissions` (
  `permissionid` BIGINT,
  `userid`       BIGINT,
  PRIMARY KEY (`permissionid`, `userid`),
  FOREIGN KEY (userid) REFERENCES users (userid)
    ON DELETE CASCADE,
  FOREIGN KEY (permissionid) REFERENCES permissions (permissionid)
    ON DELETE CASCADE
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `apgroups` (
  `tenantid`    INT,
  `groupid`     BIGINT NOT NULL AUTO_INCREMENT,
  `groupname`   VARCHAR(255)    DEFAULT NULL,
  `groupsymbol` VARCHAR(255)    DEFAULT NULL,
  PRIMARY KEY (`groupid`, `tenantid`),
  FOREIGN KEY (`tenantid`) REFERENCES tenants (tenantid)
    ON DELETE CASCADE
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `aplocations` (
  `tenantid`   INT,
  `locationid` BIGINT       NOT NULL AUTO_INCREMENT,
  `ssid`       VARCHAR(255) NOT NULL,
  `mac`        VARCHAR(255)          DEFAULT NULL,
  `bssid`      VARCHAR(255)          DEFAULT NULL,
  `longitude`  FLOAT,
  `latitude`   FLOAT,
  `groupid`    BIGINT,
  `groupname`  VARCHAR(255) NOT NULL,
  PRIMARY KEY (`locationid`, `ssid`, `mac`),
  FOREIGN KEY (tenantid) REFERENCES tenants (tenantid)
    ON DELETE CASCADE,
  FOREIGN KEY (groupid) REFERENCES apgroups (groupid)
    ON DELETE CASCADE
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `userapgroups` (
  `userid`  BIGINT,
  `groupid` BIGINT,
  PRIMARY KEY (`groupid`, `userid`),
  FOREIGN KEY (userid) REFERENCES users (userid)
    ON DELETE CASCADE,
  FOREIGN KEY (groupid) REFERENCES apgroups (groupid)
    ON DELETE CASCADE
)
  ENGINE = InnoDB;

--
-- Metrics
--
CREATE TABLE IF NOT EXISTS `metrics` (
  `tenantid` INT,
  `metricid` INT NOT NULL AUTO_INCREMENT,
  `name`     VARCHAR(255) DEFAULT NULL,
  PRIMARY KEY (`metricid`),
  FOREIGN KEY (tenantid) REFERENCES users (tenantid)
    ON DELETE CASCADE
)
  ENGINE = InnoDB;

--
-- applications
--
CREATE TABLE IF NOT EXISTS `apps` (
  `appid`     INT NOT NULL AUTO_INCREMENT,
  `tenantid`  INT,
  `name`      VARCHAR(255) DEFAULT NULL,
  `aggregate` VARCHAR(255) DEFAULT NULL,
  `createdon` TIMESTAMP,
  PRIMARY KEY (`appid`),
  FOREIGN KEY (tenantid) REFERENCES users (tenantid)
    ON DELETE CASCADE
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `appgroups` (
  `appid`     INT,
  `groupname` VARCHAR(255) DEFAULT NULL,
  FOREIGN KEY (appid) REFERENCES apps (appid)
    ON DELETE CASCADE
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `appusers` (
  `tenantid` INT,
  `appid`    INT,
  `username` VARCHAR(255) DEFAULT NULL,
  FOREIGN KEY (appid) REFERENCES apps (appid)
    ON DELETE CASCADE,
  FOREIGN KEY (tenantid) REFERENCES users (tenantid)
    ON DELETE CASCADE
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `appmetrics` (
  `appid`    INT,
  `metricid` INT,
  FOREIGN KEY (appid) REFERENCES apps (appid)
    ON DELETE CASCADE
)
  ENGINE = InnoDB;
