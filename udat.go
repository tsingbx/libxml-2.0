package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const UDATYEAR = "y"
const UDATQUARTER = "QQQQ"
const UDATABBRQUARTER = "QQQ"
const UDATYEARQUARTER = "yQQQQ"
const UDATYEARABBRQUARTER = "yQQQ"
const UDATMONTH = "MMMM"
const UDATABBRMONTH = "MMM"
const UDATNUMMONTH = "M"
const UDATYEARMONTH = "yMMMM"
const UDATYEARABBRMONTH = "yMMM"
const UDATYEARNUMMONTH = "yM"
const UDATDAY = "d"
const UDATYEARMONTHDAY = "yMMMMd"
const UDATYEARABBRMONTHDAY = "yMMMd"
const UDATYEARNUMMONTHDAY = "yMd"
const UDATWEEKDAY = "EEEE"
const UDATABBRWEEKDAY = "E"
const UDATYEARMONTHWEEKDAYDAY = "yMMMMEEEEd"
const UDATYEARABBRMONTHWEEKDAYDAY = "yMMMEd"
const UDATYEARNUMMONTHWEEKDAYDAY = "yMEd"
const UDATMONTHDAY = "MMMMd"
const UDATABBRMONTHDAY = "MMMd"
const UDATNUMMONTHDAY = "Md"
const UDATMONTHWEEKDAYDAY = "MMMMEEEEd"
const UDATABBRMONTHWEEKDAYDAY = "MMMEd"
const UDATNUMMONTHWEEKDAYDAY = "MEd"
const UDATHOUR = "j"
const UDATHOUR24 = "H"
const UDATMINUTE = "m"
const UDATHOURMINUTE = "jm"
const UDATHOUR24MINUTE = "Hm"
const UDATSECOND = "s"
const UDATHOURMINUTESECOND = "jms"
const UDATHOUR24MINUTESECOND = "Hms"
const UDATMINUTESECOND = "ms"
const UDATLOCATIONTZ = "VVVV"
const UDATGENERICTZ = "vvvv"
const UDATABBRGENERICTZ = "v"
const UDATSPECIFICTZ = "zzzz"
const UDATABBRSPECIFICTZ = "z"
const UDATABBRUTCTZ = "ZZZZ"
const UDATSTANDALONEMONTH = "LLLL"
const UDATABBRSTANDALONEMONTH = "LLL"
const UDATHOURMINUTEGENERICTZ = "jmv"
const UDATHOURMINUTETZ = "jmz"
const UDATHOURGENERICTZ = "jv"
const UDATHOURTZ = "jz"
const JPERA2019ROOT = "Reiwa"
const JPERA2019JA = "\\u4EE4\\u548C"
const JPERA2019NARROW = "R"
const UDATHASPATTERNCHARFORTIMESEPARATOR = 0

type UDateFormat unsafe.Pointer
type UDateFormatStyle c.Int

const (
	UDateFormatStyleUDATFULL           UDateFormatStyle = 0
	UDateFormatStyleUDATLONG           UDateFormatStyle = 1
	UDateFormatStyleUDATMEDIUM         UDateFormatStyle = 2
	UDateFormatStyleUDATSHORT          UDateFormatStyle = 3
	UDateFormatStyleUDATDEFAULT        UDateFormatStyle = 2
	UDateFormatStyleUDATRELATIVE       UDateFormatStyle = 128
	UDateFormatStyleUDATFULLRELATIVE   UDateFormatStyle = 128
	UDateFormatStyleUDATLONGRELATIVE   UDateFormatStyle = 129
	UDateFormatStyleUDATMEDIUMRELATIVE UDateFormatStyle = 130
	UDateFormatStyleUDATSHORTRELATIVE  UDateFormatStyle = 131
	UDateFormatStyleUDATNONE           UDateFormatStyle = -1
	UDateFormatStyleUDATPATTERN        UDateFormatStyle = -2
	UDateFormatStyleUDATIGNORE         UDateFormatStyle = -2
)

type UDateFormatField c.Int

