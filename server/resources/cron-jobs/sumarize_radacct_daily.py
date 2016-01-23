import mysql.connector
import logging
import os
import time
from datetime import date, timedelta

dbuser = 'root'
dbpass = 'root'
dbhost = 'localhost'

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
    logger = logging.getLogger('sumarize_radacct_todaily')
    hdlr = logging.FileHandler('daily.log')
    formatter = logging.Formatter('%(asctime)s %(levelname)s %(message)s')
    hdlr.setFormatter(formatter)
    logger.addHandler(hdlr)
    logger.setLevel(logging.INFO)

    today = date.today()
    from_ = today - timedelta(days=100)
    to = today.isoformat()
    initLocationDictionary()
    initGroupsDictionary()

    # updateLocationGroups((today-timedelta(days=1)).isoformat())
    dumpExistingData('portal')

    summarizeBrowserStats(from_, to, 1)
    summarizeOSStats(from_, to, 1)
    summarizeDevicesStats(from_, to, 1)
    cleanUserAgentInfo(to)
    logger.info("starting daily cron job ...")

    # conn = mysql.connector.connect(host="localhost", user="root", passwd="root", db="summary")
    # cursor = conn.cursor()
    #
    # try:
    #     cursor.callproc('sumarize_radacct_todaily')
    # except mysql.connector.Error, e:
    #     logger.error("Error while performing daily corn job %s", str(e))
    #     conn.rollback()
    # finally:
    #     cursor.close()
    #     conn.close()

    logger.info("daily cron job stopped...")


def dumpExistingData(database):
    filestamp = time.strftime('%Y-%m-%d-%I:%M')
    os.popen("mysqldump -u %s -p%s -h %s -e --opt -c %s | gzip -c > %s.gz" % (
    dbuser, dbpass, dbhost, database, database + "_" + filestamp))


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

            insertQuery = "INSERT INTO  browserstats (date, tenantid, groupname,chrome, firefox,ie,iemobile,kindle,safari,safarimobile,opera,webkit,chromemobile,other) VALUES ('%s',%d,'%s',%d,%d,%d,%d,%d,%d,%d,%d,%d,%d,%d)" % (
                date.today().isoformat(), tenantId, groupName, browsersStats['chrome'], browsersStats['firefox'],
                browsersStats['ie'],
                browsersStats['iemobile'], browsersStats['kindle'], browsersStats['safari'],
                browsersStats['safarimobile'], browsersStats['opera'], browsersStats['webkit'],
                browsersStats['chromemobile'], browsersStats['other']
            )

            try:
                dashboardcursor.execute(insertQuery)
            except Exception, e:
                print "error occurred while summarizing browser information"
                print str(e)
                dashboardconn.rollback()

        except Exception, e:
            print str(e)
            portalconn.rollback()

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

            insertQuery = "INSERT INTO osstats (date, tenantid, groupname,android, ios, windows, linux, macos, windowsmobile, other) VALUES ('%s',%d,'%s',%d,%d,%d,%d,%d,%d,%d)" % (
                date.today().isoformat(), tenantId, groupName, osStats['android'], osStats['ios'],
                osStats['windows'], osStats['linux'], osStats['macos'], osStats['windowsmobile'], osStats['other']
            )

            try:
                dashboardcursor.execute(insertQuery)
            except Exception, e:
                print "error occurred while summarizing browser information"
                print str(e)
                dashboardconn.rollback()

        except Exception, e:
            print str(e)
            portalconn.rollback()

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

            insertQuery = "INSERT INTO devicestats (date, tenantid, groupname, mobile, tablet, smarttv, wearable, embedded, other) VALUES ('%s',%d,'%s',%d,%d,%d,%d,%d,%d)" % (
                date.today().isoformat(), tenantId, groupName, devicesStats['mobile'], devicesStats['tablet'],
                devicesStats['smarttv'], devicesStats['wearable'], devicesStats['embedded'], devicesStats['other']
            )

            try:
                summarycursor.execute(insertQuery)
            except Exception, e:
                print "error occurred while summarizing browser information"
                print str(e)
                sumaryconn.rollback()

        except Exception, e:
            print str(e)
            portalconn.rollback()

    portalconn.commit()
    portalcursor.close()
    portalconn.close()

    sumaryconn.commit()
    summarycursor.close()
    sumaryconn.close()
    return groups

def updateLocationGroups(date):
    global updatequery
    radiusconn = mysql.connector.connect(host=dbhost, user=dbuser, passwd=dbpass, db="summary")
    radiuscursor = radiusconn.cursor()
    query = "SELECT calledstationid from dailyacct WHERE date >= '%s'" % (date)
    print query
    try:
        radiuscursor.execute(query)
        result = radiuscursor.fetchall()
        for row in result:
            tmp = row[0]  # calledstationid
            values = row[0].split(':')
            if (len(values) == 2):
                updatequery = ""
                bssid = ((values[0])[:14]).upper()
                try:
                    group = groups[bssid + ':' + values[1]]
                    updatequery = "UPDATE dailyacct SET ssid='%s', calledstationmac='%s', groupname='%s' WHERE calledstationid='%s'" % (
                        values[1], values[0], group, tmp)
                    print "matched"
                except KeyError, e:
                    updatequery = "UPDATE dailyacct SET ssid='%s', calledstationmac='%s' WHERE calledstationid='%s'" % (
                        values[1], values[0], tmp)
                    print "missed %s" % bssid

                radiuscursor.execute(updatequery)
    except Exception, e:
        print('errr')
        print str(e)
        radiusconn.rollback()
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
        print "Error occurred while initializing the location dictionary"
        print str(e)
        dashboardconn.rollback()

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
        print "Error occurred while initializing the location dictionary"
        print str(e)
        dashboardconn.rollback()

    dashboardconn.commit()
    dashboardcursor.close()
    dashboardconn.close()
    return groups

if __name__ == "__main__": main()
