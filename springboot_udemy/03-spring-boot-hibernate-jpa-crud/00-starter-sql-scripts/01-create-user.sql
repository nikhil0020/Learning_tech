DROP USER if exists 'springstudent'@'%' ;

CREATE USER 'springstudent'@'%' IDENTIFIED BY 'IbIbO@9938';

GRANT ALL PRIVILEGES ON * . * TO 'springstudent'@'%';