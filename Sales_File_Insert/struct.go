package salesFileInsert

import "time"

type Customer struct {
	ID         uint   `gorm:"column:id"`
	CustomerID string `gorm:"column:CustomerID"`
	Name       string `gorm:"column:Name"`
	Email      string `gorm:"column:Email"`
	Address    string `gorm:"column:Address"`
}

type Product struct {
	ID          uint   `gorm:"column:id"`
	ProductID   string `gorm:"column:ProductID"`
	ProductName string `gorm:"column:ProductName"`
	Category    string `gorm:"column:Category"`
}

type Order struct {
	ID            uint      `gorm:"column:id"`
	OrderID       string    `gorm:"column:OrderID"`
	CustomerID    string    `gorm:"column:CustomerID"`
	DateOfSale    time.Time `gorm:"column:DateOfSale"`
	Region        string    `gorm:"column:Region"`
	PaymentMethod string    `gorm:"column:PaymentMethod"`
	ShippingCost  float64   `gorm:"column:ShippingCost"`
}

type OrderDetail struct {
	ID           uint    `gorm:"column:id"`
	OrderID      string  `gorm:"column:OrderID"`
	ProductID    string  `gorm:"column:ProductID"`
	QuantitySold int     `gorm:"column:QuantitySold"`
	UnitPrice    float64 `gorm:"column:UnitPrice"`
	Discount     float64 `gorm:"column:Discount"`
}

type CommonRespStruct struct {
	Status string `json:"status"`
	ErrMsg string `json:"errMsg"`
}

type ReqStruct struct {
	Indicator string `json:"indicator"`
	FromDate  string `json:"fromDate"`
	ToDate    string `json:"toDate"`
}

type RevenueStruct struct {
	Status             string            `json:"status"`
	ErrMsg             string            `json:"errMsg"`
	Total_revenue      string            `json:"totalRevenue"`
	TotProdRevenue     []ProductRevenue  `json:"prodRevenue"`
	TotalcatRevenue    []CategoryRevenue `json:"catRevenue"`
	TotalRevenue_byreg []RegionRevenue   `json:"regRevenue"`
}

type ProductRevenue struct {
	ProductName  string  `json:"prodname" gorm:"column:prodname"`
	TotalRevenue float64 `json:"totalRevenue" gorm:"column:totalRevenue"`
}

type CategoryRevenue struct {
	Category     string  `json:"category" gorm:"column:category"`
	TotalRevenue float64 `json:"totalRevenue" gorm:"column:totalRevenue"`
}

type RegionRevenue struct {
	Region       string  `json:"region" gorm:"column:region"`
	TotalRevenue float64 `json:"totalRevenue" gorm:"column:totalRevenue"`
}
