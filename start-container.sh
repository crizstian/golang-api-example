# eval `docker-machine env dimtec`

# creates the cairo image
sh ./create-image.sh

# starts the cairo app in a container
docker run --name cairo \
-p 8500:8000 \
--env-file env \
-v "$(pwd)/src":/go/src/easycast/src \
-d cairo && \
docker ps && \
docker logs -f cairo
