// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/0xfzz/tuwitt/ent/threadcount"
)

// ThreadCount is the model entity for the ThreadCount schema.
type ThreadCount struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ReplyCount holds the value of the "reply_count" field.
	ReplyCount int `json:"reply_count,omitempty"`
	// LikeCount holds the value of the "like_count" field.
	LikeCount    int `json:"like_count,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ThreadCount) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case threadcount.FieldID, threadcount.FieldReplyCount, threadcount.FieldLikeCount:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ThreadCount fields.
func (tc *ThreadCount) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case threadcount.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tc.ID = int(value.Int64)
		case threadcount.FieldReplyCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field reply_count", values[i])
			} else if value.Valid {
				tc.ReplyCount = int(value.Int64)
			}
		case threadcount.FieldLikeCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field like_count", values[i])
			} else if value.Valid {
				tc.LikeCount = int(value.Int64)
			}
		default:
			tc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ThreadCount.
// This includes values selected through modifiers, order, etc.
func (tc *ThreadCount) Value(name string) (ent.Value, error) {
	return tc.selectValues.Get(name)
}

// Update returns a builder for updating this ThreadCount.
// Note that you need to call ThreadCount.Unwrap() before calling this method if this ThreadCount
// was returned from a transaction, and the transaction was committed or rolled back.
func (tc *ThreadCount) Update() *ThreadCountUpdateOne {
	return NewThreadCountClient(tc.config).UpdateOne(tc)
}

// Unwrap unwraps the ThreadCount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tc *ThreadCount) Unwrap() *ThreadCount {
	_tx, ok := tc.config.driver.(*txDriver)
	if !ok {
		panic("ent: ThreadCount is not a transactional entity")
	}
	tc.config.driver = _tx.drv
	return tc
}

// String implements the fmt.Stringer.
func (tc *ThreadCount) String() string {
	var builder strings.Builder
	builder.WriteString("ThreadCount(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tc.ID))
	builder.WriteString("reply_count=")
	builder.WriteString(fmt.Sprintf("%v", tc.ReplyCount))
	builder.WriteString(", ")
	builder.WriteString("like_count=")
	builder.WriteString(fmt.Sprintf("%v", tc.LikeCount))
	builder.WriteByte(')')
	return builder.String()
}

// ThreadCounts is a parsable slice of ThreadCount.
type ThreadCounts []*ThreadCount
