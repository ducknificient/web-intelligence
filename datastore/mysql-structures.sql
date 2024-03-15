SET GLOBAL time_zone = '+8:00';
SET GLOBAL time_zone = 'Asia/Jakarta';

-- uuid v1 source 
-- https://stackoverflow.com/questions/43056220/store-uuid-v4-in-mysql
-- SELECT BIN_TO_UUID(id) FROM sso.users;

-- FOR PRODUCTION RUMAHWEB 
-- id BINARY(16) NOT NULL DEFAULT (UUID()),

CREATE TABLE IF NOT EXISTS webintelligence.crawlpage(
    id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID())),
    pagesource longtext,
    link longtext,
    task longtext, 
    flag VARCHAR(50) DEFAULT 'INSERT',
    createdon TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    createdby VARCHAR(50),
    updatedon TIMESTAMP DEFAULT NULL,
    updatedby VARCHAR(50),
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS webintelligence.crawlhref(
    id BINARY(16) NOT NULL DEFAULT (UUID_TO_BIN(UUID())),
    link longtext,
    href longtext,
    task longtext,
    flag VARCHAR(50) DEFAULT 'INSERT',
    createdon TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    createdby VARCHAR(50),
    updatedon TIMESTAMP DEFAULT NULL,
    updatedby VARCHAR(50),
    PRIMARY KEY(id)
);
