FROM golang:alpine
WORKDIR /app
COPY . .
WORKDIR /app/cmd
RUN go mod tidy
RUN go build -o /time
EXPOSE 8080
EXPOSE 8081
CMD [ "/time" ]