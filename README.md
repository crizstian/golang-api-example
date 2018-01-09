## App Structure

In easycast golang we follow the MVC app structure specification, we organize our files around specifications:

*Note: the order of the files may be different but the structure must be the same.*

example:

```
├── api // api's only handle requests and responses
|   ├── auth.go
|   ├── error.go
|   ├── index.go
|   └── log_analytics.go
├── ctrls // the controlers handles all the logic and some cases db interactions
|   ├── analytics
|   |   └── log_analytics.go
|   ├── auth
|   |   ├── activate.go
|   |   ├── login.go
|   |   ├── recover.go
|   |   └── register.go
|   ├── user
|   └── test // local test files
├── conf // here we grab al the static values, like credentials, getting env variables, etc.
├── fn // the reason for separating fn in a folder is to avoid circuluar imports
├── public // here goes static and public files to be render
|   └── accountid.go
├── libs // to separte fn based on libraries
├── models // every model created represents a table in mysql db, so db interactions goes here
|   ├── account.go
|   ├── automigrate.go
|   ├── user_info.go // example: this model contains the user profile information
|   └── user.go // this model containse the email, password, token, etc
├── routes // defining routes by group
|   ├── account.go
|   ├── analytics.go
|   ├── auth.go
|   ├── index.go
|   └── user.go
├── test // global test files
├── main.go
└── main_test.go
```

In order to understand this example i suggest to start opening the files in the following order

`main.go -> routes (index.go file first) -> api -> ctrls -> models`

this is one of the most common flow of interactions

## Run development

To run the code we execute the following command:

`go run main.go`

we need to have golang installed localy or using a docker container:

but for the docker container we need some more steps, but we have a script to do that

`sh start-container.sh`, will create the docker image, and run the docker container, with

some extra configurations needed.


**Note:** to be able to link the api with a db, we need to have a db up and running, at localhost or have a docker container with a db up and running and linked to the golang container, the code is already prepared for both scenarios, we linked the golang container to a container called db, you are going to see the configurations at the `db/gorm.go` file, and uncomment the `startDB() fn` in the `main.go` file. There is a file called `models/automigrate.go` which generates automatically the tables on the database, just be sure to have a database created and the automigrate will handle the rest.


The api is ready to run the test at `src/ctrls/test` it will run some golang functions, interacting with a mongodb cluster in the cloud.


There is 2 `dockerfiles` one for development and the `Dockerfile.multi` for production env, for the multi file it needs the latest version of docker to be able to create a production ready image.

### Install this project on the next folder:

$GOPATH/src/easycast

----------
## Testing

For testing, i'll suggest to see the ginkgo and gomega libraries, its very similar to mocha and (chai or should) from nodejs.

to run the test you must be located in the path where the `test_suite_test.go` file is located, and run the `ginkgo -v` command.

*if you want to test a particular function try*

```
ginkgo -cover -v --focus=Bandwidth
```

**Note**: i suggest to run the `src/ctrls/test/log_analytics_test.go` file to see how it works, just be located at the path of the file and run the command: `ginkgo -v`.

----------

## Define API URLS

http://jsonapi.org/recommendations/


----------

## Error Handling

To handle errors in the project, we must first detect the type and level of the functions where the error occurred, the levels are the following:

- 1st Level Fn: this functions are the ones that handle the echo.Context or functions that depends from another function, so that this functions can handle the error, like sending back a message to the user.

e.g.
```
// for api fn

// in case of error
return sendError("level", msg, err)
```


- 2nd Level Fn: here are functions that are not dependent, so this functions must return the error, to the dependent functions to be able to handle the error.

e.g.

```
// functions must follow the following pattern
func superFunc(...params) (type of return, error) {}

// in case of error
return [valid value], err

// if err not exist
return [value], nil
```

### API Response pattern

For endpoint responses, we must follow the next pattern
ß
```
return c.JSON(http.StatusOK, echo.Map{
  [value | values]: ...values,
  "msg": "some message",
  "success": true | false
})
```

## Documentation

To document the api we use RAML, and we follow the RAML specifications, and to generate
the html we use a nodejs library called `raml2html` there is a sh file called `bin/generate-doc.sh`
that executes the command to generate the documentation html file.
