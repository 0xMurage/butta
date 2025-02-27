
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
	[ -f .env ] && . .env; go tool air -c .air.toml .

db\:clean-dump:
	 [ -f .env ] && . .env; ./scripts/clean-schema.sh

.PHONY: db\:dump
db\:dump:
	go tool dbmate $([[ -f .env ]] && echo '--env-file .env') dump

.PHONY: db\:migrate
db\:migrate:
	go tool dbmate $([[ -f .env ]] && echo '--env-file .env') up
	$(MAKE) db\:clean-dump

.PHONY: db\:rollback
db\:rollback:
	go tool dbmate $([[ -f .env ]] && echo '--env-file .env') rollback

.PHONY: db\:cm
db\:cm:
	 @go tool dbmate $([[ -f .env ]] && echo '--env-file .env') new $(filter-out $@,$(MAKECMDGOALS))

.PHONY: db\:seed
db\:seed:
	go tool dbmate $([[ -f .env ]] && echo '--env-file .env') --migrations-table 'seed_migrations'	--migrations-dir './database/seeders' --no-dump-schema up

.PHONY: river\:db-dump
river\:db-dump:
	 [ -f .env ] && . .env;  && ./scripts/river-db-migrations.sh

.PHONY: sqlc
sqlc:
	go tool sqlc generate

.PHONY: queue\:work
queue\:work:
	[ -f .env ] && . .env; go run cmd/console/main.go


%: #catch all command
	@: # do nothing silently