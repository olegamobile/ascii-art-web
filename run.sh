echo "Press Enter to run command \033[92m docker image build -f Dockerfile -t ascii .\033[0m and build an image"
read -p ""

docker image build -f Dockerfile -t ascii .

echo
read -p "Image is built. Press Enter to see list of images."
echo 

echo "Running command \033[92mdocker images\033[0m"
echo 
docker images

echo
read -p "Press Enter to create and run container"
echo 

echo "Running command \033[92mcontainer run -p 8080:8080 --detach --name ascii-app ascii\033[0m"
echo 
docker container run -p 8080:8080 --detach --name ascii-app ascii

echo
read -p "Container is up and running. Press Enter to see list of containers."
echo 

echo "Running command \033[92mdocker ps -a\033[0m"
echo 
docker ps -a
  

echo
read -p "Press Enter to see metadata"
echo 

echo "Running command \033[92mdocker inspect ascii-app  | jq -r '.[0].Config.Labels'\033[0m"
echo 
docker inspect ascii-app  | jq -r '.[0].Config.Labels'

echo
echo "Press Enter to run command \033[92m docker exec -it ascii-app /bin/sh\033[0m to run terminal inside the container.\nYou can use \033[92mls -l\033[0m command to see files iside the image. When you finish please type \033[92mexit\033[0m command"
read -p ""

docker exec -it ascii-app /bin/sh
