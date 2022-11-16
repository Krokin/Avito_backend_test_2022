package database

import (
	"fmt"
	"strconv"

	m "Avito_tech_test_2022/internal/models"
)

func (d *DB) DownloadReport(r *m.ReportDate) ([][]string, error) {
	q := fmt.Sprintf(`SELECT service_id, service_name, sum(cost) as total_cost FROM report_services 
		WHERE date LIKE '%d-%d-%%' GROUP BY service_id ORDER BY service_id ASC`, r.Year, r.Month)
	rows, err := d.DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := [][]string {
		{
			"service_id",
			"service_name",
			"cost",
		},
	}
	for rows.Next() {
		var (
			service_id int
			service_name string
			cost float64
		) 
		if err := rows.Scan(&service_id, &service_name, &cost); err != nil {
			return nil, err
		}
		list = append(list, []string {strconv.Itoa(service_id), service_name, strconv.FormatFloat(cost, 'f', 2, 64)})
	}

	return list, nil
}