# GO BARBER

Aplicação desenvolvida em golang usando o exemplo original em NodeJS da rocketseat.

# EXECUÇãO

## Install go 
https://golang.org/doc/install

## Install database (postgres)
``docker-compose up -d``

## Run api
``go run server.go``

## Install air (optional)
executa o projeto e atualiza automaticamente os arquivos alterados

``go get -u github.com/cosmtrek/air``

``air``

## Run test
``go test ./... -v``

## Run coverage tests
### export coverage
``go test ./... -coverprofile=coverage.out``

### analyze coverage
``go tool cover -func=coverage.out``

### analyze coverage via a browser
``go tool cover -html=coverage.out``

