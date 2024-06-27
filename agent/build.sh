#!/bin/bash

#####CONFIG HERE#####
HUB="chung123abc" # docker.io/$HUB/$NAME:$TAG
TAG="v1"
NAME="monlat-agent"
#####################

GREEN="\e[32m"
OPTION=$1
logSuccess() { echo -e "$GREEN-----$message-----";}

buildImage() {
    image=$(docker images | grep $NAME | awk '{print $1}'):$TAG
    docker rmi $image
    message="Remove old image success" && logSuccess
    docker build -t $HUB/$NAME:$TAG .
    message="Build Success" && logSuccess
}

pushImage() {
    image=$(docker images | grep $NAME | awk '{print $1}'):$TAG
    docker push $image
    message="Push Success" && logSuccess
}

if [ $OPTION == "image" ]; then
    buildImage
elif [ $OPTION == "push" ]; then
    pushImage
elif [ $OPTION == "ful" ]; then
    buildImage
    pushImage
fi