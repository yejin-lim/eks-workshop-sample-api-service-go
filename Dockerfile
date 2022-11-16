FROM ubuntu:16.04

RUN apt-get update
RUN apt-get install -y nginx
RUN echo "\ndaemon off;" >> /etc/nginx/nginx.conf
RUN chown -R www-data:www-data /var/lib/nginx
COPY index.html /usr/share/nginx/html
COPY login.html /usr/share/nginx/html
VOLUME ["/data", "/etc/nginx/site-enabled", "/var/log/nginx"]

WORKDIR /etc/nginx
CMD ["nginx"]

EXPOSE 80
EXPOSE 13000

# This is a multi-stage build. First we are going to compile and then
# create a small image for runtime.

FROM golang:1.13 as builder

RUN mkdir -p /go/src/github.com/eks-workshop-sample-api-service-go
WORKDIR /go/src/github.com/eks-workshop-sample-api-service-go
RUN useradd -u 10001 app

#ENV GIN_MODE=release
#WORKDIR /go/src/crawler-docker
#COPY main.go .
#RUN apk update && apk add --no-cache git
#RUN go get github.com/ryulitaro/crawler
#RUN go get github.com/gin-gonic/gin

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch

COPY --from=builder /go/src/github.com/eks-workshop-sample-api-service-go/main /main
COPY --from=builder /etc/passwd /etc/passwd
USER app

EXPOSE 8080
CMD ["/main"]
