SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";

--
-- Database: `summary`
--
CREATE DATABASE IF NOT EXISTS `summary`
  DEFAULT CHARACTER SET latin1
  COLLATE latin1_swedish_ci;
USE `summary`;

CREATE TABLE IF NOT EXISTS dailyacct (
  `tenantid`             INT(10)      DEFAULT NULL,
  `username`             VARCHAR(255) DEFAULT NULL,
  `date`                 DATETIME     DEFAULT NULL,
  `noofsessions`         INT(11)      DEFAULT 0,
  `totalsessionduration` INT(11)      DEFAULT 0,
  `sessionmaxduration`   INT(11)      DEFAULT 0,
  `sessionminduration`   INT(11)      DEFAULT 0,
  `sessionavgduration`   INT(11)      DEFAULT 0,
  `inputoctets`          BIGINT(20)   DEFAULT 0,
  `outputoctets`         BIGINT(20)   DEFAULT 0,
  `nasipaddress`         VARCHAR(255) DEFAULT NULL,
  `framedipaddress`      VARCHAR(255) DEFAULT NULL,
  `calledstationid`      VARCHAR(255) DEFAULT NULL,
  `ssid`                 VARCHAR(255) DEFAULT NULL,
  `calledstationmac`     VARCHAR(255) DEFAULT NULL,
  `groupname`            VARCHAR(255) DEFAULT NULL,
  `locationid`           VARCHAR(255) DEFAULT NULL,
  PRIMARY KEY (`tenantid`, `username`, `date`, `calledstationid`, `nasipaddress`)
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS monthlyacct (
  `tenantid`             INT(10)      DEFAULT NULL,
  `username`             VARCHAR(255) DEFAULT NULL,
  `date`                 DATETIME     DEFAULT NULL,
  `noofsessions`         INT(11)      DEFAULT 0,
  `totalsessionduration` INT(11)      DEFAULT 0,
  `sessionmaxduration`   INT(11)      DEFAULT 0,
  `sessionminduration`   INT(11)      DEFAULT 0,
  `sessionavgduration`   INT(11)      DEFAULT 0,
  `inputoctets`          BIGINT(20)   DEFAULT 0,
  `outputoctets`         BIGINT(20)   DEFAULT 0,
  `nasipaddress`         VARCHAR(15)  DEFAULT NULL,
  `framedipaddress`      VARCHAR(15)  DEFAULT NULL,
  `groupname`            VARCHAR(255) DEFAULT NULL,
  `locationid`           VARCHAR(255) DEFAULT NULL,
  PRIMARY KEY (`tenantid`, `username`, `date`)
)
  ENGINE = InnoDB;

--
-- User agent summary
--

CREATE TABLE IF NOT EXISTS browserstats (
  `tenantid`     INT(10)      DEFAULT NULL,
  `groupname`    VARCHAR(255) DEFAULT NULL,
  `ssid`         VARCHAR(255) DEFAULT NULL,
  `date`         DATETIME     DEFAULT NULL,
  `chrome`       INT(11)      DEFAULT 0,
  `firefox`      INT(11)      DEFAULT 0,
  `ie`           INT(11)      DEFAULT 0,
  `iemobile`     INT(11)      DEFAULT 0,
  `kindle`       INT(11)      DEFAULT 0,
  `safari`       INT(11)      DEFAULT 0,
  `safarimobile` INT(11)      DEFAULT 0,
  `opera`        INT(11)      DEFAULT 0,
  `webkit`       INT(11)      DEFAULT 0,
  `chromemobile` INT(11)      DEFAULT 0,
  `other`        INT(11)      DEFAULT 0,
  PRIMARY KEY (`date`, `groupname`, `tenantid`)
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS osstats (
  `tenantid`      INT(10)      DEFAULT NULL,
  `groupname`     VARCHAR(255) DEFAULT NULL,
  `ssid`         VARCHAR(255) DEFAULT NULL,
  `date`          DATETIME     DEFAULT NULL,
  `android`       INT(11)      DEFAULT 0,
  `ios`           INT(11)      DEFAULT 0,
  `windows`       INT(11)      DEFAULT 0,
  `linux`         INT(11)      DEFAULT 0,
  `macos`         INT(11)      DEFAULT 0,
  `windowsmobile` INT(11)      DEFAULT 0,
  `other`         INT(11)      DEFAULT 0,
  PRIMARY KEY (`date`, `groupname`, `tenantid`)
)
  ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS devicestats (
  `tenantid`  INT(10)      DEFAULT NULL,
  `groupname` VARCHAR(255) DEFAULT NULL,
  `ssid`         VARCHAR(255) DEFAULT NULL,
  `date`      DATETIME     DEFAULT NULL,
  `mobile`    INT(11)      DEFAULT 0,
  `tablet`    INT(11)      DEFAULT 0,
  `smarttv`   INT(11)      DEFAULT 0,
  `wearable`  INT(11)      DEFAULT 0,
  `embedded`  INT(11)      DEFAULT 0,
  `other`     INT(11)      DEFAULT 0,
  PRIMARY KEY (`date`, `groupname`, `tenantid`)
)
  ENGINE = InnoDB;

CREATE INDEX calledstationid_date ON dailyacct(calledstationid,date);
CREATE INDEX groupname_date ON dailyacct(groupname,date);