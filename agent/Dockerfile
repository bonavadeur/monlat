FROM alpine:latest

RUN apk update && apk add --no-cache \
    bash \
    curl

RUN wget -O /usr/bin/kubectl https://storage.googleapis.com/kubernetes-release/release/v1.23.0/bin/linux/amd64/kubectl && \
    chmod +x /usr/bin/kubectl

ENV TIME 5

COPY ping.sh .

CMD ["bash", "/ping.sh"]
