UPDATE mysql.user SET authentication_string = PASSWORD('123456') WHERE User ='root' ;
FLUSH PRIVILEGES;
CREATE DATABASE orders ;
INSERT INTO mysql.user (User,Host,authentication_string,ssl_cipher,x509_issuer,x509_subject) VALUES('dumyuser','localhost',PASSWORD('dumypassword'),'','','') ;
FLUSH PRIVILEGES;
GRANT ALL PRIVILEGES ON orders.* to 'dumyuser'@localhost;
FLUSH PRIVILEGES;
use orders ; CREATE TABLE orderinfo (iOrderId int(11) NOT NULL AUTO_INCREMENT,iDistance int(11) DEFAULT NULL, vStatus varchar(45) , dtOrder datetime DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (iOrderId)) ENGINE=InnoDB AUTO_INCREMENT=0;