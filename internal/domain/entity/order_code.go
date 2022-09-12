package entity

import (
	"strconv"
	"strings"
	"time"
)

const (
	MaxSequenceLength = 8
	SequencePadding   = "0"
)

type OrderCode struct {
	value string
}

// NewOrderCode returns a new OrderCode.
// The order code is represented by YYYYPPPPPPPP where the YYYY is the year and the PPPPPPPP is the order sequence.
func NewOrderCode(date time.Time, sequence int) OrderCode {
	return OrderCode{value: generateCode(date, sequence)}
}

func (o OrderCode) Value() string {
	return o.value
}

func generateCode(date time.Time, sequence int) string {
	yearStr := strconv.Itoa(date.Year())
	sequenceStr := strconv.Itoa(sequence)
	sequencePaddingLen := MaxSequenceLength - len(sequenceStr)
	builder := strings.Builder{}
	builder.WriteString(yearStr)
	builder.WriteString(strings.Repeat(SequencePadding, sequencePaddingLen))
	builder.WriteString(sequenceStr)
	return builder.String()
}
