# Orders Service

## Description

It's educational project.

A simple http service on **Go** that provides authentication, order and service functionality. 
You can register, authorize, create/delete/update/get services and orders.

- **Framework**: Fiber
- **DB library**: Gorm
- **Storage**: PostgresSQL
- **Architecture**: Clean
- **Docs**: Swagger

## Run

Requirements:
- Go
- make
- PostgresSQL

Don't forget to fill in the environment variables for the config (see **Config** section).

```shell
make build
./orders-service
```

## Config
Config for application populates from environment variables.

|       Name       | Description                                       |      Default       |
|:----------------:|:--------------------------------------------------|:------------------:|
|     DB_HOST      | Host of database                                  |     localhost      |
|     DB_PORT      | Port of database                                  |        5432        |
|     DB_USER      | User of database                                  |      postgres      |
|   DB_PASSWORD    | Password for database (user)                      |         -          |
|     DB_NAME      | Name of database                                  | orders_service_dev |
|    DB_SSLMODE    | Switch SSL mode for database                      |      disable       |
|    LOG_LEVEL     | Level of logging                                  |       debug        |
|   SERVER_PORT    | Port of server                                    |        3000        |
|    JWT_SECRET    | JWT Secret (may be any string)                    |         -          |
| JWT_EXPIRE_HOURS | Time after which JWT token will expire (in hours) |         24         |

## Documentation

Documentation is available at `http://localhost:3000/swagger/index.html`
