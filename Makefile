SHELL:=/bin/bash

ifdef test_run
	TEST_ARGS := -run $(test_run)
endif

migrate_up=go run main.go migrate --direction=up --step=0

run: check-modd-exists
	@modd -f ./.modd/server.modd.conf

check-modd-exists:
	@modd --version > /dev/null

migrate:
	@if [ "$(DIRECTION)" = "" ] || [ "$(STEP)" = "" ]; then\
    	$(migrate_up);\
	else\
		go run main.go migrate --direction=$(DIRECTION) --step=$(STEP);\
    fi

internal/model/mock/mock_mail_usecase.go:
	mockgen -destination=internal/model/mock/mock_mail_usecase.go -package=mock mail-service/internal/model MailUsecase
internal/model/mock/mock_mail_repository.go:
	mockgen -destination=internal/model/mock/mock_mail_repository.go -package=mock mail-service/internal/model MailRepository

mockgen: internal/model/mock/mock_mail_usecase.go \
	internal/model/mock/mock_mail_repository.go

clean:
	rm -v internal/model/mock/mock_*.go

check-cognitive-complexity:
	find . -type f -name '*.go' -not -name "mock*.go" \
      -exec gocognit -over 15 {} +

lint: check-cognitive-complexity
	golangci-lint run --print-issued-lines=false --exclude-use-default=false --enable=revive --enable=goimports  --enable=unconvert --enable=unparam --concurrency=2

test-only: check-gotest mockgen
	SVC_DISABLE_CACHING=true $(test_command)

test: lint test-only

check-gotest:
ifeq (, $(shell which richgo))
	$(warning "richgo is not installed, falling back to plain go test")
	$(eval TEST_BIN=go test)
else
	$(eval TEST_BIN=richgo test)
endif

ifdef test_run
	$(eval TEST_ARGS := -run $(test_run))
endif
	$(eval test_command=$(TEST_BIN) ./... $(TEST_ARGS) -v --cover)