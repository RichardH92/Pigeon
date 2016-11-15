## Synopsis

Pigeon is a microservice for sending template emails via a REST api. It runs on AWS and utilizes SES.


## Motivation

Pigeon allows users to swap html templates without changing backend code. It allows other services to send a template, automated email with a simple REST query.

## Installation

```
git clone https://github.com/RichardH92/Pigeon.git
docker build -t Pigeon Pigeon
docker run -p 8081:8081 -v <your cloned directory>:/go --name=Pigeon Pigeon
```

## API Reference

**Email**
----  

* **URL**

  /email

* **Method:**

  `GET`

*  **URL Params**

   **Required:**

   `type=[string]`  
	 `dest=[string]`

   **Optional:**

   Optional parameters are specified in the "Query_Vals" section(s) of the conf.json file

* **Success Response:**

  * **Code:** 200 <br />


* **Error Response:**

  * **Code:** 401 UNAUTHORIZED <br />
    **Content:** `"Invalid value for query parameter <parameter>"`

  OR

  * **Code:** 501 INTERNAL SERVER ERROR <br />
    **Content:** `"Error opening file <name>"`


* **Sample Queries For Configuration Settings:**  

	* Sample Queries <br />

		`wget "localhost:8081/email?type=Example_Email_1&aaa=123&bbb=456&dest=dest@test.com"`  
		`wget "localhost:8081/email?type=Example_Email_2&dest=dest@test.com"`

	* Config Settings


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



* **Notes:**

	* Pigeon also exports Prometheus metrics on the "/metrics" endpoint


## Tests

`go test -v`

## Contributors

To contribute, please submit a pull-request. Thanks!

## License

MIT
