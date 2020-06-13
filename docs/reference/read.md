# Update a row 
Below API specs allow to read data of a specific row in a worksheet.

### HTTP Request
`POST https://api.sheetson.com/v2/sheets/`_**sheetName**_`/`_**rowIndex**_

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

### URL Parameters
|**Parameters**| &nbsp;
|-------|---------
| keys  | `string` <br /> Comma separated list of fields to return.

### Request body
No request body is required.

### Response body
Response body contains JSON object with `rowIndex` and row data.
