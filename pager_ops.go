package pager

func (p Pager) Page() int {
	return p.page
}

func (p Pager) Limit() int {
	return p.limit
}

func (p Pager) AllRecord() int {
	return p.allRecord
}

func (p *Pager) SetAllRecord(allRecord int) {
	p.allRecord = allRecord
}

func (p Pager) GetOffset() int {
	if p.page == 0 || p.page == 1 {
		return 0
	}
	return p.limit * (p.page - 1)

}

func (p Pager) GetAllPage() int {
	if p.allRecord == 0 {
		return 0
	}

	allPage := p.allRecord / p.limit
	if p.allRecord%p.limit != 0 {
		allPage++
	}
	return allPage
}
