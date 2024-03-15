package datastore

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ducknificient/web-intelligence/go/entity"
)

func (db *PostgresDB) CrawlpageList(param entity.CrawlpageListParam) (dataList []entity.CrawlpageListData, err error) {
	prefixLog := `CrawlpageList`
	var (
		errMsg string
	)

	qResultsInt := param.Count
	if param.Count == 0 {
		param.Count = 10
	}

	qPageInt := param.Page
	if param.Page == 0 {
		qPageInt = 1
	}

	var limit = qResultsInt
	var toffset = qResultsInt
	if toffset == 0 {
		toffset = 1
	}
	var offset = (qPageInt - 1) * toffset

	fmt.Println("before query row")
	// Prepare SQL statement to check if data exists
	sqlStatement := `SELECT 
		COUNT(1) 
	FROM 
		webintelligence.tabled cp
	WHERE
		du ilike $3
	LIMIT $1 OFFSET $2`
	row := db.Conn.QueryRow(db.Ctx, sqlStatement, limit, offset, `%`+param.Search+`%`)
	var count int
	err = row.Scan(&count)
	if err != nil {
		errMsg = fmt.Sprintf("Unable to select count from webintelligence.tableD. q: %v. param: %v,%v,%v .err: %#v", sqlStatement, limit, offset, `%`+param.Search+`%`, err.Error())
		err = errors.New(errMsg)
		return dataList, err
	}

	fmt.Println("query row")

	if count == 0 {
		return dataList, err
	}

	sqlStatement = `SELECT 
	cp.du, 
	cp.u,
	cp.task
	FROM webintelligence.tabled cp
	WHERE
		du ilike $3
	ORDER BY cp.created DESC
	LIMIT $1 OFFSET $2`

	fmt.Println("before query")
	// sqlStatement = `SELECT
	// cp.pagesource,
	// cp.link,
	// cp.task
	// FROM webintelligence.crawlpage cp
	// ORDER BY cp.created DESC`
	rows, err := db.Conn.Query(db.Ctx, sqlStatement, limit, offset, `%`+param.Search+`%`)
	if err != nil {
		errMsg = fmt.Sprintf("Unable to select from webintelligence.crawlpage. q: %v. param: %v,%v,%v .err: %#v", sqlStatement, limit, offset, `%`+param.Search+`%`, err.Error())
		err = errors.New(errMsg)
		return dataList, err
	}

	fmt.Println("query")
	defer rows.Close()
	for rows.Next() {
		var (
			rs_pagesource sql.NullString
			rs_link       sql.NullString
			rs_task       sql.NullString
		)
		err := rows.Scan(&rs_pagesource, &rs_link, &rs_task)

		if err != nil {
			errMsg = fmt.Sprintf("%v Can't scan query, q:'%v', err:'%v'.", prefixLog, sqlStatement, err.Error())
			err = errors.New(errMsg)
			return dataList, err
		}

		var (
			subdataList []entity.CrawlhrefListData
		)

		subdataList = make([]entity.CrawlhrefListData, 0)

		data := entity.CrawlpageListData{
			Pagesource: rs_pagesource.String,
			Link:       rs_link.String,
			HrefList:   subdataList,
		}

		dataList = append(dataList, data)
	}

	return dataList, err
}
