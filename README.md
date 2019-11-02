# Thanos-API
An API that returns quotes from the super-villain Thanos who wants to wipe out half of all life in the universe.
<br>
<br>
This is built using GoLang and it was my first ever GoLang project, it seems to be quite a versatile language and it seems to handle a variety of tasks quite well.
<h2>How to run</h2>
Simply download this repo and then in the command line within the folder execute 

```
run thanos.go
```

<h2>Endpoints</h2>
<h4>For a random quote to be generated each time</h4>

http://localhost:8000/api/thanosapi/v1

<h4>To recieve all of the quotes at once</h4>

http://localhost/8000/api/thanosapi/v1/all

<h2>Changing variable names</h2>
<br>
By default the port is set to 8000 on localhost but can be easily changed on line 79.
<br>

```go
http.ListenAndServe(":8000", nil)
```
<h2>Output</h2>
The output for the API takes the form of 

```
All that for a drop of blood.
```

This is arguably simplier than using it by parsing JSON. This is also because of the nature of the API being only one line phrases and not requiring any more fields such as the time of retrieval or who is the author of the quote. This also allows for easy parsing due to it's heavily simplistic output.
<br><br>

#### Donate
Buy me a coffee to donate if you like
<br>
https://www.buymeacoffee.com/uiSK0Ex

<a href="https://www.buymeacoffee.com/uiSK0Ex"><img src="https://static-2.gumroad.com/res/gumroad/9026696959709/asset_previews/09c9bf14407c2a76d088f22121d0b0a9/retina/Screen_20Shot_202017-10-20_20at_2010.09.59.jpg" alt="alt text" width="251.25" height="183.75"></a>
