## Introduction
Welcome to the Sheetson API! You can use our API to access data in any Google Sheet file that you have access to. Basically, every Google Sheet file is a database, and each tab within that file is a table. Sheetson API is designed around these definition.

#### API Endpoint
All API calls must be made to `https://api.sheetson.com`

#### JSON only
All responses will be in JSON. Input data passed through the request body can be form-encoded or JSON-encoded. If using a JSON body, please specify the `Content-Type` header as `application/json`.

#### String-only
Currently, all values in row object are returned with an exception of rowIndex which is automatically added to the results set.

## Keys & Authentication
In order to protect your data, all API requests to Sheetson require an API Key which can be obtained in [Console](https://sheetson.com/console)
You need to  provide an `Authentication` header with API Key as the value.

> Authentication: Bearer `YOUR_API_KEY`

## Status Codes
Every API request to Sheetson API returns a HTTP status code that reflect the result. Following codes should be returned by Sheetson API.

#### Successful requests
Every API request to Sheetson API returns a HTTP status code that reflect the result. Following codes should be returned by Sheetson API.

|CODE |MEANING
|-----|-----
|`200 OK`| A resource has been retrieved (GET requests)
|`201 Created`|	A resource has been created (POST requests)
|`204 No Content`|	A resource has been deleted (DELETE requests)

#### Failed requests
APIs that failed to fulfill its purpose should still have `200 OK` status code, but their response body should tell more about type of error. See [Status Codes](/status-codes) for more information.


