FROM nginx:1.20.1 as builder

COPY index.html /usr/share/nginx/html/index.html


EXPOSE 80
EXPOSE 8080
EXPOSE 443

CMD ["nginx","-g","daemon off;"]
