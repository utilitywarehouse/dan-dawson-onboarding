FROM golang:latest
WORKDIR /app
COPY . .
WORKDIR /app/cmd
RUN go mod download
RUN go build -o /onboarding
EXPOSE 8080
CMD [ "/onboarding" ]