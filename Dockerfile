FROM golang:1.17


ENV GO111MODULE on


WORKDIR /home/src/github.com/ismaelpereira/rabbitmq-listener/


COPY go.mod .


COPY go.sum .


RUN go mod download 


RUN go mod verify


COPY . .


RUN go build -o receiver


RUN ls


ENTRYPOINT ["./receiver"]