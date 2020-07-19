# Creating Rows

## Example data
To get started, we prepare a sheet named **Cities** as below:
<div class='example'>

| | A | B | C | D |
|-|---|---|---|---|
|1| name | state | country | population |
|2|   |   |   |   |

</div>

!!! important "Note"
    Please prepare the same data as above in a new (or existing) spreadsheet. You will need to use your Spreadsheet ID as descibed in [Getting Started](/getting-started) in example codes.

## Add a row

To create a new row in **Cities** sheet, send a POST request to that sheet API endpoint containing the contents of the row.

=== "cURL"
    ``` shell
    curl "https://api.sheetson.com/v2/sheets/Cities" \
    -X POST \
    -H "Authorization: Bearer YOUR_API_KEY" \
    -H "X-Spreadsheet-Id: YOUR_SPREADSHEET_ID" \
    -H "Content-Type: application/json" \
    -d '{"name": "San Francisco", "state": "CA", "country": "USA", "population": "860000"}'
    
    ```

=== "Javascript"
    ``` javascript
    const fetch = require('isomorphic-fetch');
    fetch("http://api.sheetson.com/v2/sheets/Cities", {
      method: "POST",
      headers: {
        "Authorization": "Bearer YOUR_API_KEY",
        "X-Spreadsheet-Id": "YOUR_SPREADSHEET_ID",
        "Content-Type": "application/json"
      },
      body: JSON.stringify({name: 'San Francisco', state: 'CA', country: 'USA', population: 860000})
    }).then(r => r.json())
    .then(result => console.log(result))
    ```

When the creation is successful, the HTTP response is a `201 Created` and response body is a JSON object containing the `rowIndex` of the newly created row:
```
{
  "rowIndex": 2,
  "name": "San Francisco", 
  "state": "CA", 
  "country": "USA",
  "population": "860000"
}
```
The sheet content should look like below:
<div class='example'>

| |A | B | C | D |
|-|-------------|--|---|------|
|1|name|state|country|population
|2|San Francisco|CA|USA|860000

</div>
