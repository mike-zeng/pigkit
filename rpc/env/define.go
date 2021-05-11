package env

// MsgType
const(
	MsgReqType = 0x0
	MsgRespType = 0x1
	MsgHeartType = 0x2
)

func CheckMsgType(n uint8) bool {
	return n == MsgReqType || n == MsgRespType || n == MsgHeartType
}

// MagicNum 魔数版本号
const (
	MagicNum = 0x1f
)

func CheckMagicNum(n uint8) bool {
	return n ==MagicNum
}

const(
	ReqTypeSendAndRec = 0x0 // send and receive
	ReqTypeSendButNotRec = 0x1 // send but not receive
	ReqTypeClientStream = 0x2 // client stream request
)

func CheckReqType(n uint8) bool {
	return n ==ReqTypeSendAndRec || n == ReqTypeSendButNotRec || n == ReqTypeClientStream
}