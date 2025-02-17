// Code generated by github.com/mackee/go-sqlla/v2/cmd/sqlla - DO NOT EDIT.
package record

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"database/sql"
	"time"

	"github.com/mackee/go-sqlla/v2"
)

type messageSQL struct {
	where sqlla.Where
}

func NewMessageSQL() messageSQL {
	q := messageSQL{}
	return q
}

var messageAllColumns = []string{
	"\"id\"", "\"author\"", "\"text\"", "\"created_at\"", "\"updated_at\"",
}

type messageSelectSQL struct {
	messageSQL
	Columns     []string
	order       sqlla.OrderWithColumn
	limit       *uint64
	offset      *uint64
	tableAlias  string
	joinClauses []string

	additionalWhereClause func(int) (string, int, []any)
	groupByColumns        []string

	isForUpdate bool
}

func (q messageSQL) Select() messageSelectSQL {
	return messageSelectSQL{
		q,
		messageAllColumns,
		nil,
		nil,
		nil,
		"",
		nil, nil,
		nil,
		false,
	}
}

func (q messageSelectSQL) Or(qs ...messageSelectSQL) messageSelectSQL {
	ws := make([]sqlla.Where, 0, len(qs))
	for _, q := range qs {
		ws = append(ws, q.where)
	}
	q.where = append(q.where, sqlla.ExprOr(ws))
	return q
}

func (q messageSelectSQL) Limit(l uint64) messageSelectSQL {
	q.limit = &l
	return q
}

func (q messageSelectSQL) Offset(o uint64) messageSelectSQL {
	q.offset = &o
	return q
}

func (q messageSelectSQL) ForUpdate() messageSelectSQL {
	q.isForUpdate = true
	return q
}

func (q messageSelectSQL) TableAlias(alias string) messageSelectSQL {
	q.tableAlias = "\"" + alias + "\""
	return q
}

func (q messageSelectSQL) SetColumns(columns ...string) messageSelectSQL {
	q.Columns = make([]string, 0, len(columns))
	for _, column := range columns {
		if strings.ContainsAny(column, "(."+"\"") {
			q.Columns = append(q.Columns, column)
		} else {
			q.Columns = append(q.Columns, "\""+column+"\"")
		}
	}
	return q
}

func (q messageSelectSQL) JoinClause(clause string) messageSelectSQL {
	q.joinClauses = append(q.joinClauses, clause)
	return q
}

func (q messageSelectSQL) AdditionalWhereClause(clause func(int) (string, int, []any)) messageSelectSQL {
	q.additionalWhereClause = clause
	return q
}

func (q messageSelectSQL) appendColumnPrefix(column string) string {
	if q.tableAlias == "" || strings.ContainsAny(column, "(.") {
		return column
	}
	return q.tableAlias + "." + column
}

func (q messageSelectSQL) GroupBy(columns ...string) messageSelectSQL {
	q.groupByColumns = make([]string, 0, len(columns))
	for _, column := range columns {
		if strings.ContainsAny(column, "(."+"\"") {
			q.groupByColumns = append(q.groupByColumns, column)
		} else {
			q.groupByColumns = append(q.groupByColumns, "\""+column+"\"")
		}
	}
	return q
}

func (q messageSelectSQL) ID(v MessageID, exprs ...sqlla.Operator) messageSelectSQL {
	where := sqlla.ExprValue[int64]{Value: int64(v), Op: sqlla.Operators(exprs), Column: q.appendColumnPrefix("\"id\"")}
	q.where = append(q.where, where)
	return q
}

func (q messageSelectSQL) IDIn(vs ...MessageID) messageSelectSQL {
	_vs := make([]int64, 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, int64(v))
	}
	where := sqlla.ExprMultiValue[int64]{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("\"id\"")}
	q.where = append(q.where, where)
	return q
}

func (q messageSelectSQL) PkColumn(pk int64, exprs ...sqlla.Operator) messageSelectSQL {
	v := MessageID(pk)
	return q.ID(v, exprs...)
}

func (q messageSelectSQL) OrderByID(order sqlla.Order) messageSelectSQL {
	q.order = order.WithColumn(q.appendColumnPrefix("\"id\""))
	return q
}

func (q messageSelectSQL) Author(v string, exprs ...sqlla.Operator) messageSelectSQL {
	where := sqlla.ExprValue[string]{Value: v, Op: sqlla.Operators(exprs), Column: q.appendColumnPrefix("\"author\"")}
	q.where = append(q.where, where)
	return q
}

func (q messageSelectSQL) AuthorIn(vs ...string) messageSelectSQL {
	where := sqlla.ExprMultiValue[string]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("\"author\"")}
	q.where = append(q.where, where)
	return q
}

func (q messageSelectSQL) OrderByAuthor(order sqlla.Order) messageSelectSQL {
	q.order = order.WithColumn(q.appendColumnPrefix("\"author\""))
	return q
}

