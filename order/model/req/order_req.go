package req

type OrderReq struct {
	Time       string    `binding:"required" json:"time"`
	Name       string    `binding:"required" json:"name"`
	Buyers     string    `binding:"required" json:"buyers"`
	FoodName   []string  `binding:"required" json:"food_name"`
	FoodNPrice []float32 `binding:"required" json:"food_price"`
	TotalPrice int       `binding:"required" json:"total_price"`
	OrgId      int       `binding:"required" json:"org_id"`
}

type OrderSearchReq struct {
	OrgId   int    `binding:"required" json:"org_id"`
	Time    string `json:"time"`
	Status  int    `json:"status"`
	Name    string `json:"name"`
	OrderId string `json:"order_id"`
	Page    Page   `binding:"required" json:"page"`
}
