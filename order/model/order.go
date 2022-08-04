package model

type Order struct {
	Id     int64
	Time   string
	Name   string
	Buyers string
}

type Food struct {
	Name  string
	Price string
	Total int
	Pic   []string
}

type Comment struct {
	FoodName string
	Content  string
	Cname    string
	Time     string
}
