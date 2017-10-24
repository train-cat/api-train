run:
	aah run

code-quality:
	@echo "== GOLINT =="
	@find . -type d | xargs -L 1 golint
	@echo "== GO VET =="
	@find . -name "*.go" -exec go vet {} \;