func (q messageSelectSQL) Text(v string, exprs ...sqlla.Operator) messageSelectSQL {
	where := sqlla.ExprValue[string]{Value: v, Op: sqlla.Operators(exprs), Column: q.appendColumnPrefix("\"text\"")}
	q.where = append(q.where, where)
	return q
}

func (q messageSelectSQL) TextIn(vs ...string) messageSelectSQL {
	where := sqlla.ExprMultiValue[string]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("\"text\"")}
	q.where = append(q.where, where)
	return q
}

func (q messageSelectSQL) OrderByText(order sqlla.Order) messageSelectSQL {
	q.order = order.WithColumn(q.appendColumnPrefix("\"text\""))
	return q
}

func (q messageSelectSQL) CreatedAt(v time.Time, exprs ...sqlla.Operator) messageSelectSQL {
	where := sqlla.ExprValue[time.Time]{Value: v, Op: sqlla.Operators(exprs), Column: q.appendColumnPrefix("\"created_at\"")}
	q.where = append(q.where, where)
	return q
}

func (q messageSelectSQL) CreatedAtIn(vs ...time.Time) messageSelectSQL {
	where := sqlla.ExprMultiValue[time.Time]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("\"created_at\"")}
	q.where = append(q.where, where)
	return q
}

func (q messageSelectSQL) OrderByCreatedAt(order sqlla.Order) messageSelectSQL {
	q.order = order.WithColumn(q.appendColumnPrefix("\"created_at\""))
	return q
}

func (q messageSelectSQL) UpdatedAt(v time.Time, exprs ...sqlla.Operator) messageSelectSQL {
	where := sqlla.ExprValue[time.Time]{Value: v, Op: sqlla.Operators(exprs), Column: q.appendColumnPrefix("\"updated_at\"")}
	q.where = append(q.where, where)
	return q
}

func (q messageSelectSQL) UpdatedAtIn(vs ...time.Time) messageSelectSQL {
	where := sqlla.ExprMultiValue[time.Time]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: q.appendColumnPrefix("\"updated_at\"")}
	q.where = append(q.where, where)
	return q
}

func (q messageSelectSQL) OrderByUpdatedAt(order sqlla.Order) messageSelectSQL {
	q.order = order.WithColumn(q.appendColumnPrefix("\"updated_at\""))
	return q
}

func (q messageSelectSQL) ToSql() (string, []interface{}, error) {
	columns := strings.Join(q.Columns, ", ")
	wheres, offset, vs, err := q.where.ToSqlPg(0)
	if err != nil {
		return "", nil, err
	}

	tableName := "\"messages\""
	if q.tableAlias != "" {
		tableName = tableName + " AS " + q.tableAlias
		pcs := make([]string, 0, len(q.Columns))
		for _, column := range q.Columns {
			pcs = append(pcs, q.appendColumnPrefix(column))
		}
		columns = strings.Join(pcs, ", ")
	}
	query := "SELECT " + columns + " FROM " + tableName
	if len(q.joinClauses) > 0 {
		jc := strings.Join(q.joinClauses, " ")
		query += " " + jc
	}
	if wheres != "" {
		query += " WHERE" + wheres
	}
	if q.additionalWhereClause != nil {
		_query, _offset, _args := q.additionalWhereClause(offset)
		query += " " + _query
		if len(_args) > 0 {
			vs = append(vs, _args...)
		}
		offset = _offset
	}
	if len(q.groupByColumns) > 0 {
		query += " GROUP BY "
		gbcs := make([]string, 0, len(q.groupByColumns))
		for _, column := range q.groupByColumns {
			gbcs = append(gbcs, q.appendColumnPrefix(column))
		}
		query += strings.Join(gbcs, ", ")
	}
	if q.order != nil {
		_query, _ := q.order.OrderExprPg(offset)
		query += " ORDER BY " + _query
		vs = append(vs, q.order.Values()...)
	}
	if q.limit != nil {
		query += " LIMIT " + strconv.FormatUint(*q.limit, 10)
	}
	if q.offset != nil {
		query += " OFFSET " + strconv.FormatUint(*q.offset, 10)
	}

	if q.isForUpdate {
		query += " FOR UPDATE"
	}

	return query + ";", vs, nil
}

func (s Message) Select() messageSelectSQL {
	return NewMessageSQL().Select().ID(s.ID)
}
func (q messageSelectSQL) Single(db sqlla.DB) (Message, error) {
	q.Columns = messageAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return Message{}, err
	}

	row := db.QueryRow(query, args...)
	return q.Scan(row)
}

func (q messageSelectSQL) SingleContext(ctx context.Context, db sqlla.DB) (Message, error) {
	q.Columns = messageAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return Message{}, err
	}

	row := db.QueryRowContext(ctx, query, args...)
	return q.Scan(row)
}

