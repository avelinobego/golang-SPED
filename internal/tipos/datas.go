package tipos

import (
	"encoding/xml"
	"time"
)

type DataNasc time.Time

func (d DataNasc) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(time.Time(d).Format(time.DateOnly), start)
}

func MakeDataNasc(year uint32, month uint32, day uint64) DataNasc {
	return DataNasc(time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.Local))
}
