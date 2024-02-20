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

package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/avelinobego/sped/esocial/constant"
)

type id struct {
	tipo    constant.Tipo
	numero  int64
	current string
	counter int64
}

func MakeId(tipo constant.Tipo, numero int64) *id {
	return &id{tipo: tipo, numero: numero, counter: 0}
}

func (id *id) Generate(data time.Time, sb *strings.Builder) {

	year, month, day := data.Date()
	hour := data.Hour()
	minute := data.Minute()
	second := data.Second()

	for {
		result := fmt.Sprintf("ID%d%014d%0d%0d%d%d%d%d%05d",
			id.tipo,
			id.numero,
			year,
			month,
			day,
			hour,
			minute,
			second,
			id.counter)

		if result != id.current {
			id.current = result
			sb.WriteString(result)
			return
		}

		id.counter++
	}

}