func (q messageSelectSQL) All(db sqlla.DB) ([]Message, error) {
	rs := make([]Message, 0, 10)
	q.Columns = messageAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		r, err := q.Scan(rows)
		if err != nil {
			return nil, err
		}
		rs = append(rs, r)
	}
	return rs, nil
}

func (q messageSelectSQL) AllContext(ctx context.Context, db sqlla.DB) ([]Message, error) {
	rs := make([]Message, 0, 10)
	q.Columns = messageAllColumns
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		r, err := q.Scan(rows)
		if err != nil {
			return nil, err
		}
		rs = append(rs, r)
	}
	return rs, nil
}

func (q messageSelectSQL) Scan(s sqlla.Scanner) (Message, error) {
	var row Message
	err := s.Scan(
		&row.ID,
		&row.Author,
		&row.Text,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	return row, err
}

// IterContext returns iter.Seq2[Message, error] and closer.
//
// The returned Iter.Seq2 assembles and executes a query in the first iteration.
// Therefore, the first iteration may return an error in assembling or executing the query.
// Subsequent iterations read rows. Again, the read may return an error.
//
// closer is a function that closes the row reader object. Execution of this function is idempotent.
// Be sure to call it when you are done using iter.Seq2.
func (q messageSelectSQL) IterContext(ctx context.Context, db sqlla.DB) (func(func(Message, error) bool), func() error) {
	var rowClose func() error
	closer := func() error {
		if rowClose != nil {
			err := rowClose()
			rowClose = nil
			return err
		}
		return nil
	}

	q.Columns = messageAllColumns
	query, args, err := q.ToSql()
	return func(yield func(Message, error) bool) {
		if err != nil {
			var r Message
			yield(r, err)
			return
		}
		rows, err := db.QueryContext(ctx, query, args...)
		if err != nil {
			var r Message
			yield(r, err)
			return
		}
		rowClose = rows.Close
		for rows.Next() {
			r, err := q.Scan(rows)
			if !yield(r, err) {
				break
			}
		}
	}, closer
}

type messageUpdateSQL struct {
	messageSQL
	setMap  sqlla.SetMap
	Columns []string
}

func (q messageSQL) Update() messageUpdateSQL {
	return messageUpdateSQL{
		messageSQL: q,
		setMap:     sqlla.SetMap{},
	}
}

func (q messageUpdateSQL) SetID(v MessageID) messageUpdateSQL {
	q.setMap["\"id\""] = int64(v)
	return q
}

func (q messageUpdateSQL) WhereID(v MessageID, exprs ...sqlla.Operator) messageUpdateSQL {
	where := sqlla.ExprValue[int64]{Value: int64(v), Op: sqlla.Operators(exprs), Column: "\"id\""}
	q.where = append(q.where, where)
	return q
}

func (q messageUpdateSQL) WhereIDIn(vs ...MessageID) messageUpdateSQL {
	_vs := make([]int64, 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, int64(v))
	}
	where := sqlla.ExprMultiValue[int64]{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: "\"id\""}
	q.where = append(q.where, where)
	return q
}

func (q messageUpdateSQL) SetAuthor(v string) messageUpdateSQL {
	q.setMap["\"author\""] = v
	return q
}

func (q messageUpdateSQL) WhereAuthor(v string, exprs ...sqlla.Operator) messageUpdateSQL {
	where := sqlla.ExprValue[string]{Value: v, Op: sqlla.Operators(exprs), Column: "\"author\""}
	q.where = append(q.where, where)
	return q
}

func (q messageUpdateSQL) WhereAuthorIn(vs ...string) messageUpdateSQL {
	where := sqlla.ExprMultiValue[string]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "\"author\""}
	q.where = append(q.where, where)
	return q
}

func (q messageUpdateSQL) SetText(v string) messageUpdateSQL {
	q.setMap["\"text\""] = v
	return q
}

func (q messageUpdateSQL) WhereText(v string, exprs ...sqlla.Operator) messageUpdateSQL {
	where := sqlla.ExprValue[string]{Value: v, Op: sqlla.Operators(exprs), Column: "\"text\""}
	q.where = append(q.where, where)
	return q
}

func (q messageUpdateSQL) WhereTextIn(vs ...string) messageUpdateSQL {
	where := sqlla.ExprMultiValue[string]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "\"text\""}
	q.where = append(q.where, where)
	return q
}

func (q messageUpdateSQL) SetCreatedAt(v time.Time) messageUpdateSQL {
	q.setMap["\"created_at\""] = v
	return q
}

func (q messageUpdateSQL) WhereCreatedAt(v time.Time, exprs ...sqlla.Operator) messageUpdateSQL {
	where := sqlla.ExprValue[time.Time]{Value: v, Op: sqlla.Operators(exprs), Column: "\"created_at\""}
	q.where = append(q.where, where)
	return q
}

func (q messageUpdateSQL) WhereCreatedAtIn(vs ...time.Time) messageUpdateSQL {
	where := sqlla.ExprMultiValue[time.Time]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "\"created_at\""}
	q.where = append(q.where, where)
	return q
}

