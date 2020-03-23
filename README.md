<h1>Simple Go http webserver</h1>

To reproduce it with docker cmd.
1. 'docker build -t test-server .'
2. 'docker run --rm -p 8080:8080 test-server'

<h1>How to use?</h1>
<p>There are 3 endpoints in httpwebserver.go, including "/", "/list", "/add"
and below is how to query the endpoints with `curl`: </p>

<ul>
  <li> "/", with func homePage
    => $curl http://localhost:8080/</li>

<li> GET "/list" json response, with func listAll to list all the objects with it's timestamp, key and value<br>
=> $curl http://localhost:8080/list

e.g.  `[{"timestamp": "2019-12-02T06:53:32Z", "key": "a", "value": "some value"},{"timestamp": "2019-12-02T06:53:35Z", "key": "asdf", "value": "some other value"}]`</li>

<li> POST "/add" json payload, with func addOne to add one object to dummy database<br>
=> $curl -X POST -H "Content-Type: application/json" -d "{ \"key\": \"some key\",  \"value\": \"some value\"}" http://localhost:8080/add
  
e.g.  `{"key": "asdf", "value": "abcd"}`<br>
=> $curl -X POST -H "Content-Type: application/json" -d "{ \"key\": \"asdf\",  \"value\": \"abcd\"}" http://localhost:8080/add
</li></ul>
