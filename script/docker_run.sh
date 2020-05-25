docker stop wxservic
docker rmi -f wxservice:latest
docker build -t wxservice:latest -f ../Dockerfile ..
docker rmi $(docker images | grep "none" | awk '{print $3}')
docker run --name wxservise -p 7001:7001 -d wxservice 