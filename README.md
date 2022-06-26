# FACEIT Backend API Assignment

The web service store user data in a database

## Contents

- Getting Started
    - [Build With](#build-with)
    - [Requirements](#requirements)
    - [Setup](#set-up)
    - [Usage of Makefile](#usage-of-makefile)

### Build With
- Golang
- Fiber
- Kafka
- MongoDB

### Requirements

- [Docker](https://www.docker.com/products/docker-desktop)
- REST client [Postman](https://www.getpostman.com/collections/fb130c44909e4765760c) you can import collection to
  Postman via link.

## Usage of Makefile

|Command| Description                                                                                          |
|-------|------------------------------------------------------------------------------------------------------|
| up | This command will be opening Docker and preparing web service. It will be running on `localhost:3001` |
| run | Web service is ran by this command. It will be running on `localhost:3001`                           |
| unit-test | Run tests                                                                                            |
| generate-mocks | Mocks are created by this command if it is needed.                              |
