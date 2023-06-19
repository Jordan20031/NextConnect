
type User struct {
	ID       int64
	Username string
	Email    string
	Password string
	Role     string
	Pdp      []byte
}

type discussions struct {
    ID int64
    image []byte
	titre string
    description string
    nmbreDeLikes int64
    idUser int64
}

type messages struct (
    ID int64
    text string
    IDcreateur int64
    IDdiscution int64
)

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

func updateUsername(db *sql.DB, userID int64, newUsername string) error {
	updateQuery := "UPDATE users SET Username = ? WHERE ID = ?"
	_, err := db.Exec(updateQuery, newUsername, user.ID)
	if err != nil {
		return err
	}
	user.Username = newUsername
	return nil
}

func insertUser(db *sql.DB, user *User) error {
	// Vérification de l'existence d'un utilisateur avec le même nom d'utilisateur ou la même adresse e-mail
	existsQuery := "SELECT COUNT(*) FROM users WHERE Username = ? OR Email = ?"
	var count int
	err := db.QueryRow(existsQuery, user.Username, user.Email).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		// Un utilisateur avec le même nom d'utilisateur ou la même adresse e-mail existe déjà
		return fmt.Errorf("un utilisateur avec le même nom d'utilisateur ou la même adresse e-mail existe déjà")
	}

	// Insertion de l'utilisateur dans la base de données
	insertQuery := "INSERT INTO users (Username, Email, Password, Role, pdp) VALUES (?, ?, ?, ?, ?)"
	result, err := db.Exec(insertQuery, user.Username, user.Email, user.Password, user.Role, user.Pdp)
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

func getUserByUsername(db *sql.DB, username string) (User, error) {
	query := "SELECT ID, Username, Email, Password, Role, pdp FROM users WHERE Username = ?"

	var user User
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.Pdp)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func insertDiscussion(db *sql.DB, discussion *Discussions) error {
	insertQuery := "INSERT INTO discussions (image, titre, description, nmbreDeLikes, idUser) VALUES (?, ?, ?, ?, ?)"
	result, err := db.Exec(insertQuery, discussion.Image, discussion.Titre, discussion.Description, discussion.NmbreDeLikes, discussion.IDUser)
	if err != nil {
		return err
	}
	discussion.ID, _ = result.LastInsertId()
	return nil
}

func getDiscussionByID(db *sql.DB, discussionID int64) (Discussions, error) {
	query := "SELECT ID, image, titre, description, nmbreDeLikes, idUser FROM discussions WHERE ID = ?"

	var discussion Discussions
	err := db.QueryRow(query, discussionID).Scan(&discussion.ID, &discussion.Image, &discussion.Titre, &discussion.Description, &discussion.NmbreDeLikes, &discussion.IDUser)
	if err != nil {
		return Discussions{}, err
	}

	return discussion, nil
}

func insertMessage(db *sql.DB, message *Messages) error {
	insertQuery := "INSERT INTO messages (text, IDcreateur, IDdiscution) VALUES (?, ?, ?)"
	result, err := db.Exec(insertQuery, message.Text, message.IDCreateur, message.IDDiscution)
	if err != nil {
		return err
	}
	message.ID, _ = result.LastInsertId()
	return nil
}

func getMessageByID(db *sql.DB, messageID int64) (Messages, error) {
	query := "SELECT ID, text, IDcreateur, IDdiscution FROM messages WHERE ID = ?"

	var message Messages
	err := db.QueryRow(query, messageID).Scan(&message.ID, &message.Text, &message.IDCreateur, &message.IDDiscution)
	if err != nil {
		return Messages{}, err
	}

	return message, nil
}