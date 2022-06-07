package storage

type QueryBuilderHandler func(qb *QueryBuilder, result interface{}) error

type qcondition struct {
	fieldname string
	value     interface{}
}

type QueryBuilder struct {
	Conds []qcondition

	LimitResult *int
	SortOrder   int
	TakeFirst   bool
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		Conds: make([]qcondition, 0),
	}
}

func (b *QueryBuilder) Where(field string, cond interface{}) *QueryBuilder {
	b.Conds = append(b.Conds, qcondition{
		fieldname: field,
		value:     cond,
	})

	return b
}

func (b *QueryBuilder) Limit(number int) *QueryBuilder {

	b.LimitResult = &number

	return b
}

func (b *QueryBuilder) Sort(order int) *QueryBuilder {

	b.SortOrder = order

	return b
}

func (b *QueryBuilder) First() *QueryBuilder {

	b.TakeFirst = true

	return b
}
