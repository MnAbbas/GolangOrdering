# GolangOrdering
this is a sample of restful apis for ordering system

I did my best to keep it simple , clean and structured 

this project is built based on mvc models 

# Technology & Packages
I used [Golang](https://golang.org)  , [gorilla](http://www.gorillatoolkit.org) , [mysql](https://www.mysql.com) in this project

# Settup
``` bash
./bash.sh 
``` 
will install mysql if needed otherwise it will create use , database and one table
```
default user is :=dumyuser
default pass is :=dumypassword
```
# Existinig mysql
if an istance of mysql is exist , please in bash.sh provide one use data in bash.sh top of the file 
` username='YOUR_MYSQL_USERNAME' & password='YOUR_MYSQL_USERNAME'  `
to execute the sql scripts
### otherwise it will handeled by the bash file

# Google API Key
this project need a Google API key which has permission to call matrix api
for putting your own Google API Key you shoud go to config folder
```bash
 vim config/config.yaml
 ```
under default.json or production.json change 
```javascript
apikey: YOUR_API_KEY
```
or 
>you could change the value of this is bash.sh file also it will do all required tasks

# Pre-requirment
this project has its default value such as 
`dbuser` `information` , `logfile path` , `server address` `listner`  `google api key`
but in case of changing anyof them 
you must have a folder named config
with a config.yaml in it according to source 
you could be able to change it accordingly
>permission on bash.sh is important you must grant execution access to bash.sh
```bash
chmod +x bash.sh
```

# Test
I create more than 20 test cases which will be tested automaticly
just needed to go to root of project and call
``` golang
go test -v test/*.go 
```
it will print the result of tests

# Build
the binary for `ubuntu 18 TLS(amd64)`, using `GOOS=linux GOARCH=amd64` 
### build command:
``` golang 
GOOS=linux GOARCH=amd64 go build -o mydemoapp
```
parameter to build the binary
>based on your defination 

# Run
for setup and run project just need these files
- bash.sh ---- this is for setup and run project
- mydemoapp --- execatable application

:+1: its ready to go! :shipit:
