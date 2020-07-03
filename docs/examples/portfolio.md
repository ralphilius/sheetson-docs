---
title: How to build a Portfolio with Sheetson and VueJS
---
# Build a portfolio with Sheetson and VueJS
## What are we going to build?
We will build a simple portfolio website that has header, portfolio items and footer. The data is loaded from a Google Sheet using Sheetson. The stack we will use includes:

- VueJS
- TailwindCSS
- Axios (for making API request)

![Sheetson Portfolio Example](/static/examples/portfolio.png)
[View demo](/static/examples/portfolio.html)

## What do we need before we get started?

To get started, please make sure you have prepared below requirements

 - Make a copy of this [Google Sheet](https://docs.google.com/spreadsheets/d/1UibHVUlXtfO6Ik6XElPKlt1xsAb5XqTLrCMenIh1raE/copy) 
 - Follow our [Getting Started](/getting-started) guide to get spreadsheet ID and sheet name
 - A computer with Internet connection
 - Your favourite text editor 

## Spreadsheet structure
In the copied Google Sheet, you should see 2 sheets named `Meta` and `PortfolioItems`. **Meta** manages website inormation, while **PortfolioItems** controls what projects to display in the porfolio grid. 

## Importing libraries and stylesheets
Let's say we have a basic HTML structure as below:

```html
<html>
<head>

</head>
<body>

</body>
</html>
```

Add below scripts/links inside `<head>` tag of your HTML file in order to use TailwindCSS, VueJS and Axios.

```html
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/tailwindcss/1.4.0/tailwind.min.css"/>
<script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/vue/2.6.11/vue.min.js"></script>
<script type="text/javascript" src="https://cdn.jsdelivr.net/npm/axios/dist/axios.min.js"></script>
```

## Defining VueJS template
Right after the opening `<body>` tag, paste below template that will render as main HTML of the website. There are some variables that VueJS used to append data fetched from the API.

```html
<div id="app">
  <header class="bg-teal-800 p-6 shadow-lg">
    <div class="container mx-auto">
      <nav class="flex items-center justify-between flex-wrap">
        <div class="flex items-center flex-shrink-0 text-white mr-6">
          <span class="font-semibold text-xl tracking-tight">{{ meta.title }}</span>
        </div>
        <a :href="meta.readMore" class="text-white hover:text-teal-300 py-2 rounded-lg" target="_blank">Read More &#8594;</a>
      </nav>
    </div>
  </header>
  <main>
    <div class="bg-auto" style="background-image: url(https://source.unsplash.com/njUBfL1Oc3Y/1600x900)">
      <div class="bg-teal-800 bg-opacity-75 mx-auto pt-48 pb-8 text-center">
        <p class="text-3xl text-white font-bold">{{meta.introduction}}</p>
        <p class="mx-auto text-white pt-32 pb-8">{{meta.subintro}}</p>
      </div>
      
    </div>
    <div class="container mx-auto">
      <div class="flex flex-wrap -mx-3 py-8">
        <div class="w-1/1 md:w-1/2 lg:w-1/3 px-3 pb-6" v-for="item in items">
          <div class="rounded overflow-hidden shadow-lg">
            <a :href="item.link">
              <img class="w-full object-cover h-64" :src="item.screenshot" :alt="item.title">
            </a>
            <div class="px-6 py-4">
              <div class="font-bold text-xl mb-2">{{ item.title }}</div>
              <p class="text-gray-700 text-base">{{ item.description }}</p>
            </div>
            <div class="px-6 py-4">
              <span class="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2" v-for="tag in getTags(item.tags)">#{{ tag }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
  <footer>
    <div class="container mx-auto pb-8 text-gray-500">{{ meta.footer }}</div>
  </footer>
</div>
```

## Fetching data and binding VueJS
Right before closing `</body>` tag, paste below code to set up VueJS and Axios to make API call and bind data to the template we defined earlier. Remember to replace `YOUR_API_KEY` and `YOUR_SPREADSHEET_ID` with real credentials acquired from earlier steps.

```html
<script>
var app = new Vue({
  el: '#app',
  data: {
    meta: [], // Meta data from Google Sheet
    items: [] // Items data from Google Sheet
  },
  mounted () {
    // Fetch the Meta data from Sheetson API
    axios
      .get('https://api.sheetson.com/v2/sheets/Meta', {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer YOUR_API_KEY',
          'X-Sheetson-Spreadsheet-Id': 'YOUR_SPREADSHEET_ID'
        }
      })
      .then(response => (this.meta = response.data.results[0]))
    
    // Fetch the Items data from Sheetson API
    axios
      .get('https://api.sheetson.com/v2/sheets/PortfolioItems', {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer YOUR_API_KEY',
          'X-Sheetson-Spreadsheet-Id': 'YOUR_SPREADSHEET_ID'
        }
      })
      .then(response => (this.items = response.data.results))
  },
  methods: {
    // Convert comma seperated tags to an array
    getTags: function (tags) {
      return tags.split(',')
    }
  }
})
</script>
```

Save the file and you can open it using Chrome or Firefox.