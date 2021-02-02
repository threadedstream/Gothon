package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

func (a *App) saveStatisticsToDatabase(date, views, clicks, cost string) error {
	// To avoid further division by zero
	var viewsInt int = 1
	var clicksInt int = 1
	var costInt int = 1
	var err error
	if views != "" {
		viewsInt, err = strconv.Atoi(views)
		if err != nil {
			return err
		}
	}

	if clicks != "" {
		clicksInt, err = strconv.Atoi(clicks)
		if err != nil {
			return err
		}
	}

	if cost != "" {
		costInt, err = strconv.Atoi(cost)
		if err != nil {
			return err
		}
	}

	query := fmt.Sprintf("INSERT INTO statistics (date, views, clicks, cost) VALUES ('%s', %d, %d, %d)", date, viewsInt, clicksInt, costInt)
	_, err = a.Conn.Query(query)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) retrieveStatisticsFromDatabase(from, to string) ([]map[string]interface{}, error) {
	query := fmt.Sprintf("SELECT date, views, clicks, cost FROM statistics WHERE date >= '%s' AND date <= '%s' ORDER BY date", from, to)
	rows, err := a.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	m := mapify(rows)
	return m, err
}

func mapify(rows *sql.Rows) []map[string]interface{} {
	mapping := make([]map[string]interface{}, 0)
	var date time.Time
	var views int
	var clicks int
	var cost int
	var cpc float32
	var cpm float32
	for rows.Next() {
		err := rows.Scan(&date, &views, &clicks, &cost)
		if err != nil {
			fmt.Println(err)
		}
		if clicks != 0 {
			cpc = float32(cost / clicks)
		} else {
			cpc = 0
		}

		if views != 0 {
			cpm = (float32(cost / views)) * 1000
		} else {
			cpm = 0
		}

		elem := map[string]interface{}{"Date": date, "Views": views, "Clicks": clicks, "Cost": cost, "Cpc": cpc, "Cpm": cpm}
		mapping = append(mapping, elem)
	}
	return mapping
}
