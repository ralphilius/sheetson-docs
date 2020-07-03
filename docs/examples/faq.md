---
title: How to build an FAQ section with Sheetson and Bootstap 4
---
# Build an FAQ with Sheetson and Bootstap 4
## What are we going to build?
We will build a simple FAQ section that can be embeded in your website. The stack we use is:

- Bootstap 4
- jQuery
- HandlebarsJS

![Sheetson FAQ Example](/static/examples/faq.png)
[View demo](/static/examples/faq.html)

## What do we need before we get started?

To get started, please make sure you have prepared below requirements

 - Make a copy of this [Google Sheet](https://docs.google.com/spreadsheets/d/1laM094CMEigLiz4t4FrfFfAr18_hLseOX2M9m4Pai8k/copy) 
 - Follow our [Getting Started](/getting-started) guide to get spreadsheet ID and sheet name
 - A computer with Internet connection
 - Your favourite text editor 

 ## Importing libraries and stylesheets
Let's say we have a basic HTML structure as below:

```html
<html>
<head>

</head>
<body>
	<div id="faq">Loading FAQs...</div>
</body>
</html>
```

Add below scripts/links inside `<head>` tag of your HTML file in order to use Bootstap 4, JQuery and HandlebarsJS.

```html
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.5.0/css/bootstrap.min.css"/>
<script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/handlebars.js/4.0.12/handlebars.min.js"></script>
<script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.5.0/js/bootstrap.min.js" ></script>
```

## Defining HandleBarsJS template
Continue after adding above files, we will add below code which is still inside `<head>` tag. 

```html
<script id="item-template" type="text/x-handlebars-template">
		<div class="container py-3">
    <div class="row">
        <div class="col-10 mx-auto">
            <div class="accordion" id="faqExample">
              {{#each this}}
                <div class="card">
                    <div class="card-header p-2" id="heading{{rowIndex}}">
                        <h5 class="mb-0">
                            <button class="btn btn-link" type="button" data-toggle="collapse" data-target="#collapse{{rowIndex}}" aria-expanded="false" aria-controls="collapse{{rowIndex}}">
                              Q: {{question}}
                            </button>
                          </h5>
                    </div>

                    <div id="collapse{{rowIndex}}" class="collapse" aria-labelledby="heading{{rowIndex}}" data-parent="#faqExample">
                        <div class="card-body">
                            <b>Answer:</b> {{answer}}
                        </div>
                    </div>
                </div>
              {{/each}}
            </div>
        </div>
    </div>
  </div>
</script>
```

## Fetching data & bind to HTML
Right before closing `</head>` tag, paste below code to set up JQuery to make API call and bind data to the template we defined earlier. Remember to replace `YOUR_API_KEY` and `YOUR_SPREADSHEET_ID` with real credentials acquired from earlier steps.

```html
<script type="text/javascript">
		$(document).ready(function(){
      fetch('https://api.sheetson.com/v2/sheets/FAQ', {
        withCredentials: true,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer YOUR_API_KEY',
          'X-Sheetson-Spreadsheet-Id': 'YOUR_SPREADSHEET_ID'
        }
      }).then(r => r.json())
      .then(data => {
        var template = Handlebars.compile($('#item-template').html())
        $('#faq').html(template(data.results))
      })
      .catch(error => console.log(error))
    })
	</script>
```

Save the file and you can open it using Chrome or Firefox.