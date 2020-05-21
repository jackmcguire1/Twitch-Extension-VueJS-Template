# Twitch - Extension - Template - vuejs / golang


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
[twitch-developer-rig]: https://dev.twitch.tv/docs/extensions/rig/
> This repo contains a template for both a serverless EBS written in golang and the front-end for a Twitch panel extension in vueJS.

<p align="center">
    <img align="center" src="https://vuejs.org/images/logo.png" width="25%" height="25%" title="VueJS logo">
</p>

## ABOUT

This project serves as a template to create a serverless EBS and creating the front-end of a Twitch panel extension in vueJS.<br>
This extension will display a list of a broadcaster's followers in a formatted table.

### Prerequisites
> Make sure you have the following installed:

- [Git][git]
- [Go 1.14][golang]+
- [vue-cli][vue]
- [aws-cli][aws-cli]
- [aws-sam-cli][aws-sam-cli]
- [localstack][localstack] - [OPTIONAL]
- [twitch-developer-rig] - OPTIONAL
## Installation
This project is formed of two components, please see their sub-directories
> goto ./Client for the vueJS project

> goto ./EBS for the EBS.

### Twitch Extension Configuration
From your [Twitch Extension Dashboard](https://dev.twitch.tv/dashboard/extensions) you can get the following:
- Client ID
- Base64 Secret
- Extension Version
- Extension (**Broadcaster**/**Developer**) Config Version - *OPTIONAL*

### Owner ID
To get the owner ID, you will need to first create a temporary access token via [https://twitchtokengenerator.com/](https://twitchtokengenerator.com/)<br>
Next fetch your ownerID from a simple CURL command against the Twitch `/users` endpoint.<br> You'll also need the <b>client ID</b> from [https://twitchtokengenerator.com/](https://twitchtokengenerator.com/) as part of the query.

```bash
curl -H "Client-ID: <client id>" -H "Authorization: Bearer <access token>" -X GET "https://api.twitch.tv/helix/users?login=<twitch-login-name>"
```

your ownerID is the value of **'ID'**

```bash
{
	"data": [{
		"id": "35851594",
		"login": "crazyjack12",
		"display_name": "crazyjack12",
		"type": "",
		"broadcaster_type": "",
		"description": "Do What Thou Wilt",
		"profile_image_url": "https://static-cdn.jtvnw.net/jtv_user_pictures/ea0bf955-255a-4eca-ad80-460b88162910-profile_image-300x300.png",
		"offline_image_url": "",
		"view_count": 1702
	}]
}
```

## Create your own extension!
Get started and create your extension [today!](https://dev.twitch.tv/extensions).

## FAQ & SUPPORT
For any questions or suggestions please join the **'go-twitch-ext'** channel on [Discord][discord]!
> - [https://dev.twitch.tv/docs/extensions](https://dev.twitch.tv/docs/extensions)
> - [https://discord.gg/566fFzA](https://discord.gg/566fFzA) - Twitch dev community discord
> - [https://discord.gg/qe7b8da](https://discord.gg/qe7b8da) - BootstrapVue discord
> - [https://discord.gg/Q4PX2Yj](https://discord.gg/Q4PX2Yj) - Twitch API discord

## Donations
All donations are appreciated!

[![Donate](https://img.shields.io/badge/Donate-PayPal-green.svg)](http://paypal.me/crazyjack12)
