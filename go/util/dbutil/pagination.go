package dbutil

var NoLimitPagination = Pagination{
	Offset: 0,
	Limit:  ^uint32(0),
}

type Pagination struct {
	Offset uint32 `json:"offset" schema:"offset"`
	Limit  uint32 `json:"limit" schema:"limit"`
}

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func min[T ordered](x, y T) T {
	if x < y {
		return x
	} else {
		return y
	}
}

func (p *Pagination) WithMaxLimit(limit uint32) Pagination {
	pp := *p
	pp.Limit = min(pp.Limit, limit)
	return pp
}
