#!/usr/bin/env bash

docker stop $(docker ps -qa)
docker rm $(docker ps -qa)

docker rmi $(docker images -q --filter reference=dev-*:*)

docker-compose -f docker-compose-orderer.yaml -f docker-compose-peer.yaml  -f docker-compose-cli.yaml -f docker-compose-ca.yaml down --volumes --remove-orphans