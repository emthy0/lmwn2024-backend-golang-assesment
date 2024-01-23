FROM golang:alpine AS build

WORKDIR /src/app

COPY go.mod go.sum ./

ENV GO111MODULE=on
RUN go mod download
COPY . . 
RUN go build -o server .

FROM golang:alpine AS server

WORKDIR /app
COPY --from=build /src/app .

EXPOSE 4000

CMD ./server