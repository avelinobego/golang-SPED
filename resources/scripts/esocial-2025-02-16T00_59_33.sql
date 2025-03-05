
CREATE TABLE empregador
(
  id            serial4      NOT NULL GENERATED ALWAYS AS IDENTITY,
  identificador numeric(50) ,
  nome          varchar(255) NOT NULL,
  razao         varchar(255),
  PRIMARY KEY (id)
);

COMMENT ON TABLE empregador IS 'detalhes do empregador';

CREATE TABLE eventos
(
  id        serial4      NOT NULL GENERATED ALWAYS AS IDENTITY,
  nome      varchar(255) NOT NULL,
  descricao varchar(255),
  PRIMARY KEY (id)
);

COMMENT ON TABLE eventos IS 'lista dos eventos';
