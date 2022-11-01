FROM golang:1.11.1 as builder

RUN yum update
RUN yum install -y nginx
COPY index.html /var/www/html

EXPOSE 80

CMD ["nginx"]
