package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const ZLIBVERSION = "1.3.1"
const ZLIBVERNUM = 0x1310
const ZLIBVERMAJOR = 1
const ZLIBVERMINOR = 3
const ZLIBVERREVISION = 1
const ZLIBVERSUBREVISION = 0
const ZNOFLUSH = 0
const ZPARTIALFLUSH = 1
const ZSYNCFLUSH = 2
const ZFULLFLUSH = 3
const ZFINISH = 4
const ZBLOCK = 5
const ZTREES = 6
const ZOK = 0
const ZSTREAMEND = 1
const ZNEEDDICT = 2
const ZNOCOMPRESSION = 0
const ZBESTSPEED = 1
const ZBESTCOMPRESSION = 9
const ZFILTERED = 1
const ZHUFFMANONLY = 2
const ZRLE = 3
const ZFIXED = 4
const ZDEFAULTSTRATEGY = 0
const ZBINARY = 0
const ZTEXT = 1
const ZUNKNOWN = 2
const ZDEFLATED = 8
const ZNULL = 0
// llgo:type C
type AllocFunc func(Voidpf, UInt, UInt) Voidpf
// llgo:type C
type FreeFunc func(Voidpf, Voidpf)

type InternalState struct {
	Unused [8]uint8
}

type ZStreamS struct {
	NextIn   *Bytef
	AvailIn  UInt
	TotalIn  ULong
	NextOut  *Bytef
	AvailOut UInt
	TotalOut ULong
	Msg      *int8
	State    *InternalState
	Zalloc   unsafe.Pointer
	Zfree    unsafe.Pointer
	Opaque   Voidpf
	DataType c.Int
	Adler    ULong
	Reserved ULong
}
type ZStream ZStreamS
type ZStreamp *ZStream

type GzHeaderS struct {
	Text     c.Int
	Time     ULong
	Xflags   c.Int
	Os       c.Int
	Extra    *Bytef
	ExtraLen UInt
	ExtraMax UInt
	Name     *Bytef
	NameMax  UInt
	Comment  *Bytef
	CommMax  UInt
	Hcrc     c.Int
	Done     c.Int
}
type GzHeader GzHeaderS
type GzHeaderp *GzHeader
//go:linkname Inflate C.inflate
func Inflate(strm ZStreamp, flush c.Int) c.Int
//go:linkname InflateEnd C.inflateEnd
func InflateEnd(strm ZStreamp) c.Int
// llgo:type C
type InFunc func(unsafe.Pointer, **int8) c.Uint
// llgo:type C
type OutFunc func(unsafe.Pointer, *int8, c.Uint) c.Int

type GzFileS struct {
	Unused [8]uint8
}
type GzFile *GzFileS
//go:linkname Gzdopen C.gzdopen
func Gzdopen(fd c.Int, mode *int8) GzFile
//go:linkname Gzread C.gzread
func Gzread(file GzFile, buf Voidp, len c.Uint) c.Int
//go:linkname Gzwrite C.gzwrite
func Gzwrite(file GzFile, buf Voidpc, len c.Uint) c.Int
//go:linkname Gzdirect C.gzdirect
func Gzdirect(file GzFile) c.Int
//go:linkname Gzclose C.gzclose
func Gzclose(file GzFile) c.Int
//go:linkname InflateInit2 C.inflateInit2_
func InflateInit2(strm ZStreamp, windowBits c.Int, version *int8, stream_size c.Int) c.Int
