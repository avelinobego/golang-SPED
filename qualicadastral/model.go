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

package qualicadastral

import (
	"fmt"
	"math"
	"strings"
	"time"
)

type Lote struct {
	Cpf       int64
	Nis       int64
	Nome      string
	Dn        time.Time
	Uf        string
	Municipio string
	NomeMae   string
}

func (lote Lote) String() string {
	result := make([]string, 0)

	if lote.Cpf != 0 {
		result = append(result, fmt.Sprintf("%011d", lote.Cpf))
	}

	if lote.Nis != 0 {
		result = append(result, fmt.Sprintf("%011d", lote.Nis))
	}

	if lote.Nome != "" {
		result = append(result, ajuste(lote.Nome, 70))
	}

	if !lote.Dn.IsZero() {
		result = append(result, lote.Dn.Format("02012006"))
	}

	if lote.Uf != "" {
		result = append(result, ajuste(lote.Uf, 2))
	}

	if lote.Municipio != "" {
		result = append(result, ajuste(lote.Municipio, 30))
	}

	if lote.NomeMae != "" {
		result = append(result, ajuste(lote.NomeMae, 60))
	}

	return strings.Join(result, ";")
}

func ajuste(valor string, maximo int64) string {
	tamanho := int64(math.Min(float64(len(valor)), float64(maximo)))
	return valor[0:tamanho]
}
