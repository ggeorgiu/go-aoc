.PHONY: new
new:
	./scripts/generate.sh "$(day)" "$(year)"

FORMAT_FILES := grep -l -R --exclude-dir=.git --include="*.go" .
.PHONY: format
format:
	go install mvdan.cc/gofumpt golang.org/x/tools/cmd/goimports
	$(FORMAT_FILES) | xargs gofumpt -w
	$(FORMAT_FILES) | xargs goimports -w -local github.com/ggeorgiu/go-aoc
