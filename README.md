# GOLANG REST API

This use case parses the data from the XML files under /api/files and uses an ORM (object relational modeling) called Gorum to save that data into a postgres database.

## Development workflow

The stepwise workflow is as follows:
- The XML files are parsed into objects.
- A database is created using Gorum.
- Then, those XML objects are parsed into another objects which represent the tables in a database, using Gorum.
- Once the database and its tables are setup and filled with data, the api exposes its functions.
- These functions (GET,POST,...) are accessible for other services to read the XML data.
- The data is read using POSTMAN for testing purposes.

## Deployment

We will run this API locally using docker-compose. For this you need to download docker and run the following command.

```
docker-compose up --build
```

This command runs the container and it is now availble for POSTMAN or other services to access.

## Test

We use POSTMAN to access the API. The steps should be:
- Create a collection.
- Create GET requests to read the data from the exposed functions.
  - The URL to use for a request to read all clients is:
    `http://localhost:5000/clients`
  - The URL to use for a request to read all merged data is:
    `http://localhost:5000/merged/116730191637`
    Where 116730191637 is the timeLineId.

Make these requests and see the results. Good Luck :)
