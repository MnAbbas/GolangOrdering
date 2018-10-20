# GolangOrdering
this is a sample of restful apis for ordering system

I did my best to keep it simple , clean and structured 

this project is built based on mvc models 

# Technology
I used [Golang](https://golang.org)  , [gorilla](http://www.gorillatoolkit.org) , [mysql](https://www.mysql.com) in this project

# Settup
``` ./bash.sh ``` 
will install mysql if needed otherwise it will create use , database and one table
```
default user is :=dumyuser
default pass is :=dumypassword
```
# Existinig mysql
if an istance of mysql is exist , please in bash.sh provide one user & password 
to execute the sql scripts
otherwise it will handeled by the bash file

# Pre-requirment
this project has its default value such as 
dbuser information , logfile path , server address listner 
but in case of changing anyof them 
you must have a folder named config
with a config.yaml in it according to source 
you could be able to change it accordingly

# Test
I create more than 20 test cases which will be tested automaticly
just needed to go to root of project and call
``` go test test/*.go ```
it will print the result of tests

# Run
for setup and run project just need these files
- bash.sh ---- this is for setup and run project
- mydemoapp --- this is the binary for ubuntu 18 TLS(amd64), using GOOS=linux GOARCH=amd64 parameter to build the binary
>based on your defination 
- mysqlscript.sql -- this is mysql script

:+1: its ready to go! :shipit:
