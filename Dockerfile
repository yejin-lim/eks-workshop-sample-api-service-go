FROM nginx:1.20.1 as builder

RUN mkdir -p /go/src/github.com/eks-workshop-sample-api-service-go


COPY . /usr/share/nginx/html
COPY --from=builder /go/src/github.com/eks-workshop-sample-api-service-go/index.html /nginx

EXPOSE 80
EXPOSE 8080
EXPOSE 443

CMD ["nginx","-g","daemon off;"]
