
FROM golang:1.21.5-alpine AS build


WORKDIR /app


COPY . .


RUN go mod download
RUN go build -o main .


FROM alpine:latest  


WORKDIR /app


COPY --from=build /app/main .


EXPOSE 3000


CMD ["./main"]