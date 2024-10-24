FROM golang:1.22

WORKDIR /app
COPY . .
# COPY go.mod .
# COPY main.go .
# COPY functions/ functions/
# COPY banners/ banners/
# COPY home.html .

RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]
LABEL maintainer="mostafa.elhadroubi.dev@gmail.com"
LABEL version="1.0"
LABEL description="A Go web server for ASCII art"