docker stop wxservice
docker rm wxservice
docker rmi -f wxservice:latest
docker build -t wxservice:latest .
docker rmi $(docker images | grep "none" | awk '{print $3}')
docker run --name wxservice -p 7001:7001 -d wxservice 