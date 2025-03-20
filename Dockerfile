FROM golang:1.23-alpine

COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o /bin/demo
EXPOSE 8888
CMD [ "/bin/demo" ]