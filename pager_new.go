package pager

import "strconv"

const defaultPageLimit = 50

type Pager struct {
	page      int
	limit     int
	allRecord int
}

func NewPager(page string, limit string) Pager {
	pager := Pager{}

	if page == "" {
		pager.page = 1
	} else {
		if pageInt, err := strconv.Atoi(page); err != nil {
			pager.page = 1
		} else {
			pager.page = pageInt
		}
	}

	if limit == "" {
		pager.limit = defaultPageLimit
	} else {
		if limitInt, err := strconv.Atoi(limit); err != nil {
			pager.limit = defaultPageLimit
		} else {
			pager.limit = limitInt
		}
	}

	return pager
}
