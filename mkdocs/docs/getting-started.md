# Getting Started

## Preparing worksheets
It is mandatory that you should always identify the first row of your sheet as header. You can use any character or any length, but it is recommended to use alphanumeric characters only. Take this sheet below as an example
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


The fields `name`, `state`, `country` and `population` will be used as object keys. Each row afterwards are transferred to below JSON:

```
{
  "rowIndex": 2
  "name": "San Francisco",
  "state": "CA",
  "country": "USA",
  "population": "3314000"
}

```

## Sharing your spreadsheet 
Sheetson has taken security serious by not letting your spreadsheet published to the world. Instead, you only need to share your spreadsheet with our email `google@sheetson.com` as an editor, then we will handle the rest.

![](/static/sheetson_preparation_01_docked.png)

## Getting spreadsheet ID and sheet name
Spreadsheet ID and Sheet Name are required to manipulate data. Below is a guide where you can find such information.

!!! danger ""
    Sheet Name is case sensitive, so please make sure you use correct name when making API requests.

![](/static/sheetson_preparation_02.jpg)