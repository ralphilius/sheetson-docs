# Delete a sheet row 
Below API specs allow to delete a specific row in a worksheet.

### HTTP Request
`DELETE https://api.sheetson.com/v2/sheets/`_**sheetName**_`/`_**rowIndex**_

### Path Parameters
|**Parameters**| &nbsp;
|----------|------------
|`sheetName` | `string`  <br /> Name of the sheet containing data
|`rowIndex` | `number` <br />Position of the row to delete data

### Headers
|**Parameters**| &nbsp;
|-------|---------
|`Authorization`| `Bearer YOUR_API_KEY`
|`X-Spreadsheet-Id`| `string` <br/> The ID of the spreadsheet to create new a row

### Request body
No request body is required.

### Response body
No response body is returned.