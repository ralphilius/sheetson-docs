# Deleting Rows

## Example data
To get started, we prepare a sheet named **Cities** as below:
<div class='example'>

| | A             | B     | C       | D          |
|-| ------------- | ----- | ------- | ---------- |
|1| name          | state | country | population |
|2| San Francisco | CA    | USA     | 3314000    |
|3| Gotham City   | NJ    | USA     | 30000000   |

</div>

!!! important "Note"
    Please prepare the same data as above in a new (or existing) spreadsheet. You will need to use your Spreadsheet ID as descibed in [Getting Started](/getting-started) in example codes.

## Delete a row
To delete a row from **Cities** sheet, send a DELETE request to its `rowIndex` URL.

=== "cURL"
    ``` shell
    curl "http://api.sheetson.com/v2/sheets/Cities/3" \
    -X DELETE \
    -H "Authorization: Bearer {YOUR_API_KEY}" \
    -H "X-Spreadsheet-Id: {YOUR_SPREADSHEET_ID}" \
    ```

=== "Javascript"
    ``` javascript
    const fetch = require('isomorphic-fetch');
    fetch("http://api.sheetson.com/v2/sheets/Cities/3", {
      method: "DELETE",
      headers: {
        "Authorization": "Bearer {YOUR_API_KEY}",
        "X-Spreadsheet-Id": "{YOUR_SPREADSHEET_ID}"
      }
    }).then(r => r.json())
    .then(result => console.log(result))
    ```

When the deletion is successful, the HTTP response is a `204 No-Content` without any response body. The sheet content should look like below:
<div class='example'>

| |A | B | C | D |
|-|-------------|--|---|------|
|1|name|state|country|population
|2|San Francisco|CA|USA|3314000

</div>