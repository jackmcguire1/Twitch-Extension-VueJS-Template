# Twitch - Extension - Template - vuejs/ golang


[git]:      https://git-scm.com/
[golang]:   https://golang.org/
[discord]: https://discord.gg/Q4PX2Yj
[vue]: https://cli.vuejs.org/guide/installation.html
[dlv]:    https://github.com/go-delve/delve
[modules]: https://github.com/golang/go/wiki/Modules
[goLand]: https://www.jetbrains.com/go/
[golint]: https://github.com/golangci/golangci-lint
[aws-cli]: https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html
[aws-cli-config]: https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html
[aws-sam-cli]: https://github.com/awslabs/aws-sam-cli
[localstack]: https://github.com/localstack/localstack

> This repo contains a template for both a serverless EBS written in golang and the front-end for a Twitch panel extension in vueJS.

<p align="center">
    <img align="center" src="https://vuejs.org/images/logo.png" width="50%" height="50%" title="Glo Boards">
</p>

## ABOUT

This project serves as a template to create a serverless EBS and creating the front-end of a Twitch panel extension in vueJS.



### Prerequisites
> Make sure you have the following installed:

- [Git][git]
- [Go 1.14][golang]+
- [VueJS][vue]
- [aws-cli][aws-cli]
- [aws-sam-cli][aws-sam-cli]
- [localstack][localstack]- [OPTIONAL]

## Installation
This project is formed of two components, please see their sub-directories
> goto ./client for the vueJS project

> goto ./server for the EBS.

### Twitch Extension Configuration
From your [Twitch Extension Dashboard](https://dev.twitch.tv/dashboard/extensions) you can get the following:
- Client ID
- Base64 Secret
- Extension Version
- Extension (**Broadcaster**/**Developer**) Config Version - *OPTIONAL*

### Owner ID
To get the owner ID, you will need to execute a simple CURL command against the Twitch `/users` endpoint. You'll need your extension <b>client ID</b> as part of the query (this will be made consistent with the Developer Rig shortly, by using _owner name_).

```bash
curl -H "Client-ID: <client id>" -X GET "https://api.twitch.tv/helix/users?login=<owner name>"
```
## Create your own extension!
Get started and create your extension [today!](https://dev.twitch.tv/extensions).

## FAQ & SUPPORT
For any questions or suggestions please join the **'go-twitch-ext'** channel on [Discord][discord]!
> - [https://dev.twitch.tv/docs/extensions]()
> - [https://discord.gg/566fFzA]() - Twitch dev community discord
> - [https://discord.gg/qe7b8da]() - BootstrapVue discord
> - [https://discord.gg/Q4PX2Yj]() - Twitch API discord


## License
The source code for go-twitch-ext is released under the [MIT License][MIT].

## Donations
All donations are appreciated!

[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](http://paypal.me/crazyjack12)
