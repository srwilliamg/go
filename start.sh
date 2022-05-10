#!/bin/sh
docker build -t golang-container . && docker run golang-container