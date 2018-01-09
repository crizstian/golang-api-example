#!/usr/bin/env bash

# removes any cairo container if exists
docker rm -f cairo

# removes cairo image if exists
docker rmi cairo

# clean the docker env from idle images
docker image prune -y

# clean the docker env from idle volumes
docker volume prune -y

# it builds the dockerfile image
docker build -t cairo .
