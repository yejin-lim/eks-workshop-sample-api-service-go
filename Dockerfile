FROM centos:7 

RUN yum update
RUN yum install -y nginx
COPY index.html /var/www/html

EXPOSE 80

CMD ["nginx"]
