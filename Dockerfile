FROM golang:latest

RUN mkdir /app
RUN mkdir /data

COPY . /app

WORKDIR /app

# RUN apk add --no-cache libwebp libwebp-tools
RUN apt-get update 
RUN yes | apt-get install libwebp-dev

# GET DEPS
RUN go get -u github.com/kolesa-team/go-webp
RUN go mod download
# RUN go get github.com/gorilla/mux
# RUN go get github.com/joho/godotenv
# RUN go get github.com/fsnotify/fsnotify
# RUN go get github.com/kolesa-team/go-webp
# RUN go get golang.org/x/net/webdav

RUN go build -o server /app/cmd/api/main.go

CMD ["/app/server"]
