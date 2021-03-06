WiP
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)
## Unit Tests
[![Latest release](https://badgen.net/github/release/Naereen/Strapdown.js)](https://github.com/golang/mock)

For autogenerate the mock ill use mockgen and gomock, in order to autogenerate the mock code in reflect mode.
Set following cmd at top of rocket.go:
```
//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/rsh456/grpc-services/internal/rocket/internal/rocket Store
```
destination: name of autogenerated file
package: package name
path: dir of interface we want to mock and run
```
go generate ./...
```

<p align="center">
<img src="https://github.com/rsh456/grpc-services/blob/master/images/gogenerate.PNG" border="10"/>
</p>

## Implementing Database Package
[![Latest release](https://badgen.net/github/release/Naereen/Strapdown.js)](https://github.com/jmoiron/sqlx)
 
Run a Postgres database using docker
Use SQLX package for database interactions

## Migrations
If can't run migrations the applications shouldn't start.
up: interpreted as the happy path
down: interpreted as the sad path
This will create the structure of the database, create and alter tables columns. 

## Docker Compose
Will use docker compose which will contain both the definition for the database and file for the  application with the environment variables in it.

Docker file top 5 lines represent the builder docker container
<p align="center">
<img src="https://github.com/rsh456/grpc-services/blob/master/images/docker-1.PNG" border="10"/>
</p>

Define the lightweight production image in which be running the application
<p align="center">
<img src="https://github.com/rsh456/grpc-services/blob/master/images/docker-2.PNG" border="10"/>
</p>
