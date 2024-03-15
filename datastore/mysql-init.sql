-- mysql -h 172.24.1.200 --user=root --password
-- mysql -h 172.24.1.200 --user=ssouser --password

DROP USER IF EXISTS 's2user'@'172.24.1.1';
CREATE USER IF NOT EXISTS 's2user'@'172.24.1.200' IDENTIFIED BY '1234';
-- CREATE USER IF NOT EXISTS 's2user'@'172.24.1.1' IDENTIFIED WITH mysql_native_password BY '1234';

DROP DATABASE IF EXISTS webintelligence;
CREATE DATABASE IF NOT EXISTS webintelligence;
GRANT ALL PRIVILEGES ON webintelligence.* TO 's2user'@'172.24.1.200';
FLUSH PRIVILEGES;

REVOKE ALL PRIVILEGES FROM 's2user'@'172.24.1.1';

-- CREATE USER IF NOT EXISTS 'user'@'192.168.1.1' IDENTIFIED BY 'password';
-- GRANT ALL PRIVILEGES ON table.* TO 'user'@'192.168.1.1' WITH GRANT OPTION;