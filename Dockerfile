FROM golang:1.21-alpine as build
WORKDIR /app
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .
RUN go build

FROM golang:1.21-alpine
WORKDIR /app
COPY --from=build /app/calculationimg /app/calculationimg
COPY ./LatinmodernmathRegular.otf /app

CMD ["./calculationimg"]
