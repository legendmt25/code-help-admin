FROM golang:1.24.1-alpine3.21 as builder

COPY . /build

WORKDIR /build

RUN go mod download

RUN CGO_ENABLED=0 go build ./proxy/cmd/main.go


FROM scratch

ARG SOURCE_DIRECTORY

WORKDIR /app

COPY --from=builder /build/main ./main

EXPOSE 4000

ENTRYPOINT ["./main"]