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


bee api testapi
cd testapi
go get testapi
bee run -downdoc=true -gendoc=true
bee run


bee generate model sample_mst -fields="name:int"


----------------------



