# Thanos-API
An API that returns quotes from the super-villain Thanos who wants to wipe out half of all life in the universe.
<br>
<br>
This is build using GoLang and it was my first ever GoLang project, it seems to be quite a versatile language and it seems to handle a variety of tasks quite well.
<h3>How to run</h3>
Simply download this repo and then in the command line within the folder execute 

```
run thanos.go
```

<h3>Endpoints</h3>
<h5>for a random quote to be generated each time</h5>

http://localhost:8000/api/thanosapi/v1

<h5>To recieve all of the quotes at once</h5>

http://localhost/8000/api/thanosapi/v1/all

<h3>Changing variable names</h3>
<br>
By default the port is set to 8000 on localhost but can be easily changed on line 79.
<br>

```go
http.ListenAndServe(":8000", nil)
```
<h4>Output</h4>
The output for the API takes the form of 

```
All that for a drop of blood.
```

This is arguably simplier than using it by parsing JSON. This is also because of the nature of the API being only one line phrases and not requiring any more fields such as the time of retrieval or who is the author of the quote. This also allows for easy parsing due to it's heavily simplistic output.
