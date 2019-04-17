package pager

import "testing"

func TestOffset(t *testing.T) {
	cases := []struct {
		page   string
		limit  string
		offset int
	}{
		{"", "", 0},
		{"1", "20", 0},
		{"2", "20", 20},
		{"3", "20", 40},
	}
	for _, c := range cases {
		pager := NewPager(c.page, c.limit)
		if pager.GetOffset() != c.offset {
			t.Errorf("page = %s, limit = %s, offset = %d,  want %d", c.page, c.limit, pager.GetOffset(), c.offset)
		}
	}
}

func TestAllPage(t *testing.T) {
	cases := []struct {
		limit     string
		allRecord int
		allPage   int
	}{
		{"", 100, 2}, // default limit is 50
		{"20", 100, 5},
		{"20", 26, 2},
		{"20", 19, 1},
	}
	for _, c := range cases {
		pager := NewPager("1", c.limit)
		pager.SetAllRecord(c.allRecord)
		if pager.GetAllPage() != c.allPage {
			t.Errorf("limit = %s, allRecord = %d, allPage = %d,  want %d",
				c.limit, c.allRecord, pager.GetAllPage(), c.allPage)
		}
	}
}
