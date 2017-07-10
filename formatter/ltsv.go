package formatter

import "bytes"

type Ltsv struct {
}

func (l *Ltsv) Format(targets []FormatTargetInterface) []byte {
	b := bytes.Buffer{}
	for _, target := range targets {
		b.WriteString(target.Key())
		b.WriteString(":")
		b.WriteString(target.Val())
		b.WriteString("\t")
	}
	b.WriteString("\n")
	return b.Bytes()
}

func (l *Ltsv) ConvertToFormatTarget(key, val string) FormatTargetInterface {
	return &FormatTarget{key: key, val: val}
}

type FormatTarget struct {
	key string
	val string
}

func (f *FormatTarget) Key() string {
	return f.key
}

func (f *FormatTarget) Val() string {
	return f.val
}
