# go



go install github.com/beego/bee/v2@latest

bee -h
2022/08/10 00:56:00.804 [D]  init global config instance failed. If you do not use this, just ignore it.  open conf/app.conf: no such file or directory
2022/08/10 00:56:00 INFO     ▶ 0001 Getting bee latest version...
2022/08/10 00:56:01 INFO     ▶ 0002 Your bee are up to date
Bee is a Fast and Flexible tool for managing your Beego Web Application.

You are using bee for beego v2.x. If you are working on beego v1.x, please downgrade version to bee v1.12.0

USAGE
    bee command [arguments]

AVAILABLE COMMANDS

    version     Prints the current Bee version
    migrate     Runs database migrations
    api         Creates a Beego API application
    bale        Transforms non-Go files to Go source files
    fix         Fixes your application by making it compatible with newer versions of Beego
    pro         Source code generator
    dev         Commands which used to help to develop beego and bee
    dlv         Start a debugging session using Delve
    dockerize   Generates a Dockerfile for your Beego application
    generate    Source code generator
    hprose      Creates an RPC application based on Hprose and Beego frameworks
    new         Creates a Beego application
    pack        Compresses a Beego application into a single file
    rs          Run customized scripts
    run         Run the application by starting a local development server
    server      serving static content over HTTP on port
    update      Update Bee

Use bee help [command] for more information about a command.

ADDITIONAL HELP TOPICS


Use bee help [topic] for more information about that topic.



bee api myapp
cd myapp
go get myapp
bee run


----

go get -u github.com/beego/bee
go get -u github.com/astaxie/beego


bee api testapi
cd testapi
go get testapi
bee run -downdoc=true -gendoc=true
bee run


bee generate model sample_mst -fields="name:int"


----------------------


go install github.com/beego/bee/v2@latest
bee api testapi
cd testapi
go get testapi
bee run

bee generate scaffold comment -fields="content:string" -driver=mysql -conn="root:password@tcp(mysql:3306)/test_db"


bee api sample -driver=mysql -conn=root:password@tcp"(mysql:3306)"/test_db

----------------------

root@1a38f9c1e98f:/go/src# bee help api
2022/08/10 02:38:58.134 [D]  init global config instance failed. If you do not use this, just ignore it.  open conf/app.conf: no such file or directory
USAGE
  bee api [appname]

OPTIONS
  -beego
      set beego version,only take effect by go mod

  -conn
      Connection string used by the driver to connect to a database instance.

  -driver
      Database driver. Either mysql, postgres or sqlite.

  -gopath
      Support go path,default false

  -tables
      List of table names separated by a comma.

DESCRIPTION
  The command 'api' creates a Beego API application.
  now default supoort generate a go modules project.

  Example:
      $ bee api [appname] [-tables=""] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]  [-gopath=false] [-beego=v1.12.3]

  If 'conn' argument is empty, the command will generate an example API application. Otherwise the command
  will connect to your database and generate models based on the existing tables.

  The command 'api' creates a folder named [appname] with the following structure:

            ├── main.go
            ├── go.mod
            ├── conf
            │     └── app.conf
            ├── controllers
            │     └── object.go
            │     └── user.go
            ├── routers
            │     └── router.go
            ├── tests
            │     └── default_test.go
            └── models
                  └── object.go
                  └── user.go

--------------------

bee api testapi -driver=mysql -conn=root:password@tcp"(mysql:3306)"/test_db

--------------------
root@1a38f9c1e98f:/go/src/testapi# bee help dockerize
USAGE
  bee dockerize

OPTIONS
  -expose=8080
      Port(s) to expose in the Docker container.

  -image=library/golang
      Set the base image of the Docker container.

DESCRIPTION
  Dockerize generates a Dockerfile for your Beego Web Application.
  The Dockerfile will compile, get the dependencies with godep, and set the entrypoint.

  Example:
    $ bee dockerize -expose="3000,80,25"

bee generate scaffold comment -fields="content:string"
--------------------
root@1a38f9c1e98f:/go/src/testapi# bee help generate
USAGE
  bee generate [command]

