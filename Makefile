
BUILD_OUTPUT_FILENAME ?= butta

# platform-specific settings
GOOS := $(shell go env GOOS)

ifeq ($(GOOS),windows)
	ifneq ($(suffix $(BUILD_OUTPUT_FILENAME)),.exe)
		BUILD_OUTPUT_FILENAME := $(addsuffix .exe,$(BUILD_OUTPUT_FILENAME))
	endif
endif

.PHONY: build
build:
	go build -o dist/api/$(BUILD_OUTPUT_FILENAME) $(BUILD_FLAGS) ./cmd/api
	go build -o dist/console/$(BUILD_OUTPUT_FILENAME) $(BUILD_FLAGS) ./cmd/console

.PHONY: dev
dev:
	test -f .env && . .env; air -c .air.toml .

.PHONY: db\:dump
db\:dump:
	dbmate $([[ -f .env ]] && echo '--env-file .env') dump

.PHONY: db\:migrate
db\:migrate:
	dbmate $([[ -f .env ]] && echo '--env-file .env') up

.PHONY: db\:rollback
db\:rollback:
	dbmate $([[ -f .env ]] && echo '--env-file .env') rollback

.PHONY: db\:cm
db\:cm:
	 @dbmate $([[ -f .env ]] && echo '--env-file .env') new $(filter-out $@,$(MAKECMDGOALS))


.PHONY: db\:seed
db\:seed:
	dbmate $([[ -f .env ]] && echo '--env-file .env') --migrations-table 'seed_migrations'	--migrations-dir './database/seeders' --no-dump-schema up


.PHONY: install
install:
	@# install Air verse
	go install 'github.com/air-verse/air@latest'
	@# install db-mate to manage migrations
	go install 'github.com/amacneil/dbmate/v2@v2.25'


%: #catch all command
	@: # do nothing silently