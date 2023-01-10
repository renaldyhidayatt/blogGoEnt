// Code generated by ent, DO NOT EDIT.

package post

const (
	// Label holds the string label denoting the post type in the database.
	Label = "post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldSlug holds the string denoting the slug field in the database.
	FieldSlug = "slug"
	// FieldImg holds the string denoting the img field in the database.
	FieldImg = "img"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeCategories holds the string denoting the categories edge name in mutations.
	EdgeCategories = "categories"
	// Table holds the table name of the post in the database.
	Table = "posts"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "user_posts"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// CategoriesTable is the table that holds the categories relation/edge. The primary key declared below.
	CategoriesTable = "post_categories"
	// CategoriesInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoriesInverseTable = "categories"
)

// Columns holds all SQL columns for post fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldSlug,
	FieldImg,
	FieldBody,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "post_id"}
	// CategoriesPrimaryKey and CategoriesColumn2 are the table columns denoting the
	// primary key for the categories relation (M2M).
	CategoriesPrimaryKey = []string{"post_id", "category_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}