
--summarize daily stats
-- username,count, total-sessiontime,maximum-session-time, min-session-time, input, output,
SELECT UserName,COUNT(*),SUM(AcctSessionTime),MAX(AcctSessionTime),MIN(AcctSessionTime),SUM(AcctInputOctets),SUM(AcctOutputOctets),NASIPAddress FROM radacct
WHERE AcctStopTime >= '2015-08-15 05:28:51' AND AcctStopTime <= '2015-08-16 15:28:47' GROUP BY UserName,NASIPAddress;

--Total Active sessions


-- Delete sessions from the radacct table which are older than
-- $back_days

$date = NOW() - $back-days
LOCK TABLES radacct WRITE
DELETE FROM radacct WHERE AcctStopTime < '$date' AND AcctStopTime IS NOT NULL ;
UNLOCK TABLES;


-- Clean stale open sessions from the radacct table.
-- we only clean up sessions which are older than $back_days
DELETE FROM $sql_accounting_table WHERE AcctStopTime IS NULL AND AcctStartTime < '$date'

---
-- Portal DB
---

--get total number of unique users
SELECT COUNT(DISTINCT username) FROM accounting

-- users can connect through different APS different hotspots
