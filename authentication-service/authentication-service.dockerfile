FROM alpine:latest

RUN mkdir /app

COPY authApp /app

CMD [ "/app/authApp" ]

#When we run thins, should first of all, should build the code on one docker image, and must create much smaller docker image and copy over executable