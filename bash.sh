#!/bin/bash
####################################
#
# Backup to NFS mount script.
#
####################################
# default username if exists
username="root"
# default password if exists
password="mydumypass"
# crete dumy user for demo
dumyuser="dumyuser"
# crete dumy password for demo
dumypassword="dumypassword"

# Checking there is mysql or not
mysqlpkg=$(dpkg-query -W mysql | wc  -l)


install_mysql () {
    echo "Install of Mysql just begun"
    date
    echo

    # Install MySQL
    # Install the MySQL server by using the Ubuntu package manager:
    sudo apt-get update
    sudo apt-get install mysql-server
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
    echo "Install of Mysql just finished"
    date
    echo

    run_script
}

run_script () {
    echo "Preparing mysql just begun"
    date
    echo

    # Set the root password
    execsql=$(printf "UPDATE mysql.user SET authentication_string = PASSWORD(%s) WHERE User =%s" $password $username)
    mysql -u root -p --execute=$execsql
    echo "User root Updated!"
    date
    echo

    # Create orders database    
    mysql -u root -p$password --execute="CREATE DATABASE orders "
    echo "Creating Database!!"
    date
    echo

    # Add a database user named demouser with pass demopassword
    execsql=$(printf "INSERT INTO mysql.user (User,Host,authentication_string,ssl_cipher,x509_issuer,x509_subject) VALUES(%s,'localhost',PASSWORD(%s),'','','')" $username $password )
    mysql -u root -p$password --execute=$execsql
    echo "create user!!!"
    date
    echo

    # Grant database user permissions
    execsql=$(printf "GRANT ALL PRIVILEGES ON orders.* to %s@localhost;" $username)
    mysql -u root -p$password --execute=$execsql
    echo "Grant Privilages!"
    date
    echo

    # to apply changes on privilges
    mysql -u root -p$password --execute="FLUSH PRIVILEGES;"
    echo "Grant Privilages!"
    date
    echo

    # Create orderinfo table
    mysql -u root -p$password --execute="use orders ; CREATE TABLE orderinfo (iOrderId int(11) NOT NULL AUTO_INCREMENT,iDistance int(11) DEFAULT NULL, vStatus varchar(45) , dtOrder datetime DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (iOrderId)) ENGINE=InnoDB AUTO_INCREMENT=0;"
    echo "Everything DOne! , ready to go "
    date
    echo
}

runapplication (){
    echo "Esecute the application , go everythere with Golang "
    date
    echo
}

# Check the status of installation mysql on Server
# if needed it will be installed
# otherwise it will execute the script
if [ $mysqlpkg -eq 0 ] ; then
    install_mysql
else
    run_script
fi
