#!/bin/bash

set -e

# Run this as `sudo build.sh` if it gives you permission denied

docker build . -t groupie-tracker

echo "Docker image built, run me locally with: docker run -p 8080:8080 groupie-tracker"