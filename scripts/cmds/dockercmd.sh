sudo docker run -it ubuntu bash
sudo docker run hello-world

sudo docker ps -l
sudo docker rm d85b8bf3dc33
sudo docker rm -v $(docker ps -aq -f status=exited)
sudo docker rm -i d85b8bf3dc33

docker container prune

sudo docker build -t pifagor-app .
sudo docker run -p 8080:8081 -it pifagor-app