OPTIONS
  -conn
      Connection string used by the SQLDriver to connect to a database instance.

  -ctrlDir
      Controller directory. Bee scans this directory and its sub directory to generate routers

  -ddl
      Generate DDL Migration

  -driver
      Database SQLDriver. Either mysql, postgres or sqlite.

  -fields
      List of table Fields.

  -level
      Either 1, 2 or 3. i.e. 1=models; 2=models and controllers; 3=models, controllers and routers.

  -routersFile
      Routers file. If not found, Bee create a new one. Bee will truncates this file and output routers info into this file

  -routersPkg
      router's package. Default is routers, it means that "package routers" in the generated file

  -tables
      List of table names separated by a comma.

DESCRIPTION
  ▶ To scaffold out your entire application:

     $ bee generate scaffold [scaffoldname] [-fields="title:string,body:text"] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"]

  ▶ To generate a Model based on fields:

     $ bee generate model [modelname] [-fields="name:type"]

  ▶ To generate a controller:

     $ bee generate controller [controllerfile]

  ▶ To generate a CRUD view:

     $ bee generate view [viewpath]

  ▶ To generate a migration file for making database schema updates:

     $ bee generate migration [migrationfile] [-fields="name:type"]

  ▶ To generate swagger doc file:

     $ bee generate docs

    ▶ To generate swagger doc file:

     $ bee generate routers [-ctrlDir=/path/to/controller/directory] [-routersFile=/path/to/routers/file.go] [-routersPkg=myPackage]

  ▶ To generate a test case:

     $ bee generate test [routerfile]

  ▶ To generate appcode based on an existing database:

     $ bee generate appcode [-tables=""] [-driver=mysql] [-conn="root:@tcp(127.0.0.1:3306)/test"] [-level=3]

--------------------


bee generate scaffold comment -fields="content:string" -driver=mysql -conn="root:password@tcp(mysql:3306)/test_db"


--------------------
go mod init beego
go install github.com/beego/bee/v2@latest
go get -u github.com/go-sql-driver/mysql
go get -u github.com/jmoiron/sqlx


            go get "github.com/go-sql-driver/mysql"
            go get "github.com/beego/bee"
            go get "github.com/astaxie/beego"
            go get "github.com/astaxie/beego/session"
            go get "golang.org/x/crypto/bcrypt"
            go get "github.com/PuerkitoBio/goquery"
            go get "github.com/microcosm-cc/bluemonday"


bee api sample -driver=mysql -conn=root:password@tcp"(mysql:3306)"/test_db
cd sample
go get sample


bee generate scaffold post -fields="title:string,body:text" -driver=mysql


--------------------
go mod init beego
go install github.com/beego/bee/v2@latest
go get -u github.com/go-sql-driver/mysql
go get -u github.com/jmoiron/sqlx

bee api sample -driver=mysql -conn="root:password@tcp(mysql:3306)/test_db"
cd sample
go get sample
bee generate scaffold post -fields="title:string,body:text" -driver=mysql  -conn="root:password@tcp(db_mysql)/test_db"


bee migrate -driver=mysql -conn="root:password@tcp(db_mysql)/test_db"



--------------------
go mod init beego
go install github.com/beego/bee/v2@latest
go get -u github.com/go-sql-driver/mysql
go get -u github.com/jmoiron/sqlx

bee api sample -driver=mysql -conn="root:password@tcp(mysql:3306)/test_db"

bee generate scaffold post -fields="title:string,body:text"
go get sample/database/migrations
go mod tidy

bee migrate -driver=mysql -conn="root:password@tcp(mysql:3306)/test_db"


--------------------
go mod init beego
go install github.com/beego/bee/v2@latest
go get -u github.com/go-sql-driver/mysql
go get -u github.com/jmoiron/sqlx

bee api sample -driver=mysql -conn="root:password@tcp(mysql:3306)/test_db"
cd sample
go mod tidy

bee generate scaffold post -fields="title:string,body:text"
go get sample/database/migrations
go mod tidy

bee migrate -driver=mysql -conn="root:password@tcp(mysql:3306)/test_db"


--------------------
go mod init beego
go install github.com/beego/bee/v2@latest
go get -u github.com/go-sql-driver/mysql
go get -u github.com/jmoiron/sqlx

bee api sample -driver=mysql -conn="root:password@tcp(mysql:3306)/test_db"
cd sample
go mod tidy

bee generate scaffold post -fields="title:string,body:text" -driver=mysql -conn="root:password@tcp(mysql:3306)/test_db"

bee generate scaffold post2 -fields="title:string,body:text" -driver=mysql -conn="root:password@tcp(mysql:3306)/test_db"



