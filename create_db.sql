-- create_db.sql

CREATE DATABASE api_cotacao;

\c api_cotacao

CREATE TABLE IF NOT EXISTS trades (
    id SERIAL PRIMARY KEY,
    hora_fechamento VARCHAR(500),
    data_negocio DATE,
    codigo_instrumento VARCHAR(300),
    preco_negocio DECIMAL(10, 2),
    quantidade_negociada INT  
);
