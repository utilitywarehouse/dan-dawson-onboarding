FROM golang:alpine
WORKDIR /app
COPY . .
WORKDIR /app/cmd
RUN go mod tidy
RUN go build -o /time
EXPOSE 8080
CMD [ "/time" ]