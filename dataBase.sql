CREATE SCHEMA IF NOT EXISTS `desafio_backend`;
USE `desafio_backend`;

DROP TABLE IF EXISTS consultas;
DROP TABLE IF EXISTS dentistas;
DROP TABLE IF EXISTS pacientes;

CREATE TABLE `dentistas` (
`id`            INT NOT NULL AUTO_INCREMENT,
`nome`          VARCHAR(25) NOT NULL,
`sobrenome`     VARCHAR(50) NOT NULL,
`matricula`     VARCHAR(10) NOT NULL UNIQUE,
    PRIMARY KEY (id)
) ENGINE = innodb;

CREATE TABLE `pacientes` (
`id`            INT NOT NULL AUTO_INCREMENT,
`nome`          VARCHAR(25) NOT NULL,
`sobrenome`     VARCHAR(50) NOT NULL,
`rg`            VARCHAR(10) NOT NULL UNIQUE,
`dataCadastro`  DATETIME NOT NULL,
    PRIMARY KEY (id)
) ENGINE = innodb;

CREATE TABLE `consultas` (
`id`            INT NOT NULL AUTO_INCREMENT,
`descricao`     VARCHAR(250) NOT NULL,
`dataConsulta`  DATETIME NOT NULL,
`dentistaId`      INT NOT NULL,
`pacienteId`      INT NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT fk_dentistas
        FOREIGN KEY (dentistaId)
        REFERENCES dentistas(id),
    CONSTRAINT fk_pacientes
        FOREIGN KEY (pacienteId)
        REFERENCES pacientes(id)
) ENGINE = innodb;