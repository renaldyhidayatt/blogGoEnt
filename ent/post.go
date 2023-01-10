// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/renaldyhidayatt/blogGoEnt/ent/post"
)

// Post is the model entity for the Post schema.
type Post struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Slug holds the value of the "slug" field.
	Slug string `json:"slug,omitempty"`
	// Img holds the value of the "img" field.
	Img string `json:"img,omitempty"`
	// Body holds the value of the "body" field.
	Body string `json:"body,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PostQuery when eager-loading is set.
	Edges PostEdges `json:"edges"`
}

// PostEdges holds the relations/edges for other nodes in the graph.
type PostEdges struct {
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// Categories holds the value of the categories edge.
	Categories []*Category `json:"categories,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// CategoriesOrErr returns the Categories value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) CategoriesOrErr() ([]*Category, error) {
	if e.loadedTypes[1] {
		return e.Categories, nil
	}
	return nil, &NotLoadedError{edge: "categories"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Post) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case post.FieldID:
			values[i] = new(sql.NullInt64)
		case post.FieldTitle, post.FieldSlug, post.FieldImg, post.FieldBody:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Post", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Post fields.
func (po *Post) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case post.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int(value.Int64)
		case post.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				po.Title = value.String
			}
		case post.FieldSlug:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field slug", values[i])
			} else if value.Valid {
				po.Slug = value.String
			}
		case post.FieldImg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field img", values[i])
			} else if value.Valid {
				po.Img = value.String
			}
		case post.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				po.Body = value.String
			}
		}
	}
	return nil
}

// QueryUsers queries the "users" edge of the Post entity.
func (po *Post) QueryUsers() *UserQuery {
	return (&PostClient{config: po.config}).QueryUsers(po)
}

// QueryCategories queries the "categories" edge of the Post entity.
func (po *Post) QueryCategories() *CategoryQuery {
	return (&PostClient{config: po.config}).QueryCategories(po)
}

// Update returns a builder for updating this Post.
// Note that you need to call Post.Unwrap() before calling this method if this Post
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Post) Update() *PostUpdateOne {
	return (&PostClient{config: po.config}).UpdateOne(po)
}

// Unwrap unwraps the Post entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Post) Unwrap() *Post {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Post is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Post) String() string {
	var builder strings.Builder
	builder.WriteString("Post(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	builder.WriteString("title=")
	builder.WriteString(po.Title)
	builder.WriteString(", ")
	builder.WriteString("slug=")
	builder.WriteString(po.Slug)
	builder.WriteString(", ")
	builder.WriteString("img=")
	builder.WriteString(po.Img)
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(po.Body)
	builder.WriteByte(')')
	return builder.String()
}

// Posts is a parsable slice of Post.
type Posts []*Post

func (po Posts) config(cfg config) {
	for _i := range po {
		po[_i].config = cfg
	}
}