# Create a new row
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
