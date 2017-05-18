package robot

//IPacketModle use by netpactket
type IPacketModle interface {
	Serialize() []byte
}
