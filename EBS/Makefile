clean:
	@rm -rf ./.aws-sam/

build:
	sam build

run-api: build
	sam local start-api \
	--skip-pull-image \
	--log-file ./output.log

run-api-debug: build
	sam local start-api \
	--skip-pull-image \
	--debugger-path . \
	-d 5859 \
	--debug-args "-delveAPI=2" \
	--log-file ./output.log

validate:
	sam validate

delve:
	GOARCH=amd64 GOOS=linux go build -o ./dlv github.com/go-delve/delve/cmd/dlv