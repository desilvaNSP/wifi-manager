import mysql.connector
import logging
import os
import time
from datetime import date, timedelta

dbuser = os.environ.get('SUMMARY_DB_USERNAME', 'root')
dbpass = os.environ.get('SUMMARY_DB_PASSWORD', 'root')
dbhost = os.environ.get('SUMMARY_DB_HOST', 'localhost')
wifiserverdir = os.environ.get('WIFI_SERVER_DIR', '/home/anuruddha/git/wifi-manager/server/')
tenantid = 1

logger = ""
groups = {}
groups2 = []
updatequery = ""
browsers_ = {'chrome': 'chrome', 'firefox': 'firefox', 'ie': 'ie', 'iemobile': 'iemobile', 'kindle': 'kindle',
             'mobile safari': 'safarimobile',
             'webkit': 'webkit', 'opera': 'opera', 'android browser': 'chromemobile'}
browsersStats = {'chrome': 0, 'firefox': 0, 'ie': 0, 'kindle': 0, 'safari': 0, 'safarimobile': 0, 'iemobile': 0,
                 'webkit': 0, 'opera': 0, 'chromemobile': 0, 'other': 0}
os_ = {'android': 'android', 'windows': 'windows', 'linux': 'linux', 'ios': 'ios', 'mac os': 'macos',
       'windows phone': 'windowsmobile'}
osStats = {'android': 0, 'windows': 0, 'linux': 0, 'ios': 0, 'macos': 0,
           'windowsmobile': 0, 'other': 0}
devices_ = {'mobile': 'mobile', 'tablet': 'tablet', 'smarttv': 'smarttv', 'wearable': 'wearable',
            'embedded': 'embedded'}
devicesStats = {'mobile': 0, 'tablet': 0, 'smarttv': 0, 'wearable': 0, 'embedded': 0, 'other': 0}


def main():
    global logger
    logger = logging.getLogger('summarize_radacct_todaily')
    hdlr = logging.FileHandler(wifiserverdir + 'logs/daily.log')
    formatter = logging.Formatter('%(asctime)s %(levelname)s %(message)s')
    hdlr.setFormatter(formatter)
    logger.addHandler(hdlr)
    logger.setLevel(logging.INFO)

    today = date.today()
    from_ = (today - timedelta(days=200)).isoformat()
    to = today.isoformat()

    initLocationDictionary()
    initGroupsDictionary()

    logger.info("-----------------  Starting daily cron job  --------------------")

    logger.info("Dumping portal database")
    dumpExistingData('portal')

    logger.info("Dumping radius database")
    dumpExistingData('radius')

    logger.info("Starting Browser stat summarizer    [START]")
    summarizeBrowserStats(from_, to, tenantid)
    logger.info("Browser stat summary completed      [OK]")

    logger.info("Starting OS stat summarizer         [START]")
    summarizeOSStats(from_, to, tenantid)
    logger.info("OS stat summary completed           [OK]")

    logger.info("Starting Device stat summarizer     [START]")
    summarizeDevicesStats(from_, to, tenantid)
    logger.info("Device stat summary completed       [OK]")

    logger.info("Starting cleaning useragentinfo     [START]")
    cleanUserAgentInfo(to)
    logger.info("Cleaning useragentinfo completed    [OK]")

    logger.info("Starting Accounting summarizer      [START]")
    conn = mysql.connector.connect(host=dbhost, user=dbuser, passwd=dbpass, db="summary")
    cursor = conn.cursor()

    logger.info("Starting summarizing procedure      [START]")
    try:
        cursor.callproc('summarize_radacct_todaily')
    except mysql.connector.Error, e:
        conn.rollback()
        logger.error("Error occurred executing account summarizing procedure cron job : %s" % str(e))
        logger.info("----------------- Daily cron job stopped [FAILED] ---------------------")
        raise
    conn.commit()
    cursor.close()
    conn.close()

    logger.info("Summarize procedure completed       [OK]")

    logger.info("Starting location group update      [START]")
    updateLocationGroups(from_)
    logger.info("Location group update completed     [OK]")

    logger.info("Accounting summarizer completed    [OK]")
    logger.info("----------------- Daily cron job stopped [PASS] ---------------------")


def dumpExistingData(database):
    filestamp = time.strftime('%Y-%m-%d-%I:%M')
    os.popen("mysqldump -u %s -p%s -h %s -e --opt -c %s | gzip -c > %s.gz" % (
        dbuser, dbpass, dbhost, database, wifiserverdir + "/resources/cron-jobs/" + database + "_" + filestamp))


def cleanUserAgentInfo(date):
    portalconn = mysql.connector.connect(host=dbhost, user=dbuser, passwd=dbpass, db="portal")
    portalcursor = portalconn.cursor()
    query = "DELETE FROM useragentinfo WHERE date<='%s'" % (date)
    try:
        portalcursor.execute(query)
    except Exception, e:
        print "Error occurred while cleaning the useragentinfo table"
        print str(e)
        portalconn.rollback()

    portalconn.commit()
    portalcursor.close()
    portalconn.close()


