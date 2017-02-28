FROM golang

ADD task_list.go /
ADD views/* /views/

RUN go build -o /task_list /task_list.go

EXPOSE 8080
ENTRYPOINT ["/task_list"]
