docker build -t hub-botbot:latest .
docker service rm hub_botbot
docker stack up -c ./service.yaml hub
echo [LOG] waiting 5 seconds to deploy
sleep 5
echo [LOG] done, listing up services
docker service ls