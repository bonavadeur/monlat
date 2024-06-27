# Latency between Kubernetes Nodes Monitoring

## 1. Introduction

This is a monitoring system give us the ability of monitoring latency between Kubernetes Nodes. Metrics then is exported to Prometheus with metrics `latency_between_nodes`

![alt](images/model.jpg)


## 2. Installation

### 2.1. Build Image Monlat-Agent
Rebuild image of `monlat-agent` component if you need. Let's take a quick view in [agent/build.sh](agent/build.sh). Then, rebuild image if you need.

```bash
cd agent
vi build.sh
...
#####CONFIG HERE#####
HUB="chung123abc" # docker.io/$HUB/$NAME:$TAG
TAG="v1"
NAME="monlat-agent"
#####################
...

chmod +x build.sh
./build.sh image # build from Golang code to Docker image
./build.sh push # push Docker image to Docker Hub
cd ..
```

### 2.2. Build Image Monlat

Rebuild image of `monlat` component if you need. Let's take a quick view in [src/build.sh](src/build.sh). Then, rebuild image if you need.

```bash
cd src
vi build.sh
...
#####CONFIG HERE#####
HUB="bonavadeur" # docker.io/$HUB/$NAME:$TAG
TAG="latest"
NAME="monlat"
#####################
...

chmod +x build.sh
./build.sh image # build from Golang code to Docker image
./build.sh push # push Docker image to Docker Hub
cd ..
```

### 2.3. Some other small changes

Specify image you use to run `monlat` and specified Nodes you want to deploy it in `Deployment` in file [manifest/monlat.yaml](manifest/monlat.yaml)

### 2.4. Install!

```bash
kubectl apply -f manifest/rbac.yaml
kubectl apply -f manifest/monlat-agents.yaml
kubectl apply -f manifest/monlat.yaml
```

Assume that Your Kubernetes run Prometheus before. Metrics are updated to prometheus immediately with metrics `latency_between_nodes`.

## 3. Contributeurs

My colaborators who write this system is [chungtd203338](https://github.com/chungtd203338). His original code is [monitor-delay-nodes-k8s](https://github.com/chungtd203338/monitor-delay-nodes-k8s)
