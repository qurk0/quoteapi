# Переменные для компиляции приложения
APP_NAME = app
MAIN = cmd/main.go
BUILD_DIR = build

# Список архитектур, под которые makefile сможет собрать проект
OS_LIST = linux windows darwin
ARCH = amd64

.PHONY: all clean help $(OS_LIST)

all: $(OS_LIST)

$(OS_LIST):
	@echo "Собираем под $@..."
	GOOS=$@ GOARCH=$(ARCH) go build -o $(BUILD_DIR)/$(APP_NAME)_$@$(if $(filter $@, windows),.exe,) $(MAIN)

clean:
	rm -rf $(BUILD_DIR)

help:
	@echo "Список доступных команд:"
	@echo "  make           - Собрать проект под все платформы"
	@echo "  make clean     - Удалить папку build с собранными там проектами"
	@echo "  make {Ваша ОС} - Собрать проект под конкретную ОС."
	@echo "  'make linux'   - Пример сборки проекта под Linux"
	@echo "  Доступные ОС:  $(OS_LIST)"