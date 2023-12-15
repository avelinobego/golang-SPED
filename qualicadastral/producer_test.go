// Copyright (C) 2023 Avelino Bego
//
// This file is part of SPED.
//
// SPED is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 2 of the License, or
// (at your option) any later version.
//
// SPED is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with SPED.  If not, see <http://www.gnu.org/licenses/>.

package qualicadastral_test

import (
	"testing"
	"time"

	qualicadastral "github.com/avelinobego/sped/qualificacao"
)

func TestFile(t *testing.T) {
	lotes := []qualicadastral.Lote{}

	for i := 0; i < 10; i++ {
		lotes = append(lotes, qualicadastral.Lote{
			Cpf:       16944301890 + int64(i),
			Nome:      "Avelino de Almeida Bego",
			Nis:       int64(i + 1),
			Dn:        time.Now(),
			Uf:        "SPPPPP",
			NomeMae:   "Maria Aparecida de Almeida Navas",
			Municipio: "Viamão",
		})
	}
	qualicadastral.MakeFile("saida.txt", lotes)
}
