FROM golang:alpine3.18 AS build-stage

WORKDIR /app/src/duck-cook-recipe

ENV GOPATH=/app

COPY . .

RUN chmod +x /app/src/duck-cook-recipe

RUN go test ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o /duck-cook-recipe

FROM alpine:latest AS build-release-stage

ENV TZ=America/Sao_Paulo

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /app/src/duck-cook-recipe

RUN chmod +x /app/src/duck-cook-recipe

COPY --from=build-stage /duck-cook-recipe /app/src/duck-cook-recipe/duck-cook-recipe

EXPOSE 8080

ENTRYPOINT ["/app/src/duck-cook-recipe/duck-cook-recipe"]