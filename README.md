<h1>Simple Go http webserver</h1>

To run it with docker cmd.
1. 'docker build -t test-server .'
2. 'docker run --rm -p 8080:8080 test-server'

There are 3 endpoints in httpwebserver.go, including "/", "/list", "/add"
And below is how to query the endpoints with `curl`:

1. For "/" with func homePage
=> $curl http://localhost:8080/

2. For "/list" with func listAll to list all the objects
=> $curl http://localhost:8080/list

like  `[{"timestamp": "2019-12-02T06:53:32Z", "key": "a", "value": "some value"},{"timestamp": "2019-12-02T06:53:35Z", "key": "asdf", "value": "some other value"}]`

3. For "/add" with func addOne to add one object to dummy database
=> $curl -X POST -H "Content-Type: application/json" -d "{ \"key\": \"some key\",  \"value\": \"some value\"}" http://localhost:8080/add

e.g.  `{"key": "asdf", "value": "abcd"}`
=> $curl -X POST -H "Content-Type: application/json" -d "{ \"key\": \"asdf\",  \"value\": \"abcd\"}" http://localhost:8080/add
