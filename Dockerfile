# Build stage
FROM golang:alpine AS build-env
WORKDIR /src
ADD . /src
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -ldflags '-w -s' -a -installsuffix cgo -o /src/run

# Final stage
FROM alpine
WORKDIR /var/app
COPY --from=build-env /src/run /var/app/run
ENTRYPOINT /var/app/run
EXPOSE 5000