def summarizeBrowserStats(from_, to, tenantId):
    portalconn = mysql.connector.connect(host=dbhost, user=dbuser, passwd=dbpass, db="portal")
    portalcursor = portalconn.cursor()

    dashboardconn = mysql.connector.connect(host=dbhost, user=dbuser, passwd=dbpass, db="summary")
    dashboardcursor = dashboardconn.cursor()

    for groupName in groups2:
        summarizeQuery = "SELECT LOWER(browser), count(*) AS num FROM useragentinfo WHERE tenantid=%d AND groupname='%s' AND date >='%s' AND date <'%s' GROUP BY browser" % (
            tenantId, groupName, from_, to)
        try:
            portalcursor.execute(summarizeQuery)
            result = portalcursor.fetchall()

            for row in result:
                browser = row[0].lower()
                count = row[1]
                if browser in browsers_:
                    browsersStats[browsers_[browser]] = count
                else:
                    browsersStats['other'] = count

            insertQuery = """INSERT INTO  browserstats (date, tenantid, groupname,chrome, firefox,ie,iemobile, kindle,
                             safari,safarimobile,opera,webkit,chromemobile,other)
                             VALUES ('%s',%d,'%s',%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d)ON DUPLICATE KEY UPDATE chrome = chrome + VALUES(chrome), firefox = firefox + VALUES(firefox),
                              ie = ie + VALUES(ie), iemobile = iemobile + VALUES(iemobile), kindle = kindle + VALUES(kindle),
                              safari = safari + VALUES(safari), safarimobile = safarimobile + VALUES(safarimobile),
                              opera = opera + VALUES(opera), webkit = webkit + VALUES(webkit),
                              chromemobile = chromemobile + VALUES(chromemobile),
                              other = other + VALUES(other)""" % (
                date.today().isoformat(), tenantId, groupName, browsersStats['chrome'], browsersStats['firefox'],
                browsersStats['ie'],
                browsersStats['iemobile'], browsersStats['kindle'], browsersStats['safari'],
                browsersStats['safarimobile'], browsersStats['opera'], browsersStats['webkit'],
                browsersStats['chromemobile'], browsersStats['other']
            )
            try:
                dashboardcursor.execute(insertQuery)
            except Exception, e:
                dashboardconn.rollback()
                raise


        except Exception, e:
            portalconn.rollback()
            logger.error("Error occurred while preparing Browser summary : %s" % str(e))
            logger.info("----------------- Daily cron job stopped [FAILED] ---------------------")
            raise

    portalconn.commit()
    portalcursor.close()
    portalconn.close()

    dashboardconn.commit()
    dashboardcursor.close()
    dashboardconn.close()
    return groups


def summarizeOSStats(from_, to, tenantId):
    portalconn = mysql.connector.connect(host=dbhost, user=dbuser, passwd=dbpass, db="portal")
    portalcursor = portalconn.cursor()

    dashboardconn = mysql.connector.connect(host=dbhost, user=dbuser, passwd=dbpass, db="summary")
    dashboardcursor = dashboardconn.cursor()

    for groupName in groups2:
        summarizeQuery = "SELECT LOWER(os), count(*) AS num FROM useragentinfo WHERE tenantid=%d AND groupname='%s' AND date >='%s' AND date <'%s' GROUP BY os" % (
            tenantId, groupName, from_, to)
        try:
            portalcursor.execute(summarizeQuery)
            result = portalcursor.fetchall()

            for row in result:
                os = row[0].lower()
                count = row[1]
                if os in os_:
                    osStats[os_[os]] = count
                else:
                    osStats['other'] = count

            insertQuery = """INSERT INTO osstats (date, tenantid, groupname,android, ios, windows, linux, macos,
                            windowsmobile, other) VALUES ('%s',%d,'%s',%d,%d,%d,%d,%d,%d,%d)
                            ON DUPLICATE KEY UPDATE android = android + VALUES(android), ios = ios + VALUES(ios),
                            windows = windows + VALUES(windows), linux = linux + VALUES(linux), macos = macos + VALUES(macos),
                            windowsmobile = windowsmobile + VALUES(windowsmobile), other = other + VALUES(other)""" % (
                date.today().isoformat(), tenantId, groupName, osStats['android'], osStats['ios'],
                osStats['windows'], osStats['linux'], osStats['macos'], osStats['windowsmobile'], osStats['other']
            )

            try:
                dashboardcursor.execute(insertQuery)
            except Exception, e:
                dashboardconn.rollback()
                raise

        except Exception, e:
            portalconn.rollback()
            logger.error("Error occurred while preparing OS summary : %s" % str(e))
            logger.info("----------------- Daily cron job stopped [FAILED] ---------------------")
            raise

    portalconn.commit()
    portalcursor.close()
    portalconn.close()

    dashboardconn.commit()
    dashboardcursor.close()
    dashboardconn.close()
    return groups


