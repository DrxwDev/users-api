package users

type User struct {
	ID        string
	Email     string
	Password  string
	CreatedAt string
	UpdatedAt string
}

type CreateUserParams struct {
	Email    string
	Password string
}

type UserRow struct {
	ID        string
	Email     string
	CreatedAt string
	UpdatedAt string
}
