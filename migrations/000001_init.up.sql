CREATE TABLE file
(
    uuid UUID         NOT NULL PRIMARY KEY UNIQUE,
    name VARCHAR(255) NOT NULL,
    size INT          NOT NULL
);

CREATE TABLE file_statistic
(
    file_uuid     UUID REFERENCES file (uuid) NOT NULL UNIQUE,
    request_count INT                         NOT NULL DEFAULT 0
);
