#!/bin/bash

docker build -t "$docker_username/sentence:latest" -t "$docker_username/sentence:1.0-${GITHUB_SHA::4}" ./app
