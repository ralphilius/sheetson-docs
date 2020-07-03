

# Error Codes
All error responses will contain an object with below attributes:

|CODE	|MEANING
|-----|-------
|`400 Bad Request`  |Invalid request parameters
|`401 Unauthorized`	|Incorrect or missing API key, spreadsheetId or sheetName

All error responses will contain a JSON with `code` and `message` fields that explains the error.

|Name                          |Code |Meaning
|------------------------------|-----|------------
|SpreadsheetIdNotFoundException|101  |Spreadsheet ID is not available
|NoHeaderRowException          |102  |Spreadsheet does not have a header row
|UpdateRowNotFoundException    |103  |Row is not found to update
|InvalidBodyException          |104  |Request body is invalid. This could be because you send an empty or malformed body.
|WorksheetNotFoundException    |105  |The requesting worksheet is not availble in the spreadsheet. This could be because you use a wrong name or case-sensitive name.
