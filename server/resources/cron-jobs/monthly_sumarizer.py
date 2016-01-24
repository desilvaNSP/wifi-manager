import mysql.connector
import logging
import sys

def main():
    logger = logging.getLogger('sumarize_radacct_todaily')
    hdlr = logging.FileHandler('monthly.log')
    formatter = logging.Formatter('%(asctime)s %(levelname)s %(message)s')
    hdlr.setFormatter(formatter)
    logger.addHandler(hdlr)
    logger.setLevel(logging.INFO)

    logger.info("starting monthly cron job ...")
    conn = mysql.connector.connect(host= "localhost",user="root",passwd="root",db="radius_tmp")
    cursor = conn.cursor()

    try:
        cursor.callproc('sumarize_dailyacct_tomonthly')
    except mysql.connector.Error, e:
        logger.error("Error while performing monthly corn job %s",str(e))
        conn.rollback()
    finally:
        cursor.close()
        conn.close()
    logger.info("monthly cron job stopped...")

if __name__ == "__main__": main()