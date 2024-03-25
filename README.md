# sbarter_be_base_examples

This repository provides base project examples for backend services. It's crucial to note that this project serves as a guideline to ensure consistent standards across services, thereby enhancing debugging and maintenance processes. While these examples predominantly utilize `gin-gonic` for HTTP implementations, the structures and principles showcased here can seamlessly adapt to support GRPC implementations as well.

## Base Structure

The structure showcased in this example underscores meticulous organization aimed at promoting code reuse and maintainability across a spectrum of microservices. It's important to emphasize that this structure is presented purely for illustrative purposes.

### Directories

The project is compartmentalized into the following directories:

- `sbartererrors`: Reserved for components pertinent to error handling and logging the error stack.

- `sbarterlog`: Houses components essential for managing the project's foundational logging mechanisms.

- `sbarternetwork`: Contains components facilitating call stack management for streamlined debugging processes.

- `sbarterservice`: Encompasses microservice-related components including base handlers, middlewares, models, and responses.

- `sbarterutils`: Dedicated to utility functions catering to a variety of needs such as data masking, UUID generation, cryptography, and more.

## Response Examples

To explore the structure of the example response objects, refer to the `http_response.go` file located in the `responses` directory. Please note that these examples represent the structure of response objects and do not contain actual implementation details.

### Success Response (HTTP 200 OK):

```json
{
  "success": true,
  "data": {},
  "errors": null
}
```

### Success Response with CallStack for easier debugging (HTTP 200 OK):

```json
{
  "success": true,
  "data": {},
  "errors": null,
  "callstack": {
    "sequenceId": 6,
    "direction": -1,
    "correlationId": "a6a8bd65-60b8-4b12-8d57-9936e77162f6",
    "timestamp": 1711295866375506700,
    "project": "service_proxy_api",
    "function": "handlers.(*SystemUserHandler).FindAllSystemUsers",
    "path": "",
    "line": 69,
    "previous": {
      "sequenceId": 5,
      "direction": -1,
      "correlationId": "a6a8bd65-60b8-4b12-8d57-9936e77162f6",
      "timestamp": 1711295866375500300,
      "project": "service_proxy_api",
      "function": "managers.(*SystemUserManager).FindAllSystemUsers",
      "path": "",
      "line": 55,
      "previous": {
        "sequenceId": 4,
        "direction": -1,
        "correlationId": "a6a8bd65-60b8-4b12-8d57-9936e77162f6",
        "timestamp": 1711295866375492400,
        "project": "service_core",
        "function": "service.(*coreService).CreateGenericRoute",
        "path": "/service_core@v0.0.1/src/shared/service/service.go",
        "line": 72,
        "previous": {
          "sequenceId": 3,
          "direction": 1,
          "correlationId": "a6a8bd65-60b8-4b12-8d57-9936e77162f6",
          "timestamp": 1711295866359399700,
          "project": "service_core",
          "function": "service.(*coreService).CreateGenericRoute",
          "path": "/service_core@v0.0.1/src/shared/service/service.go",
          "line": 61,
          "previous": {
            "sequenceId": 2,
            "direction": 1,
            "correlationId": "a6a8bd65-60b8-4b12-8d57-9936e77162f6",
            "timestamp": 1711295866359392500,
            "project": "service_proxy",
            "function": "managers.(*SystemUserManager).FindAllSystemUsers",
            "path": "",
            "line": 35,
            "previous": {
              "sequenceId": 1,
              "direction": 1,
              "correlationId": "a6a8bd65-60b8-4b12-8d57-9936e77162f6",
              "timestamp": 1711295866356060200,
              "project": "service_proxy",
              "function": "handlers.(*SystemUserHandler).FindAllSystemUsers",
              "path": "",
              "line": 61
            }
          }
        }
      }
    }
  }
}
```

### Error Response (HTTP 400 Bad Request):

