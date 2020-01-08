#!/bin/bash

docker push "$docker_username/sentence:1.0-${GITHUB_SHA::4}" 
docker push "$docker_username/sentence:latest" &
wait