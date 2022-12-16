FROM golang:1.20rc1-alpine3.17

WORKDIR /project

COPY go.mod go.sum ./
RUN git config --global user.email "" \
    && git config --global user.password "" \
    && git config --global user.name "" \
    && GO111MODULE="on" \
    && GOSUMDB=off \
    && GOPROXY="https://goproxy.cn,direct" \
    && env \
    && time go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    && go build -o main