func (q messageUpdateSQL) SetUpdatedAt(v time.Time) messageUpdateSQL {
	q.setMap["\"updated_at\""] = v
	return q
}

func (q messageUpdateSQL) WhereUpdatedAt(v time.Time, exprs ...sqlla.Operator) messageUpdateSQL {
	where := sqlla.ExprValue[time.Time]{Value: v, Op: sqlla.Operators(exprs), Column: "\"updated_at\""}
	q.where = append(q.where, where)
	return q
}

func (q messageUpdateSQL) WhereUpdatedAtIn(vs ...time.Time) messageUpdateSQL {
	where := sqlla.ExprMultiValue[time.Time]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "\"updated_at\""}
	q.where = append(q.where, where)
	return q
}

func (q messageUpdateSQL) ToSql() (string, []interface{}, error) {
	var err error
	var s interface{} = Message{}
	if t, ok := s.(messageDefaultUpdateHooker); ok {
		q, err = t.DefaultUpdateHook(q)
		if err != nil {
			return "", []interface{}{}, err
		}
	}
	setColumns, offset, svs, err := q.setMap.ToUpdateSqlPg(0)
	if err != nil {
		return "", []interface{}{}, err
	}
	wheres, _, wvs, err := q.where.ToSqlPg(offset)
	if err != nil {
		return "", []interface{}{}, err
	}

	query := "UPDATE " + "\"messages\"" + " SET" + setColumns
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", append(svs, wvs...), nil
}
func (s Message) Update() messageUpdateSQL {
	return NewMessageSQL().Update().WhereID(s.ID)
}

func (q messageUpdateSQL) Exec(db sqlla.DB) ([]Message, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	qq := q.messageSQL

	return qq.Select().All(db)
}

func (q messageUpdateSQL) ExecContext(ctx context.Context, db sqlla.DB) ([]Message, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	qq := q.messageSQL

	return qq.Select().AllContext(ctx, db)
}

type messageDefaultUpdateHooker interface {
	DefaultUpdateHook(messageUpdateSQL) (messageUpdateSQL, error)
}

type messageInsertSQL struct {
	messageSQL
	setMap  sqlla.SetMap
	Columns []string
}

func (q messageSQL) Insert() messageInsertSQL {
	return messageInsertSQL{
		messageSQL: q,
		setMap:     sqlla.SetMap{},
	}
}

func (q messageInsertSQL) ValueID(v MessageID) messageInsertSQL {
	q.setMap["\"id\""] = int64(v)
	return q
}

func (q messageInsertSQL) ValueAuthor(v string) messageInsertSQL {
	q.setMap["\"author\""] = v
	return q
}

func (q messageInsertSQL) ValueText(v string) messageInsertSQL {
	q.setMap["\"text\""] = v
	return q
}

func (q messageInsertSQL) ValueCreatedAt(v time.Time) messageInsertSQL {
	q.setMap["\"created_at\""] = v
	return q
}

func (q messageInsertSQL) ValueUpdatedAt(v time.Time) messageInsertSQL {
	q.setMap["\"updated_at\""] = v
	return q
}

func (q messageInsertSQL) ToSql() (string, []any, error) {
	query, _, vs, err := q.messageInsertSQLToSqlPg(0)
	if err != nil {
		return "", []any{}, err
	}
	return query + " RETURNING " + "\"id\"" + ";", vs, nil
}

func (q messageInsertSQL) messageInsertSQLToSqlPg(offset int) (string, int, []any, error) {
	var err error
	var s interface{} = Message{}
	if t, ok := s.(messageDefaultInsertHooker); ok {
		q, err = t.DefaultInsertHook(q)
		if err != nil {
			return "", 0, []any{}, err
		}
	}
	qs, offset, vs, err := q.setMap.ToInsertSqlPg(offset)
	if err != nil {
		return "", 0, []any{}, err
	}

	query := "INSERT INTO " + "\"messages\"" + " " + qs
	return query, offset, vs, nil
}

func (q messageInsertSQL) Exec(db sqlla.DB) (Message, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return Message{}, err
	}
	row := db.QueryRow(query, args...)
	var pk MessageID
	if err := row.Scan(&pk); err != nil {
		return Message{}, err
	}
	return NewMessageSQL().Select().ID(pk).Single(db)
}

func (q messageInsertSQL) ExecContext(ctx context.Context, db sqlla.DB) (Message, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return Message{}, err
	}
	row := db.QueryRowContext(ctx, query, args...)
	var pk MessageID
	if err := row.Scan(&pk); err != nil {
		return Message{}, err
	}
	return NewMessageSQL().Select().ID(pk).SingleContext(ctx, db)
}

func (q messageInsertSQL) ExecContextWithoutSelect(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	return result, err
}

type messageDefaultInsertHooker interface {
	DefaultInsertHook(messageInsertSQL) (messageInsertSQL, error)
}

