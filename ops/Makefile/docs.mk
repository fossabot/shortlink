# DOCS TASKS ===========================================================================================================
.PHONY: docs
docs: ## Generate documentation
	@swag init \
		-g server.go \
		--dir ./internal/services/api/application/http-chi \
		--output docs/api \
		--parseDependency
