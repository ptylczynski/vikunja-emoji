#!/bin/bash

go build -o vikunja-emoji
docker build -t ghcr.io/ptylczynski/vikunja-emoji:latest .
docker push ghcr.io/ptylczynski/vikunja-emoji:latest