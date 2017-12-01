run:
	AAH_HOST=127.0.0.1 AAH_PORT=8080 AAH_ENV=dev aah run -e dev

code-quality:
	@echo "== GOLINT =="
	@find . -type d | xargs -L 1 golint
	@echo "== GO VET =="
	@find . -name "*.go" -exec go vet {} \;
