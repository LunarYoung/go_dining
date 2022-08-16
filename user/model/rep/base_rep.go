package rep

type BaseRep struct {
	Code int
	Msg  string
}

type Date struct {
	Content interface{}
	PageRep PageRep
}

type NoPageRep struct {
	Code int
	Msg  string
	Date interface{}
}

func NewNoPageRep() (rep NoPageRep) {
	rep.Msg = "ok"
	rep.Code = 200
	return rep
}

type Rep struct {
	Code int
	Msg  string
	Date Date
}

type Token struct {
	Code  int
	Token string
}

type PageRep struct {
	PageIndex int
	PageSize  int
	ItemTotal int
	PageTotal int
}

func NewBSSRep() BaseRep {
	var r BaseRep
	r.Code = 200
	r.Msg = "success"
	return r
}

func NewBSERep() BaseRep {
	var r BaseRep
	return r
}
