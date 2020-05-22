This guide explains how to use simple CRUD API requests with data within a worksheet as below:

 - [Create a new row](#create-a-new-row)
 - [Retrieve a specific row](#retrieve-a-row)
 - [Update a specific row](#update-a-row)
 - [Delete a specific row](#delete-a-row)

If you want to know how to retrieve complex data, see [Advanced Queries](/queries/)

## Create a new row
Below API specs allow to create a new row at the last row of a worksheet.

### HTTP Request
`POST https://api.sheetson.com/v2/sheets/{sheetName}`

### Path Parameters
|**Parameters**| &nbsp;
|----------|------------
|`sheetName` | `string`  <br /> Name of the sheet containing data

### Headers
|**Parameters**| &nbsp;
|-------|---------
|`Authorization`| `Bearer {YOUR_API_KEY}`
|`X-Spreadsheet-Id`| `string` <br/> The ID of the spreadsheet to create new a row

### Request body
Request body contains object with keys are worksheet headers and values to add into cell.

### Response body
Response body contains JSON object with `rowIndex` and requesting data.

### Examples (Request)
  
=== "cURL"
    ``` shell
    curl "http://api.sheetson.com/v2/sheets/Demo" \
    -X POST \
    -H "API-Key: {YOUR_API_KEY}" \
    -H "X-Spreadsheet-Id: 1h-Eet6qTsPrdL5IVwiAvPXH47wjEludJVpwrpdCWEuM" \
    -H "Content-Type: application/json" \
    -d '{"email" : "something@example.com"}'
    ```

=== "Javascript"
    ``` javascript
    const fetch = require('isomorphic-fetch');
    fetch("http://api.sheetson.com/v2/sheets/Demo", {
      method: "POST",
      headers: {
        "Authorization": "Bearer {YOUR_API_KEY}",
        "X-Spreadsheet-Id": "1h-Eet6qTsPrdL5IVwiAvPXH47wjEludJVpwrpdCWEuM"
      },
      body: JSON.stringify({email: "something@example.com"})
    }).then(r => r.json())
    .then(result => console.log(result))
    ```

### Example (JSON Response)
```
{
  "rowIndex": 1234,
  "email": "something@example.com"
}
```

## Retrieve multiple rows
## Retrieve a row
Below API specs allow to retrieve a row data at a specific index of a worksheet.

### HTTP Request
`GET https://api.sheetson.com/v2/sheets/{sheetName}/{rowIndex}`

### Path Parameters
|**Parameters**| &nbsp;
|----------|------------
|`sheetName`| `string`  <br /> Name of the sheet containing data
|`rowIndex` | `number` <br />Position of the row to retrieve data

### Headers
|**Parameters**| &nbsp;
|-------|---------
|`Authorization`| `Bearer {YOUR_API_KEY}`
|`X-Spreadsheet-Id`| `string` <br/> The ID of the spreadsheet to create new a row

### Request body
No request body is required.

### Response body
Response body contains JSON object with `rowIndex` and all cell data that has headers.

### Examples (Request)
  
=== "cURL"
    ``` shell
    curl "http://api.sheetson.com/v2/sheets/Demo/1" \
    -H "API-Key: {YOUR_API_KEY}" \
    -H "X-Spreadsheet-Id: 1h-Eet6qTsPrdL5IVwiAvPXH47wjEludJVpwrpdCWEuM" \
    ```

=== "Javascript"
    ``` javascript
    const fetch = require('isomorphic-fetch');
    fetch("http://api.sheetson.com/v2/sheets/Demo", {
      headers: {
        "Authorization": "Bearer {YOUR_API_KEY}",
        "X-Spreadsheet-Id": "1h-Eet6qTsPrdL5IVwiAvPXH47wjEludJVpwrpdCWEuM"
      }
    }).then(r => r.json())
    .then(result => console.log(result))
    ```

### Example (JSON Response)
```
{
  "rowIndex": 1234,
  "email": "something@example.com",
  "username": "iamcool",
  "phone": "+12983736165",
  "avatar": "https://source.unsplash.com/random"

}
```
## Update a row 
Below API specs allow to update data of a specific row of a worksheet.

### HTTP Request
`POST https://api.sheetson.com/v2/sheets/{sheetName}/{rowIndex}`

### Path Parameters
|**Parameters**| &nbsp;
|----------|------------
|`sheetName` | `string`  <br /> Name of the sheet containing data
|`rowIndex` | `number` <br />Position of the row to retrieve data

### Headers
|**Parameters**| &nbsp;
|-------|---------
|`Authorization`| `Bearer {YOUR_API_KEY}`
|`X-Spreadsheet-Id`| `string` <br/> The ID of the spreadsheet to create new a row

### Request body
Request body contains object with keys are worksheet headers and values to update into cell.

### Response body
Response body contains JSON object with `rowIndex` and requesting data.

### Examples (Request)
  
=== "cURL"
    ``` shell
    curl "http://api.sheetson.com/v2/sheets/Demo/1234" \
    -X POST \
    -H "API-Key: {YOUR_API_KEY}" \
    -H "X-Spreadsheet-Id: 1h-Eet6qTsPrdL5IVwiAvPXH47wjEludJVpwrpdCWEuM" \
    -H "Content-Type: application/json" \
    -d '{"email" : "something@example.com"}'
    ```

=== "Javascript"
    ``` javascript
    const fetch = require('isomorphic-fetch');
    fetch("http://api.sheetson.com/v2/sheets/Demo", {
      method: "POST",
      headers: {
        "Authorization": "Bearer {YOUR_API_KEY}",
        "X-Spreadsheet-Id": "1h-Eet6qTsPrdL5IVwiAvPXH47wjEludJVpwrpdCWEuM"
      },
      body: JSON.stringify({email: "something@example.com"})
    }).then(r => r.json())
    .then(result => console.log(result))
    ```

### Example (JSON Response)
```
{
  "rowIndex": 1234,
  "email": "something@example.com"
}
```

## Remove a row 