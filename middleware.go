package pager

import (
	"context"
	"net/http"
)

type pagerContextKey uint

const PagerCtxKey pagerContextKey = 0

type sortContextKey uint

const SortCtxKey sortContextKey = 0

const keyPage = "page"
const keyLimit = "limit"
const keySort = "sort"

func InjectPager(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		// pager
		page := query.Get(keyPage)
		limit := query.Get(keyLimit)
		pager := NewPager(page, limit)

		// sort
		sort := query.Get(keySort)
		sorting := NewSort(sort)

		ctx := r.Context()
		ctx = context.WithValue(ctx, PagerCtxKey, pager)
		ctx = context.WithValue(ctx, SortCtxKey, sorting)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
