include .env
export

# Все команды
# Команда help
help:
	@echo "Доступные команды:"
	@echo "  swag          - Запускает генерацию swagger файла"
	@echo "  lint          - Устанавливает golangci-lint и запускает линтер"
	@echo "  lint-fast          - Устанавливает golangci-lint и запускает быстрый линтер"
	@echo "  mock          - Запускает генерацию мок make mock SRC_FILE='' DEST_FILE='' PACKAGE=''"
	@echo "  run_tests          - Запускает unit и интеграционные тесты и генерирует покрытие"
	@echo "  unit_test         - Запускает unit test"
	@echo "  integration_test   - Запускает интеграционные тесты"
	@echo "  coverage           - Объединяет покрытия и генерирует отчеты"
	@echo "  clean              - Удаляет временные файлы покрытия"
	@echo "  help               - Выводит это сообщение"

.PHONY: run_tests unit_test integration_test coverage clean help swag mock lint install-linter lint-fast

PROJECT_DIR = $(shell pwd)
PROJECT_BIN = $(PROJECT_DIR)/bin
$(shell [ -f bin ] || mkdir -p $(PROJECT_BIN))
PATH := $(PROJECT_BIN):$(PATH)

# linter
GOLANG_CI_LINT = $(PROJECT_BIN)/golangci-lint
GOLANG_CI_LINT_VERSION = v1.64.8

# Директории для тестов
UNIT_TEST_DIR := ./internal/...
INTEGRATION_TEST_DIR := ./internal/integration-tests/...

# Файлы покрытия
COVER_UNIT := cover_unit.out
COVER_INTEGRATION := cover_integration.out
COVER_MERGED := cover.out
COVER_HTML := cover.html

# Пути к пакетам для -cover pkg
# Включает все пакеты
COVER_PKG_INTEGRATION := ./internal/...

# Сформировать swagger файл
swag:
	swag init -g cmd/app/main.go

# Основная команда для запуска всех тестов и получения покрытия
run_tests: coverage

# Объединение покрытий и генерация отчетов
coverage: unit_test integration_test
	@echo "Merging coverage profiles..."
	@echo "mode: set" > $(COVER_MERGED)
	@cat $(COVER_UNIT) | grep -v "mode:" >> $(COVER_MERGED)
	@cat $(COVER_INTEGRATION) | grep -v "mode:" >> $(COVER_MERGED)
	@echo "Merged coverage saved to $(COVER_MERGED)"

	@echo "Generating coverage report..."
	go tool cover -func=$(COVER_MERGED)
	go tool cover -html=$(COVER_MERGED) -o $(COVER_HTML)
	@echo "HTML coverage report saved to $(COVER_HTML)"


# Запуск unit тестов
unit_test:
	@echo "Running unit tests..."
	go clean -testcache
	go test --short -v -cover -race $(UNIT_TEST_DIR) -coverprofile=$(COVER_UNIT).tmp
	@cat $(COVER_UNIT).tmp | grep -v "_mock.go" > $(COVER_UNIT)
	@rm $(COVER_UNIT).tmp
	@echo "Unit tests coverage saved to $(COVER_UNIT)"

# Запуск интеграционных тестов
integration_test:
	@echo "Running integration tests..."
	go clean -testcache
	go test -v -cover -race $(INTEGRATION_TEST_DIR) -coverprofile=$(COVER_INTEGRATION).tmp -coverpkg=$(COVER_PKG_INTEGRATION)
	@cat $(COVER_INTEGRATION).tmp > $(COVER_INTEGRATION)
	@rm $(COVER_INTEGRATION).tmp
	@echo "Integration tests coverage saved to $(COVER_INTEGRATION)"

# Очистка временных файлов
clean:
	@rm -f $(COVER_UNIT) $(COVER_INTEGRATION) $(COVER_MERGED) $(COVER_HTML)
	@echo "Cleaned up coverage files."

# Сформировать mock файл
# SRC_FILE -  файл генерации dao.go
# DEST_FILE - путь, куда и в какой файл сгенерировать mocks/dao_mock.go
# PACKAGE - название пакета файла mock_dao
# make mock SRC_FILE=internal/repository/repo.go DEST_FILE=mocks/repo_mock.go PACKAGE=mocks
mock:
	@if [ -z "$(SRC_FILE)" ]; then echo "SRC_FILE is not set"; exit 1; fi
	@if [ -z "$(DEST_FILE)" ]; then echo "DEST_FILE is not set"; exit 1; fi
	@if [ -z "$(PACKAGE)" ]; then echo "PACKAGE is not set"; exit 1; fi
	mockgen -source=$(SRC_FILE) -destination=$(DEST_FILE) -package=$(PACKAGE)

# Запуск линтера
lint: .install-linter
	@echo "Running golangci-lint..."
	$(GOLANG_CI_LINT) run --fix ./... --config=./.golangci.yml
	@echo "Success"

lint-fast: .install-linter
	@echo "Running golangci-lint --fast ..."
	$(GOLANG_CI_LINT) run ./... --fast --config=./.golangci.yml
	@echo "Success"
.install-linter:
	@echo "Install golangci-lint..."
	[ -f $(PROJECT_BIN)/golangci-lint ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(PROJECT_BIN) $(GOLANG_CI_LINT_VERSION)
