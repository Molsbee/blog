# Build Stage
FROM golang:1.12 as build-env

RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN yarn --cwd ./frontend/ build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o blog

FROM scratch
COPY --from=build-env /app/frontend/dist /frontend/dist
COPY --from=build-env /app/database-migrations /database-migrations
COPY --from=build-env /app/blog /blog
EXPOSE 8080
ENTRYPOINT ["/blog"]