type messageInsertSQLToSqler interface {
	messageInsertSQLToSqlPg(offset int) (string, int, []any, error)
}

type messageBulkInsertSQL struct {
	insertSQLs []messageInsertSQL
}

func (q messageSQL) BulkInsert() *messageBulkInsertSQL {
	return &messageBulkInsertSQL{
		insertSQLs: []messageInsertSQL{},
	}
}

func (q *messageBulkInsertSQL) Append(iqs ...messageInsertSQL) {
	q.insertSQLs = append(q.insertSQLs, iqs...)
}

func (q *messageBulkInsertSQL) messageInsertSQLToSqlPg(offset int) (string, int, []any, error) {
	if len(q.insertSQLs) == 0 {
		return "", 0, []any{}, fmt.Errorf("sqlla: This messageBulkInsertSQL's InsertSQL was empty")
	}
	iqs := make([]messageInsertSQL, len(q.insertSQLs))
	copy(iqs, q.insertSQLs)

	var s interface{} = Message{}
	if t, ok := s.(messageDefaultInsertHooker); ok {
		for i, iq := range iqs {
			var err error
			iq, err = t.DefaultInsertHook(iq)
			if err != nil {
				return "", 0, []any{}, err
			}
			iqs[i] = iq
		}
	}

	sms := make(sqlla.SetMaps, 0, len(q.insertSQLs))
	for _, iq := range q.insertSQLs {
		sms = append(sms, iq.setMap)
	}

	query, offset, vs, err := sms.ToInsertSqlPg(offset)
	if err != nil {
		return "", 0, []any{}, err
	}
	return "INSERT INTO " + "\"messages\"" + " " + query, offset, vs, nil
}

func (q *messageBulkInsertSQL) ToSql() (string, []any, error) {
	query, _, vs, err := q.messageInsertSQLToSqlPg(0)
	if err != nil {
		return "", []any{}, err
	}
	return query + " RETURNING " + "\"id\"" + ";", vs, nil
}
func (q *messageBulkInsertSQL) ExecContext(ctx context.Context, db sqlla.DB) ([]Message, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	pks := make([]MessageID, 0, len(q.insertSQLs))
	for rows.Next() {
		var pk MessageID
		if err := rows.Scan(&pk); err != nil {
			return nil, err
		}
		pks = append(pks, pk)
	}
	return NewMessageSQL().Select().IDIn(pks...).AllContext(ctx, db)
}
func (q *messageBulkInsertSQL) ExecContextWithoutSelect(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	return result, err
}

type messageInsertOnConflictDoNothingSQL struct {
	insertSQL messageInsertSQLToSqler
}

func (q messageInsertSQL) OnConflictDoNothing() messageInsertOnConflictDoNothingSQL {
	return messageInsertOnConflictDoNothingSQL{
		insertSQL: q,
	}
}

func (q messageInsertOnConflictDoNothingSQL) ToSql() (string, []any, error) {
	query, _, vs, err := q.insertSQL.messageInsertSQLToSqlPg(0)
	if err != nil {
		return "", nil, err
	}
	query += " ON CONFLICT DO NOTHING"
	query += " RETURNING " + "\"id\""
	return query + ";", vs, nil

}

func (q messageInsertOnConflictDoNothingSQL) ExecContext(ctx context.Context, db sqlla.DB) (Message, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return Message{}, err
	}
	row := db.QueryRowContext(ctx, query, args...)
	var pk MessageID
	if err := row.Scan(&pk); err != nil {
		return Message{}, err
	}
	return NewMessageSQL().Select().ID(pk).SingleContext(ctx, db)

}

func (q messageInsertOnConflictDoNothingSQL) ExecContextWithoutSelect(ctx context.Context, db sqlla.DB) (sql.Result, error) {

	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	return result, err

}

type messageInsertOnConflictDoUpdateSQL struct {
	insertSQL             messageInsertSQLToSqler
	onConflictDoUpdateMap sqlla.SetMap
	target                string
}

func (q messageInsertSQL) OnConflictDoUpdate(target string) messageInsertOnConflictDoUpdateSQL {
	return messageInsertOnConflictDoUpdateSQL{
		insertSQL:             q,
		onConflictDoUpdateMap: sqlla.SetMap{},
		target:                target,
	}
}

func (q messageInsertOnConflictDoUpdateSQL) ValueOnUpdateID(v MessageID) messageInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"id\""] = int64(v)
	return q
}

func (q messageInsertOnConflictDoUpdateSQL) RawValueOnUpdateID(v sqlla.SetMapRawValue) messageInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"id\""] = v
	return q
}

func (q messageInsertOnConflictDoUpdateSQL) SameOnUpdateID() messageInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"id\""] = sqlla.SetMapRawValue(`"excluded".` + "\"id\"")
	return q
}

func (q messageInsertOnConflictDoUpdateSQL) ValueOnUpdateAuthor(v string) messageInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"author\""] = v
	return q
}

