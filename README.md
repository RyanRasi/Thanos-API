# Thanos-API
An API that returns quotes from the super-villain Thanos who wants to wipe out half of all life in the universe.
<br><br>
<h3>How to run</h3>
Simply download this repo and then in the command line execute 
<br>

```
run thanos.go
```

<br>
Go to <http://localhost/8000/api/thanosapi/v1>
<br>
For a random quote to be generated each time.
<br>
Alternatively you can go to http://localhost/8000/api/thanosapi/v1/all
<br>
To recieve all of the quotes at once.
<br>
<h3>Changing variable names</h3>
<br>
By default the port is set to 8000 on localhost but can be easily changed on line 79.
<br>

```go
http.ListenAndServe(":8000", nil)
```

<br>

The URL
