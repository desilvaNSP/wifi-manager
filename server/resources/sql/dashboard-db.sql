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
  `domain`    VARCHAR(255) DEFAULT NULL UNIQUE,
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

CREATE TABLE IF NOT EXISTS `userssids` (
  `userid` BIGINT,
  `ssid`   VARCHAR(255) DEFAULT NULL,
  PRIMARY KEY (`userid`, `ssid`),
  FOREIGN KEY (userid) REFERENCES users (userid)
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
  FOREIGN KEY (tenantid) REFERENCES tenants (tenantid)
    ON DELETE CASCADE
)
  ENGINE = InnoDB;

--
-- Radius NAS
--
CREATE TABLE IF NOT EXISTS radiusservers (
  InsId        INT(11) NOT NULL AUTO_INCREMENT,
  tenantid     INT(11)          DEFAULT NULL,
  dbhostname   VARCHAR(45)      DEFAULT NULL,
  dbhostip     VARCHAR(45)      DEFAULT NULL,
  dbschemaname VARCHAR(45)      DEFAULT NULL,
  dbport       INT(11)          DEFAULT NULL,
  dbusername   VARCHAR(45)      DEFAULT NULL,
  dbpassword   VARCHAR(45)      DEFAULT NULL,
  status       VARCHAR(45)      DEFAULT NULL,
  PRIMARY KEY (InsId)
)
  ENGINE = InnoDB
  AUTO_INCREMENT = 0;
--
-- applications
--
CREATE TABLE IF NOT EXISTS `apps` (
  `appid`          INT NOT NULL AUTO_INCREMENT,
  `tenantid`       INT,
  `name`           VARCHAR(255) DEFAULT NULL,
  `aggregate`      VARCHAR(255) DEFAULT NULL,
  `filtercriteria` VARCHAR(255) DEFAULT NULL,
  `createdon`      TIMESTAMP,
  PRIMARY KEY (`appid`),
  FOREIGN KEY (tenantid) REFERENCES tenants (tenantid)
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
  FOREIGN KEY (tenantid) REFERENCES tenants (tenantid)
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

CREATE TABLE IF NOT EXISTS `appfilterparams` (
  `appid`     INT,
  `parameter` VARCHAR(255) DEFAULT NULL,
  FOREIGN KEY (appid) REFERENCES apps (appid)
    ON DELETE CASCADE
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS appacls (
  appid INT,
  acl   VARCHAR(255) DEFAULT NULL,
  FOREIGN KEY (appid) REFERENCES apps (appid)
    ON DELETE CASCADE
)
  ENGINE = InnoDB;

-- Inserting default data set
INSERT IGNORE INTO tenants (domain, status)
VALUES ('isl.com', 'active');

INSERT IGNORE INTO users (tenantid, username, password, email, status)
VALUES (1, 'admin', '$2a$10$FesfnIBKqhH2MuF1hmss0umXNrrx28AW1E4re9OCAwib3cIOKBz3C', 'admin@isl.com', 'active');

INSERT IGNORE INTO permissions (permissionid, tenantid, name, action)
VALUES
  (1, 1, 'wifi_location', 'read'),
  (2, 1, 'wifi_location', 'write'),
  (3, 1, 'wifi_location', 'execute'),
  (4, 1, 'wifi_users', 'read'),
  (5, 1, 'wifi_users', 'write'),
  (6, 1, 'wifi_users', 'execute'),
  (7, 1, 'dashboard_users', 'read'),
  (8, 1, 'dashboard_users', 'write'),
  (9, 1, 'dashboard_users', 'execute');

INSERT IGNORE INTO userpermissions (userid, permissionid)
VALUES (1, 1),
  (1, 2),
  (1, 3),
  (1, 4),
  (1, 5),
  (1, 6),
  (1, 7),
  (1, 8),
  (1, 9);