func (q messageInsertOnConflictDoUpdateSQL) RawValueOnUpdateAuthor(v sqlla.SetMapRawValue) messageInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"author\""] = v
	return q
}

func (q messageInsertOnConflictDoUpdateSQL) SameOnUpdateAuthor() messageInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"author\""] = sqlla.SetMapRawValue(`"excluded".` + "\"author\"")
	return q
}

func (q messageInsertOnConflictDoUpdateSQL) ValueOnUpdateText(v string) messageInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"text\""] = v
	return q
}

func (q messageInsertOnConflictDoUpdateSQL) RawValueOnUpdateText(v sqlla.SetMapRawValue) messageInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"text\""] = v
	return q
}

func (q messageInsertOnConflictDoUpdateSQL) SameOnUpdateText() messageInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"text\""] = sqlla.SetMapRawValue(`"excluded".` + "\"text\"")
	return q
}

func (q messageInsertOnConflictDoUpdateSQL) ValueOnUpdateCreatedAt(v time.Time) messageInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"created_at\""] = v
	return q
}

func (q messageInsertOnConflictDoUpdateSQL) RawValueOnUpdateCreatedAt(v sqlla.SetMapRawValue) messageInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"created_at\""] = v
	return q
}

func (q messageInsertOnConflictDoUpdateSQL) SameOnUpdateCreatedAt() messageInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"created_at\""] = sqlla.SetMapRawValue(`"excluded".` + "\"created_at\"")
	return q
}

func (q messageInsertOnConflictDoUpdateSQL) ValueOnUpdateUpdatedAt(v time.Time) messageInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"updated_at\""] = v
	return q
}

func (q messageInsertOnConflictDoUpdateSQL) RawValueOnUpdateUpdatedAt(v sqlla.SetMapRawValue) messageInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"updated_at\""] = v
	return q
}

func (q messageInsertOnConflictDoUpdateSQL) SameOnUpdateUpdatedAt() messageInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"updated_at\""] = sqlla.SetMapRawValue(`"excluded".` + "\"updated_at\"")
	return q
}

func (q messageInsertOnConflictDoUpdateSQL) ToSql() (string, []any, error) {
	var err error
	var s any = Message{}
	if t, ok := s.(messageDefaultInsertOnConflictDoUpdateHooker); ok {
		q, err = t.DefaultInsertOnConflictDoUpdateHook(q)
		if err != nil {
			return "", nil, err
		}
	}

	query, offset, vs, err := q.insertSQL.messageInsertSQLToSqlPg(0)
	if err != nil {
		return "", nil, err
	}

	os, _, ovs, err := q.onConflictDoUpdateMap.ToUpdateSqlPg(offset)
	if err != nil {
		return "", nil, err
	}
	query += " ON CONFLICT (" + q.target + ") DO UPDATE SET" + os
	vs = append(vs, ovs...)
	query += " RETURNING " + "\"id\""

	return query + ";", vs, nil
}

func (q messageInsertOnConflictDoUpdateSQL) ExecContext(ctx context.Context, db sqlla.DB) (Message, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return Message{}, err
	}
	row := db.QueryRowContext(ctx, query, args...)
	var pk MessageID
	if err := row.Scan(&pk); err != nil {
		return Message{}, err
	}
	return NewMessageSQL().Select().ID(pk).SingleContext(ctx, db)

}

func (q messageInsertOnConflictDoUpdateSQL) ExecContextWithoutSelect(ctx context.Context, db sqlla.DB) (sql.Result, error) {

	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	return result, err

}

type messageDefaultInsertOnConflictDoUpdateHooker interface {
	DefaultInsertOnConflictDoUpdateHook(messageInsertOnConflictDoUpdateSQL) (messageInsertOnConflictDoUpdateSQL, error)
}

type messageBulkInsertOnConflictDoNothingSQL struct {
	insertSQL messageInsertSQLToSqler
}

func (q *messageBulkInsertSQL) OnConflictDoNothing() messageBulkInsertOnConflictDoNothingSQL {
	return messageBulkInsertOnConflictDoNothingSQL{
		insertSQL: q,
	}
}

func (q messageBulkInsertOnConflictDoNothingSQL) ToSql() (string, []any, error) {
	query, _, vs, err := q.insertSQL.messageInsertSQLToSqlPg(0)
	if err != nil {
		return "", nil, err
	}
	query += " ON CONFLICT DO NOTHING"
	query += " RETURNING " + "\"id\""
	return query + ";", vs, nil

}

func (q messageBulkInsertOnConflictDoNothingSQL) ExecContext(ctx context.Context, db sqlla.DB) ([]Message, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	pks := make([]MessageID, 0)
	for rows.Next() {
		var pk MessageID
		if err := rows.Scan(&pk); err != nil {
			return nil, err
		}
		pks = append(pks, pk)
	}

	return NewMessageSQL().Select().IDIn(pks...).AllContext(ctx, db)

}

