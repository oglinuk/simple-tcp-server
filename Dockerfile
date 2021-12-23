FROM golang:1.17
ADD . /tcp-server
WORKDIR /tcp-server
RUN go build -o tcp-server
EXPOSE 9001
CMD ["./tcp-server"]
