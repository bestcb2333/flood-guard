#!/bin/bash
go build -o floodguard . && \
docker build --pull=false -t Sauce233/floodguard . && \
docker-compose -f /srv/floodguard/docker-compose.yml up -d && \
docker logs -f floodguard
