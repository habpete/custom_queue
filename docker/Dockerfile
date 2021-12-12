FROM golang:alpine

WORKDIR /go/src

RUN apk update && apk add git
RUN git clone https://github.com/habpete/custom_queue.git

WORKDIR custom_queue

RUN go build -o /main ./cmd

EXPOSE 8082

CMD ["./main"]