func (q messageBulkInsertOnConflictDoNothingSQL) ExecContextWithoutSelect(ctx context.Context, db sqlla.DB) (sql.Result, error) {

	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	return result, err

}

type messageBulkInsertOnConflictDoUpdateSQL struct {
	insertSQL             messageInsertSQLToSqler
	onConflictDoUpdateMap sqlla.SetMap
	target                string
}

func (q *messageBulkInsertSQL) OnConflictDoUpdate(target string) messageBulkInsertOnConflictDoUpdateSQL {
	return messageBulkInsertOnConflictDoUpdateSQL{
		insertSQL:             q,
		onConflictDoUpdateMap: sqlla.SetMap{},
		target:                target,
	}
}

func (q messageBulkInsertOnConflictDoUpdateSQL) ValueOnUpdateID(v MessageID) messageBulkInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"id\""] = int64(v)
	return q
}

func (q messageBulkInsertOnConflictDoUpdateSQL) RawValueOnUpdateID(v sqlla.SetMapRawValue) messageBulkInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"id\""] = v
	return q
}

func (q messageBulkInsertOnConflictDoUpdateSQL) SameOnUpdateID() messageBulkInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"id\""] = sqlla.SetMapRawValue(`"excluded".` + "\"id\"")
	return q
}

func (q messageBulkInsertOnConflictDoUpdateSQL) ValueOnUpdateAuthor(v string) messageBulkInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"author\""] = v
	return q
}

func (q messageBulkInsertOnConflictDoUpdateSQL) RawValueOnUpdateAuthor(v sqlla.SetMapRawValue) messageBulkInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"author\""] = v
	return q
}

func (q messageBulkInsertOnConflictDoUpdateSQL) SameOnUpdateAuthor() messageBulkInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"author\""] = sqlla.SetMapRawValue(`"excluded".` + "\"author\"")
	return q
}

func (q messageBulkInsertOnConflictDoUpdateSQL) ValueOnUpdateText(v string) messageBulkInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"text\""] = v
	return q
}

func (q messageBulkInsertOnConflictDoUpdateSQL) RawValueOnUpdateText(v sqlla.SetMapRawValue) messageBulkInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"text\""] = v
	return q
}

func (q messageBulkInsertOnConflictDoUpdateSQL) SameOnUpdateText() messageBulkInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"text\""] = sqlla.SetMapRawValue(`"excluded".` + "\"text\"")
	return q
}

func (q messageBulkInsertOnConflictDoUpdateSQL) ValueOnUpdateCreatedAt(v time.Time) messageBulkInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"created_at\""] = v
	return q
}

func (q messageBulkInsertOnConflictDoUpdateSQL) RawValueOnUpdateCreatedAt(v sqlla.SetMapRawValue) messageBulkInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"created_at\""] = v
	return q
}

func (q messageBulkInsertOnConflictDoUpdateSQL) SameOnUpdateCreatedAt() messageBulkInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"created_at\""] = sqlla.SetMapRawValue(`"excluded".` + "\"created_at\"")
	return q
}

func (q messageBulkInsertOnConflictDoUpdateSQL) ValueOnUpdateUpdatedAt(v time.Time) messageBulkInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"updated_at\""] = v
	return q
}

func (q messageBulkInsertOnConflictDoUpdateSQL) RawValueOnUpdateUpdatedAt(v sqlla.SetMapRawValue) messageBulkInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"updated_at\""] = v
	return q
}

func (q messageBulkInsertOnConflictDoUpdateSQL) SameOnUpdateUpdatedAt() messageBulkInsertOnConflictDoUpdateSQL {
	q.onConflictDoUpdateMap["\"updated_at\""] = sqlla.SetMapRawValue(`"excluded".` + "\"updated_at\"")
	return q
}

func (q messageBulkInsertOnConflictDoUpdateSQL) ToSql() (string, []any, error) {
	var s any = Message{}
	if t, ok := s.(messageDefaultInsertOnConflictDoUpdateHooker); ok {
		sq := messageInsertOnConflictDoUpdateSQL{
			insertSQL:             q.insertSQL,
			onConflictDoUpdateMap: q.onConflictDoUpdateMap,
			target:                q.target,
		}
		sq, err := t.DefaultInsertOnConflictDoUpdateHook(sq)
		if err != nil {
			return "", nil, err
		}
		q.insertSQL = sq.insertSQL
		q.onConflictDoUpdateMap = sq.onConflictDoUpdateMap
		q.target = sq.target
	}

	query, offset, vs, err := q.insertSQL.messageInsertSQLToSqlPg(0)
	if err != nil {
		return "", nil, err
	}

	os, _, ovs, err := q.onConflictDoUpdateMap.ToUpdateSqlPg(offset)
	if err != nil {
		return "", nil, err
	}
	query += " ON CONFLICT (" + q.target + ") DO UPDATE SET" + os
	vs = append(vs, ovs...)
	query += " RETURNING " + "\"id\""

	return query + ";", vs, nil
}

