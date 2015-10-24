
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";

--
-- Database: `portal`
--
CREATE DATABASE IF NOT EXISTS `dashboard` DEFAULT CHARACTER SET latin1 COLLATE latin1_swedish_ci;
USE `dashboard`;

-- --------------------------------------------------------
--
-- Table structure for table `accounting`
--

DROP TABLE IF EXISTS `login`;

CREATE TABLE IF NOT EXISTS `dashboarduser` (
  `username` varchar(255) DEFAULT NULL,
  `password` VARCHAR(255) DEFAULT NULL,
  `email` VARCHAR(255) DEFAULT NULL,
  `activated` BIT DEFAULT 0,
  PRIMARY KEY(`username`)
) ENGINE=InnoDB;

-- --------------------------------------------------------

INSERT INTO dashboarduser (username, password, email)
VALUES('anu','anu123', 'anuruddha@gmail.com');



--
-- INSERT int(11)O radcheck (username,attribute,op,value) VALUES ('test','CLeartext-Password',':=','test')
--