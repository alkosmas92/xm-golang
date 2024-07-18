CREATE TABLE IF NOT EXISTS users (
                                     user_id TEXT PRIMARY KEY,
                                     username TEXT UNIQUE,
                                     password TEXT,
                                     firstname TEXT,
                                     lastname TEXT
);
CREATE TABLE IF NOT EXISTS company (
                                       CompanyID TEXT PRIMARY KEY,
                                       name TEXT NOT NULL UNIQUE,
                                       description TEXT,
                                       amountofemployees INTEGER NOT NULL,
                                       registered BOOLEAN NOT NULL,
                                       type TEXT NOT NULL
);