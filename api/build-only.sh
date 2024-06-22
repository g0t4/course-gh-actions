#!/bin/bash

my_app_version=$(git describe --tag)

docker compose build --pull --build-arg "MY_APP_VERSION=$my_app_version"
