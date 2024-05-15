package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/RichSvK/Stock_Holder_Composition_Go/configs"
	"github.com/RichSvK/Stock_Holder_Composition_Go/models"
)

func FindDataByCode(code string) []models.Stock {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	sql_query := "SELECT * FROM Stocks WHERE `Code` = ? ORDER BY `Date`"
	statement, err := configs.PoolDB.PrepareContext(ctx, sql_query)
	if err != nil {
		fmt.Println("Fail to export because", err.Error())
		return nil
	}
	defer statement.Close()

	rows, err := statement.QueryContext(ctx, code)
	if err != nil {
		fmt.Println("Fail to export because", err.Error())
		return nil
	}
	defer rows.Close()

	var stock models.Stock
	var listStock []models.Stock
	for rows.Next() {
		err = rows.Scan(&stock.Date, &stock.Kode, &stock.LocalIS, &stock.LocalCP, &stock.LocalPF,
			&stock.LocalIB, &stock.LocalID, &stock.LocalMF, &stock.LocalSC, &stock.LocalFD, &stock.LocalOT,
			&stock.ForeignIS, &stock.ForeignCP, &stock.ForeignPF, &stock.ForeignIB, &stock.ForeignID,
			&stock.ForeignMF, &stock.ForeignSC, &stock.ForeignFD, &stock.ForeignOT)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		listStock = append(listStock, stock)
	}
	return listStock
}

func InsertData(stock models.Stock) error {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	sql_query := "INSERT INTO Stocks VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	statement, err := configs.PoolDB.PrepareContext(ctx, sql_query)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer statement.Close()

	_, err = statement.ExecContext(ctx, stock.Date, stock.Kode, stock.LocalIS, stock.LocalCP, stock.LocalPF, stock.LocalIB, stock.LocalID, stock.LocalMF, stock.LocalSC, stock.LocalFD, stock.LocalOT,
		stock.ForeignIS, stock.ForeignCP, stock.ForeignPF, stock.ForeignIB, stock.ForeignID, stock.ForeignMF, stock.ForeignSC, stock.ForeignFD, stock.ForeignOT)
	return err
}
