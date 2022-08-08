package req

type MenuReq struct {
	Uuid             int    `json:"uuid"`
	Score            int    `json:"score"`
	HistoricalSales  int    `json:"historical_sales"` //历史销量
	Name             string `binding:"required" json:"name"`
	Image            string `binding:"required" json:"image"`
	Price            int    `binding:"required" json:"price"`
	MaxSell          int    `json:"max_sell"` //月销量
	SalesVolumeMonth int    `json:"sales_volume_month"`
	OrgId            int    `binding:"required" json:"org_id"`         //企业id
	Classification   string `binding:"required" json:"classification"` //菜分类

}

type MenuSearchReq struct {
	Score            int    `json:"score"`
	Name             string ` json:"name"`
	Price            int    ` json:"price"`
	SalesVolumeMonth int    ` json:"sales_volume_month"`
	OrgId            int    ` json:"org_id"`         //企业id
	Classification   string ` json:"classification"` //菜分类
	Status           int    `json:"status"`
	Page             Page   ` json:"page"`
}

type MenuChangeReq struct {
	OrgId   int     `binding:"required" json:"org_id"`
	Name    string  `json:"name"`
	Uuid    string  `binding:"required"  json:"uuid"`
	Status  int     `json:"status"`
	Price   float32 `json:"price"`
	MaxSell int     `json:"max_sell"`
}
