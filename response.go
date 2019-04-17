package pager

type PagerResponse struct {
	Page      int
	Limit     int
	AllPage   int
	AllRecord int
}

func NewPagerResponse(pager *Pager) *PagerResponse {
	return &PagerResponse{
		Page:      pager.Page(),
		Limit:     pager.Limit(),
		AllPage:   pager.GetAllPage(),
		AllRecord: pager.AllRecord(),
	}
}
