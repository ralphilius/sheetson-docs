# Paginate & Filter Data

Paginating and data filter is required in many applications to display data that is relevant to users. With Sheetson API, there are many ways to narrow down massive data.

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

!!! tip "Note"
    Please prepare the same data as above in a new (or existing) spreadsheet. You will need to use your Spreadsheet ID as descibed in [Getting Started](/getting-started) in example codes.

## Search data

!!! danger ""
    Search data is only available in paid plans.

We can look up data in a worksheet in several ways by using `where` URL parameter. The value of the where parameter should be encoded JSON. To look up cities in USA:

=== "cURL"
    ``` shell
    curl "http://api.sheetson.com/v2/sheets/Cities" \
    -G --data-urlencode 'where={"country":"USA}' \
    -H "Authorization: Bearer {YOUR_API_KEY}" \
    -H "X-Spreadsheet-Id: {YOUR_SPREADSHEET_ID}" \
    ```

=== "Javascript"
    ``` javascript
    const fetch = require('isomorphic-fetch');
    fetch(`https://api.sheetson.com/v2/sheets/Cities?${encodeURIComponent('where={"country":"USA"}')}`, {
      headers: {
        "Authorization": "Bearer {YOUR_API_KEY}",
        "X-Spreadsheet-Id": "{YOUR_SPREADSHEET_ID}"
      }
    }).then(r => r.json())
    .then(result => console.log(result))
    ```

See this lis below for all supported operations in `where` parameter:
<div class='light-table'> 

| Key | Operation
|-----|------------
|$lt  | Less Than
|$lte | Less Than or Equal To
|$gt  | Greater Than
|$gte | Greater Than or Equal To
|$eq  | Equal To
|$neq | Not Equal To

</div>

For example, to retrieve cities with population between 10,000,000 and 30,000,000:

=== "cURL"
    ``` shell
    curl "http://api.sheetson.com/v2/sheets/Cities" \
    -G --data-urlencode 'where={"population": {"$gte": 10000000, "$lte": 30000000}}' \
    -H "Authorization: Bearer {YOUR_API_KEY}" \
    -H "X-Spreadsheet-Id: {YOUR_SPREADSHEET_ID}" \
    ```

=== "Javascript"
    ``` javascript
    const fetch = require('isomorphic-fetch');
    fetch(`https://api.sheetson.com/v2/sheets/Cities?${encodeURIComponent('where={"population": {"$gte": 10000000, "$lte": 30000000}}')}`, {
      headers: {
        "Authorization": "Bearer {YOUR_API_KEY}",
        "X-Spreadsheet-Id": "{YOUR_SPREADSHEET_ID}"
      }
    }).then(r => r.json())
    .then(result => console.log(result))
    ```


## Order data

!!! danger ""
    Order data is only available in paid plans.

By default, rows are returned by order displayed in a worksheet. We can use the order parameter to specify a field to sort by. Prefixing with a negative sign reverses the order. In order to retrieve cities by population ascending order:

=== "cURL"
    ``` shell
    curl "http://api.sheetson.com/v2/sheets/Cities" \
    -G --data-urlencode 'order=population' \
    -H "Authorization: Bearer {YOUR_API_KEY}" \
    -H "X-Spreadsheet-Id: {YOUR_SPREADSHEET_ID}" \
    ```

=== "Javascript"
    ``` javascript
    const fetch = require('isomorphic-fetch');
    fetch(`https://api.sheetson.com/v2/sheets/Cities?${encodeURIComponent('order=population')}`, {
      headers: {
        "Authorization": "Bearer {YOUR_API_KEY}",
        "X-Spreadsheet-Id": "{YOUR_SPREADSHEET_ID}"
      }
    }).then(r => r.json())
    .then(result => console.log(result))
    ```

To get cities by population in descending order:

=== "cURL"
    ``` shell
    curl "http://api.sheetson.com/v2/sheets/Cities" \
    -G --data-urlencode 'order=-population' \
    -H "Authorization: Bearer {YOUR_API_KEY}" \
    -H "X-Spreadsheet-Id: {YOUR_SPREADSHEET_ID}" \
    ```

=== "Javascript"
    ``` javascript
    const fetch = require('isomorphic-fetch');
    fetch(`https://api.sheetson.com/v2/sheets/Cities?${encodeURIComponent('order=-population')}`, {
      headers: {
        "Authorization": "Bearer {YOUR_API_KEY}",
        "X-Spreadsheet-Id": "{YOUR_SPREADSHEET_ID}"
      }
    }).then(r => r.json())
    .then(result => console.log(result))
    ```

## Paginate data
By default, each time we request multiple rows, the maximum number of rows to return is 24. We can also increase this upto 100 by specifying `limit` parameter in the URL. Combining with `skip`, we can paginate between set of rows.

=== "cURL"
    ``` shell
    curl "http://api.sheetson.com/v2/sheets/Cities" \
    -G \
    --data-urlencode 'skip=100' \
    --data-urlencode 'limit=100' \
    -H "Authorization: Bearer {YOUR_API_KEY}" \
    -H "X-Spreadsheet-Id: {YOUR_SPREADSHEET_ID}" \
    ```

=== "Javascript"
    ``` javascript
    const fetch = require('isomorphic-fetch');
    fetch(`https://api.sheetson.com/v2/sheets/Cities?${encodeURIComponent('skip=100,limit=100')}`, {
      headers: {
        "Authorization": "Bearer {YOUR_API_KEY}",
        "X-Spreadsheet-Id": "{YOUR_SPREADSHEET_ID}"
      }
    }).then(r => r.json())
    .then(result => console.log(result))
    ```

!!! tip
    Along with `results`, we also include `hasNextPage` field in response data in order to let you know if there's more data after current set. This field always returns either `true` or `false`.

## Restrict fields returned
To save bandwidth, we can choose to return only needed fields by using `keys` parameter. 

=== "cURL"
    ``` shell
    curl "http://api.sheetson.com/v2/sheets/Cities" \
    -G \
    --data-urlencode 'keys=name,country' \
    -H "Authorization: Bearer {YOUR_API_KEY}" \
    -H "X-Spreadsheet-Id: {YOUR_SPREADSHEET_ID}" \
    ```

=== "Javascript"
    ``` javascript
    const fetch = require('isomorphic-fetch');
    fetch(`https://api.sheetson.com/v2/sheets/Cities?${encodeURIComponent('keys=name,country')}`, {
      headers: {
        "Authorization": "Bearer {YOUR_API_KEY}",
        "X-Spreadsheet-Id": "{YOUR_SPREADSHEET_ID}"
      }
    }).then(r => r.json())
    .then(result => console.log(result))
    ```
