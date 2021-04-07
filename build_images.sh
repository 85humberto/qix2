#! /bin/bash
docker build --build-arg KIND=cliente -t 85humberto/qix2-cliente .
docker build --build-arg KIND=balanceador -t 85humberto/qix2-balanceador .
docker build --build-arg KIND=server -t 85humberto/qix2-server .