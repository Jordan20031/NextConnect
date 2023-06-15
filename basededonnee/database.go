
type User struct {
	ID       int64
	Username string
	Email    string
	Password string
	Role     string
}

func updateRole(db *sql.DB, user *User, newRole string) error {
	updateRole := "UPDATE users SET Role = ? WHERE ID = ?"
	_, err := db.Exec(updateRole, newRole, user.ID)
	if err != nil {
		return err
	}
	user.Role = newRole
	return nil
}

func deleteUser(db *sql.DB, user *User) error {
	deleteUser := "DELETE FROM users WHERE ID = ?"
	_, err := db.Exec(deleteUser, user.ID)
	return err
}

func insertUser(db *sql.DB, user *User) error {
	insertUser := "INSERT INTO users (Username, Email, Password, Role) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(insertUser, user.Username, user.Email, user.Password, user.Role)
	if err != nil {
		return err
	}
	user.ID, _ = result.LastInsertId()
	return nil
}

func insertUser(db *sql.DB, user *User) error {
	insertUser := "INSERT INTO users (Username, Email, Password, Role) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(insertUser, user.Username, user.Email, user.Password, user.Role)
	if err != nil {
		return err
	}
	user.ID, _ = result.LastInsertId()
	return nil
}

func updatePassword(db *sql.DB, user *User, newPassword string) error {
	updatePassword := "UPDATE users SET Password = ? WHERE ID = ?"
	_, err := db.Exec(updatePassword, newPassword, user.ID)
	if err != nil {
		return err
	}
	user.Password = newPassword
	return nil
}

