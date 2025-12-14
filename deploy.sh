docker build -t botbot-discord:latest .
docker service rm botbot_discord
docker stack up -c ./service.yaml botbot
docker service ls