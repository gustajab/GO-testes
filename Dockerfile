#definindo imagem
FROM golang:1.21.1 AS build

WORKDIR /app

COPY go.mod ./
COPY main.go ./

# compilando o código-fonte go e criando executavel server
RUN go build -o /server

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /server /server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/server" ]