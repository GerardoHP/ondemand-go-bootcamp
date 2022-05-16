################################### ECHO GO SERVER ###################################
FROM golang:1.18

WORKDIR /usr/src/pokemon-service

# pre-copy/cache go.mod for predownloading all dependencies and only reaload required
COPY go.mod go.sum ./
RUN go mod download && go mod verify 

COPY . .
RUN go build -v -o /usr/local/bin/service ./.
COPY pokemons.csv /usr/local/bin/

EXPOSE 8080

CMD ["service"]
