#!/bin/bash
####################################
#
# This is a bash file to start my demo application.
#
####################################
# current user privlage username if exists
username="root"
# default password if exists
password="123456"

# for dumyuser test there is default
dumyuser="dumyuser"
# for dumyuser test there is default
dumypass="dumypassword"
# this is for apikey configuration
apikey="AIzaSyBzuM7atg360ClN4hmao7J3Y0UbvxSrkx8"

qtation="'"
RED='\033[0;31m' # Red Color
NC='\033[0m' # No Color
BLUE='\033[0;34m' # Blue Color
# make a directory name config
mkdir config

# creating config file and its default contents

echo "appname: testGolang

server :
    address : :8080

db:
    address:     localhost
    user:        ${dumyuser}
    dbname :     orders
    password:   ${dumypass}
    port:       3306

logpath: /tmp/log

apikey: ${apikey}
"   > config/config.yaml


echo "CREATE DATABASE orders ;
INSERT INTO mysql.user (User,Host,authentication_string,ssl_cipher,x509_issuer,x509_subject) VALUES(${qtation}${dumyuser}${qtation},'localhost',PASSWORD(${qtation}${dumypass}${qtation}),'','','') ;
FLUSH PRIVILEGES;
GRANT ALL PRIVILEGES ON orders.* to ${qtation}${dumyuser}${qtation}@localhost;
FLUSH PRIVILEGES;
use orders ; CREATE TABLE orderinfo (iOrderId int(11) NOT NULL AUTO_INCREMENT,iDistance int(11) DEFAULT NULL, vStatus varchar(45) , dtOrder datetime DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (iOrderId)) ENGINE=InnoDB AUTO_INCREMENT=0;" > _tempfilescript.sql

# Checking there is mysql or not
mysqlpkg=$(dpkg -l | grep mysql | wc -l)

install_mysql () {
    echo "Install of Mysql just begun"
    date
    echo

    # Install MySQL
    # Install the MySQL server by using the Ubuntu package manager:
    sudo apt-get update
    sudo apt-get -y install mysql-server
    # Allow remote access
    # Run the following command to allow remote access to the mysql server:
    sudo ufw allow mysql
    # Start the MySQL service
    # After the installation is complete, you can start the database service by running the following command. 
    # If the service is already started, a message informs you that the service is already running:
    systemctl start mysql
    # Launch at reboot
    # To ensure that the database server launches after a reboot, run the following command:
    systemctl enable mysql
    # This is running the script and create the schema
    printf "${RED}Install of Mysql just finished${NC}\n"
    date
    echo

    run_script
}

run_script () {
    echo "Preparing mysql just begun"
    date
    echo

    # Set the root password
    mysql -u $username -p$password -s < _tempfilescript.sql
    printf "${BLUE}Database Updateded${NC}\n!"
    date
    echo
}

runapplication (){
    printf "${RED}Esecute the application , ready to go${NC}\n"
    date
    echo
    nohup ./mydemoapp&
    exit 0 
}

# Check the status of installation mysql on Server
# if needed it will be installed
# otherwise it will execute the script
if [ $mysqlpkg -eq 0 ] ; then
    install_mysql
else
    run_script
fi
rm _tempfilescript.sql
runapplication
