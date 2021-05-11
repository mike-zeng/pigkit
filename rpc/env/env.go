package env

const(
	Version = 0x0
)

func CheckVersion(n uint8)bool  {
	return n<=Version
}