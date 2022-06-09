package storage

type QueryBuilderHandler func(qb *QueryBuilder, result interface{}) error
type StoreHandlerFunction func(item interface{}) error

type qcondition struct {
	Fieldname string
	Value     interface{}
}

type QueryBuilder struct {
	Conds []qcondition

	LimitResult *int
	SortOrder   int
	SortField   *string
	TakeFirst   bool
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		Conds: make([]qcondition, 0),
	}
}

func (b *QueryBuilder) Where(field string, cond interface{}) *QueryBuilder {
	b.Conds = append(b.Conds, qcondition{
		Fieldname: field,
		Value:     cond,
	})

	return b
}

func (b *QueryBuilder) Limit(number int) *QueryBuilder {

	b.LimitResult = &number

	return b
}

func (b *QueryBuilder) Sort(fieldname string, order int) *QueryBuilder {

	b.SortOrder = order
	b.SortField = &fieldname

	return b
}

func (b *QueryBuilder) First() *QueryBuilder {

	b.TakeFirst = true

	return b
}