```json
{
  "success": false,
  "data": null,
  "errors": [
    {
      "code": 400,
      "message": "bad request",
      "type": "BAD_REQUEST_ERROR"
    }
  ],
  "errorstack": {
    "code": 400,
    "message": "bad request",
    "timestamp": 1711294659411757000,
    "runtimeinfo": {
      "project": "service_name",
      "function": "handlers.(*UserHandler).FindAll",
      "path": "/service_name/src/internal/handlers/user.go",
      "line": 42
    },
    "previous": {
      "code": 2001,
      "message": "limit - Value must be greater than 0",
      "timestamp": 1711294659411749000,
      "runtimeinfo": {
        "project": "service_name",
        "function": "handlers.(*UserHandler).FindAll",
        "path": "/service_name/src/internal/handlers/user.go",
        "line": 42
      }
    }
  }
}
```

## Logging Examples:

To explore the structure for logging, refer to the `base_handler.go` file located in the `handlers` directory. This file defines the logging structure and format used throughout the examples.

### With error:

```json
{
  "ClientIP": "::1",
  "ClientUserAgent": "PostmanRuntime/7.36.3",
  "CorrelationID": "431d539f-72be-4f92-9fe0-352d7b4e0d36",
  "CorrelationService": "service_name",
  "CorrelationTimeStart": "1711294659411441000",
  "Duration": 0.75,
  "Env": "local",
  "Referer": "",
  "Response": "{'errors': [{'code': 400,'message': 'bad request','type': 'BAD_REQUEST_ERROR'}],'errorstack': {'code': 400,'message': 'bad request','previous': {'code': 1,'message': 'limit - Value must be greater than 0','runtimeinfo': {'function': 'handlers.(*UserHandler).FindAll','line': 42,'path': '/service_name/src/internal/handlers/user.go','project': 'service_name'},'timestamp': 1711294659411749000},'runtimeinfo': {'function': 'handlers.(*UserHandler).FindAll','line': 42,'path': '/service_name/src/internal/handlers/user.go','project': 'service_name'},'timestamp': 1711294659411757000},'success': false}",
  "Service": "service_name",
  "StatusCode": 200,
  "Timestamp": 1711294659412382000,
  "Type": "Response",
  "level": "debug",
  "msg": "POST /api/v1/UserHandler.FindAll",
  "time": "2024-03-24T16:37:39+01:00"
}
```

### With masking:

```json
{
  "ClientIP": "::1",
  "ClientUserAgent": "PostmanRuntime/7.36.3",
  "CorrelationID": "4878ecc5-41bc-419a-8a3a-56af7dd9f5a2",
  "CorrelationService": "service-core",
  "CorrelationTimeStart": "1711295111109049000",
  "Duration": 27.55,
  "Env": "local",
  "LocalEnv": "generic",
  "Referer": "",
  "Response": "{'data': {'id': 1, 'name': '***'},'success': true}",
  "Service": "service-core",
  "StatusCode": 200,
  "Timestamp": 1711295111137449000,
  "Type": "Response",
  "level": "debug",
  "msg": "POST /api/v1/SystemUserHandler.FindByID",
  "time": "2024-03-24T16:45:11+01:00"
}
```

## GET Version:

All microservices should include a GET version endpoint to facilitate communication with the infrastructure and other components.

```json
{
  "success": true,
  "data": {
    "ApplicationInfo": {
      "Name": "Service Core",
      "Version": "v1.0.0",
      "Service": "service-core",
      "Environment": "local",
      "CommitDate": "24-03-2024 09:00:00",
      "BuildDate": "24-03-2024 09:00:00",
      "BuildHash": "1cc034654e864eb494560b7b7fbcb29d383cfa2c"
    }
  },
  "errors": null
}
```

## Middlewares:

Middlewares play a crucial role in enhancing the functionality and security of our backend services. They intercept incoming requests and responses, allowing us to perform various operations.

For detailed information on our example middlewares, please refer to the following files:

- `correlation_id.go`: This middleware manages correlation IDs, ensuring traceability and debugging capabilities across distributed systems.

- `cors.go`: The CORS (Cross-Origin Resource Sharing) middleware facilitates secure communication between client-side web applications and our server, enforcing proper access control policies.

- `request_logger.go`: The request logger middleware captures detailed information about incoming requests, providing valuable insights for monitoring and debugging purposes.
