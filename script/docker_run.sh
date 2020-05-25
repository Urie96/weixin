docker stop wxservice:latest
docker image rm -f wxservice:latest
docker build -t wxservice:latest -f ../Dockerfile ..
docker rmi $(docker images | grep "none" | awk '{print $3}')
docker run -p 7001:7001 -d wxservice