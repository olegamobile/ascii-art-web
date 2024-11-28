echo "Press Enter to stop the container "
read -p ""
echo "Running command \033[92mdocker stop ascii-app\033[0m"
echo 
docker stop ascii-app

echo 
echo "Running command \033[92mdocker ps -a\033[0m to check status of the container"
echo 
docker ps -a

echo
echo "Press Enter to clean garbage (remove all unused objects) "
read -p ""
echo "Running command \033[92mdocker system prune -a --volumes\033[0m"
echo 

docker system prune -a --volumes

echo 
echo "Running command \033[92mdocker ps -a\033[0m to ensure no containers are left"
echo 
docker ps -a

echo 
echo "Running command \033[92mdocker images\033[0m to ensure no images are left"
echo 
docker images