package NewCustStructures

import (
	"github.com/uptrace/go-clickhouse/ch"
	"time"
)

type GeneralTrafficFields struct {
	Probe          string    `ch:"Probe"`
	Country        string    `ch:"Country"`
	TypeConnection uint8     `ch:"TypeConnection"`
	ProbeTimestamp time.Time `ch:"ProbeTimestamp, type:DateTime"`
}

type CustomerTrafficDown struct {
	Idx                uint32 `ch:"Idx"`
	Period             uint32 `ch:"Period"`
	APN                string `ch:"APN"`
	Customer           string `ch:"Customer"`
	AccessType         uint8  `ch:"AccessType"`
	Terminal           string `ch:"Terminal"`
	AddrIp             string `ch:"AddrIp"`
	Status             uint32 `ch:"Status"`
	AS                 string `ch:"AS"`
	IpSGSN             string `ch:"IpSGSN"`
	MaxTcpSession      uint32 `ch:"MaxTcpSession"`
	NPacket            uint32 `ch:"NPacket"`
	NByte              uint64 `ch:"NByte"`
	NByteTcp           uint64 `ch:"NByteTcp"`
	NByteUdp           uint64 `ch:"NByteUdp"`
	ByteApp0           uint64 `ch:"ByteApp0"`
	PackApp0           uint32 `ch:"PackApp0"`
	ByteApp1           uint64 `ch:"ByteApp1"`
	PackApp1           uint32 `ch:"PackApp1"`
	ByteApp2           uint64 `ch:"ByteApp2"`
	PackApp2           uint32 `ch:"PackApp2"`
	ByteApp3           uint64 `ch:"ByteApp3"`
	PackApp3           uint32 `ch:"PackApp3"`
	ByteApp4           uint64 `ch:"ByteApp4"`
	PackApp4           uint32 `ch:"PackApp4"`
	ByteApp5           uint64 `ch:"ByteApp5"`
	PackApp5           uint32 `ch:"PackApp5"`
	ByteApp6           uint64 `ch:"ByteApp6"`
	PackApp6           uint32 `ch:"PackApp6"`
	ByteApp7           uint64 `ch:"ByteApp7"`
	PackApp7           uint32 `ch:"PackApp7"`
	ByteApp8           uint64 `ch:"ByteApp8"`
	PackApp8           uint32 `ch:"PackApp8"`
	ByteApp9           uint64 `ch:"ByteApp9"`
	PackApp9           uint32 `ch:"PackApp9"`
	ByteApp10          uint64 `ch:"ByteApp10"`
	PackApp10          uint32 `ch:"PackApp10"`
	ByteApp11          uint64 `ch:"ByteApp11"`
	PackApp11          uint32 `ch:"PackApp11"`
	ByteApp12          uint64 `ch:"ByteApp12"`
	PackApp12          uint32 `ch:"PackApp12"`
	ByteApp13          uint64 `ch:"ByteApp13"`
	PackApp13          uint32 `ch:"PackApp13"`
	ByteApp14          uint64 `ch:"ByteApp14"`
	PackApp14          uint32 `ch:"PackApp14"`
	ByteApp15          uint64 `ch:"ByteApp15"`
	PackApp15          uint32 `ch:"PackApp15"`
	ByteApp16          uint64 `ch:"ByteApp16"`
	PackApp16          uint32 `ch:"PackApp16"`
	ByteApp17          uint64 `ch:"ByteApp17"`
	PackApp17          uint32 `ch:"PackApp17"`
	ByteApp18          uint64 `ch:"ByteApp18"`
	PackApp18          uint32 `ch:"PackApp18"`
	OsNum              uint8  `ch:"OsNum"`
	Fingerprint        string `ch:"Fingerprint"`
	Device             string `ch:"Device"`
	LocInfo            string `ch:"LocInfo"`
	ServproviderField0 string `ch:"ServproviderField0"`
	ServproviderField1 string `ch:"ServproviderField1"`
	IdxCnx             uint32 `ch:"IdxCnx"`
	IdxMainCnx         uint32 `ch:"IdxMainCnx"`
	MaxVol10s          uint32 `ch:"MaxVol10s"`
	NbSecActive        uint32 `ch:"NbSecActive"`
	GeneralTrafficFields
	ch.CHModel `ch:"table:CustomerConnectionDown,partition:toYYYYMM(timestamp)"`
}

