FROM node:16.19.0


WORKDIR /Template_Echo

COPY go.mod ./

RUN go mod download

COPY . .

# RUN npm run build

EXPOSE 80

CMD go run main.go