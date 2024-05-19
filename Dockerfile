# define go ver
FROM golang:1.22.0 AS builder
LABEL stage=builder

# set directory inside the container
WORKDIR /kreditplus-xyz/

# copy mod sum file to dir container
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./main.go

CMD ["./main"]
# ENTRYPOINT ["main.go"]