type CustomerTrafficUp struct {
	Idx           uint32 `ch:"Idx"`
	Period        uint32 `ch:"Period"`
	APN           string `ch:"APN"`
	Customer      string `ch:"Customer"`
	AccessType    uint8  `ch:"AccessType"`
	Terminal      string `ch:"Terminal"`
	AddrIp        string `ch:"AddrIp"`
	Status        uint32 `ch:"Status"`
	AS            string `ch:"AS"`
	IpSGSN        string `ch:"IpSGSN"`
	MaxTcpSession uint32 `ch:"MaxTcpSession"`
	NPacket       uint32 `ch:"NPacket"`
	NByte         uint64 `ch:"NByte"`
	NByteTcp      uint64 `ch:"NByteTcp"`
	NByteUdp      uint64 `ch:"NByteUdp"`
	ByteApp0      uint64 `ch:"ByteApp0"`
	PackApp0      uint32 `ch:"PackApp0"`
	ByteApp1      uint64 `ch:"ByteApp1"`
	PackApp1      uint32 `ch:"PackApp1"`
	ByteApp2      uint64 `ch:"ByteApp2"`
	PackApp2      uint32 `ch:"PackApp2"`
	ByteApp3      uint64 `ch:"ByteApp3"`
	PackApp3      uint32 `ch:"PackApp3"`
	ByteApp4      uint64 `ch:"ByteApp4"`
	PackApp4      uint32 `ch:"PackApp4"`
	ByteApp5      uint64 `ch:"ByteApp5"`
	PackApp5      uint32 `ch:"PackApp5"`
	ByteApp6      uint64 `ch:"ByteApp6"`
	PackApp6      uint32 `ch:"PackApp6"`
	ByteApp7      uint64 `ch:"ByteApp7"`
	PackApp7      uint32 `ch:"PackApp7"`
	ByteApp8      uint64 `ch:"ByteApp8"`
	PackApp8      uint32 `ch:"PackApp8"`
	ByteApp9      uint64 `ch:"ByteApp9"`
	PackApp9      uint32 `ch:"PackApp9"`
	ByteApp10     uint64 `ch:"ByteApp10"`
	PackApp10     uint32 `ch:"PackApp10"`
	ByteApp11     uint64 `ch:"ByteApp11"`
	PackApp11     uint32 `ch:"PackApp11"`
	ByteApp12     uint64 `ch:"ByteApp12"`
	PackApp12     uint32 `ch:"PackApp12"`
	ByteApp13     uint64 `ch:"ByteApp13"`
	PackApp13     uint32 `ch:"PackApp13"`
	ByteApp14     uint64 `ch:"ByteApp14"`
	PackApp14     uint32 `ch:"PackApp14"`
	ByteApp15     uint64 `ch:"ByteApp15"`
	PackApp15     uint32 `ch:"PackApp15"`
	ByteApp16     uint64 `ch:"ByteApp16"`
	PackApp16     uint32 `ch:"PackApp16"`
	ByteApp17     uint64 `ch:"ByteApp17"`
	PackApp17     uint32 `ch:"PackApp17"`
	ByteApp18     uint64 `ch:"ByteApp18"`
	PackApp18     uint32 `ch:"PackApp18"`
	OsNum         uint8  `ch:"OsNum"`
	Fingerprint   string `ch:"Fingerprint"`
	Device        string `ch:"Device"`
	LocInfo       string `ch:"LocInfo"`
	IdxCnx        uint32 `ch:"IdxCnx"`
	IdxMainCnx    uint32 `ch:"IdxMainCnx"`
	MaxVol10s     uint32 `ch:"MaxVol10s"`
	GeneralTrafficFields
	ch.CHModel `ch:"table:CustomerConnectionUp,partition:toYYYYMM(timestamp)"`
}
