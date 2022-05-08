.PHONY: generate-mocks

.DEFAULT_GOAL:=help

help: ## Prints the help about targets.
	@printf "Usage:             make [\033[34mtarget\033[0m]\n"
	@printf "Default:           \033[34m%s\033[0m\n" $(.DEFAULT_GOAL)
	@printf "Targets:\n"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf " \033[34m%-17s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

generate-mocks: ## Auto-generate with mockgen tool (1.6.0)
	mockgen -source=./internal/nodes/attacker_factory.go -destination ./internal/mock/nodes/attacker_factory.go -package mock_nodes