FROM golang:1.11.1 as builder
#FROM $AWS_ACCOUNT_ID.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/$REPOSITORY_NAME:$IMAGE_TAG
#FROM Centos:7

COPY index.html /var/www/html

EXPOSE 80

CMD ["nginx"]
