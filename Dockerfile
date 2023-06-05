FROM golang:latest

WORKDIR /app

# Copier les fichiers du projet dans le conteneur
COPY . .

# Compiler le projet Go
RUN go build -o app

# Définir la commande d'exécution
CMD ["./app"]
