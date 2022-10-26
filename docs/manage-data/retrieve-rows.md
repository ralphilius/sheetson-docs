# Retrieving Rows

## Example data
To get started, we prepare a sheet named **Cities** as below:
<div class='example'>

|  | A               | B     | C       | D          |
|--| --------------- | ----- | ------- | ---------- |
|1 | name            | state | country | population |
|2 | San Francisco   | CA    | USA     | 3314000    |
|3 | Los Angeles     | CA    | USA     | 12458000   |
|4 | Washington, D.C.| null  | USA     | 5207000    |
|5 | Tokyo           | null  | Japan   | 37400068   |
|6 | Shanghai        | null  | China   | 25582000   |
|7 | Paris           | null  | France  | 10901000   |
|8 | São Paulo       | null  | Brazil  | 21650000   |
|9 | Madrid          | null  | Spain   | 6497000    |
|10| Toronto         | null  | Canada  | 6082000	  |
|11| Chicago         | IL    | USA     | 8864000    |

</div>

!!! important "Note"
    Please prepare the same data as above in a new (or existing) spreadsheet. You will need to use your Spreadsheet ID as descibed in [Getting Started](/getting-started) in example codes.

##  Get a row
You can retrieve a row's contents by sending a GET request to the `rowIndex` URL. For example, to retrieve Los Angeles information in **Cities** sheet:

=== "cURL"
    ``` shell
    curl "https://api.sheetson.com/v2/sheets/Cities/3" \
    -G \
    --data-urlencode 'apiKey=YOUR_API_KEY' \
    --data-urlencode 'spreadsheetId=YOUR_SPREADSHEET_ID' \
    ```

=== "Javascript"
    ``` javascript
    const fetch = require('isomorphic-fetch');
    const params = {
      apiKey: "YOUR_API_KEY",
      spreadsheetId: "YOUR_SPREADSHEET_ID"
    }
    const url = = new URL("https://api.sheetson.com/v2/sheets/Cities/3");
    Object.keys(params).forEach(key => url.searchParams.append(key, encodeURIComponent(params[key])));
    fetch(url).then(r => r.json())
      .then(result => console.log(result))
    ```

The response body is a JSON object containing all the header fields, plus the `rowIndex` field:
```
{
  "rowIndex": 3,
  "name": "Los Angeles", 
  "state": "CA", 
  "country": "USA",
  "population": "12458000"
}
```
## Get multiple rows
You can retrieve multiple objects at once by sending a `GET` request to the class URL. Without any URL parameters, this simply lists objects in the class:

=== "cURL"
    ``` shell
    curl "https://api.sheetson.com/v2/sheets/Cities" \
    -H "Authorization: Bearer YOUR_API_KEY" \
    -H "X-Spreadsheet-Id: YOUR_SPREADSHEET_ID" \
    ```

=== "Javascript"
    ``` javascript
    const fetch = require('isomorphic-fetch');
    const params = {
      apiKey: "YOUR_API_KEY",
      spreadsheetId: "YOUR_SPREADSHEET_ID"
    }
    const url = = new URL("https://api.sheetson.com/v2/sheets/Cities");
    Object.keys(params).forEach(key => url.searchParams.append(key, encodeURIComponent(params[key])));
    fetch(url).then(r => r.json())
      .then(result => console.log(result))
    ```

The return value is a JSON object that contains a `results` field with a JSON array that lists the objects
```
{
  results: [
    {
      "rowIndex": 2,
      "name": "San Francisco", 
      "state": "CA", 
      "country": "USA",
      "population": "860000"
    },
    {
      "rowIndex": 3,
      "name": "Los Angeles", 
      "state": "CA", 
      "country": "USA",
      "population": "12458000"
    }
    ...
  ]
}
```

## Filter rows
There are several ways to filter rows based on specific criteria. Refer to [Paginate & refine data](/manage-data/paginate-filter) in order to perform below operations:

 - Paginate between multiple results set
 - Limit number of rows returned
 - Order rows by specific field
 - Search for rows with specific criteria
