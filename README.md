API for an event, speaker, and organization relationship.

Currently, only event is built out.

Notes:

Used this article as design foundation: 
https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1#.hgk4poe7y

Used this for some Golang API basics: 
http://thenewstack.io/make-a-restful-json-api-go/

Used this for idiomatic/safe Go DB transactions:
https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.2.html

Recognized areas of improvement:

Error Handling 	- many of the errors are redundant / don't handle the issue at the source

Responses		- many of the header error responses aren't correct/convention currently just
                  defaulting to StatusUnprocessableEntity until I figure out which codes I 
                  actually want to use

Logging 		- should create/incorporate a logging pkg (need to research what deserves 
				  to be logged)

Queries 		- many redundant/inefficient queries. Many terrible cases, where if one 
                  query succeeds & another fails, the first query is not rolled back

Testing			- need to create mocks, tests are non-existent rn

```
shetalks
	cmd
		shetalks			// Router executable 
			main.go
	mysql					// MySQL implementation
		eventService.go
	routes					// Register router
		router
	shetalks.go 			// Domain logic
```

> The router can take on any service configured to any database
> The handlers can take on service interfaces configured to any database 
