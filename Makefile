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

# internal/model/mock/mock_user_usecase.go:
# 	mockgen -destination=internal/model/mock/mock_user_usecase.go -package=mock github.com/Himatro2021/API/internal/model UserUsecase
# internal/model/mock/mock_user_repository.go:
# 	mockgen -destination=internal/model/mock/mock_user_repository.go -package=mock github.com/Himatro2021/API/internal/model UserRepository
# internal/model/mock/mock_absent_usecase.go:
# 	mockgen -destination=internal/model/mock/mock_absent_usecase.go -package=mock github.com/Himatro2021/API/internal/model AbsentUsecase
# internal/model/mock/mock_absent_repository.go:
# 	mockgen -destination=internal/model/mock/mock_absent_repository.go -package=mock github.com/Himatro2021/API/internal/model AbsentRepository
# internal/model/mock/mock_auth_usecase.go:
# 	mockgen -destination=internal/model/mock/mock_auth_usecase.go -package=mock github.com/Himatro2021/API/internal/model AuthUsecase
# internal/model/mock/mock_session_repository.go:
# 	mockgen -destination=internal/model/mock/mock_session_repository.go -package=mock github.com/Himatro2021/API/internal/model SessionRepository

# mockgen: internal/model/mock/mock_session_repository.go \
# 	internal/model/mock/mock_user_usecase.go \
# 	internal/model/mock/mock_user_repository.go \
# 	internal/model/mock/mock_absent_usecase.go \
# 	internal/model/mock/mock_absent_repository.go \
# 	internal/model/mock/mock_auth_usecase.go

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