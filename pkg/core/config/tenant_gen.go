package config

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *TenantConfiguration) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "DATABASE_URL":
			z.DBConnectionStr, err = dc.ReadString()
			if err != nil {
				return
			}
		case "API_KEY":
			z.APIKey, err = dc.ReadString()
			if err != nil {
				return
			}
		case "MASTER_KEY":
			z.MasterKey, err = dc.ReadString()
			if err != nil {
				return
			}
		case "APP_NAME":
			z.AppName, err = dc.ReadString()
			if err != nil {
				return
			}
		case "TokenStore":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			for zb0002 > 0 {
				zb0002--
				field, err = dc.ReadMapKeyPtr()
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "TOKEN_STORE_SECRET":
					z.TokenStore.Secret, err = dc.ReadString()
					if err != nil {
						return
					}
				case "TOKEN_STORE_EXPIRY":
					z.TokenStore.Expiry, err = dc.ReadInt64()
					if err != nil {
						return
					}
				default:
					err = dc.Skip()
					if err != nil {
						return
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *TenantConfiguration) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 5
	// write "DATABASE_URL"
	err = en.Append(0x85, 0xac, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x55, 0x52, 0x4c)
	if err != nil {
		return
	}
	err = en.WriteString(z.DBConnectionStr)
	if err != nil {
		return
	}
	// write "API_KEY"
	err = en.Append(0xa7, 0x41, 0x50, 0x49, 0x5f, 0x4b, 0x45, 0x59)
	if err != nil {
		return
	}
	err = en.WriteString(z.APIKey)
	if err != nil {
		return
	}
	// write "MASTER_KEY"
	err = en.Append(0xaa, 0x4d, 0x41, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x4b, 0x45, 0x59)
	if err != nil {
		return
	}
	err = en.WriteString(z.MasterKey)
	if err != nil {
		return
	}
	// write "APP_NAME"
	err = en.Append(0xa8, 0x41, 0x50, 0x50, 0x5f, 0x4e, 0x41, 0x4d, 0x45)
	if err != nil {
		return
	}
	err = en.WriteString(z.AppName)
	if err != nil {
		return
	}
	// write "TokenStore"
	// map header, size 2
	// write "TOKEN_STORE_SECRET"
	err = en.Append(0xaa, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x82, 0xb2, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54)
	if err != nil {
		return
	}
	err = en.WriteString(z.TokenStore.Secret)
	if err != nil {
		return
	}
	// write "TOKEN_STORE_EXPIRY"
	err = en.Append(0xb2, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x59)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.TokenStore.Expiry)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TenantConfiguration) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "DATABASE_URL"
	o = append(o, 0x85, 0xac, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x55, 0x52, 0x4c)
	o = msgp.AppendString(o, z.DBConnectionStr)
	// string "API_KEY"
	o = append(o, 0xa7, 0x41, 0x50, 0x49, 0x5f, 0x4b, 0x45, 0x59)
	o = msgp.AppendString(o, z.APIKey)
	// string "MASTER_KEY"
	o = append(o, 0xaa, 0x4d, 0x41, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x4b, 0x45, 0x59)
	o = msgp.AppendString(o, z.MasterKey)
	// string "APP_NAME"
	o = append(o, 0xa8, 0x41, 0x50, 0x50, 0x5f, 0x4e, 0x41, 0x4d, 0x45)
	o = msgp.AppendString(o, z.AppName)
	// string "TokenStore"
	// map header, size 2
	// string "TOKEN_STORE_SECRET"
	o = append(o, 0xaa, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x82, 0xb2, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54)
	o = msgp.AppendString(o, z.TokenStore.Secret)
	// string "TOKEN_STORE_EXPIRY"
	o = append(o, 0xb2, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x59)
	o = msgp.AppendInt64(o, z.TokenStore.Expiry)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TenantConfiguration) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "DATABASE_URL":
			z.DBConnectionStr, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "API_KEY":
			z.APIKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "MASTER_KEY":
			z.MasterKey, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "APP_NAME":
			z.AppName, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "TokenStore":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			for zb0002 > 0 {
				zb0002--
				field, bts, err = msgp.ReadMapKeyZC(bts)
				if err != nil {
					return
				}
				switch msgp.UnsafeString(field) {
				case "TOKEN_STORE_SECRET":
					z.TokenStore.Secret, bts, err = msgp.ReadStringBytes(bts)
					if err != nil {
						return
					}
				case "TOKEN_STORE_EXPIRY":
					z.TokenStore.Expiry, bts, err = msgp.ReadInt64Bytes(bts)
					if err != nil {
						return
					}
				default:
					bts, err = msgp.Skip(bts)
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TenantConfiguration) Msgsize() (s int) {
	s = 1 + 13 + msgp.StringPrefixSize + len(z.DBConnectionStr) + 8 + msgp.StringPrefixSize + len(z.APIKey) + 11 + msgp.StringPrefixSize + len(z.MasterKey) + 9 + msgp.StringPrefixSize + len(z.AppName) + 11 + 1 + 19 + msgp.StringPrefixSize + len(z.TokenStore.Secret) + 19 + msgp.Int64Size
	return
}
