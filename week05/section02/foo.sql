PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE customers(id INTEGER PRIMARY KEY, name TEXT);
INSERT INTO customers VALUES(1,'Widgets Inc');
INSERT INTO customers VALUES(2,'foobar inc');
COMMIT;

