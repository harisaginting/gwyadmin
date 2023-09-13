FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /app/gwyn
COPY . .
COPY ./frontend /app/gwyn/frontend
COPY .env /app/gwyn/.env
RUN apk update && apk add tzdata
ENV TZ=Asia/Jakarta
RUN GOOS=linux go build -ldflags="-s -w" -o ./gwyn-app ./main.go

FROM alpine:3.13
RUN apk --no-cache add ca-certificates
WORKDIR /go/bin
COPY --from=build /app/gwyn /go/bin
EXPOSE 4000
RUN apk update && apk add tzdata
ENV TZ=Asia/Jakarta
ENTRYPOINT /go/bin/gwyn-app --port 4000