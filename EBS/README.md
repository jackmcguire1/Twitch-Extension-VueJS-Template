# EBS
> This directory contains the EBS for the Twitch Extension

[git]:    https://git-scm.com/
[golang]: https://golang.org/
[dlv]:    https://github.com/go-delve/delve
[modules]: https://github.com/golang/go/wiki/Modules
[goLand]: https://www.jetbrains.com/go/
[golint]: https://github.com/golangci/golangci-lint
[aws-cli]: https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html
[aws-cli-config]: https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html
[aws-sam-cli]: https://github.com/awslabs/aws-sam-cli

## About
This Directory contains a serverless EBS for the Twitch extension <br>
This project is to be run via AWS SAM, please see below for<br>
instructions about running locally


## SETUP
Please edit the template.yml file and supply your own
values where 'XXX' is found.

## Running locally with debugging
This sections shows how to run the API endpoints locally

> This project uses AWS SAM to run the API endpoints locally with debugging optionally enabled

> - By default debugging is turned on and will run on port :**5859**

> - check output.log for server logs

```shell
  make run-api
 ```
 
 **OR**
 
 ```shell
  make run-api-debug
 ```

### API Endpoints

**API Endpoints supported**:
- Followers

## Development

To develop `EBS` or interact with its source code in any meaningful way, be
sure you have the following installed:

### Prerequisites

- [Git][git]
- [Go 1.12][golang]+
- [golangCI-Lint][golint]
- [Delve Debugger][dlv]
- [AWS CLI][aws-cli]
- [AWS SAM CLI][aws-sam-cli]

>You will need to activate [Modules][modules] for your version of [GO][golang], 

> by setting the `GO111MODULE=on` environment variable set, or by enabling go-mod in [GoLand][goLand]

### [golangCI-Lint][golint]
```shell
curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $GOPATH/bin latest
```

### [Delve Debugger][dlv]
```shell
make delve
```
**OR**
```shell
GOARCH=amd64 GOOS=linux go build -o ./dlv github.com/go-delve/delve/cmd/dlv
```

### [AWS CLI Configuration][aws-cli-config]
> Make sure you configure the AWS CLI
- AWS Access Key ID
- AWS Secret Access Key
- Default region 'us-east-1'
```shell
aws configure
```


## Build
> How to build the individual lambdas

> **Note***:- Executables are placed in 'dist/handlers/*/main' 

```shell
make build
```

**OR**

- **Lint the project FIRST**
> **Note***:-This may throw build **errors**, which **may** not break anything!

```Shell 
golangci-lint run
```
## Package & Deploy
> This section describe how to validate, compile and deploy the latest back-end stack.

### Validate
> Validate the template.yml before packaging

> NOTE*:- requires AWS IAM permissions

```shell
make validate
```

**OR**

```shell
sam validate
```

### Package
> How to compile template.yaml after updating the back-end stack definition.

```shell 
sam package --template-file template.yml --s3-bucket {your-s3-bucket-name} --output-template-file packaged.yaml
```

**OR**

```shell 
aws cloudformation package --template-file template.yml --output-template-file packaged.yaml --s3-bucket {your-s3-bucket-name}
```

### Deploy
> How to deploy the complied packged.yaml to update the back-end stack.

```shell
sam deploy --template-file ./packaged.yaml --stack-name {your-stackname-here} --capabilities CAPABILITY_IAM
```

**OR**

```shell 
aws cloudformation deploy --template-file ./packaged.yaml --stack-name {your-stackname-here} --capabilities CAPABILITY_IAM
```

## Contributors

This project exists thanks to **all** the people who contribute.

## Donations
All donations are appreciated!

[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](http://paypal.me/crazyjack12)