// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: accountcreateoperation.go

package common

// MarshalJSON marshal bytes to json - template
//func (j *SignMessageOperation) MarshalJSON() ([]byte, error) {
//	var buf fflib.Buffer
//	if j == nil {
//		buf.WriteString("null")
//		return buf.Bytes(), nil
//	}
//	err := j.MarshalJSONBuf(&buf)
//	if err != nil {
//		return nil, err
//	}
//	return buf.Bytes(), nil
//}

// MarshalJSONBuf marshal buff to json - template
//func (j *SignMessageOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
//	if j == nil {
//		buf.WriteString("null")
//		return nil
//	}
//	var err error
//	var obj []byte
//	_ = obj
//	_ = err
//	buf.WriteString(`{ "registrar":`)
//
//	{
//
//		obj, err = j.Registrar.MarshalJSON()
//		if err != nil {
//			return err
//		}
//		buf.Write(obj)
//
//	}
//	buf.WriteString(`,"referrer":`)
//
//	{
//
//		obj, err = j.Referrer.MarshalJSON()
//		if err != nil {
//			return err
//		}
//		buf.Write(obj)
//
//	}
//	buf.WriteString(`,"referrer_percent":`)
//	fflib.FormatBits2(buf, uint64(j.ReferrerPercent), 10, false)
//	/* Struct fall back. type=types.Authority kind=struct */
//	buf.WriteString(`,"owner":`)
//	err = buf.Encode(&j.Owner)
//	if err != nil {
//		return err
//	}
//	/* Struct fall back. type=types.Authority kind=struct */
//	buf.WriteString(`,"active":`)
//	err = buf.Encode(&j.Active)
//	if err != nil {
//		return err
//	}
//	buf.WriteString(`,"name":`)
//	fflib.WriteJsonString(buf, string(j.Name))
//	/* Struct fall back. type=types.AccountCreateExtensions kind=struct */
//	buf.WriteString(`,"extensions":`)
//	err = buf.Encode(&j.Extensions)
//	if err != nil {
//		return err
//	}
//	/* Struct fall back. type=types.AccountOptions kind=struct */
//	buf.WriteString(`,"options":`)
//	err = buf.Encode(&j.Options)
//	if err != nil {
//		return err
//	}
//	buf.WriteByte(',')
//	if j.Fee != nil {
//		if true {
//			/* Struct fall back. type=types.AssetAmount kind=struct */
//			buf.WriteString(`"fee":`)
//			err = buf.Encode(j.Fee)
//			if err != nil {
//				return err
//			}
//			buf.WriteByte(',')
//		}
//	}
//	buf.Rewind(1)
//	buf.WriteByte('}')
//	return nil
//}

const (
	ffjtSignMessageOperationbase = iota
	ffjtSignMessageOperationnosuchkey

	ffjtSignMessageOperationRegistrar

	ffjtSignMessageOperationReferrer

	ffjtSignMessageOperationReferrerPercent

	ffjtSignMessageOperationOwner

	ffjtSignMessageOperationActive

	ffjtSignMessageOperationName

	ffjtSignMessageOperationExtensions

	ffjtSignMessageOperationOptions

	ffjtSignMessageOperationFee
)

var ffjKeySignMessageOperationRegistrar = []byte("registrar")

var ffjKeySignMessageOperationReferrer = []byte("referrer")

var ffjKeySignMessageOperationReferrerPercent = []byte("referrer_percent")

var ffjKeySignMessageOperationOwner = []byte("owner")

var ffjKeySignMessageOperationActive = []byte("active")

var ffjKeySignMessageOperationName = []byte("name")

var ffjKeySignMessageOperationExtensions = []byte("extensions")

var ffjKeySignMessageOperationOptions = []byte("options")

var ffjKeySignMessageOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
//func (j *SignMessageOperation) UnmarshalJSON(input []byte) error {
//	fs := fflib.NewFFLexer(input)
//	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
//}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
//func (j *SignMessageOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
//	var err error
//	currentKey := ffjtSignMessageOperationbase
//	_ = currentKey
//	tok := fflib.FFTok_init
//	wantedTok := fflib.FFTok_init
//
//mainparse:
//	for {
//		tok = fs.Scan()
//		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
//		if tok == fflib.FFTok_error {
//			goto tokerror
//		}
//
//		switch state {
//
//		case fflib.FFParse_map_start:
//			if tok != fflib.FFTok_left_bracket {
//				wantedTok = fflib.FFTok_left_bracket
//				goto wrongtokenerror
//			}
//			state = fflib.FFParse_want_key
//			continue
//
//		case fflib.FFParse_after_value:
//			if tok == fflib.FFTok_comma {
//				state = fflib.FFParse_want_key
//			} else if tok == fflib.FFTok_right_bracket {
//				goto done
//			} else {
//				wantedTok = fflib.FFTok_comma
//				goto wrongtokenerror
//			}
//
//		case fflib.FFParse_want_key:
//			// json {} ended. goto exit. woo.
//			if tok == fflib.FFTok_right_bracket {
//				goto done
//			}
//			if tok != fflib.FFTok_string {
//				wantedTok = fflib.FFTok_string
//				goto wrongtokenerror
//			}
//
//			kn := fs.Output.Bytes()
//			if len(kn) <= 0 {
//				// "" case. hrm.
//				currentKey = ffjtSignMessageOperationnosuchkey
//				state = fflib.FFParse_want_colon
//				goto mainparse
//			} else {
//				switch kn[0] {
//
//				case 'a':
//
//					if bytes.Equal(ffjKeySignMessageOperationActive, kn) {
//						currentKey = ffjtSignMessageOperationActive
//						state = fflib.FFParse_want_colon
//						goto mainparse
//					}
//
//				case 'e':
//
//					if bytes.Equal(ffjKeySignMessageOperationExtensions, kn) {
//						currentKey = ffjtSignMessageOperationExtensions
//						state = fflib.FFParse_want_colon
//						goto mainparse
//					}
//
//				case 'f':
//
//					if bytes.Equal(ffjKeySignMessageOperationFee, kn) {
//						currentKey = ffjtSignMessageOperationFee
//						state = fflib.FFParse_want_colon
//						goto mainparse
//					}
//
//				case 'n':
//
//					if bytes.Equal(ffjKeySignMessageOperationName, kn) {
//						currentKey = ffjtSignMessageOperationName
//						state = fflib.FFParse_want_colon
//						goto mainparse
//					}
//
//				case 'o':
//
//					if bytes.Equal(ffjKeySignMessageOperationOwner, kn) {
//						currentKey = ffjtSignMessageOperationOwner
//						state = fflib.FFParse_want_colon
//						goto mainparse
//
//					} else if bytes.Equal(ffjKeySignMessageOperationOptions, kn) {
//						currentKey = ffjtSignMessageOperationOptions
//						state = fflib.FFParse_want_colon
//						goto mainparse
//					}
//
//				case 'r':
//
//					if bytes.Equal(ffjKeySignMessageOperationRegistrar, kn) {
//						currentKey = ffjtSignMessageOperationRegistrar
//						state = fflib.FFParse_want_colon
//						goto mainparse
//
//					} else if bytes.Equal(ffjKeySignMessageOperationReferrer, kn) {
//						currentKey = ffjtSignMessageOperationReferrer
//						state = fflib.FFParse_want_colon
//						goto mainparse
//
//					} else if bytes.Equal(ffjKeySignMessageOperationReferrerPercent, kn) {
//						currentKey = ffjtSignMessageOperationReferrerPercent
//						state = fflib.FFParse_want_colon
//						goto mainparse
//					}
//
//				}
//
//				if fflib.SimpleLetterEqualFold(ffjKeySignMessageOperationFee, kn) {
//					currentKey = ffjtSignMessageOperationFee
//					state = fflib.FFParse_want_colon
//					goto mainparse
//				}
//
//				if fflib.EqualFoldRight(ffjKeySignMessageOperationOptions, kn) {
//					currentKey = ffjtSignMessageOperationOptions
//					state = fflib.FFParse_want_colon
//					goto mainparse
//				}
//
//				if fflib.EqualFoldRight(ffjKeySignMessageOperationExtensions, kn) {
//					currentKey = ffjtSignMessageOperationExtensions
//					state = fflib.FFParse_want_colon
//					goto mainparse
//				}
//
//				if fflib.SimpleLetterEqualFold(ffjKeySignMessageOperationName, kn) {
//					currentKey = ffjtSignMessageOperationName
//					state = fflib.FFParse_want_colon
//					goto mainparse
//				}
//
//				if fflib.SimpleLetterEqualFold(ffjKeySignMessageOperationActive, kn) {
//					currentKey = ffjtSignMessageOperationActive
//					state = fflib.FFParse_want_colon
//					goto mainparse
//				}
//
//				if fflib.SimpleLetterEqualFold(ffjKeySignMessageOperationOwner, kn) {
//					currentKey = ffjtSignMessageOperationOwner
//					state = fflib.FFParse_want_colon
//					goto mainparse
//				}
//
//				if fflib.AsciiEqualFold(ffjKeySignMessageOperationReferrerPercent, kn) {
//					currentKey = ffjtSignMessageOperationReferrerPercent
//					state = fflib.FFParse_want_colon
//					goto mainparse
//				}
//
//				if fflib.SimpleLetterEqualFold(ffjKeySignMessageOperationReferrer, kn) {
//					currentKey = ffjtSignMessageOperationReferrer
//					state = fflib.FFParse_want_colon
//					goto mainparse
//				}
//
//				if fflib.EqualFoldRight(ffjKeySignMessageOperationRegistrar, kn) {
//					currentKey = ffjtSignMessageOperationRegistrar
//					state = fflib.FFParse_want_colon
//					goto mainparse
//				}
//
//				currentKey = ffjtSignMessageOperationnosuchkey
//				state = fflib.FFParse_want_colon
//				goto mainparse
//			}
//
//		case fflib.FFParse_want_colon:
//			if tok != fflib.FFTok_colon {
//				wantedTok = fflib.FFTok_colon
//				goto wrongtokenerror
//			}
//			state = fflib.FFParse_want_value
//			continue
//		case fflib.FFParse_want_value:
//
//			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
//				switch currentKey {
//
//				case ffjtSignMessageOperationRegistrar:
//					goto handle_Registrar
//
//				case ffjtSignMessageOperationReferrer:
//					goto handle_Referrer
//
//				case ffjtSignMessageOperationReferrerPercent:
//					goto handle_ReferrerPercent
//
//				case ffjtSignMessageOperationOwner:
//					goto handle_Owner
//
//				case ffjtSignMessageOperationActive:
//					goto handle_Active
//
//				case ffjtSignMessageOperationName:
//					goto handle_Name
//
//				case ffjtSignMessageOperationExtensions:
//					goto handle_Extensions
//
//				case ffjtSignMessageOperationOptions:
//					goto handle_Options
//
//				case ffjtSignMessageOperationFee:
//					goto handle_Fee
//
//				case ffjtSignMessageOperationnosuchkey:
//					err = fs.SkipField(tok)
//					if err != nil {
//						return fs.WrapErr(err)
//					}
//					state = fflib.FFParse_after_value
//					goto mainparse
//				}
//			} else {
//				goto wantedvalue
//			}
//		}
//	}
//
//handle_Registrar:
//
///* handler: j.Registrar type=types.GrapheneID kind=struct quoted=false*/
//
//	{
//		if tok == fflib.FFTok_null {
//
//		} else {
//
//			tbuf, err := fs.CaptureField(tok)
//			if err != nil {
//				return fs.WrapErr(err)
//			}
//
//			err = j.Registrar.UnmarshalJSON(tbuf)
//			if err != nil {
//				return fs.WrapErr(err)
//			}
//		}
//		state = fflib.FFParse_after_value
//	}
//
//	state = fflib.FFParse_after_value
//	goto mainparse
//
//handle_Referrer:
//
///* handler: j.Referrer type=types.GrapheneID kind=struct quoted=false*/
//
//	{
//		if tok == fflib.FFTok_null {
//
//		} else {
//
//			tbuf, err := fs.CaptureField(tok)
//			if err != nil {
//				return fs.WrapErr(err)
//			}
//
//			err = j.Referrer.UnmarshalJSON(tbuf)
//			if err != nil {
//				return fs.WrapErr(err)
//			}
//		}
//		state = fflib.FFParse_after_value
//	}
//
//	state = fflib.FFParse_after_value
//	goto mainparse
//
//handle_ReferrerPercent:
//
///* handler: j.ReferrerPercent type=types.UInt16 kind=uint16 quoted=false*/
//
//	{
//		if tok == fflib.FFTok_null {
//
//		} else {
//
//			tbuf, err := fs.CaptureField(tok)
//			if err != nil {
//				return fs.WrapErr(err)
//			}
//
//			err = j.ReferrerPercent.UnmarshalJSON(tbuf)
//			if err != nil {
//				return fs.WrapErr(err)
//			}
//		}
//		state = fflib.FFParse_after_value
//	}
//
//	state = fflib.FFParse_after_value
//	goto mainparse
//
//handle_Owner:
//
///* handler: j.Owner type=types.Authority kind=struct quoted=false*/
//
//	{
//		/* Falling back. type=types.Authority kind=struct */
//		tbuf, err := fs.CaptureField(tok)
//		if err != nil {
//			return fs.WrapErr(err)
//		}
//
//		err = json.Unmarshal(tbuf, &j.Owner)
//		if err != nil {
//			return fs.WrapErr(err)
//		}
//	}
//
//	state = fflib.FFParse_after_value
//	goto mainparse
//
//handle_Active:
//
///* handler: j.Active type=types.Authority kind=struct quoted=false*/
//
//	{
//		/* Falling back. type=types.Authority kind=struct */
//		tbuf, err := fs.CaptureField(tok)
//		if err != nil {
//			return fs.WrapErr(err)
//		}
//
//		err = json.Unmarshal(tbuf, &j.Active)
//		if err != nil {
//			return fs.WrapErr(err)
//		}
//	}
//
//	state = fflib.FFParse_after_value
//	goto mainparse
//
//handle_Name:
//
///* handler: j.Name type=string kind=string quoted=false*/
//
//	{
//
//		{
//			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
//				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
//			}
//		}
//
//		if tok == fflib.FFTok_null {
//
//		} else {
//
//			outBuf := fs.Output.Bytes()
//
//			j.Name = string(string(outBuf))
//
//		}
//	}
//
//	state = fflib.FFParse_after_value
//	goto mainparse
//
//handle_Extensions:
//
///* handler: j.Extensions type=types.AccountCreateExtensions kind=struct quoted=false*/
//
//	{
//		/* Falling back. type=types.AccountCreateExtensions kind=struct */
//		tbuf, err := fs.CaptureField(tok)
//		if err != nil {
//			return fs.WrapErr(err)
//		}
//
//		err = json.Unmarshal(tbuf, &j.Extensions)
//		if err != nil {
//			return fs.WrapErr(err)
//		}
//	}
//
//	state = fflib.FFParse_after_value
//	goto mainparse
//
//handle_Options:
//
///* handler: j.Options type=types.AccountOptions kind=struct quoted=false*/
//
//	{
//		/* Falling back. type=types.AccountOptions kind=struct */
//		tbuf, err := fs.CaptureField(tok)
//		if err != nil {
//			return fs.WrapErr(err)
//		}
//
//		err = json.Unmarshal(tbuf, &j.Options)
//		if err != nil {
//			return fs.WrapErr(err)
//		}
//	}
//
//	state = fflib.FFParse_after_value
//	goto mainparse
//
//handle_Fee:
//
///* handler: j.Fee type=types.AssetAmount kind=struct quoted=false*/
//
//	{
//		/* Falling back. type=types.AssetAmount kind=struct */
//		tbuf, err := fs.CaptureField(tok)
//		if err != nil {
//			return fs.WrapErr(err)
//		}
//
//		err = json.Unmarshal(tbuf, &j.Fee)
//		if err != nil {
//			return fs.WrapErr(err)
//		}
//	}
//
//	state = fflib.FFParse_after_value
//	goto mainparse
//
//wantedvalue:
//	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
//wrongtokenerror:
//	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
//tokerror:
//	if fs.BigError != nil {
//		return fs.WrapErr(fs.BigError)
//	}
//	err = fs.Error.ToError()
//	if err != nil {
//		return fs.WrapErr(err)
//	}
//	panic("ffjson-generated: unreachable, please report bug.")
//done:
//
//	return nil
//}
