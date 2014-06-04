package paging

import (
	"database/sql"
	"net/http"
	"strconv"
)

type ResultsPage struct {
	Page     int           `json:"page"`
	PageSize int           `json:"page_size"`
	Results  []interface{} `json:"results"`
}

type pageScanFunc func(rows *sql.Rows) (interface{}, error)

func defaultPagingParam(strValue string, defaultValue uint64) uint64 {
	if len(strValue) > 0 {
		val, err := strconv.Atoi(strValue)
		if err != nil {
			return uint64(val)
		}
	}

	return defaultValue
}

func PagingParams(r *http.Request) (page, pageSize uint64) {
	pageSize = defaultPagingParam(r.Form.Get("page"), 20)
	page = defaultPagingParam(r.Form.Get("page_size"), 0)
	return
}

func MakePage(pageNumber int, pageSize int, rows *sql.Rows, scanFn pageScanFunc) (page *ResultsPage, err error) {
	page = &ResultsPage{
		Page:     pageNumber,
		PageSize: pageSize,
		Results:  make([]interface{}, 0),
	}

	for rows.Next() {
		record, err := scanFn(rows)
		if err == nil {
			page.Results = append(page.Results, record)
		}
	}

	err = rows.Err()
	return
}
