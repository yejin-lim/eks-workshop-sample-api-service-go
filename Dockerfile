FROM nginx:1.20.1

COPY . /usr/share/nginx/html


EXPOSE 80
EXPOSE 443

CMD ["nginx","-g","daemon off;"]
