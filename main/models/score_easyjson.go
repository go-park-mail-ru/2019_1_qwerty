// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson2c1ef15cDecode20191QwertyMainModels(in *jlexer.Lexer, out *Score) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Place":
			out.Place = uint64(in.Uint64())
		case "Name":
			out.Name = string(in.String())
		case "Score":
			out.Points = uint64(in.Uint64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2c1ef15cEncode20191QwertyMainModels(out *jwriter.Writer, in Score) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Place\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.Place))
	}
	{
		const prefix string = ",\"Name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Score\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.Points))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Score) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2c1ef15cEncode20191QwertyMainModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Score) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2c1ef15cEncode20191QwertyMainModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Score) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2c1ef15cDecode20191QwertyMainModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Score) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2c1ef15cDecode20191QwertyMainModels(l, v)
}
