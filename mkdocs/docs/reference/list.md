# Update a row 
Below API specs allow to read data of multiple rows in a worksheet.

### HTTP Request
`POST https://api.sheetson.com/v2/sheets/`_**sheetName**_

### Path Parameters
|**Parameters**| &nbsp;
|----------|------------
|`sheetName` | `string`  <br /> Name of the sheet containing data

### Headers
|**Parameters**| &nbsp;
|-------|---------
|`Authorization`| `Bearer {YOUR_API_KEY}`
|`X-Spreadsheet-Id`| `string` <br/> The ID of the spreadsheet to create new a row

### URL Parameters
|**Parameters**| &nbsp;
|-------|---------
| limit | `number` <br /> Default to `24`, but can we change upto `100`. Used to limit number of rows returned.
| skip  | `number` <br /> Default to `0`. Used to skip rows in returned data.
| order | `string` <br /> Used to sort result set.
| keys  | `string` <br /> Comma separated list of fields to return.
| where | `string` <br /> Criteria to filter rows.

### Request body
No request body is required.

### Response body
Response body contains JSON object with `results` field contains multiple rows data.
