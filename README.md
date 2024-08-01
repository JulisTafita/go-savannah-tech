
# Golang Take-Home Assignment: GitHub API Data Fetching and Service

The goal of this project is to build a service in Golang that fetches data from GitHub's public APIs to retrieve repository information and commits. The service will save this data in a persistent store ( MySql database ) and continuously monitor the repository for changes. It also provides mechanisms for querying specific data from the database.


## Getting started

To run go-savannah-tech on local :

1- Install golang v1.22 from source:

🔗 https://go.dev/doc/install

2- Install latest MySql version

🔗 https://dev.mysql.com/downloads/installer/

3- Clone go-savannah-tech repository

```bash
  git clone https://github.com/JulisTafita/go-savannah-tech.git

```

4-  Install go-savannah-tech dependancies

```bash
  cd go-savannah-tech
  go mod tidy
```

5- Create and restore savannah_tech database

You can find the full schema and ready to use mysql script in ``/database``.
It will create a database called ``savannah_tech`` and two table ``i_repository`` and ``i_commit``.

6- Ensure that config.toml is correctly configured with your settings.

7- Run

To run the application : ``go build`` then ``./go-savannahTech`` or ``go run main.go``

8- Run test

To run test : ``go test ./...``


## configuration file

You can find the configuration file ( ``config.toml`` ) in the root directory.
It is splitted into four parts:

A- 🛢️ database

Here, you can setup your database connection:

``user_name`` : the name of the database user, ``root`` is set by default

``user_password`` : the password of the database user.

``name`` : the database name, by default it's ``savannah_tech``

``host`` : the host of the database. ``localhost`` by default

``port`` : the port of the database, ``3306`` by default

B- 👨🏻‍💻 GitHub

Here, you can configure how we interact with GitHub API.

``api_endpoint`` : the base url for the github api.

``user_token`` : the github user token, this field is not required if you want to deal with public repositories. But if you want to get private repositories informations you have to fill this field.
You can generate your token here: https://github.com/settings/tokens .

``repository_name`` : this field is required.

C - ⚙️ Option

Option is the Configuration for data pulling options.

``pulling_cron_job_spec`` : The cron schedule expression for periodically pulling data.
The base schema is ``@every 01h00m00s``.

Example cron expression: "@every 00h00m10s" will execute the job every 10 seconds.

``pulling_start_date`` : The commit date on which we will start pulling in ISO 8601 format.
example : 2023-09-14 12:11:00 .Field not required, leave with empty string if not needed.

``pulling_end_date`` :   The commit date on which we will end pulling in ISO 8601 format.
example : 2023-09-14 12:11:00 .Field not required, leave with empty string if not needed.

``reset_collection`` : boolean, true if you want to remove all commit records from the database before pulling.

D - 🚀 Server

This is a web server with two endpoints:

1- ``GET /commit-author`` : Get the top N commit authors by commit counts from the  database.

* Query Parameters:

  `number_of_author`: Number of top authors to return.


* How it works ?

To get the top N commit authors by commit counts from the database, query the desired repository first, then all its commit by a `JOIN` operation.
Then group commit by `author_login` and `COUNT()` the total rows number of this author.

This is an  example query that get the 5 first authors :
````
  SELECT
		    ic.author_name as 'name',
		    ic.author_email as 'email',
		    ic.author_login as 'login',
		    count(ic.id) as 'commit_count' 
		FROM i_repository ir
			JOIN i_commit ic On ir.id = ic.i_repository_id
		WHERE ir.name = ? group by ic.author_login ORDER BY count(ic.id) DESC LIMIT 5
````


* example of use : ``http://localhost:8089/commit-author?number_of_author=1``



2- ``GET /repository-commit`` : Retrieve commits of a repository by repository name from the database

* Query Parameters:

  `repository`: The name of the repository to fetch commits from.

* example : ``http://localhost:8089/repository-commit?repository=abeg-rebuild``



In the config.toml:

`run_web_server` : boolean, true if you want to run the web server.

`web_server_host` :  The hostname or IP address on which the web server will listen

`web_server_port` : The port number on which the web server will listen



