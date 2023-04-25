package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"sort"
	"strings"
)

//自定义formatter
type Formatter struct {
	PrintKey bool
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	b := &bytes.Buffer{}
	b.WriteString(entry.Time.Format("[2006/01/02 15:04:05.000]"))

	level := strings.ToUpper(entry.Level.String())

	b.WriteString(" [" + level[:1] + "]")

	if len(entry.Data) != 0 {
		fields := make([]string, 0, len(entry.Data))
		for field := range entry.Data {
			fields = append(fields, field)
		}

		sort.Strings(fields)

		for _, field := range fields {
			if f.PrintKey {
				fmt.Fprintf(b, " [%s:%v]", field, entry.Data[field])
			} else {
				fmt.Fprintf(b, " [%v]", entry.Data[field])
			}
		}
	}

	if entry.HasCaller() {
		fmt.Fprintf(b, " [%s:%d]", entry.Caller.File, entry.Caller.Line)
	}

	b.WriteString(" " + entry.Message)

	b.WriteByte('\n')

	return b.Bytes(), nil
}
