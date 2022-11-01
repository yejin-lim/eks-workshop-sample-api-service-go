FROM nginx:1.20.1

COPY . /usr/share/nginx/html


EXPOSE 80

CMD ["nginx","daemon off;"]
