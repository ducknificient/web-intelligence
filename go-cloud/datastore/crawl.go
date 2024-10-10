package datastore

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"go-cloud/entity"
)

func (d *PostgresDB) CrawlpageList(param entity.CrawlpageListParam) (dataList []entity.CrawlpageListData, err error) {
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

	// Prepare SQL statement to check if data exists
	sqlStatement := `SELECT 
		COUNT(1) 
	FROM 
		webintelligence.tabled cp
	WHERE
		du ilike $3
	LIMIT $1 OFFSET $2`
	row := d.conn.QueryRow(d.ctx, sqlStatement, limit, offset, `%`+param.Search+`%`)
	var count int
	err = row.Scan(&count)
	if err != nil {
		errMsg = fmt.Sprintf("Unable to select count from webintelligence.tableD. q: %v. param: %v,%v,%v .err: %#v", sqlStatement, limit, offset, `%`+param.Search+`%`, err.Error())
		err = errors.New(errMsg)
		return dataList, err
	}

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

	d.logger.Info(sqlStatement)
	rows, err := d.conn.Query(d.ctx, sqlStatement, limit, offset, `%`+param.Search+`%`)
	if err != nil {
		errMsg = fmt.Sprintf("Unable to select from webintelligence.crawlpage. q: %v. param: %v,%v,%v .err: %#v", sqlStatement, limit, offset, `%`+param.Search+`%`, err.Error())
		err = errors.New(errMsg)
		return dataList, err
	}

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

func (d *PostgresDB) CrawlpageListParsed(param entity.CrawlpageListParam) (dataList []entity.CrawlpageListParsedData, err error) {
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

	// Prepare SQL statement to check if data exists
	sqlStatement := `SELECT 
		COUNT(1) 
	FROM webintelligence.crawlpage cp
	WHERE task ='JATIMPROV'	
	AND link LIKE 'https://jatimprov%'
	AND link LIKE '%/berita/%'
	AND cp.pagesource like $3
	LIMIT $1 OFFSET $2`
	row := d.conn.QueryRow(d.ctx, sqlStatement, limit, offset, `%`+param.Search+`%`)
	var count int
	err = row.Scan(&count)
	if err != nil {
		errMsg = fmt.Sprintf("Unable to select count from webintelligence.tableD. q: %v. param: %v,%v,%v .err: %#v", sqlStatement, limit, offset, `%`+param.Search+`%`, err.Error())
		err = errors.New(errMsg)
		return dataList, err
	}

	if count == 0 {
		return dataList, err
	}

	sqlStatement = `SELECT 
	COALESCE(cp.link,''),
	COALESCE(cp.pagesource,''),
	substring(
			cp.pagesource
			FROM
			'<meta property="og:title" content="(.*?)">'
		) AS document_metatitle,
	COALESCE(substring(
		cp.pagesource
		FROM
		'<meta property="og:description" content="(.*?)">'
	),'') AS document_metacontent,
	COALESCE(substring(
		cp.pagesource
		FROM
		'<div class="pr-bg pr-bg-white"></div>
	<h3>(.*?)</h3>'
	),'') AS document_title,
	COALESCE(substring(
		cp.pagesource
		FROM
		'<div class="parallax-header"> <a href="#">(.*?)</a>'
	),'') AS document_date,
	COALESCE(substring(
		cp.pagesource
		FROM
		'<span>Kategori : </span><a href="#">(.*?)</a> </div>'
	),'') AS document_category,
	COALESCE(substring(
		cp.pagesource
		FROM
		'<li><span><i class="fal fa-eye"></i>(.*?)</span></li>'
	),'') AS document_totalview,
	COALESCE(substring(
		cp.pagesource
		FROM
		'<li><span><i class="fal fa-hashtag"></i>(.*?)</span></li>'
	),'') AS document_hashtag,
	COALESCE(substring(
		cp.pagesource
		FROM
		'<p><p>(.*?)</p></p>'
	),'') AS document_content,
	COALESCE(substring(
		cp.pagesource
		FROM
		'<span>Berita Terkait</span>(.*?)</div>
	</div>
	</div>
	</div>
	</div>
	</div>'
	),'') AS document_relatednews
	FROM webintelligence.crawlpage cp
	WHERE task ='JATIMPROV'	
	AND link LIKE 'https://jatimprov%'
	AND link LIKE '%/berita/%'
	AND cp.pagesource like $3
	ORDER BY cp.created DESC
	LIMIT $1 OFFSET $2
	`

	d.logger.Info(sqlStatement)
	rows, err := d.conn.Query(d.ctx, sqlStatement, limit, offset, `%`+param.Search+`%`)
	if err != nil {
		errMsg = fmt.Sprintf("Unable to select from webintelligence.crawlpage. q: %v. param: %v,%v,%v .err: %#v", sqlStatement, limit, offset, `%`+param.Search+`%`, err.Error())
		err = errors.New(errMsg)
		return dataList, err
	}

	defer rows.Close()
	for rows.Next() {
		var (
			rs_link       sql.NullString
			rs_pagesource sql.NullString
			rs_task       sql.NullString

			rs_document_metatitle   sql.NullString
			rs_document_metacontent sql.NullString
			rs_document_title       sql.NullString

			rs_document_date      sql.NullString
			rs_document_category  sql.NullString
			rs_document_totalview sql.NullString

			rs_document_hashtag     sql.NullString
			rs_document_content     sql.NullString
			rs_document_relatednews sql.NullString
		)
		err := rows.Scan(&rs_link, &rs_pagesource, &rs_document_metatitle, &rs_document_metacontent,
			&rs_document_title, &rs_document_date, &rs_document_category, &rs_document_totalview,
			&rs_document_hashtag, &rs_document_content, &rs_document_relatednews)

		if err != nil {
			errMsg = fmt.Sprintf("%v Can't scan query, q:'%v', err:'%v'.", prefixLog, sqlStatement, err.Error())
			err = errors.New(errMsg)
			return dataList, err
		}

		// clean up data
		var (
			newcontent string
		)
		// clean up content

		// remove strong tag
		newcontent = strings.ReplaceAll(rs_document_content.String, "<strong>", "")
		newcontent = strings.ReplaceAll(newcontent, "</strong>", "")

		// remove p tag
		newcontent = strings.ReplaceAll(newcontent, "<p>", "")
		newcontent = strings.ReplaceAll(newcontent, "</p>", "")

		data := entity.CrawlpageListParsedData{
			// Pagesource:  rs_pagesource.String,
			Link:        rs_link.String,
			Task:        rs_task.String,
			Metatitle:   rs_document_metatitle.String,
			Metacontent: rs_document_metacontent.String,
			Title:       rs_document_title.String,
			Date:        rs_document_date.String,
			Category:    rs_document_category.String,
			TotalView:   rs_document_totalview.String,
			Hashtag:     rs_document_hashtag.String,
			Content:     newcontent,
			RelatedNews: rs_document_relatednews.String,
		}

		dataList = append(dataList, data)
	}

	return dataList, err
}

func (d *PostgresDB) GetLatestSeedUrl(task string, seedurl string) (latest_seedurl string, err error) {

	latest_seedurl = seedurl

	// Prepare SQL statement to check if data exists in tableD
	var sqlStatement string

	sqlStatement = `SELECT 
		COUNT(*) 
	FROM webintelligence.crawlhref WHERE task = $1`
	row := d.conn.QueryRow(d.ctx, sqlStatement, task)
	var count int
	err = row.Scan(&count)
	if err != nil {
		return latest_seedurl, err
	}

	// If count > 0, data exists in tableD
	if count > 0 {

		var rs_latest_seedurl sql.NullString

		sqlStatement = `select link from webintelligence.crawlhref ch
		WHERE task = $1
		order by ch.created desc 
		limit 1`
		row := d.conn.QueryRow(d.ctx, sqlStatement, task)
		err = row.Scan(&rs_latest_seedurl)
		if err != nil {
			return seedurl, err
		}
		latest_seedurl = rs_latest_seedurl.String

	}

	return latest_seedurl, err
}

func (d *PostgresDB) GetExistingQueue(task string) (dataList []string, err error) {

	prefixLog := `CrawlpageList`
	var (
		errMsg string
	)

	// Prepare SQL statement to check if data exists
	sqlStatement := `select count (1) from (
		select 
			ch.href,
			min(ch.created) as timestamp
		from 
			webintelligence.crawlhref ch
		inner join webintelligence.crawlpage cp on cp.link = ch.link 
		where
		ch.href not in (
			select cp.link from webintelligence.crawlpage cp 
		)
		and ch.task = $1
		GROUP BY ch.href
		ORDER BY timestamp ASC
	)`
	row := d.conn.QueryRow(d.ctx, sqlStatement, task)
	var count int
	err = row.Scan(&count)
	if err != nil {
		errMsg = fmt.Sprintf("Unable to select count from webintelligence.crawlhref. q: %v. param: %v .err: %#v", sqlStatement, task, err.Error())
		err = errors.New(errMsg)
		return dataList, err
	}

	if count == 0 {
		return dataList, err
	}

	sqlStatement = `SELECT
		ch.href,
		min(ch.created) as timestamp
	from 
		webintelligence.crawlhref ch
	inner join webintelligence.crawlpage cp on cp.link = ch.link 
	where
	ch.href not in (
		select cp.link from webintelligence.crawlpage cp 
	)
	and ch.task = $1
	GROUP BY ch.href
	ORDER BY timestamp ASC`

	d.logger.Info(sqlStatement)
	rows, err := d.conn.Query(d.ctx, sqlStatement, task)
	if err != nil {
		errMsg = fmt.Sprintf("Unable to select from webintelligence.crawlpage. q: %v. param: %v .err: %#v", sqlStatement, task, err.Error())
		err = errors.New(errMsg)
		return dataList, err
	}

	defer rows.Close()
	for rows.Next() {
		var (
			rs_href      sql.NullString
			rs_timestamp sql.NullString
		)
		err := rows.Scan(&rs_href, &rs_timestamp)

		if err != nil {
			errMsg = fmt.Sprintf("%v Can't scan query, q:'%v', err:'%v'.", prefixLog, sqlStatement, err.Error())
			err = errors.New(errMsg)
			return dataList, err
		}

		dataList = append(dataList, rs_href.String)
	}

	return dataList, err
}

func (d *PostgresDB) GetAlodokterListParsed(param entity.AlodokterListDataParsedParam) (dataList []entity.AlodokterListDataParsedData, err error) {
	prefixLog := `GetAlodokterListParsed`
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

	// Prepare SQL statement to check if data exists
	sqlStatement := `SELECT 
		COUNT(1) 
		FROM (
			SELECT 
				cp.link as document_link,
				COALESCE(substring(
					cp.pagesource
					FROM
					'<html><body>(.*?)</body></html>'
				),'') AS document_content,
				COALESCE(substring(
					cp.pagesource
					FROM
					'const pageType = "(.*?)"'
				),'') AS document_type,
				COALESCE(substring(
					cp.pagesource
					FROM
					'<meta name="description" content="(.*?)"'
				),'') AS document_description,
				COALESCE(substring(
					cp.pagesource
					FROM
					'<meta name="keywords" content="(.*?)"'
				),'') AS document_keywords,
				COALESCE(substring(
					cp.pagesource
					FROM
					'<meta property="og:image" itemprop="image" content="(.*?)"'
				),'') AS document_image
			FROM 
				webintelligence.crawlpage cp
			WHERE 
				task ='ALODOKTER'
			ORDER BY cp.created DESC
			)`
	row := d.conn.QueryRow(d.ctx, sqlStatement)
	var count int
	err = row.Scan(&count)
	if err != nil {
		errMsg = fmt.Sprintf("Unable to select count from webintelligence.tableD. q: %v. param: %v,%v,%v .err: %#v", sqlStatement, limit, offset, `%`+param.Search+`%`, err.Error())
		err = errors.New(errMsg)
		return dataList, err
	}

	if count == 0 {
		return dataList, err
	}

	sqlStatement = `SELECT 	
		document_link,
		document_content, 
		document_type,
		document_description,
		document_keywords,
		document_image
	FROM (
	SELECT 
		cp.link as document_link,
		COALESCE(substring(
			cp.pagesource
			FROM
			'<html><body>(.*?)</body></html>'
		),'') AS document_content,
		COALESCE(substring(
			cp.pagesource
			FROM
			'const pageType = "(.*?)"'
		),'') AS document_type,
		COALESCE(substring(
			cp.pagesource
			FROM
			'<meta name="description" content="(.*?)"'
		),'') AS document_description,
		COALESCE(substring(
			cp.pagesource
			FROM
			'<meta name="keywords" content="(.*?)"'
		),'') AS document_keywords,
		COALESCE(substring(
			cp.pagesource
			FROM
			'<meta property="og:image" itemprop="image" content="(.*?)"'
		),'') AS document_image
	FROM 
		webintelligence.crawlpage cp
	WHERE 
		task ='ALODOKTER'
	ORDER BY cp.created DESC
	)
	`

	d.logger.Info(sqlStatement)
	rows, err := d.conn.Query(d.ctx, sqlStatement)
	if err != nil {
		errMsg = fmt.Sprintf("Unable to select from webintelligence.crawlpage. q: %v. param: %v,%v,%v .err: %#v", sqlStatement, limit, offset, `%`+param.Search+`%`, err.Error())
		err = errors.New(errMsg)
		return dataList, err
	}

	defer rows.Close()
	for rows.Next() {
		var (
			rs_document_link        sql.NullString
			rs_document_content     sql.NullString
			rs_document_type        sql.NullString
			rs_document_description sql.NullString
			rs_document_keywords    sql.NullString
			rs_document_image       sql.NullString
		)
		err := rows.Scan(&rs_document_link, &rs_document_content, &rs_document_type, &rs_document_description,
			&rs_document_keywords, &rs_document_image)

		if err != nil {
			errMsg = fmt.Sprintf("%v Can't scan query, q:'%v', err:'%v'.", prefixLog, sqlStatement, err.Error())
			err = errors.New(errMsg)
			return dataList, err
		}

		// clean up data
		var (
		// newcontent string
		)
		// clean up content

		// remove strong tag
		// newcontent = strings.ReplaceAll(rs_document_content.String, "<strong>", "")
		// newcontent = strings.ReplaceAll(newcontent, "</strong>", "")

		// remove p tag
		// newcontent = strings.ReplaceAll(newcontent, "<p>", "")
		// newcontent = strings.ReplaceAll(newcontent, "</p>", "")

		data := entity.AlodokterListDataParsedData{
			DocumentLink:        rs_document_link.String,
			DocumentContent:     rs_document_content.String,
			DocumentType:        rs_document_type.String,
			DocumentDescription: rs_document_description.String,
			DocumentKeywords:    rs_document_keywords.String,
			DocumentImage:       rs_document_image.String,
		}

		dataList = append(dataList, data)
	}

	return dataList, err
}
