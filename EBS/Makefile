clean:
	@rm -rf ./.aws-sam/

build:
	sam build

run-api: build
	sam local start-api \
	--skip-pull-image \
	-d 5859 \
	--debugger-path . \
	--debug-args="-delveAPI=2" \
	--log-file ./output.log \
	--docker-network localstack

invoke: build
	sam local invoke \
	--skip-pull-image \
	--debug-port 5859 \
	--debugger-path . \
	--debug-args="-delveAPI=2" \
	--event ./payloads/usersettingsupdates.json \
	--log-file ./output.log \
	--docker-network gmg_default \
	Lambda

validate:
	sam validate

delve:
	GOARCH=amd64 GOOS=linux go build -o ./dlv github.com/go-delve/delve/cmd/dlv