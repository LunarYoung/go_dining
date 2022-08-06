package model

type Order struct {
	OrderId    string    `json:"order_id"`
	Time       string    `json:"time"`
	Name       string    `json:"name"`
	Buyers     string    `json:"buyers"`
	FoodName   []string  `json:"food_name"`
	FoodNPrice []float32 `binding:"required" json:"food_price"`
	TotalPrice int       `json:"total_price"`
	OrgId      int       ` json:"org_id"`
	Status     int8      `json:"status"`
	PickUp     string    `json:"pick_up"`
}

type Food struct {
	Name  string
	Price string
	Total int
	Pic   []string
}

type Comment struct {
	FoodName string `json:"food_name"`
	Content  string `json:"content"`
	Cname    string `json:"cname"`
	Time     string `json:"time"`
}
