FROM cosmtrek/air:v1.15.1
RUN go mod tidy
RUN go mod vendor