func (q messageBulkInsertOnConflictDoUpdateSQL) ExecContext(ctx context.Context, db sqlla.DB) ([]Message, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	pks := make([]MessageID, 0)
	for rows.Next() {
		var pk MessageID
		if err := rows.Scan(&pk); err != nil {
			return nil, err
		}
		pks = append(pks, pk)
	}

	return NewMessageSQL().Select().IDIn(pks...).AllContext(ctx, db)

}

func (q messageBulkInsertOnConflictDoUpdateSQL) ExecContextWithoutSelect(ctx context.Context, db sqlla.DB) (sql.Result, error) {

	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	result, err := db.ExecContext(ctx, query, args...)
	return result, err

}

type messageDeleteSQL struct {
	messageSQL
}

func (q messageSQL) Delete() messageDeleteSQL {
	return messageDeleteSQL{
		q,
	}
}

func (q messageDeleteSQL) ID(v MessageID, exprs ...sqlla.Operator) messageDeleteSQL {
	where := sqlla.ExprValue[int64]{Value: int64(v), Op: sqlla.Operators(exprs), Column: "\"id\""}
	q.where = append(q.where, where)
	return q
}

func (q messageDeleteSQL) IDIn(vs ...MessageID) messageDeleteSQL {
	_vs := make([]int64, 0, len(vs))
	for _, v := range vs {
		_vs = append(_vs, int64(v))
	}
	where := sqlla.ExprMultiValue[int64]{Values: _vs, Op: sqlla.MakeInOperator(len(vs)), Column: "\"id\""}
	q.where = append(q.where, where)
	return q
}

func (q messageDeleteSQL) Author(v string, exprs ...sqlla.Operator) messageDeleteSQL {
	where := sqlla.ExprValue[string]{Value: v, Op: sqlla.Operators(exprs), Column: "\"author\""}
	q.where = append(q.where, where)
	return q
}

func (q messageDeleteSQL) AuthorIn(vs ...string) messageDeleteSQL {
	where := sqlla.ExprMultiValue[string]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "\"author\""}
	q.where = append(q.where, where)
	return q
}

func (q messageDeleteSQL) Text(v string, exprs ...sqlla.Operator) messageDeleteSQL {
	where := sqlla.ExprValue[string]{Value: v, Op: sqlla.Operators(exprs), Column: "\"text\""}
	q.where = append(q.where, where)
	return q
}

func (q messageDeleteSQL) TextIn(vs ...string) messageDeleteSQL {
	where := sqlla.ExprMultiValue[string]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "\"text\""}
	q.where = append(q.where, where)
	return q
}

func (q messageDeleteSQL) CreatedAt(v time.Time, exprs ...sqlla.Operator) messageDeleteSQL {
	where := sqlla.ExprValue[time.Time]{Value: v, Op: sqlla.Operators(exprs), Column: "\"created_at\""}
	q.where = append(q.where, where)
	return q
}

func (q messageDeleteSQL) CreatedAtIn(vs ...time.Time) messageDeleteSQL {
	where := sqlla.ExprMultiValue[time.Time]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "\"created_at\""}
	q.where = append(q.where, where)
	return q
}

func (q messageDeleteSQL) UpdatedAt(v time.Time, exprs ...sqlla.Operator) messageDeleteSQL {
	where := sqlla.ExprValue[time.Time]{Value: v, Op: sqlla.Operators(exprs), Column: "\"updated_at\""}
	q.where = append(q.where, where)
	return q
}

func (q messageDeleteSQL) UpdatedAtIn(vs ...time.Time) messageDeleteSQL {
	where := sqlla.ExprMultiValue[time.Time]{Values: vs, Op: sqlla.MakeInOperator(len(vs)), Column: "\"updated_at\""}
	q.where = append(q.where, where)
	return q
}

func (q messageDeleteSQL) ToSql() (string, []interface{}, error) {
	wheres, _, vs, err := q.where.ToSqlPg(0)
	if err != nil {
		return "", nil, err
	}

	query := "DELETE FROM " + "\"messages\""
	if wheres != "" {
		query += " WHERE" + wheres
	}

	return query + ";", vs, nil
}

func (q messageDeleteSQL) Exec(db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

func (q messageDeleteSQL) ExecContext(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	return db.ExecContext(ctx, query, args...)
}
func (s Message) Delete(db sqlla.DB) (sql.Result, error) {
	query, args, err := NewMessageSQL().Delete().ID(s.ID).ToSql()
	if err != nil {
		return nil, err
	}
	return db.Exec(query, args...)
}

func (s Message) DeleteContext(ctx context.Context, db sqlla.DB) (sql.Result, error) {
	query, args, err := NewMessageSQL().Delete().ID(s.ID).ToSql()
	if err != nil {
		return nil, err
	}
	return db.ExecContext(ctx, query, args...)
}
