package commons
import (
	"net/http"
	"wislabs.wifi.manager/utils"
	"strconv"
	"strings"
	"gopkg.in/gorp.v1"
)

func Fetch(r *http.Request, database string, table string, totalRecordCountQuery string, columns []string, result interface{}) (int64, int64) {
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
	constructedFilterQuery += filter(r, columns)
	query += constructedFilterQuery
	query += order(r, columns)
	query += limit(r)

	_, err = dbMap.Select(result, query)
	if err != nil {
		panic(err.Error())  // Just for example purpose.
	}
	filteredRecordCount := getRecordCount(dbMap, "SELECT COUNT(*) " + constructedFilterQuery)
	totalRecordsCount := getRecordCount(dbMap, totalRecordCountQuery)
	return filteredRecordCount, totalRecordsCount
}

func limit(r *http.Request) string {
	limit := "LIMIT 1,100" // default limit
	if (len(r.FormValue("start")) > 0 && len(r.FormValue("length")) > 0) {
		start := r.FormValue("start")
		length := r.FormValue("length")
		limit = " LIMIT " + start + "," + length
	}
	return limit
}

func order(r *http.Request, columns[]string) string {
	orderingColumn := r.FormValue("order[0][column]")
	orderingDirection := r.FormValue("order[0][dir]")

	order := " ORDER BY " + columns[0] + " ASC "
	if (len(orderingColumn) > 0 && len(orderingDirection) > 0) {
		oc, _ := strconv.Atoi(orderingColumn)
		order = " ORDER BY " + columns[oc] + " " + strings.ToUpper(orderingDirection) + " "
	}
	return order
}

func filter(r *http.Request, columns[]string) string {
	filter := " WHERE "
	var filters []string
	for i := 0; i < len(columns); i++ {
		columnSearchValue := r.FormValue("columns[" + strconv.Itoa(i) + "][search][value]")
		columnData := r.FormValue("columns[" + strconv.Itoa(i) + "][data]")
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
	return ""
}

func getRecordCount(dbMap *gorp.DbMap, query string) int64 {
	var recordsCount int64
	smtOut, err := dbMap.Db.Prepare(query)
	defer smtOut.Close()
	err = smtOut.QueryRow().Scan(&recordsCount)
	if err != nil {
		panic(err.Error())
	}
	return recordsCount
}