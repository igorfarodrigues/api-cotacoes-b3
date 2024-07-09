-- create_db.sql

CREATE DATABASE api_cotacao;

\c api_cotacao

BEGIN;

-- Verificação de existência e criação das tabelas
DROP TABLE IF EXISTS trades;
DROP TABLE IF EXISTS tickers;

CREATE TABLE tickers (
    id SERIAL PRIMARY KEY,
    codigo_instrumento VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE trades (
    id SERIAL PRIMARY KEY,
    hora_fechamento VARCHAR(100),
    data_negocio DATE,
    codigo_instrumento VARCHAR(100) NOT NULL,
    preco_negocio DECIMAL(10, 2),
    quantidade_negociada INT,
    FOREIGN KEY (codigo_instrumento) REFERENCES tickers (codigo_instrumento)
);

-- Criação de índices
DROP INDEX IF EXISTS idx_ticker_date;
CREATE INDEX idx_ticker_date ON trades (codigo_instrumento, data_negocio);

COMMIT;