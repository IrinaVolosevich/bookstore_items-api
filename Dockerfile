FROM golang:1.12.13

ENV REPO_URL=bookstore_items-api
ENV GOPATH=/home/irina/go
ENV APP_PATH=$GOPATH/src/$REPO_URL
ENV WORK_PATH=$APP_PATH/src

COPY src $WORK_PATH
WORKDIR $WORK_PATH

RUN go build -o items-api .

EXPOSE 8081

CMD ["./items-api"]