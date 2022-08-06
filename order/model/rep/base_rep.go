package rep

type BaseRep struct {
	Code int
	Msg  string
}

type Date struct {
	Content interface{}
	PageRep PageRep
}

type Rep struct {
	Code int
	Msg  string
	Date Date
}

func NewRep() Rep {
	var r Rep
	r.Code = 200
	r.Msg = "ok"
	return r
}

type Token struct {
	Code  int
	Token string
}

type PageRep struct {
	PageIndex int64
	PageSize  int64
	ItemTotal int64
	PageTotal int64
}

func NewBSSRep() BaseRep {
	var r BaseRep
	r.Code = 200
	r.Msg = "success"
	return r
}

func NewBSERep() BaseRep {
	var r BaseRep
	r.Code = 500
	return r
}
func NewBSEJRep() BaseRep {
	var r BaseRep
	r.Code = 500
	r.Msg = "传参不完整，请检查"
	return r
}
