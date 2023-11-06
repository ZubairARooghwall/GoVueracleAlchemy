CREATE TABLE USERS (
    USERID INTEGER PRIMARY KEY,
    USERNAME VARCHAR2(30) UNIQUE,
    EMAIL VARCHAR2(255),
    PASSWORD VARCHAR2(64),
    SALT RAW(64),
    USERROLE VARCHAR2(10),
    EDUCATION VARCHAR2(15),
    CREATIONTIME TIMESTAMP
);

CREATE TABLE PROFILES (
    PROFILEID VARCHAR2(255) PRIMARY KEY,
    USERID CONSTRAINT (USERID) REFERENCES USERS(USERID),
    PROFILEPICTURE VARCHAR2(255),
    STATUS BOOLEAN,
    BIOGRAPHY VARCHAR2(1000)
);

CREATE TABLE USERSESSIONS (
    SESSIONID VARCHAR2(255) PRIMARY KEY,
    USERID CONSTRAINT (USERID) REFERENCES USERS(USERID),
    SESSIONTOKEN VARCHAR2(64),
    EXPIRYDATE DATE,
    USERIPADDRESS VARCHAR2(15),
    BROWSER VARCHAR2(20)
);

CREATE TABLE FILES (
    FILEID VARCHAR2(500),
    FILENAME VARCHAR(255),
    UNIQUENAME VARCHAR(500) TYPE VARCHAR2(10),
    FILESIZE INTEGER,
    CREATIONTIME TIMESTAMP,
    OWNER CONSTRAINT (USERID) REFERENCES USERS(USERID),
    FOLDER CONSTRAINT (FOLDERID) REFERENCES FOLDERS(FOLDERID),
    TAG CONSTRAINT (TAGID) REFERENCES TAG(TAGID),
);

CREATE TABLE PERMISSIONS (
    PERMISSIONID VARCHAR2(255) PRIMARY KEY,
    FILEID CONSTRAINT (FILEID) REFERENCES FILES(FILESID),
    USERID CONSTRAINT (USERID) REFERENCES USERS(USERID),
    PERMISSIONTYPE CHAR(1),
    CreationTime TIMESTAMP
);

CREATE TABLE FOLDERS (
    FOLDERID VARCHAR2(500) PRIMARY KEY,
    USERID CONSTRAINT (USERID) REFERENCES USERS(USERID),
    FOLDERNAME VARCHAR2(255) UNIQUE,
    CREATIONDATE TIMESTAMP,
    FOLDERSIZE INTEGER
);

CREATE TABLE ACTIVITYLOG (
    ACTIVITYID INTEGER PRIMARY KEY,
    ACTIVITY VARCHAR2(10),
    CREATIONTIME TIMESTAMP,
    USERID CONSTRAINT (USERID) REFERENCES USERS(USERID),
    FILEID CONSTRAINT (FILEID) REFERENCES FILES(FILEID),
);

CREATE TABLE TAGS (
    TAGNAME VARCHAR2(255) PRIMARY KEY,
    CREATIONTIME TIMESTAMP,
    COLOR VARCHAR2(7)
);

CREATE TABLE MEETINGS (
    MEETINGID VARCHAR2(255) PRIMARY KEY,
    TITLE VARCHAR2(255),
    DESCRIPTION VARCHAR2(500),
    ORGANIZERID CONSTRAINT (USERID) REFERENCES USERS(USERID),
    STARTTIME TIMESTAMP,
    ENDTIME TIMESTAMP
    URL VARCHAR2(1000) UNIQUE
);