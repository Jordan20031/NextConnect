
func canDeleteFromDB(user User) bool {
	return user.Role == "admin" || user.Role == "connected"
}

func canWriteToDB(user User) bool {
	return user.Role == "connected" || user.Role == "admin"
}

func deleteAccount(db *sql.DB, user User, userID int64) error {
	if !canDeleteFromDB(user) {
		return fmt.Errorf("seuls les utilisateurs avec le rôle 'admin' peuvent supprimer des comptes autre que soi")
	}

	if user.Role == "connected" && user.ID != userID {
		return fmt.Errorf("vous n'êtes autorisé à supprimer que votre propre compte")
	}

	// Suppression du compte de l'utilisateur dans la base de données
	// ... (votre code de suppression ici)

	// Suppression des messages associés à l'ID utilisateur
	// ... (votre code de suppression ici)

	// Suppression des discussions associées à l'ID utilisateur
	// ... (votre code de suppression ici)

	return nil
}

func deleteDiscussion(db *sql.DB, user User, discussionID int64) error {
	if !canDeleteFromDB(user) {
		return fmt.Errorf("seuls les utilisateurs avec le rôle 'admin' peuvent supprimer des discussions")
	}

	if user.Role == "connected" {
		// Vérifier si l'utilisateur connecté est autorisé à supprimer la discussion en fonction de son ID
		// Vous devez implémenter la logique correspondante pour vérifier si l'ID de l'utilisateur connecté est lié à la discussion

		// Si l'utilisateur connecté n'est pas autorisé à supprimer la discussion, renvoyer une erreur
		// return fmt.Errorf("vous n'êtes pas autorisé à supprimer cette discussion")
	}

	// Suppression de la discussion de la base de données en utilisant l'ID spécifié
	// ... (votre code de suppression ici)

	return nil
}

func deleteMessage(db *sql.DB, user User, messageID int64) error {
	if !canDeleteFromDB(user) {
		return fmt.Errorf("seuls les utilisateurs avec le rôle 'admin' peuvent supprimer des messages")
	}

	if user.Role == "connected" {
		// Vérifier si l'utilisateur connecté est autorisé à supprimer le message en fonction de son ID
		// Vous devez implémenter la logique correspondante pour vérifier si l'ID de l'utilisateur connecté est lié au message

		// Si l'utilisateur connecté n'est pas autorisé à supprimer le message, renvoyer une erreur
		// return fmt.Errorf("vous n'êtes pas autorisé à supprimer ce message")
	}

	// Suppression du message de la base de données en utilisant l'ID spécifié
	// ... (votre code de suppression ici)

	return nil
}

func createDiscussion(db *sql.DB, user User, discussion discussions) error {
	if user.Role != "admin" && user.Role != "connected" {
		return fmt.Errorf("seuls les utilisateurs avec les rôles 'admin' ou 'connected' peuvent créer une discussion")
	}

	// Vérifier si l'utilisateur a les autorisations nécessaires pour créer une discussion supplémentaire
	// Vous pouvez ajouter des conditions supplémentaires selon vos besoins

	// Insérer la nouvelle discussion dans la base de données
	insertDiscussion := "INSERT INTO discussions (image, titre, description, nmbreDeLikes, idUser) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(insertDiscussion, discussion.image, discussion.titre, discussion.description, discussion.nmbreDeLikes, discussion.idUser)
	if err != nil {
		return err
	}

	return nil
}

func createMessage(db *sql.DB, user User, message messages) error {
	if user.Role != "admin" && user.Role != "connected" {
		return fmt.Errorf("seuls les utilisateurs avec les rôles 'admin' ou 'connected' peuvent créer un message")
	}

	// Vérifier si l'utilisateur a les autorisations nécessaires pour créer un message supplémentaire
	// Vous pouvez ajouter des conditions supplémentaires selon vos besoins

	// Insérer le nouveau message dans la base de données
	insertMessage := "INSERT INTO messages (text, IDcreateur, IDdiscution) VALUES (?, ?, ?)"
	_, err := db.Exec(insertMessage, message.text, message.IDcreateur, message.IDdiscution)
	if err != nil {
		return err
	}

	return nil
}