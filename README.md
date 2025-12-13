# botbot

A Discord bot to check in for daily work session.

<div align="center">

<img src="./1.png" width="100%" />

</div>

Study, coding, gym... you name it.

## run

create discord app and issue bot token.

and then fill `.env.local`.

```sh
SERVER_ID=""
DISCORD_BOT_TOKEN=""
MEMBER_IDS=""
CHANNEL_ID=""
```

in local development,

```sh
go run main.go
```

to deploy the bot, build the image.

```sh
docker build -t botbot:latest .
```

and then in terminal,

```sh
docker
./docker.run.sh SERVER_ID DISCORD_BOT_TOKEN MEMBER_IDS CHANNEL_ID
```

or just build executable with go build.

## tmi

+80% of codes written by chatgpt.

## license

MIT.
