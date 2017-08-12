all: deps test

ifdef CHANGE
BENCHMARK_LOG=benchmark_change.log
else
BENCHMARK_LOG=benchmark_master.log
endif

DEPS = \
	github.com/golang/dep/cmd/dep \
	github.com/kisielk/errcheck \
	github.com/uber/go-torch \
	golang.org/x/perf/cmd/benchstat

tool-deps: ## Install tools deps
	go get $(DEPS)

tool-deps-up: ## Update tools deps
	go get -u $(DEPS)

setup: tool-deps

deps: setup ## Install deps to vendor
	dep ensure
	dep prune

deps-up: ## Update deps
	dep ensure -update

test: ## Test
	go test -race -cpu 1,2,4

test-bench: ## Test
	go test -race -cpu 1,2,4 -count 5 -benchmem -bench . | tee -a $(BENCHMARK_LOG)

stats: ## Benchmark statistics
	benchstat benchmark_master.log benchmark_change.log

clean: ## Clean up
	@rm *.log
	@rm -rf vendor

help:
	@printf "\033[36mHelp: \033[0m\n"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36mmake %-20s\033[0m%s\n", $$1, $$2}'

.PHONY: all setup deps dep depup test test-bench help