## Synopsis

Pigeon is a microservice for sending templated emails via a REST api. It runs on AWS and utilizes SES.

## Code Example

Show what the library does as concisely as possible, developers should be able to figure out **how** your project solves their problem by looking at the code example. Make sure the API you are showing off is obvious, and that your code is short and concise.

## Motivation

Pigeon allows users to swap html templates without changing backend code. It allows other services to send a templated, automated email with a simple REST query.

## Installation

```
git clone https://github.com/RichardH92/Pigeon.git
docker build -t Pigeon Pigeon
docker run -p 8081:8081 -v <your cloned directory>:/go --name=Pigeon Pigeon
```

## API Reference

Pigeon's API is dynamically created from the configuration settings.  

For example,i for the following configuration settings
```json
{
	"Port" : ":8081",
	"Source_Email" : "test@test.com",
	"Region" : "us-east-1",
	"Emails" : [
		{
			"Type_Name" : "Example_Email_1",
			"HTML_File_Name" : "example_email_1.html",
			"Query_Vals" : ["aaa", "bbb"],
			"Subject" : "Example_Subject_1"
		},
		{
			"Type_Name" : "Example_Email_2",
			"HTML_File_Name" : "example_email_2.html",
			"Subject" : "Example_Subject_2"
		}

	]
}
```

Pigeon will generate the following API endpoints  
```
Test
```


## Tests

Describe and show how to run the tests with code examples.

## Contributors

Let people know how they can dive into the project, include important links to things like issue trackers, irc, twitter accounts if applicable.

## License

MIT
