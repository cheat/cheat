# Start docker daemon
docker -d

# start a container with an interactive shell
docker run -ti <image_name> /bin/bash

# "shell" into a running container (docker-1.3+)
docker exec -ti <container_name> bash

# inspect a running container
docker inspect <container_name> (or <container_id>)

# Get the process ID for a container
# Source: https://github.com/jpetazzo/nsenter
docker inspect --format {{.State.Pid}} <container_name_or_ID>

# List the current mounted volumes for a container (and pretty print)
# Source:
# http://nathanleclaire.com/blog/2014/07/12/10-docker-tips-and-tricks-that-will-make-you-sing-a-whale-song-of-joy/
docker inspect --format='{{json .Volumes}}' <container_id> | python -mjson.tool

# list currently running containers
docker ps

# list all containers
docker ps -a

# list all images
docker images