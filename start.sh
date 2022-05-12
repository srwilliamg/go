#!/bin/sh
docker build -t golang-image . && docker run --name golang-container golang-image && docker rm -f golang-container