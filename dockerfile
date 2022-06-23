FROM golang:latest

ENV APP_HOME /app/DompetKilat
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"
COPY . "$APP_HOME/"

RUN go mod download
RUN go mod verify
RUN go build .

EXPOSE 8000
CMD ["./DompetKilat-SimpleWeb"]