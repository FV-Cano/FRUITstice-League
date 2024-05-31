package globals

import "time"

type Fruit struct {
	Name      string    `json:"productName"`
	Stock     int       `json:"stock"`
	UnitPrice float64   `json:"unitPrice"`
	ExpDate   time.Time `json:"expireDate"`
}

type Sale struct {
	ClientName 	string 		`json:"clientName"`
	ProductName string 		`json:"productName"`
	Quantity 	int 		`json:"quantity"`
	Total 		float64 	`json:"total"`
	SaleDate 	time.Time 	`json:"saleDate"`
}

type SaleRequest struct {
	ClientName  string  `json:"clientName"`
	ProductName string  `json:"productName"`
	Quantity    int     `json:"quantity"`
}

type ReportRequest struct {
	Type string `json:"type"`
}

var FruitInventory []Fruit
var Sales []Sale
var Today time.Time = time.Now()
var Week time.Time = Today.AddDate(0, 0, -7)
var Month time.Time = Today.AddDate(0, -1, 0)
var LowStock int = 5