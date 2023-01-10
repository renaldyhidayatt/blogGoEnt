// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "slug", Type: field.TypeString},
		{Name: "img", Type: field.TypeString},
		{Name: "body", Type: field.TypeString},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "firstname", Type: field.TypeString},
		{Name: "lastname", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// PostCategoriesColumns holds the columns for the "post_categories" table.
	PostCategoriesColumns = []*schema.Column{
		{Name: "post_id", Type: field.TypeInt},
		{Name: "category_id", Type: field.TypeInt},
	}
	// PostCategoriesTable holds the schema information for the "post_categories" table.
	PostCategoriesTable = &schema.Table{
		Name:       "post_categories",
		Columns:    PostCategoriesColumns,
		PrimaryKey: []*schema.Column{PostCategoriesColumns[0], PostCategoriesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_categories_post_id",
				Columns:    []*schema.Column{PostCategoriesColumns[0]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "post_categories_category_id",
				Columns:    []*schema.Column{PostCategoriesColumns[1]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserPostsColumns holds the columns for the "user_posts" table.
	UserPostsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "post_id", Type: field.TypeInt},
	}
	// UserPostsTable holds the schema information for the "user_posts" table.
	UserPostsTable = &schema.Table{
		Name:       "user_posts",
		Columns:    UserPostsColumns,
		PrimaryKey: []*schema.Column{UserPostsColumns[0], UserPostsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_posts_user_id",
				Columns:    []*schema.Column{UserPostsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_posts_post_id",
				Columns:    []*schema.Column{UserPostsColumns[1]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CategoriesTable,
		PostsTable,
		UsersTable,
		PostCategoriesTable,
		UserPostsTable,
	}
)

func init() {
	PostCategoriesTable.ForeignKeys[0].RefTable = PostsTable
	PostCategoriesTable.ForeignKeys[1].RefTable = CategoriesTable
	UserPostsTable.ForeignKeys[0].RefTable = UsersTable
	UserPostsTable.ForeignKeys[1].RefTable = PostsTable
}
