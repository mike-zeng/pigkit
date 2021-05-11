package request

type PigReq struct {
	ServerName string
	MetaData map[string]interface{}
	SerializationType int
	Content interface{}
}

func (req *PigReq) Bytes() []byte {
	return nil
}