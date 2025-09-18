PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE widgets(id integer primary key, name text);
INSERT INTO widgets VALUES(1,'bar');
COMMIT;

