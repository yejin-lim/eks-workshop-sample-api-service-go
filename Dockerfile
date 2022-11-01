FROM golang:1.11.1 as builder


COPY index.html /var/www/html

EXPOSE 80

CMD ["nginx"]
