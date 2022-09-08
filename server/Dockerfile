# Stage 1
FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
RUN mkdir /build 
ADD . /build/
WORKDIR /build
RUN go get -d -v
RUN go build -o frontend .
# Stage 2
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/ /app/
WORKDIR /app
CMD ["./frontend"]
# Stage 3
FROM postgres
ENV POSTGRES_PASSWORD admin
ENV POSTGRES_DB dummydb
ENV PGDATA /var/lib/postgresql/data/mydb
COPY world.sql /docker-entrypoint-initdb.d/