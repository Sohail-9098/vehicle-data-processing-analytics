# Vehicle Data Processing and Analytics

This service validates, processes, and stores vehicle telemetry data in a PostgreSQL database. 
It also exposes telemetry data through a REST endpoint.

## Table of Contents

- [Overview](#overview)
- [Setup](#setup)
  - [Prerequisites](#prerequisites)
  - [Database Setup](#database-setup)
  - [Service Setup](#service-setup)
- [Running the Service](#running-the-service)
- [Endpoints](#endpoints)
- [Configuration](#configuration)
- [License](#license)

## Overview

The Vehicle Data Processing and Analytics service is responsible for:
- Receiving telemetry data from vehicles via a gRPC server.
- Validating and storing telemetry data in a PostgreSQL database.
- Providing access to telemetry data through a REST endpoint.

## Setup

### Prerequisites

- Go (for building and running the service)
- PostgreSQL database

### Database Setup

This service connects to an instance of postgresql database running on Aiven cloud.
Please ensure your you have access to aiven cloud and environment variable AIVEN_CREDENTIALS set.

### Service Setup

1. **Clone the Repository**

   ```sh
   git clone https://github.com/Sohail-9098/vehicle-data-processing-analytics.git
   cd vehicle-data-processing-analytics
   ```

2. **Set Up Environment Variables**

   Create a `.env` file in the root directory of the project and add the following:

   ```env
   AIVEN_CREDENTIALS=postgres://user:password@aivencloud:5432/vehicle-data-analytics
   ```

3. **Install Dependencies**

   ```sh
   go mod download
   ```

4. **Build the Service**

   ```sh
   go build -o vehicle-data-processor main.go
   ```

## Running the Service

Start the service:

```sh
./vehicle-data-processor
```

The service will be available at `http://localhost:8080`.

## Endpoints

- **GET /getTelemetry**

  Retrieve telemetry data.

  ```sh
  curl http://localhost:8080/getTelemetry
  ```
