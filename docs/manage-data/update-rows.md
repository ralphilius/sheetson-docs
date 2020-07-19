# Updating Rows

## Example data
To get started, we prepare a sheet named **Cities** as below:
<div class='example'>

| |A | B | C | D |
|-|-------------|--|---|------|
|1|name|state|country|population
|2|San Francisco|CA|USA|860000

</div>

!!! important "Note"
    Please prepare the same data as above in a new (or existing) spreadsheet. You will need to use your Spreadsheet ID as descibed in [Getting Started](/getting-started) in example codes.

## Update a row
To change the data on a row that already exists, send a `PUT` request to the `rowIndex` URL. Any fields you don’t specify will remain unchanged, so you can update just a subset of the row's data. 

For example, to change population of San Francisco in **Cities** sheet.

=== "cURL"
    ``` shell
    curl "https://api.sheetson.com/v2/sheets/Cities/2" \
    -X PUT \
    -H "Authorization: Bearer YOUR_API_KEY" \
    -H "X-Spreadsheet-Id: YOUR_SPREADSHEET_ID" \
    -H "Content-Type: application/json" \
    -d '{"population" : 3314000}'
    ```

=== "Javascript"
    ``` javascript
    const fetch = require('isomorphic-fetch');
    fetch("http://api.sheetson.com/v2/sheets/Cities/2", {
      method: "PUT",
      headers: {
        "Authorization": "Bearer YOUR_API_KEY",
        "X-Spreadsheet-Id": "YOUR_SPREADSHEET_ID",
        "Content-Type": "application/json"
      },
      body: JSON.stringify({"population" : 3314000})
    }).then(r => r.json())
    .then(result => console.log(result))
    ```

The response body is a JSON object containing requesting data

```
{
  "rowIndex": 2,
  "population": "3314000"
}
```

The sheet content should look like below:
<div class='example'>

| |A | B | C | D |
|-|-------------|--|---|------|
|1|name|state|country|population
|2|San Francisco|CA|USA|3314000

</div>
