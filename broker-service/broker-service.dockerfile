FROM alpine:latest

RUN mkdir /app

COPY brokerApp /app

CMD [ "/app/brokerApp" ]

#When we run thins, should first of all, should build the code on one docker image, and must create much smaller docker image and copy over executable