#!/bin/sh

SERVER_ID="$1"
DISCORD_BOT_TOKEN="$2"
MEMBER_IDS="$3"
CHANNEL_ID="$4"

docker run -d \
  -e SERVER_ID="$SERVER_ID" \
  -e DISCORD_BOT_TOKEN="$DISCORD_BOT_TOKEN" \
  -e MEMBER_IDS="$MEMBER_IDS" \
  -e CHANNEL_ID="$CHANNEL_ID" \
  --name botbot \
  botbot