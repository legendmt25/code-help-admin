FROM golang:1.22.1-alpine3.19 as builder

ARG SOURCE_DIRECTORY

COPY . /build

WORKDIR /build/$SOURCE_DIRECTORY

RUN go mod download

RUN CGO_ENABLED=0 go build -o ./main


FROM scratch

ARG SOURCE_DIRECTORY

WORKDIR /app

COPY --from=builder /build/$SOURCE_DIRECTORY/main ./main

EXPOSE 8080

ENTRYPOINT ["./main"]