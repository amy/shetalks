
**IMPORTANT:**
Please read through this before providing feedback / comments. These are recognized areas that need to be iterated upon. The first iteration was written in one day between my classes, so THIS IS A VERY ROUGH SKELETON.

**OVERVIEW:** 
This is meant to be a sample API. Currently, it only has one resource: events.
The database sets up an events, speaker, and organization table along with pertinent pivot tables (speaker_event, organization_event, organization_speaker). For available endpoints, see routes package. 

**NOTES:**
* The handlers accept a service interface so that DB access is abstracted
* Routers accept service interfaces configured to an DB
* Handlers are typecasted HandlerFunc to avoid defining a Handler struct type with a DB field.

**NEXT STEPS:**
* Combine routes + handlers packages into a http package. Think more about packages and what functions are exportable. Maybe just put everything into the shetalks package? 
* Abstract out handler logic (Ex: encode / decode -> just attach whatever encoder / decoder you want as well as header format) 
* API Versioning
* Docker compose out the server
* Revisit whether I actually want a pointer or a value passed along as either the receiver or in parameters
* Issue during Create, where if speaker doesn't exist, you can still assign the uncreated speaker to the event. Need to look at atomic transactions. 
* Address everything below

**RECOGNIZED AREAS OF IMPROVEMENT:**

* Error Handling - many of the errors are redundant / don't handle the issue at the source

* Responses - many of the header error responses aren't correct/convention currently just defaulting to StatusUnprocessableEntity until I figure out which codes I actually want to use

* Logging - should create/incorporate a logging pkg (need to research what deserves to be logged)

* Queries - many redundant/inefficient queries. Many terrible cases, where if one query succeeds & another fails, the first query is not rolled back. Need to look into atomic transactions

* Testing - need to create mocks & more tests. 

```
shetalks
	cmd
		shetalks			// Router executable 
			main.go
    handlers
    	eventHandlers.go
        eventHandlers_test.go
    mock
    	mock.go
	mysql					// MySQL implementation
		eventService.go
	routes					// Register router
		router
	shetalks.go 			// Domain logic
```
