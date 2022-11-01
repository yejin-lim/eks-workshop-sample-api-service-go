# This is a multi-stage build. First we are going to compile and then
# create a small image for runtime.
FROM nginx:1.20.1

COPY static-html-directory /usr/share/nginx/html

COPY --from=builder /go/src/github.com/eks-workshop-sample-api-service-go/index.html /index
COPY --from=builder /etc/passwd /etc/passwd

EXPOSE 80
CMD ["/index"]
