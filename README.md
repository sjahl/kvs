# KVS

This is a simple REST-ful in-memory key/value store. I wrote this as a means to learn a little bit about how golang can respond to HTTP requests, and how you can parse and generate JSON data. I don't really expect it to be useful to anyone.

## Endpoints

`/set`

Given a JSON data structure with a "key" and a "value" via a POST, add the k/v pair to the store with the "key" as the key, and the "value" as the value. Keys must be strings, but values can be anything that JSON supports

```
curl --data '{"key": "foobar", "value": ["one", "two", "three"]}' -XPOST localhost:8000/set
```

`/get`

Given a JSON data structure with a "key" via a GET, try to retrieve that key from the k/v store and return the bare value

```
curl --data '{"key": "foobar"}' -XGET localhost:8000/get
```


`/all`

A GET request with no data will return a JSON blob of the entire k/v store

```
curl -XGET localhost:8000/all
```

`/drop`

A POST request with a JSON data structure containing a "key" will try to remove the item in the k/v store that has that key

```
curl --data '{"key": "baz"}' -XPOST localhost:8000/drop
```