const (
	UDateFormatFieldUDATERAFIELD                        UDateFormatField = 0
	UDateFormatFieldUDATYEARFIELD                       UDateFormatField = 1
	UDateFormatFieldUDATMONTHFIELD                      UDateFormatField = 2
	UDateFormatFieldUDATDATEFIELD                       UDateFormatField = 3
	UDateFormatFieldUDATHOUROFDAY1FIELD                 UDateFormatField = 4
	UDateFormatFieldUDATHOUROFDAY0FIELD                 UDateFormatField = 5
	UDateFormatFieldUDATMINUTEFIELD                     UDateFormatField = 6
	UDateFormatFieldUDATSECONDFIELD                     UDateFormatField = 7
	UDateFormatFieldUDATFRACTIONALSECONDFIELD           UDateFormatField = 8
	UDateFormatFieldUDATDAYOFWEEKFIELD                  UDateFormatField = 9
	UDateFormatFieldUDATDAYOFYEARFIELD                  UDateFormatField = 10
	UDateFormatFieldUDATDAYOFWEEKINMONTHFIELD           UDateFormatField = 11
	UDateFormatFieldUDATWEEKOFYEARFIELD                 UDateFormatField = 12
	UDateFormatFieldUDATWEEKOFMONTHFIELD                UDateFormatField = 13
	UDateFormatFieldUDATAMPMFIELD                       UDateFormatField = 14
	UDateFormatFieldUDATHOUR1FIELD                      UDateFormatField = 15
	UDateFormatFieldUDATHOUR0FIELD                      UDateFormatField = 16
	UDateFormatFieldUDATTIMEZONEFIELD                   UDateFormatField = 17
	UDateFormatFieldUDATYEARWOYFIELD                    UDateFormatField = 18
	UDateFormatFieldUDATDOWLOCALFIELD                   UDateFormatField = 19
	UDateFormatFieldUDATEXTENDEDYEARFIELD               UDateFormatField = 20
	UDateFormatFieldUDATJULIANDAYFIELD                  UDateFormatField = 21
	UDateFormatFieldUDATMILLISECONDSINDAYFIELD          UDateFormatField = 22
	UDateFormatFieldUDATTIMEZONERFCFIELD                UDateFormatField = 23
	UDateFormatFieldUDATTIMEZONEGENERICFIELD            UDateFormatField = 24
	UDateFormatFieldUDATSTANDALONEDAYFIELD              UDateFormatField = 25
	UDateFormatFieldUDATSTANDALONEMONTHFIELD            UDateFormatField = 26
	UDateFormatFieldUDATQUARTERFIELD                    UDateFormatField = 27
	UDateFormatFieldUDATSTANDALONEQUARTERFIELD          UDateFormatField = 28
	UDateFormatFieldUDATTIMEZONESPECIALFIELD            UDateFormatField = 29
	UDateFormatFieldUDATYEARNAMEFIELD                   UDateFormatField = 30
	UDateFormatFieldUDATTIMEZONELOCALIZEDGMTOFFSETFIELD UDateFormatField = 31
	UDateFormatFieldUDATTIMEZONEISOFIELD                UDateFormatField = 32
	UDateFormatFieldUDATTIMEZONEISOLOCALFIELD           UDateFormatField = 33
	UDateFormatFieldUDATRELATEDYEARFIELD                UDateFormatField = 34
	UDateFormatFieldUDATAMPMMIDNIGHTNOONFIELD           UDateFormatField = 35
	UDateFormatFieldUDATFLEXIBLEDAYPERIODFIELD          UDateFormatField = 36
	UDateFormatFieldUDATTIMESEPARATORFIELD              UDateFormatField = 37
	UDateFormatFieldUDATFIELDCOUNT                      UDateFormatField = 38
)

type UDateFormatBooleanAttribute c.Int

const (
	UDateFormatBooleanAttributeUDATPARSEALLOWWHITESPACE          UDateFormatBooleanAttribute = 0
	UDateFormatBooleanAttributeUDATPARSEALLOWNUMERIC             UDateFormatBooleanAttribute = 1
	UDateFormatBooleanAttributeUDATPARSEPARTIALLITERALMATCH      UDateFormatBooleanAttribute = 2
	UDateFormatBooleanAttributeUDATPARSEMULTIPLEPATTERNSFORMATCH UDateFormatBooleanAttribute = 3
	UDateFormatBooleanAttributeUDATBOOLEANATTRIBUTECOUNT         UDateFormatBooleanAttribute = 4
)

