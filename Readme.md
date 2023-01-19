
## Install go packages

```shell
go mod tidy
```

## Set ngrok

```shell
ngrok http 8080
```

## Line Setting
[Line Developer Console](https://developers.line.biz/console/)
Set Line Message API Webhook URL
### Allow bot to join group chats
- Disabled
### Auto-reply messages
- Disabled
### Greeting messages
- Disabled
### Webhook Url
```shell
https://<<ngrok uri>>/lineMsg/save
```
