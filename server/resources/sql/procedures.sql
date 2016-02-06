USE `summary`;
DROP PROCEDURE IF EXISTS clean_radacct;
DROP PROCEDURE IF EXISTS clean_dailyacct;
DROP PROCEDURE IF EXISTS summarize_radacct_todaily;
DROP PROCEDURE IF EXISTS sumarize_dailyacct_tomonthly;

DELIMITER //
CREATE PROCEDURE `clean_radacct`(cleandate DATE)
  BEGIN
    START TRANSACTION;
    DELETE FROM radius.radacct
    WHERE AcctStopTime IS NULL OR AcctStopTime <= (cleandate + INTERVAL 1 DAY);
    COMMIT;
  END//

DELIMITER //
CREATE PROCEDURE `clean_dailyacct`(cleandate DATE)
  BEGIN
    START TRANSACTION;
    DELETE FROM summary.dailyacct
    WHERE date <= (cleandate + INTERVAL 2 DAY);
    COMMIT;
  END//

DELIMITER //
CREATE PROCEDURE `summarize_radacct_todaily`()
  BEGIN
    DECLARE startdate DATE DEFAULT NOW();
    DECLARE enddate DATE DEFAULT NOW();

    START TRANSACTION;

    DELETE FROM radius.radacct
    WHERE AcctStopTime IS NULL;

    SELECT MIN(acctstarttime)
    INTO startdate
    FROM radius.radacct;
    SELECT MAX(acctstoptime)
    INTO enddate
    FROM radius.radacct;

    WHILE startdate <= enddate
    DO
      INSERT INTO summary.dailyacct (tenantid,
                                     username,
                                     date,
                                     noofsessions,
                                     totalsessionduration,
                                     sessionmaxduration,
                                     sessionminduration,
                                     sessionavgduration,
                                     inputoctets,
                                     outputoctets,
                                     nasipaddress,
                                     framedipaddress,
                                     calledstationid)
        SELECT
          1,
          UserName,
          startdate,
          COUNT(*),
          SUM(AcctSessionTime),
          MAX(AcctSessionTime),
          MIN(AcctSessionTime),
          AVG(AcctSessionTime),
          SUM(AcctInputOctets),
          SUM(AcctOutputOctets),
          NASIPAddress,
          framedipaddress,
          calledstationid
        FROM radius.radacct
        WHERE AcctStopTime >= startdate AND AcctStopTime < (startdate + INTERVAL 1 DAY)
        GROUP BY UserName, calledstationid, nasipaddress
      ON DUPLICATE KEY UPDATE noofsessions = noofsessions + VALUES(noofsessions),
        totalsessionduration               = totalsessionduration + VALUES(totalsessionduration),
        inputoctets                        = inputoctets + VALUES(inputoctets),
        outputoctets                       = outputoctets + VALUES(outputoctets);
      SET startdate = DATE_ADD(startdate, INTERVAL 1 DAY);
    END WHILE;
    COMMIT;
    CALL clean_radacct(enddate);
  END//

DELIMITER //
CREATE PROCEDURE `sumarize_dailyacct_tomonthly`()
  BEGIN
    DECLARE startdate DATE DEFAULT NOW();
    DECLARE enddate DATE DEFAULT NOW();

    START TRANSACTION;

    SELECT MIN(date)
    INTO startdate
    FROM summary.dailyacct;
    SELECT MAX(date)
    INTO enddate
    FROM summary.dailyacct;
    WHILE startdate <= enddate
    DO
      INSERT INTO summary.monthlyacct (tenantid,
                                       username,
                                       date,
                                       noofsessions,
                                       totalsessionduration,
                                       sessionmaxduration,
                                       sessionminduration,
                                       sessionavgduration,
                                       inputoctets,
                                       outputoctets,
                                       nasipaddress,
                                       framedipaddress,
                                       location)
        SELECT
          1,
          UserName,
          startdate,
          COUNT(*),
          SUM(totalsessionduration),
          MAX(sessionmaxduration),
          MIN(sessionminduration),
          AVG(sessionavgduration),
          SUM(inputoctets),
          SUM(outputoctets),
          NASIPAddress,
          framedipaddress,
          location
        FROM summary.dailyacct
        WHERE `date` >= startdate AND `date` < (startdate + INTERVAL 30 DAY)
        GROUP BY UserName, NASIPAddress;
      SET startdate = DATE_ADD(startdate, INTERVAL 30 DAY);
    END WHILE;
    COMMIT;
    CALL clean_dailyacct(enddate);
  END//

