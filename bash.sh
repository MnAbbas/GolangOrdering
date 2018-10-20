#!/bin/bash
####################################
#
# This is a bash file to start my demo application.
#
####################################
# current user privlage username if exists
username="root"
# default password if exists
password=""

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
    execsql="UPDATE mysql.user SET authentication_string = PASSWORD('123456') WHERE User ='root'" 
    mysql -u $username -p$password -s < mysqlscript.sql
    echo "User root Updated!"
    date
    echo
}

runapplication (){
    echo "Esecute the application , go everythere with Golang "
    date
    echo
    ./mydemoapp
}

# Check the status of installation mysql on Server
# if needed it will be installed
# otherwise it will execute the script
if [ $mysqlpkg -eq 0 ] ; then
    install_mysql
else
    run_script
fi
runapplication
