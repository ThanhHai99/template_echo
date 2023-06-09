<h1 style="color:skyblue;">Template Echo</h1>

## Feature branch

${\color{darkgrey}\textbf{Note}}$ &#58; &emsp;
${\color{red}\textbf{Open}}$ &emsp;
${\color{orange}\textbf{InProress}}$ &emsp;
${\color{green}\textbf{Done}}$ &emsp;

| Branch            | Status                               | Detail                                                                  |
| ----------------- | ------------------------------------ | ----------------------------------------------------------------------- |
| master            | ${\color{green}\textbf{Done}}$       | Origin, Logger, Config, Dockerfile, docker-compose, Precommit, Prettier |
| session           | ${\color{red}\textbf{Open}}$         | Session                                                                 |
| cookie            | ${\color{red}\textbf{Open}}$         | Cookie                                                                  |
| cqrs              | ${\color{red}\textbf{Open}}$         | CQRS                                                                    |
| configs           | ${\color{green}\textbf{Done}}$       | Configs                                                                 |
| microservice      | ${\color{red}\textbf{Open}}$         | PubSub, NATs, Microservice                                              |
| gorm.postgresql   | ${\color{green}\textbf{Done}}$       | GORM, PostgreSQL                                                        |
| gorm.mysql        | ${\color{red}\textbf{Open}}$         | GORM, MySQL                                                             |
| prisma.postgresql | ${\color{red}\textbf{Open}}$         | Prisma, PostgreSQL                                                      |
| caching.redis     | ${\color{green}\textbf{Done}}$       | Redis                                                                   |
| caching.memcached | ${\color{red}\textbf{Open}}$         | Memcached                                                               |
| queuing.rabbitmq  | ${\color{red}\textbf{Open}}$         | RabbitMQ                                                                |
| queuing.kafka     | ${\color{red}\textbf{Open}}$         | Kafka                                                                   |
| mail              | ${\color{red}\textbf{Open}}$         | Send mail                                                               |
| rbac              | ${\color{red}\textbf{Open}}$         | Authentication, Authorization, Permission                               |
| oauth             | ${\color{red}\textbf{Open}}$         |                                                                         |
| cicd              | ${\color{red}\textbf{Open}}$         | CI/CD                                                                   |
| nx                | ${\color{red}\textbf{Open}}$         | NX Workspace                                                            |
| elasticsearch     | ${\color{red}\textbf{Open}}$         | Elasticsearch                                                           |
| grpc              | ${\color{red}\textbf{Open}}$         | gRPC                                                                    |
| firebase          | ${\color{red}\textbf{Open}}$         |                                                                         |
| graphql           | ${\color{orange}\textbf{InProress}}$ | GraphQL                                                                 |
| interceptor       | ${\color{red}\textbf{Open}}$         | Interceptor                                                             |
| image             | ${\color{red}\textbf{Open}}$         | Upload, Download                                                        |
| video             | ${\color{red}\textbf{Open}}$         | Upload, Download, (Streaming)                                           |
| clean             | ${\color{red}\textbf{Open}}$         | Clean architecture                                                      |
| solid             | ${\color{red}\textbf{Open}}$         |                                                                         |
| observer          | ${\color{red}\textbf{Open}}$         | Observer Pattern                                                        |
| facade            | ${\color{red}\textbf{Open}}$         | Facede Pattern                                                          |
| proxy             | ${\color{red}\textbf{Open}}$         | Proxy Pattern                                                           |
| simple_factory    | ${\color{red}\textbf{Open}}$         | Simple Factory Pattern                                                  |
| factory_method    | ${\color{red}\textbf{Open}}$         | Factory Method Pattern                                                  |
| singleton         | ${\color{red}\textbf{Open}}$         | Singleton Pattern                                                       |
| prototype         | ${\color{red}\textbf{Open}}$         | Prototype Pattern                                                       |
| security.request  | ${\color{red}\textbf{Open}}$         | Encode request, Decode request                                          |
| security.response | ${\color{red}\textbf{Open}}$         | Encode response, Decode response                                        |

## Version

Echo `v4.10.2`<br/>
Go `v1.20.3`<br/>

## Description

[Echo](https://echo.labstack.com/) framework Golang starter repository.

## Installation

```bash
$ go mod download
```

## Running the app

```bash
# development
$ go run main.go

# watch mode
$ go run main.go

# production mode
$ go run main.go
```

## Test

```bash
# unit tests
$ npm run test

# e2e tests
$ npm run test:e2e

# test coverage
$ npm run test:cov
```

## Support

Echo is an MIT-licensed open source project. It can grow thanks to the sponsors and support by the amazing backers. If you'd like to join them, please [read more here](https://docs.nestjs.com/support).

## Stay in touch

- Author - [Kamil Myśliwiec](https://kamilmysliwiec.com)
- Website - [https://nestjs.com](https://nestjs.com/)
- Twitter - [@nestframework](https://twitter.com/nestframework)

## License

Echo is [MIT licensed](LICENSE).
