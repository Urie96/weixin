# docker-compose down
# docker rmi -f wxservice:latest
# docker build -t wxservice:latest .
# docker rmi $(docker images | grep "none" | awk '{print $3}')
docker-compose up -d
docker rmi $(docker images | grep "none" | awk '{print $3}')