
CREATE TABLE empregador
(
  id            serial       NOT NULL,
  identificador numeric(14) ,
  nome          varchar(255) NOT NULL,
  razao         varchar(255),
  PRIMARY KEY (id)
);

COMMENT ON TABLE empregador IS 'detalhes do empregador';

CREATE TABLE evento_empregador_trabalhador
(
  id             serial,
  id_empregador  serial NOT NULL,
  id_trabalhador serial NOT NULL,
  PRIMARY KEY (id)
);

COMMENT ON TABLE evento_empregador_trabalhador IS 'relação entre evento, empregador e trabalhador';

CREATE TABLE eventos
(
  id   serial      ,
  nome varchar(255) NOT NULL,
  PRIMARY KEY (id)
);

COMMENT ON TABLE eventos IS 'lista dos eventos';

CREATE TABLE trabalhador
(
  id            serial     ,
  identificador numeric(11),
  PRIMARY KEY (id)
);

COMMENT ON TABLE trabalhador IS 'dados do trabalhador';

ALTER TABLE evento_empregador_trabalhador
  ADD CONSTRAINT FK_empregador_TO_evento_empregador_trabalhador
    FOREIGN KEY (id_empregador)
    REFERENCES empregador (id);

ALTER TABLE evento_empregador_trabalhador
  ADD CONSTRAINT FK_trabalhador_TO_evento_empregador_trabalhador
    FOREIGN KEY (id_trabalhador)
    REFERENCES trabalhador (id);
