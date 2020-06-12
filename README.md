#docker 

## xóa toàn bộ image 
docker rmi -f $(docker images -a -q)

## build image 
docker build -t app .

## docker run 
docker run -p 80:8081 -it app

# study_golang
