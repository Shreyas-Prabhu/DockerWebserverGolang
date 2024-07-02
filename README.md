Steps to run 

1)docker build -t goapp .
2)docker run -it -d -p 4000:4000 goapp
3)acces http://localhost:4000/hello
4)docker ps
5)docker images
6)docker stop container_id
7)docker rm container_id
8)docker rmi image_id
