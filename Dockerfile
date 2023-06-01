# Utilisez l'image de base de Go
FROM golang:latest

# Définissez le répertoire de travail
WORKDIR /app

# Copiez les fichiers go.mod et go.sum
COPY go.mod go.sum ./

# Exécutez go mod download pour télécharger les dépendances
RUN go mod download

# Copiez le reste des fichiers du projet
COPY . .

# Compilez le serveur Go
RUN go build -o main .

# Exposez le port sur lequel votre serveur Go écoute
EXPOSE 8080

# Démarrez le serveur Go
CMD ["./main"]
