FROM alpine:latest 

RUN mkdir /app

COPY ./expenseApp /app

CMD ["/app/expenseApp"]
