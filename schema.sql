CREATE TABLE IF NOT EXISTS gebruiker (
    klant_id serial PRIMARY KEY,
    gebruikersnaam VARCHAR (50) UNIQUE NOT NULL,
    wachtwoord VARCHAR (50) NOT NULL,
    email VARCHAR (255) UNIQUE NOT NULL,
    created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS opdracht (
    vraag_id serial PRIMARY KEY,
    vaknaam VARCHAR(20) UNIQUE NOT NULL,
    lessnaam VARCHAR(20) NOT NULL,
    tijd_besteed INT,
    klant_id INT NOT NULL
    );

CREATE TABLE IF NOT EXISTS studentenpas (
    klant_id INT,
    pas_id serial PRIMARY KEY
    );

CREATE TABLE IF NOT EXISTS cursus (
    leraar_id INT,
    cursusnaam VARCHAR(20) NOT NULL,
    cursus_id serial PRIMARY KEY,
    klant_id INT
    );

CREATE TABLE IF NOT EXISTS deskundige (
    leraar_id serial PRIMARY KEY,
    naam VARCHAR(20) NOT NULL,
    e_mail VARCHAR(20) UNIQUE NOT NULL
    );