def summarizeDevicesStats(from_, to, tenantId):
    portalconn = mysql.connector.connect(host=dbhost, user=dbuser, passwd=dbpass, db="portal")
    portalcursor = portalconn.cursor()

    sumaryconn = mysql.connector.connect(host=dbhost, user=dbuser, passwd=dbpass, db="summary")
    summarycursor = sumaryconn.cursor()

    for groupName in groups2:
        summarizeQuery = "SELECT LOWER(device), count(*) AS num FROM useragentinfo WHERE tenantid=%d AND groupname='%s' AND date >='%s' AND date <'%s' GROUP BY device" % (
            tenantId, groupName, from_, to)
        try:
            portalcursor.execute(summarizeQuery)
            result = portalcursor.fetchall()

            for row in result:
                device = row[0].lower()
                count = row[1]
                if device in devices_:
                    devicesStats[devices_[device]] = count
                else:
                    devicesStats['other'] = count

            insertQuery = """INSERT INTO devicestats (date, tenantid, groupname, mobile, tablet, smarttv, wearable,
                          embedded, other) VALUES ('%s',%d,'%s',%d,%d,%d,%d,%d,%d) ON DUPLICATE KEY UPDATE
                          mobile = mobile + VALUES(mobile) , tablet = tablet + VALUES(tablet),
                          smarttv = smarttv + VALUES(smarttv), wearable = wearable + VALUES(wearable),
                          other = other + VALUES(other)""" % (
                date.today().isoformat(), tenantId, groupName, devicesStats['mobile'], devicesStats['tablet'],
                devicesStats['smarttv'], devicesStats['wearable'], devicesStats['embedded'], devicesStats['other']
            )

            try:
                summarycursor.execute(insertQuery)
            except Exception, e:
                sumaryconn.rollback()
                raise

        except Exception, e:
            portalconn.rollback()
            logger.error("Error occurred while preparing Device summary : %s" % str(e))
            logger.info("----------------- Daily cron job stopped [FAILED] ---------------------")
            raise

    portalconn.commit()
    portalcursor.close()
    portalconn.close()

    sumaryconn.commit()
    summarycursor.close()
    sumaryconn.close()
    return groups


# TODO : replace key error with 'in' check
def updateLocationGroups(date):
    global updatequery
    radiusconn = mysql.connector.connect(host=dbhost, user=dbuser, passwd=dbpass, db="summary")
    radiuscursor = radiusconn.cursor()
    query = "SELECT calledstationid from dailyacct WHERE date >= '%s'" % (date)
    try:
        radiuscursor.execute(query)
        result = radiuscursor.fetchall()
        for row in result:
            tmp = row[0]  # calledstationid
            values = row[0].split(':')
            # ZD, SZ case
            if (len(values) == 2):
                updatequery = ""
                bssid = ((values[0])[:14]).upper()
                try:
                    group = groups[bssid + ':' + values[1]]
                    updatequery = "UPDATE dailyacct SET ssid='%s', calledstationmac='%s', groupname='%s' WHERE calledstationid='%s' AND date >= '%s'" % (
                        values[1], values[0], group, tmp, date)
                except KeyError, e:
                    updatequery = "UPDATE dailyacct SET ssid='%s', calledstationmac='%s' WHERE calledstationid='%s' AND date >= '%s'" % (
                        values[1], values[0], tmp, date)
            else: # MKT case
                updatequery = "UPDATE dailyacct SET groupname='%s' WHERE calledstationid='%s' AND date >= '%s'" % (
                    group, tmp, date)
            radiuscursor.execute(updatequery)
    except Exception, e:
        radiusconn.rollback()
        logger.error("Error occurred while updating location groups : %s" % str(e))
        logger.info("----------------- Daily cron job stopped [FAILED] ---------------------")
        raise
    radiusconn.commit()
    radiuscursor.close()
    radiusconn.close()


def initGroupsDictionary():
    dashboardconn = mysql.connector.connect(host=dbhost, user=dbuser, passwd=dbpass, db="dashboard")
    dashboardcursor = dashboardconn.cursor()
    query = "SELECT groupname from apgroups WHERE tenantid=%d GROUP BY groupname, tenantid" % (1)

    global groups2
    try:
        dashboardcursor.execute(query)
        result = dashboardcursor.fetchall()
        for row in result:
            groups2.append(row[0])

    except Exception, e:
        dashboardconn.rollback()
        logger.error("Error occurred initializing group dictionary : %s" % str(e))
        logger.info("----------------- Daily cron job stopped [FAILED] ---------------------")
        raise

    dashboardconn.commit()
    dashboardcursor.close()
    dashboardconn.close()

    return groups


def initLocationDictionary():
    dashboardconn = mysql.connector.connect(host=dbhost, user=dbuser, passwd=dbpass, db="dashboard")
    dashboardcursor = dashboardconn.cursor()
    query = "SELECT mac, ssid, groupname from aplocations"

    global groups
    try:
        dashboardcursor.execute(query)
        result = dashboardcursor.fetchall()
        for row in result:
            key = ((row[0])[:14]).upper() + ':' + row[1]
            groups[key] = row[2]

    except Exception, e:
        dashboardconn.rollback()
        logger.error("Error occurred while initializing location dictionary : %s" % str(e))
        logger.info("----------------- Daily cron job stopped [FAILED] ---------------------")
        raise

    dashboardconn.commit()
    dashboardcursor.close()
    dashboardconn.close()
    return groups


if __name__ == "__main__": main()
