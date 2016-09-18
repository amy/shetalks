
**IMPORTANT**
Please read through this before providing feedback / comments. These are recognized 
areas that need to be iterated upon. The first iteration was written in one day between
my classes, so THIS IS A VERY ROUGH SKELETON.

Overview: 
This is meant to be a sample API. Currently, it only has one resource, events.
The database sets up an events, speaker, and organization table along with pertinent
pivot tables (speaker_event, organization_event, organization_speaker). 

Design:
> The handlers accepts a service interface so that DB access is abstracted
> Routers accept service interfaces 
> Handlers are closures

Next Steps:
> Combine routes + handlers packages into a http package. Think more about packages and what functions are exportable. Maybe just put everything into the shetalks package? 
> Abstract out handler logic (Ex: encode / decode -> just attach whatever encoder / decoder you want as well as header format), Include a context. 
> API Versioning
> Docker compose out the server
> Revisit whether I actually want a pointer or a value passed along as either the receiver or in parameters
> Issue during Create, where if speaker doesn't exist, you can still assign the uncreated speaker to the event
> Address everything below

Recognized areas of improvement:

Error Handling 	- many of the errors are redundant / don't handle the issue at the source

Responses		- many of the header error responses aren't correct/convention currently just
                  defaulting to StatusUnprocessableEntity until I figure out which codes I 
                  actually want to use

Logging 		- should create/incorporate a logging pkg (need to research what deserves 
				  to be logged)

Queries 		- many redundant/inefficient queries. Many terrible cases, where if one 
                  query succeeds & another fails, the first query is not rolled back. Need 
                  to look into atomic transactions

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

Used this article as design foundation: 
https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1#.hgk4poe7y

Used this for some Golang API basics: 
http://thenewstack.io/make-a-restful-json-api-go/

Used this for idiomatic/safe Go DB transactions:
https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.2.html
