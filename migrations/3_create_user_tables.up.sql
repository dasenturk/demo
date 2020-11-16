CREATE TABLE petowners (
    id UUID PRIMARY KEY,
    ownername TEXT NOT NULL,
    surname TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    cityname TEXT NOT NULL
);

CREATE TABLE pets(
    id UUID PRIMARY KEY,
    petname TEXT NOT NULL,
    pettype TEXT NOT NULL,
    breed_type TEXT NOT NULL,
    owner_id UUID NOT NULL REFERENCES petowners(id) ON DELETE CASCADE,
    gender TEXT NOT NULL,
    weight INT NOT NULL,
    age INT NOT NULL
);

