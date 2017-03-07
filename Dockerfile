FROM golang

RUN go get github.com/go-sql-driver/mysql

ADD task_list.go /
ADD templates/* /templates/

RUN go build -o /task_list /task_list.go

EXPOSE 8080
ENTRYPOINT ["/task_list"]
