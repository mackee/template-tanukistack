package record

import (
	"context"
	"fmt"
	"time"

	"github.com/carlmjohnson/errorx"
	"github.com/mackee/go-genddl/index"
	"github.com/mackee/go-sqlla/v2"
	"github.com/samber/lo"
)

type MessageID int64

//genddl:table messages
//sqlla:table messages
//sqlla:plugin table get=ID create=Author,Text
//sqlla:plugin slice
//sqlla:plugin timeHooks create=CreatedAt,UpdatedAt update=UpdatedAt
type Message struct {
	ID        MessageID `db:"id,primarykey,autoincrement" json:"id"`
	Author    string    `db:"author" json:"author"`
	Text      string    `db:"text" json:"text"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (m Message) _schemaIndex(methods index.Methods) []index.Definition {
	return []index.Definition{
		methods.Complex(m.Author, m.CreatedAt),
	}
}

func (m *MessageTable) ListByMessages(ctx context.Context, db sqlla.DB, startID int64) (_resp Messages, err error) {
	defer errorx.Trace(&err)
	q := NewMessageSQL().Select()
	if startID > 0 {
		q = q.ID(MessageID(startID), sqlla.OpLess)
	}
	q = q.OrderByID(sqlla.Desc).Limit(10)
	rows, err := q.AllContext(ctx, db)
	if err != nil {
		return nil, fmt.Errorf("failed to list Messages: %w", err)
	}
	return lo.ToSlicePtr(rows), nil
}
