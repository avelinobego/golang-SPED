package common

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
