<h1 style="color:skyblue;">Template Echo</h1>

## Feature branch

| Branch         | Status | Detail                                                                  |
| -------------- | ------ | ----------------------------------------------------------------------- |
| master         | Open   | Origin, Logger, Config, Dockerfile, docker-compose, Precommit, Prettier |
| cqrs           | Open   | CQRS                                                                    |
| configs        | Open   | Configs                                                                 |
| microservice   | Open   | PubSub, NATs, Microservice                                              |
| typeorm        | Open   | TypeORM, PostgreSQL                                                     |
| prisma         | Open   | Prisma                                                                  |
| caching        | Open   | Redis                                                                   |
| queuing        | Open   | RabbitMQ, (kafka)                                                       |
| mail           | Open   | Send mail                                                               |
| rbac           | Open   | Authentication, Authorization, Permission                               |
| oauth          | Open   |                                                                         |
| cicd           | Open   | CI/CD                                                                   |
| nx             | Open   | NX Workspace                                                            |
| elasticsearch  | Open   |                                                                         |
| mongoose       | Open   | Mongoose, MongoDB                                                       |
| grpc           | Open   |                                                                         |
| firebase       | Open   |                                                                         |
| graphql        | Open   | GraphQL                                                                 |
| interceptor    | Open   | Interceptor                                                             |
| image          | Open   | Upload, Download                                                        |
| video          | Open   | Upload, Download, (Streaming)                                           |
| clean          | Open   | Clean architecture                                                      |
| solid          | Open   |                                                                         |
| observer       | Open   | [Design-Pattern] Observer Pattern                                       |
| facade         | Open   | [Design-Pattern] Facede Pattern                                         |
| proxy          | Open   | [Design-Pattern] Proxy Pattern                                          |
| simple_factory | Open   | [Design-Pattern] Simple Factory Pattern                                 |
| factory_method | Open   | [Design-Pattern] Factory Method Pattern                                 |
| singleton      | Open   | [Design-Pattern] Singleton Pattern                                      |
| prototype      | Open   | [Design-Pattern] Prototype Pattern                                      |

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
