
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

db\:clean-dump:
	 [[ -f .env ]] &&  source .env && ./scripts/clean-schema.sh

.PHONY: db\:dump
db\:dump:
	dbmate $([[ -f .env ]] && echo '--env-file .env') dump

.PHONY: db\:migrate
db\:migrate:
	dbmate $([[ -f .env ]] && echo '--env-file .env') up
	$(MAKE) db\:clean-dump

.PHONY: db\:rollback
db\:rollback:
	dbmate $([[ -f .env ]] && echo '--env-file .env') rollback

.PHONY: db\:cm
db\:cm:
	 @dbmate $([[ -f .env ]] && echo '--env-file .env') new $(filter-out $@,$(MAKECMDGOALS))

.PHONY: db\:seed
db\:seed:
	dbmate $([[ -f .env ]] && echo '--env-file .env') --migrations-table 'seed_migrations'	--migrations-dir './database/seeders' --no-dump-schema up

.PHONY: river\:db-dump
river\:db-dump:
	 [[ -f .env ]] &&  source .env && ./scripts/river-db-migrations.sh

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: queue\:work
queue\:work:
	[ -f .env ] && . .env; go run cmd/console/main.go

.PHONY: install
install:
	@# install Air verse
	go install 'github.com/air-verse/air@latest'
	@# install db-mate to manage migrations
	go install 'github.com/amacneil/dbmate/v2@v2.25'
	@# install river cmd client for migrations export
	go install 'github.com/riverqueue/river/cmd/river@v0.16'
	@# sqlc to generate type safe code from sql
	go install 'github.com/sqlc-dev/sqlc/cmd/sqlc@latest'



%: #catch all command
	@: # do nothing silently