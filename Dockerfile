FROM go:1.20.3

WORKDIR /Template_Echo

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -a -ldflags="-s -w -X main.version=$VERSION" -o /app ./main.go