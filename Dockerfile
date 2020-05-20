# Build Stage
FROM golang:1.12 as build-env

RUN curl https://deb.nodesource.com/setup_12.x | bash
RUN curl https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add -
RUN echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list

RUN set -xe && \
    apt-get update >/dev/null && \
    apt-get -y install nodejs yarn

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