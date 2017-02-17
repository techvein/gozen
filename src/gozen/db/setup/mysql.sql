
CREATE USER 'gozen'@'localhost' IDENTIFIED BY 'rootpass';
GRANT ALL PRIVILEGES ON gozen.* TO 'gozen'@'localhost' IDENTIFIED BY 'rootpass';

DROP DATABASE IF EXISTS gozen;
CREATE DATABASE gozen;