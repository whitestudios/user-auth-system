# Variáveis
BINARY_NAME=auth-service
MAIN=./cmd/server/main.go

# ----------------------------------------
# Comandos principais
# ----------------------------------------

.PHONY: build
build:
	go build -o bin/$(BINARY_NAME) $(MAIN)

.PHONY: run
run:
	go run $(MAIN)

.PHONY: clean
clean:
	rm -rf bin/

.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: deps
deps:
	go mod tidy

.PHONY: migrate
migrate:
	go run ./database/migrations/main.go

# ----------------------------------------
# Ajuda
# ----------------------------------------

.PHONY: help
help:
	@echo "Comandos disponíveis:"
	@echo "  make build     - Compila o binário"
	@echo "  make run       - Roda a aplicação"
	@echo "  make test      - Roda os testes"
	@echo "  make fmt       - Formata o código"
	@echo "  make lint      - Roda linter (golangci-lint)"
	@echo "  make deps      - Organiza dependências"
	@echo "  make clean     - Remove arquivos de build"
	@echo "  make migrate   - Executa migrações (ex: arquivo main.go de migração)"
