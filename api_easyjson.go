// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package compat

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	types "go/types"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC1cedd36DecodeGithubComSmolaGocompat(in *jlexer.Lexer, out *Symbol) {
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
		case "package":
			out.Package = string(in.String())
		case "name":
			out.Name = string(in.String())
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
func easyjsonC1cedd36EncodeGithubComSmolaGocompat(out *jwriter.Writer, in Symbol) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"package\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Package))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Symbol) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeGithubComSmolaGocompat(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Symbol) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeGithubComSmolaGocompat(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Symbol) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeGithubComSmolaGocompat(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Symbol) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeGithubComSmolaGocompat(l, v)
}
func easyjsonC1cedd36DecodeGithubComSmolaGocompat1(in *jlexer.Lexer, out *Signature) {
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
		case "Params":
			if in.IsNull() {
				in.Skip()
				out.Params = nil
			} else {
				in.Delim('[')
				if out.Params == nil {
					if !in.IsDelim(']') {
						out.Params = make([]*Definition, 0, 8)
					} else {
						out.Params = []*Definition{}
					}
				} else {
					out.Params = (out.Params)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Definition
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Definition)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Params = append(out.Params, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Results":
			if in.IsNull() {
				in.Skip()
				out.Results = nil
			} else {
				in.Delim('[')
				if out.Results == nil {
					if !in.IsDelim(']') {
						out.Results = make([]*Definition, 0, 8)
					} else {
						out.Results = []*Definition{}
					}
				} else {
					out.Results = (out.Results)[:0]
				}
				for !in.IsDelim(']') {
					var v2 *Definition
					if in.IsNull() {
						in.Skip()
						v2 = nil
					} else {
						if v2 == nil {
							v2 = new(Definition)
						}
						(*v2).UnmarshalEasyJSON(in)
					}
					out.Results = append(out.Results, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Variadic":
			out.Variadic = bool(in.Bool())
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
func easyjsonC1cedd36EncodeGithubComSmolaGocompat1(out *jwriter.Writer, in Signature) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Params\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Params == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Params {
				if v3 > 0 {
					out.RawByte(',')
				}
				if v4 == nil {
					out.RawString("null")
				} else {
					(*v4).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Results\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Results == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Results {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Variadic\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Variadic))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Signature) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeGithubComSmolaGocompat1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Signature) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeGithubComSmolaGocompat1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Signature) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeGithubComSmolaGocompat1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Signature) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeGithubComSmolaGocompat1(l, v)
}
func easyjsonC1cedd36DecodeGithubComSmolaGocompat2(in *jlexer.Lexer, out *Package) {
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
		case "path":
			out.Path = string(in.String())
		case "objects":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.Objects = make(map[string]*Object)
				} else {
					out.Objects = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v7 *Object
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(Object)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					(out.Objects)[key] = v7
					in.WantComma()
				}
				in.Delim('}')
			}
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
func easyjsonC1cedd36EncodeGithubComSmolaGocompat2(out *jwriter.Writer, in Package) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"path\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Path))
	}
	{
		const prefix string = ",\"objects\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Objects == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v8First := true
			for v8Name, v8Value := range in.Objects {
				if v8First {
					v8First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v8Name))
				out.RawByte(':')
				if v8Value == nil {
					out.RawString("null")
				} else {
					(*v8Value).MarshalEasyJSON(out)
				}
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Package) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeGithubComSmolaGocompat2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Package) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeGithubComSmolaGocompat2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Package) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeGithubComSmolaGocompat2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Package) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeGithubComSmolaGocompat2(l, v)
}
func easyjsonC1cedd36DecodeGithubComSmolaGocompat3(in *jlexer.Lexer, out *Object) {
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
		case "symbol":
			(out.Symbol).UnmarshalEasyJSON(in)
		case "type":
			out.Type = DeclarationType(in.String())
		case "definition":
			if in.IsNull() {
				in.Skip()
				out.Definition = nil
			} else {
				if out.Definition == nil {
					out.Definition = new(Definition)
				}
				(*out.Definition).UnmarshalEasyJSON(in)
			}
		case "methods":
			if in.IsNull() {
				in.Skip()
				out.Methods = nil
			} else {
				in.Delim('[')
				if out.Methods == nil {
					if !in.IsDelim(']') {
						out.Methods = make([]*Func, 0, 8)
					} else {
						out.Methods = []*Func{}
					}
				} else {
					out.Methods = (out.Methods)[:0]
				}
				for !in.IsDelim(']') {
					var v9 *Func
					if in.IsNull() {
						in.Skip()
						v9 = nil
					} else {
						if v9 == nil {
							v9 = new(Func)
						}
						(*v9).UnmarshalEasyJSON(in)
					}
					out.Methods = append(out.Methods, v9)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonC1cedd36EncodeGithubComSmolaGocompat3(out *jwriter.Writer, in Object) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"symbol\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Symbol).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	if in.Definition != nil {
		const prefix string = ",\"definition\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Definition).MarshalEasyJSON(out)
	}
	if len(in.Methods) != 0 {
		const prefix string = ",\"methods\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v10, v11 := range in.Methods {
				if v10 > 0 {
					out.RawByte(',')
				}
				if v11 == nil {
					out.RawString("null")
				} else {
					(*v11).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Object) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeGithubComSmolaGocompat3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Object) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeGithubComSmolaGocompat3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Object) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeGithubComSmolaGocompat3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Object) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeGithubComSmolaGocompat3(l, v)
}
func easyjsonC1cedd36DecodeGithubComSmolaGocompat4(in *jlexer.Lexer, out *Func) {
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
		case "Name":
			out.Name = string(in.String())
		case "Params":
			if in.IsNull() {
				in.Skip()
				out.Params = nil
			} else {
				in.Delim('[')
				if out.Params == nil {
					if !in.IsDelim(']') {
						out.Params = make([]*Definition, 0, 8)
					} else {
						out.Params = []*Definition{}
					}
				} else {
					out.Params = (out.Params)[:0]
				}
				for !in.IsDelim(']') {
					var v12 *Definition
					if in.IsNull() {
						in.Skip()
						v12 = nil
					} else {
						if v12 == nil {
							v12 = new(Definition)
						}
						(*v12).UnmarshalEasyJSON(in)
					}
					out.Params = append(out.Params, v12)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Results":
			if in.IsNull() {
				in.Skip()
				out.Results = nil
			} else {
				in.Delim('[')
				if out.Results == nil {
					if !in.IsDelim(']') {
						out.Results = make([]*Definition, 0, 8)
					} else {
						out.Results = []*Definition{}
					}
				} else {
					out.Results = (out.Results)[:0]
				}
				for !in.IsDelim(']') {
					var v13 *Definition
					if in.IsNull() {
						in.Skip()
						v13 = nil
					} else {
						if v13 == nil {
							v13 = new(Definition)
						}
						(*v13).UnmarshalEasyJSON(in)
					}
					out.Results = append(out.Results, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Variadic":
			out.Variadic = bool(in.Bool())
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
func easyjsonC1cedd36EncodeGithubComSmolaGocompat4(out *jwriter.Writer, in Func) {
	out.RawByte('{')
	first := true
	_ = first
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
		const prefix string = ",\"Params\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Params == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Params {
				if v14 > 0 {
					out.RawByte(',')
				}
				if v15 == nil {
					out.RawString("null")
				} else {
					(*v15).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Results\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Results == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v16, v17 := range in.Results {
				if v16 > 0 {
					out.RawByte(',')
				}
				if v17 == nil {
					out.RawString("null")
				} else {
					(*v17).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Variadic\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Variadic))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Func) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeGithubComSmolaGocompat4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Func) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeGithubComSmolaGocompat4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Func) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeGithubComSmolaGocompat4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Func) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeGithubComSmolaGocompat4(l, v)
}
func easyjsonC1cedd36DecodeGithubComSmolaGocompat5(in *jlexer.Lexer, out *Field) {
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
		case "Name":
			out.Name = string(in.String())
		case "Type":
			if in.IsNull() {
				in.Skip()
				out.Type = nil
			} else {
				if out.Type == nil {
					out.Type = new(Definition)
				}
				(*out.Type).UnmarshalEasyJSON(in)
			}
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
func easyjsonC1cedd36EncodeGithubComSmolaGocompat5(out *jwriter.Writer, in Field) {
	out.RawByte('{')
	first := true
	_ = first
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
		const prefix string = ",\"Type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Type == nil {
			out.RawString("null")
		} else {
			(*in.Type).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Field) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeGithubComSmolaGocompat5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Field) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeGithubComSmolaGocompat5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Field) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeGithubComSmolaGocompat5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Field) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeGithubComSmolaGocompat5(l, v)
}
func easyjsonC1cedd36DecodeGithubComSmolaGocompat6(in *jlexer.Lexer, out *Definition) {
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
		case "type":
			out.Type = Type(in.String())
		case "symbol":
			if in.IsNull() {
				in.Skip()
				out.Symbol = nil
			} else {
				if out.Symbol == nil {
					out.Symbol = new(Symbol)
				}
				(*out.Symbol).UnmarshalEasyJSON(in)
			}
		case "elem":
			if in.IsNull() {
				in.Skip()
				out.Elem = nil
			} else {
				if out.Elem == nil {
					out.Elem = new(Definition)
				}
				(*out.Elem).UnmarshalEasyJSON(in)
			}
		case "key":
			if in.IsNull() {
				in.Skip()
				out.Key = nil
			} else {
				if out.Key == nil {
					out.Key = new(Definition)
				}
				(*out.Key).UnmarshalEasyJSON(in)
			}
		case "len":
			out.Len = int64(in.Int64())
		case "chandir":
			out.ChanDir = types.ChanDir(in.Int())
		case "fields":
			if in.IsNull() {
				in.Skip()
				out.Fields = nil
			} else {
				in.Delim('[')
				if out.Fields == nil {
					if !in.IsDelim(']') {
						out.Fields = make([]*Field, 0, 8)
					} else {
						out.Fields = []*Field{}
					}
				} else {
					out.Fields = (out.Fields)[:0]
				}
				for !in.IsDelim(']') {
					var v18 *Field
					if in.IsNull() {
						in.Skip()
						v18 = nil
					} else {
						if v18 == nil {
							v18 = new(Field)
						}
						(*v18).UnmarshalEasyJSON(in)
					}
					out.Fields = append(out.Fields, v18)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "functions":
			if in.IsNull() {
				in.Skip()
				out.Functions = nil
			} else {
				in.Delim('[')
				if out.Functions == nil {
					if !in.IsDelim(']') {
						out.Functions = make([]*Func, 0, 8)
					} else {
						out.Functions = []*Func{}
					}
				} else {
					out.Functions = (out.Functions)[:0]
				}
				for !in.IsDelim(']') {
					var v19 *Func
					if in.IsNull() {
						in.Skip()
						v19 = nil
					} else {
						if v19 == nil {
							v19 = new(Func)
						}
						(*v19).UnmarshalEasyJSON(in)
					}
					out.Functions = append(out.Functions, v19)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "signature":
			if in.IsNull() {
				in.Skip()
				out.Signature = nil
			} else {
				if out.Signature == nil {
					out.Signature = new(Signature)
				}
				(*out.Signature).UnmarshalEasyJSON(in)
			}
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
func easyjsonC1cedd36EncodeGithubComSmolaGocompat6(out *jwriter.Writer, in Definition) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	if in.Symbol != nil {
		const prefix string = ",\"symbol\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Symbol).MarshalEasyJSON(out)
	}
	if in.Elem != nil {
		const prefix string = ",\"elem\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Elem).MarshalEasyJSON(out)
	}
	if in.Key != nil {
		const prefix string = ",\"key\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Key).MarshalEasyJSON(out)
	}
	if in.Len != 0 {
		const prefix string = ",\"len\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Len))
	}
	if in.ChanDir != 0 {
		const prefix string = ",\"chandir\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.ChanDir))
	}
	if len(in.Fields) != 0 {
		const prefix string = ",\"fields\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v20, v21 := range in.Fields {
				if v20 > 0 {
					out.RawByte(',')
				}
				if v21 == nil {
					out.RawString("null")
				} else {
					(*v21).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.Functions) != 0 {
		const prefix string = ",\"functions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v22, v23 := range in.Functions {
				if v22 > 0 {
					out.RawByte(',')
				}
				if v23 == nil {
					out.RawString("null")
				} else {
					(*v23).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Signature != nil {
		const prefix string = ",\"signature\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Signature).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Definition) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeGithubComSmolaGocompat6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Definition) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeGithubComSmolaGocompat6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Definition) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeGithubComSmolaGocompat6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Definition) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeGithubComSmolaGocompat6(l, v)
}
func easyjsonC1cedd36DecodeGithubComSmolaGocompat7(in *jlexer.Lexer, out *API) {
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
		case "Packages":
			if in.IsNull() {
				in.Skip()
				out.Packages = nil
			} else {
				in.Delim('[')
				if out.Packages == nil {
					if !in.IsDelim(']') {
						out.Packages = make([]*Package, 0, 8)
					} else {
						out.Packages = []*Package{}
					}
				} else {
					out.Packages = (out.Packages)[:0]
				}
				for !in.IsDelim(']') {
					var v24 *Package
					if in.IsNull() {
						in.Skip()
						v24 = nil
					} else {
						if v24 == nil {
							v24 = new(Package)
						}
						(*v24).UnmarshalEasyJSON(in)
					}
					out.Packages = append(out.Packages, v24)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Reachable":
			if in.IsNull() {
				in.Skip()
				out.Reachable = nil
			} else {
				in.Delim('[')
				if out.Reachable == nil {
					if !in.IsDelim(']') {
						out.Reachable = make([]*Object, 0, 8)
					} else {
						out.Reachable = []*Object{}
					}
				} else {
					out.Reachable = (out.Reachable)[:0]
				}
				for !in.IsDelim(']') {
					var v25 *Object
					if in.IsNull() {
						in.Skip()
						v25 = nil
					} else {
						if v25 == nil {
							v25 = new(Object)
						}
						(*v25).UnmarshalEasyJSON(in)
					}
					out.Reachable = append(out.Reachable, v25)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonC1cedd36EncodeGithubComSmolaGocompat7(out *jwriter.Writer, in API) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Packages\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Packages == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v26, v27 := range in.Packages {
				if v26 > 0 {
					out.RawByte(',')
				}
				if v27 == nil {
					out.RawString("null")
				} else {
					(*v27).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Reachable\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Reachable == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v28, v29 := range in.Reachable {
				if v28 > 0 {
					out.RawByte(',')
				}
				if v29 == nil {
					out.RawString("null")
				} else {
					(*v29).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v API) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC1cedd36EncodeGithubComSmolaGocompat7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v API) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC1cedd36EncodeGithubComSmolaGocompat7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *API) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC1cedd36DecodeGithubComSmolaGocompat7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *API) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC1cedd36DecodeGithubComSmolaGocompat7(l, v)
}
