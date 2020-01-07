#! /bin/bash
docker_username=$docker_username docker-compose -p ci up --build --exit-code-from sut