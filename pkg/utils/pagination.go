package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"strconv"
)

const (
	defaultSize = 10
)

// Pagination query params
type PaginationQuery struct {
	Size    int    `json:"size,omitempty"`
	Page    int    `json:"page,omitempty"`
	OrderBy string `json:"order_by,omitempty"`
}

func (q *PaginationQuery) SetPage(pageParam string) error {
	if pageParam == "" {
		q.Size = 0
		return nil
	}

	n, err := strconv.Atoi(pageParam)
	if err != nil {
		return errors.Wrap(err, "PaginationQuery.GetPaginationFromCtx.SetPage")
	}
	q.Page = n

	return nil
}

func (q *PaginationQuery) GetOrderBy() string {
	return q.OrderBy
}

func (q *PaginationQuery) SetSize(sizeParam string) error {
	if sizeParam == "" {
		q.Size = defaultSize
		return nil
	}
	s, err := strconv.Atoi(sizeParam)
	if err != nil {
		return errors.Wrap(err, "PaginationQuery.GetPaginationFromCtx.SetSize")
	}
	q.Size = s

	return nil
}

func (q *PaginationQuery) SetOrderBy(orderByParam string) {
	q.OrderBy = orderByParam
}

// Get pagination query from structure
func GetPaginationFromCtx(c echo.Context) (*PaginationQuery, error) {
	q := &PaginationQuery{}
	if err := q.SetPage(c.QueryParam("page")); err != nil {
		return nil, err
	}
	if err := q.SetSize(c.QueryParam("size")); err != nil {
		return nil, err
	}

	q.SetOrderBy(c.QueryParam("order_by"))

	return q, nil
}
