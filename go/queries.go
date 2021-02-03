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
	var costFloat float32 = 1.0
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
		costFloat, err = costToFloat32(cost)
		if err != nil {
			return err
		}
	}

	query := fmt.Sprintf("INSERT INTO statistics (date, views, clicks, cost) VALUES ('%s', %d, %d, %f)", date, viewsInt, clicksInt, costFloat)
	_, err = a.Conn.Query(query)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) retrieveStatisticsFromDatabase(from, to, orderBy string) ([]map[string]interface{}, error) {
	query := fmt.Sprintf("SELECT date, views, clicks, cost FROM statistics WHERE date >= '%s' AND date <= '%s' ORDER BY %s", from, to, orderBy)
	rows, err := a.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	m := mapify(rows)

	return m, err
}

func (a *App) deleteAllStatisticsFromDatabase() error {
	_, err := a.Conn.Query("DELETE FROM statistics")
	if err != nil {
		return err
	}
	return nil
}

func mapify(rows *sql.Rows) []map[string]interface{} {
	mapping := make([]map[string]interface{}, 0)
	var date time.Time
	var views int
	var clicks int
	var cost float32
	var cpc float32
	var cpm float32
	for rows.Next() {
		err := rows.Scan(&date, &views, &clicks, &cost)
		if err != nil {
			fmt.Println(err)
		}
		if clicks != 0 {
			cpc = cost / float32(clicks)
		} else {
			cpc = 0
		}

		if views != 0 {
			cpm = (cost / float32(views)) * 1000
		} else {
			cpm = 0
		}

		elem := map[string]interface{}{"Date": date,
			"Views":  views,
			"Clicks": clicks,
			"Cost":   float32ToCost(cost),
			"Cpc":    float32ToCost(cpc),
			"Cpm":    float32ToCost(cpm)}
		mapping = append(mapping, elem)
	}
	return mapping
}
