package commons
import (
	"net/http"
	"wislabs.wifi.manager/utils"
	"strconv"
	"strings"
	"gopkg.in/gorp.v1"
	log "github.com/Sirupsen/logrus"
)

/**
 * Perform the SQL queries needed for server-side processing requested,
 * utilising the helper functions limit(), order() and
 * filter() among others. The returned array(result) is ready to be encoded as JSON
 * in response to an SSP request, or can be modified if needed before
 * sending back to the client.
 *
 *  @param  request http request Data sent to server by DataTables
 *  @param  database
 *  @param  table
 *  @param  string totalRecordCountQuery query that should be used to get the total records count of the table
 *  @param  columns  array, elements in the order it appears in DataTables
 *  @return filteredRecordCount  search result count
 *  @return totalRecordsCount  total number of records which the filtering applied
 *  @return error if an error occurred
 */
func Fetch(request *http.Request, database string, table string, totalRecordCountQuery string, columns []string, result interface{}) (int64, int64, error) {
	dbMap := utils.GetDBConnection(database);
	defer dbMap.Db.Close()
	var err error
	query := "SELECT "

	for index, element := range columns {
		query += element
		if (index + 1 != len(columns)) {
			query += ","
		}
	}
	constructedFilterQuery := ""
	constructedFilterQuery += " FROM " + table
	constructedFilterQuery += filter(request, columns)
	query += constructedFilterQuery
	query += order(request, columns)
	query += limit(request)

	_, err = dbMap.Select(result, query)

	filteredRecordCount, _ := getRecordCount(dbMap, "SELECT COUNT(*) " + constructedFilterQuery)
	totalRecordsCount, _ := getRecordCount(dbMap, totalRecordCountQuery)
	return filteredRecordCount, totalRecordsCount, err
}

/**
 * Paging
 *
 * Construct the LIMIT clause for server-side processing SQL query
 *
 *  @param  request http request Data sent to server by DataTables
 *  @param  columns array, elements in the order it appears in DataTables
 *  @return string SQL limit clause
 */
func limit(request *http.Request) string {
	limit := "LIMIT 1,100" // default limit
	if (len(request.FormValue("start")) > 0 && len(request.FormValue("length")) > 0) {
		start := request.FormValue("start")
		length := request.FormValue("length")
		limit = " LIMIT " + start + "," + length
	}
	return limit
}

/**
 * Ordering
 *
 * Construct the ORDER BY clause for server-side processing SQL query
 *
 *  @param  request http request Data sent to server by DataTables
 *  @param  columns array, elements in the order it appears in DataTables
 *  @return string SQL order by clause
 */
func order(request *http.Request, columns[]string) string {
	orderingColumn := request.FormValue("order[0][column]")
	orderingDirection := request.FormValue("order[0][dir]")

	order := " ORDER BY " + columns[0] + " ASC "
	if (len(orderingColumn) > 0 && len(orderingDirection) > 0) {
		oc, _ := strconv.Atoi(orderingColumn)
		order = " ORDER BY " + columns[oc] + " " + strings.ToUpper(orderingDirection) + " "
	}
	return order
}

/**
 * Searching / Filtering
 *
 * Construct the WHERE clause for server-side processing SQL query.
 *
 * NOTE this does not match the built-in DataTables filtering which does it
 * word by word on any field. It's possible to do here performance on large
 * databases would be very poor
 *
 *  @param  request http request Data sent to server by DataTables
 *  @param  columns array, elements in the order it appears in DataTables
 *  @return string SQL where clause
 */
func filter(request *http.Request, columns[]string) string {
	filter := " WHERE "
	var filters []string
	for i := 0; i < len(columns); i++ {
		columnSearchValue := request.FormValue("columns[" + strconv.Itoa(i) + "][search][value]")
		columnData := request.FormValue("columns[" + strconv.Itoa(i) + "][data]")
		if (len(columnSearchValue) > 0) {
			filters = append(filters, columnData + " LIKE '" + columnSearchValue + "%' ")
		}
	}
	for index, element := range filters {
		filter += element
		if (index + 1 != len(filters)) {
			filter += " AND "
		}
	}
	if (filter != " WHERE ") {
		return filter
	}
	return "" //return an empty filter if no [search][value] present
}

/**
 * Records count
 * Returns the records count using the SQL query passed
 *
 *  @param  gorp DBmap
 *  @return int count
 */
func getRecordCount(dbMap *gorp.DbMap, query string) (int64, error) {
	var recordsCount int64
	smtOut, err := dbMap.Db.Prepare(query)
	defer smtOut.Close()
	err = smtOut.QueryRow().Scan(&recordsCount)
	if err != nil {
		log.Error("Error occured while getting the records count for Query : " + query + (err.Error()))
	}
	return recordsCount, err
}