type UDateFormatHourCycle c.Int

const (
	UDateFormatHourCycleUDATHOURCYCLE11 UDateFormatHourCycle = 0
	UDateFormatHourCycleUDATHOURCYCLE12 UDateFormatHourCycle = 1
	UDateFormatHourCycleUDATHOURCYCLE23 UDateFormatHourCycle = 2
	UDateFormatHourCycleUDATHOURCYCLE24 UDateFormatHourCycle = 3
)

type UDateFormatSymbolType c.Int

const (
	UDateFormatSymbolTypeUDATERAS                      UDateFormatSymbolType = 0
	UDateFormatSymbolTypeUDATMONTHS                    UDateFormatSymbolType = 1
	UDateFormatSymbolTypeUDATSHORTMONTHS               UDateFormatSymbolType = 2
	UDateFormatSymbolTypeUDATWEEKDAYS                  UDateFormatSymbolType = 3
	UDateFormatSymbolTypeUDATSHORTWEEKDAYS             UDateFormatSymbolType = 4
	UDateFormatSymbolTypeUDATAMPMS                     UDateFormatSymbolType = 5
	UDateFormatSymbolTypeUDATLOCALIZEDCHARS            UDateFormatSymbolType = 6
	UDateFormatSymbolTypeUDATERANAMES                  UDateFormatSymbolType = 7
	UDateFormatSymbolTypeUDATNARROWMONTHS              UDateFormatSymbolType = 8
	UDateFormatSymbolTypeUDATNARROWWEEKDAYS            UDateFormatSymbolType = 9
	UDateFormatSymbolTypeUDATSTANDALONEMONTHS          UDateFormatSymbolType = 10
	UDateFormatSymbolTypeUDATSTANDALONESHORTMONTHS     UDateFormatSymbolType = 11
	UDateFormatSymbolTypeUDATSTANDALONENARROWMONTHS    UDateFormatSymbolType = 12
	UDateFormatSymbolTypeUDATSTANDALONEWEEKDAYS        UDateFormatSymbolType = 13
	UDateFormatSymbolTypeUDATSTANDALONESHORTWEEKDAYS   UDateFormatSymbolType = 14
	UDateFormatSymbolTypeUDATSTANDALONENARROWWEEKDAYS  UDateFormatSymbolType = 15
	UDateFormatSymbolTypeUDATQUARTERS                  UDateFormatSymbolType = 16
	UDateFormatSymbolTypeUDATSHORTQUARTERS             UDateFormatSymbolType = 17
	UDateFormatSymbolTypeUDATSTANDALONEQUARTERS        UDateFormatSymbolType = 18
	UDateFormatSymbolTypeUDATSTANDALONESHORTQUARTERS   UDateFormatSymbolType = 19
	UDateFormatSymbolTypeUDATSHORTERWEEKDAYS           UDateFormatSymbolType = 20
	UDateFormatSymbolTypeUDATSTANDALONESHORTERWEEKDAYS UDateFormatSymbolType = 21
	UDateFormatSymbolTypeUDATCYCLICYEARSWIDE           UDateFormatSymbolType = 22
	UDateFormatSymbolTypeUDATCYCLICYEARSABBREVIATED    UDateFormatSymbolType = 23
	UDateFormatSymbolTypeUDATCYCLICYEARSNARROW         UDateFormatSymbolType = 24
	UDateFormatSymbolTypeUDATZODIACNAMESWIDE           UDateFormatSymbolType = 25
	UDateFormatSymbolTypeUDATZODIACNAMESABBREVIATED    UDateFormatSymbolType = 26
	UDateFormatSymbolTypeUDATZODIACNAMESNARROW         UDateFormatSymbolType = 27
	UDateFormatSymbolTypeUDATNARROWQUARTERS            UDateFormatSymbolType = 28
	UDateFormatSymbolTypeUDATSTANDALONENARROWQUARTERS  UDateFormatSymbolType = 29
)

type UDateFormatSymbols struct {
	Unused [8]uint8
}
// llgo:type C
type UDateFormatOpener func(UDateFormatStyle, UDateFormatStyle, *int8, *UChar, int32, *UChar, int32, *UErrorCode) *UDateFormat
