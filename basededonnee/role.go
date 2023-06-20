package basededonnee

func canDeleteFromDB(user User) bool {
	return user.Role == "admin" || user.Role == "connected"
}

func canWriteToDB(user User) bool {
	return user.Role == "connected" || user.Role == "admin"
}

func deleteAccount(db *sql.DB, user User, userID int64) error {
	if !canDeleteFromDB(user) {
		return fmt.Errorf("seuls les utilisateurs avec le rôle 'admin' ou 'connected' peuvent supprimer des comptes")
	}

	switch user.Role {
	case "connected":
		if user.ID != userID {
			return fmt.Errorf("vous n'êtes autorisé à supprimer que votre propre compte")
		} else {
			// Suppression du compte de l'utilisateur dans la base de données
			err := deleteUserFromDB(db, userID)
			if err != nil {
				return err
			}

			// Suppression des messages associés à l'ID utilisateur
			err = deleteMessagesByUserID(db, userID)
			if err != nil {
				return err
			}

			// Suppression des discussions associées à l'ID utilisateur
			err = deleteDiscussionsByUserID(db, userID)
			if err != nil {
				return err
			}
		}
		break

	case "admin":
		// Suppression du compte de l'utilisateur dans la base de données
		err := deleteUserFromDB(db, userID)
		if err != nil {
			return err
		}

		// Suppression des messages associés à l'ID utilisateur
		err = deleteMessagesByUserID(db, userID)
		if err != nil {
			return err
		}

		// Suppression des discussions associées à l'ID utilisateur
		err = deleteDiscussionsByUserID(db, userID)
		if err != nil {
			return err
		}
		break
	}

	return nil
}

func deleteDiscussion(db *sql.DB, user User, discussionID int64) error {
	if !canDeleteFromDB(user) {
		return fmt.Errorf("seuls les utilisateurs avec le rôle 'admin' ou 'connected' peuvent supprimer des discution")
	}

	switch user.Role {
	case "connected":
		if user.ID != userID {
			return fmt.Errorf("vous n'êtes autorisé à supprimer que vos propres discutions")
		} else {
			// Suppression des messages associés à l'ID discussion
			err = deleteMessagesBydiscussionID(db, discussionID)
			if err != nil {
				return err
			}

			// Suppression des discussions associées à l'ID discussion
			err = deleteDiscussionsBydiscussionID(db, discussionID)
			if err != nil {
				return err
			}
		}
		break

	case "admin":
		// Suppression des messages associés à l'ID discussion
		err = deleteMessagesBydiscussionID(db, discussionID)
		if err != nil {
			return err
		}

		// Suppression des discussions associées à l'ID discussion
		err = deleteDiscussionsBydiscussionID(db, discussionID)
		if err != nil {
			return err
		}
		break
	}

	return nil
}

func deleteMessage(db *sql.DB, user User, messageID int64) error {
	if !canDeleteFromDB(user) {
		return fmt.Errorf("seuls les utilisateurs avec le rôle 'admin' ou 'connected' peuvent supprimer des discution")
	}

	switch user.Role {
	case "connected":
		if user.ID != userID {
			return fmt.Errorf("vous n'êtes autorisé à supprimer que vos propres discutions")
		} else {
			// Suppression des messages associés à l'ID discussion
			err = deleteMessagesBymessageID(db, messageID)
			if err != nil {
				return err
			}
		}
		break

	case "admin":
		// Suppression des messages associés à l'ID discussion
		err = deleteMessagesBymessageID(db, messageID)
		if err != nil {
			return err
		}
	}

	return nil
}

func createDiscussion(db *sql.DB, user User, image []byte, titre string, description string) error {
	if user.Role != "admin" && user.Role != "connected" {
		return fmt.Errorf("seuls les utilisateurs avec les rôles 'admin' ou 'connected' peuvent créer une discussion")
	}

	// Initialisation de la variable discussion
	discussion := discussions{
		ID:          0,             // L'ID sera automatiquement généré par la base de données
		image:       image,
		titre:       titre,
		description: description,
		nmbreDeLikes: 0,
		idUser:      user.userID
	}

	// Insérer la nouvelle discussion dans la base de données
	err := insertDiscussion(db, &discussion)
	if err != nil {
		return err
	}

	return nil
}

func createMessage(db *sql.DB, user User, text string, IDdiscution int64) error {
	if user.Role != "admin" && user.Role != "connected" {
		return fmt.Errorf("seuls les utilisateurs avec les rôles 'admin' ou 'connected' peuvent créer un message")
	}

	// Générer un nouveau message
	message := messages{
		ID:          0,             // L'ID sera automatiquement généré par la base de données
		Text:        text,
		idUser:      user.userID
		IDDiscution: IDdiscution,
	}

	// Insérer le nouveau message dans la base de données
	err := insertMessage(db, &message)
	if err != nil {
		return err
	}

	return nil
}