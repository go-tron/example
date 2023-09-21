#!/usr/bin/env bash
export GOOS=linux
export GOARCH=amd64
go build  -o ./build/dory-mall

commitId=$(git rev-parse --short HEAD)
version=$(git describe --tags --abbrev=0)
export commitId
export version
export tag=feverzy/repo:dory-mall-$version-$commitId

docker build -t "$tag" .
docker push "$tag"

echo finished


