package model

type Menu struct {
	Uuid             string `json:"uuid"`
	Score            int    `json:"score"`
	HistoricalSales  int    `json:"historical_sales"` //历史销量
	Name             string `json:"name"`
	Image            string `json:"image"`
	Price            int    `json:"price"`
	MaxSell          int    `json:"max_sell"`           //最多销多少
	SalesVolumeMonth int    `json:"sales_volume_month"` //月销量
	OrgId            int    `json:"org_id"`             //企业id
	Status           int    `json:"status"`
	Classification   string `json:"classification"` //菜分类

}
type MenuClassification struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
