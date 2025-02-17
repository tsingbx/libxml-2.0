package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const ULOCCHINESE = "zh"
const ULOCENGLISH = "en"
const ULOCFRENCH = "fr"
const ULOCGERMAN = "de"
const ULOCITALIAN = "it"
const ULOCJAPANESE = "ja"
const ULOCKOREAN = "ko"
const ULOCSIMPLIFIEDCHINESE = "zh_CN"
const ULOCTRADITIONALCHINESE = "zh_TW"
const ULOCCANADA = "en_CA"
const ULOCCANADAFRENCH = "fr_CA"
const ULOCCHINA = "zh_CN"
const ULOCPRC = "zh_CN"
const ULOCFRANCE = "fr_FR"
const ULOCGERMANY = "de_DE"
const ULOCITALY = "it_IT"
const ULOCJAPAN = "ja_JP"
const ULOCKOREA = "ko_KR"
const ULOCTAIWAN = "zh_TW"
const ULOCUK = "en_GB"
const ULOCUS = "en_US"
const ULOCLANGCAPACITY = 12
const ULOCCOUNTRYCAPACITY = 4
const ULOCFULLNAMECAPACITY = 157
const ULOCSCRIPTCAPACITY = 6
const ULOCKEYWORDSCAPACITY = 96
const ULOCKEYWORDANDVALUESCAPACITY = 100
const ULOCKEYWORDSEPARATOR = "@"
const ULOCKEYWORDSEPARATORUNICODE = 0x40
const ULOCKEYWORDASSIGN = "="
const ULOCKEYWORDASSIGNUNICODE = 0x3D
const ULOCKEYWORDITEMSEPARATOR = ";"
const ULOCKEYWORDITEMSEPARATORUNICODE = 0x3B

type ULocDataLocaleType c.Int

const (
	ULocDataLocaleTypeULOCACTUALLOCALE        ULocDataLocaleType = 0
	ULocDataLocaleTypeULOCVALIDLOCALE         ULocDataLocaleType = 1
	ULocDataLocaleTypeULOCREQUESTEDLOCALE     ULocDataLocaleType = 2
	ULocDataLocaleTypeULOCDATALOCALETYPELIMIT ULocDataLocaleType = 3
)

type ULocAvailableType c.Int

const (
	ULocAvailableTypeULOCAVAILABLEDEFAULT           ULocAvailableType = 0
	ULocAvailableTypeULOCAVAILABLEONLYLEGACYALIASES ULocAvailableType = 1
	ULocAvailableTypeULOCAVAILABLEWITHLEGACYALIASES ULocAvailableType = 2
	ULocAvailableTypeULOCAVAILABLECOUNT             ULocAvailableType = 3
)

type ULayoutType c.Int

const (
	ULayoutTypeULOCLAYOUTLTR     ULayoutType = 0
	ULayoutTypeULOCLAYOUTRTL     ULayoutType = 1
	ULayoutTypeULOCLAYOUTTTB     ULayoutType = 2
	ULayoutTypeULOCLAYOUTBTT     ULayoutType = 3
	ULayoutTypeULOCLAYOUTUNKNOWN ULayoutType = 4
)

type UAcceptResult c.Int

const (
	UAcceptResultULOCACCEPTFAILED   UAcceptResult = 0
	UAcceptResultULOCACCEPTVALID    UAcceptResult = 1
	UAcceptResultULOCACCEPTFALLBACK UAcceptResult = 2
)
