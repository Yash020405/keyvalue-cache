# Key-Value Cache

## Overview
This project implements an in-memory key-value cache service in Go, designed for high-performance storage and retrieval of key-value pairs. The service exposes two endpoints:
- `POST /put` to insert or update key-value pairs
- `GET /get` to retrieve values by key

The cache is designed to be lightweight and efficient, with optimizations for high throughput under concurrent load. The service is containerized using Docker and tested using Locust for load testing.

## Features
- Thread-safe in-memory key-value storage
- REST API for `PUT` and `GET` operations
- Dockerized for easy deployment
- Load testing setup with Locust

## Setup and Installation
### Prerequisites
- Docker installed on your system
- Go installed for local development (optional)

### Building and Running the Docker Container
To build and run the Docker container:

#### Build the Docker image
```bash
 docker build -t thetallinnov8r/keyvalue-cache:v1.0.0 .
```

#### Run the container
```bash 
docker run -p 7171:7171 thetallinnov8r/keyvalue-cache:v1.0.0 
```
## API Endpoints
### PUT `/put`
#### Request:
```json
{
  "key": "exampleKey",
  "value": "exampleValue"
}
```
#### Response:
```json
{
  "status": "OK",
  "message": "Key inserted/updated successfully."
}
```

### GET `/get?key={key}`
#### Response (Success):
```json
{
  "status": "OK",
  "key": "exampleKey",
  "value": "exampleValue"
}
```

#### Response (Key Not Found):
```json
{
  "status": "ERROR",
  "message": "Key not found."
}
```

## Load Testing with Locust

To stress-test the cache service, use Locust:

```sh
locust -f locustfile.py --host=http://localhost:7171
```
This script simulates concurrent GET and PUT requests to evaluate performance.

## Optimizations Implemented
### Efficient Synchronization:

* Utilized sync.RWMutex for read/write operations to maximize concurrency.

#### String Length Validation:

* Limits keys and values to 256 ASCII characters to prevent excessive memory usage.

#### Pre-generated Key Pool for Load Testing:

* Enhances test consistency by reducing the impact of random key selection.

#### Lightweight Docker Image:

* Uses a multi-stage build to minimize final image size and improve deployment speed.

#### Zero Wait Time in Load Testing:

* Ensures that requests are sent at maximum throughput during stress tests.





## License

[MIT](https://choosealicense.com/licenses/mit/)