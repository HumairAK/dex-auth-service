FROM golang:1.16-alpine  AS build
WORKDIR /das
COPY . .
RUN go build -o /das/das-exec

FROM golang:1.16-alpine
WORKDIR $HOME
COPY --from=build /das ./das
EXPOSE 8080
CMD [ "./das/das-exec" ]