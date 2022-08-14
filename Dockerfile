FROM  golang:1.18-alpine3.15

WORKDIR /app

COPY . .

RUN go build -o todo-list

EXPOSE 3030

CMD ./todo-list