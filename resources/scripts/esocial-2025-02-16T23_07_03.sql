
CREATE TABLE empregador
(
  id            serial4     ,
  identificador numeric(14) ,
  nome          varchar(255) NOT NULL,
  razao         varchar(255),
  PRIMARY KEY (id)
);

COMMENT ON TABLE empregador IS 'detalhes do empregador';

CREATE TABLE evento_empregador_trabalhador
(
  id             serial ,
  id_empregador  numeric,
  id_trabalhador numeric,
  PRIMARY KEY (id)
);

COMMENT ON TABLE evento_empregador_trabalhador IS 'relação entre evento, empregador e trabalhador';

CREATE TABLE eventos
(
  id   serial4     ,
  nome varchar(255) NOT NULL,
  PRIMARY KEY (id)
);

COMMENT ON TABLE eventos IS 'lista dos eventos';

CREATE TABLE trabalhador
(
  id            serial4    ,
  identificador numeric(11),
  PRIMARY KEY (id)
);

COMMENT ON TABLE trabalhador IS 'dados do trabalhador';
