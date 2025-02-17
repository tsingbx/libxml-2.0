package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const UBIDIDEFAULTLTR = 0xfe
const UBIDIDEFAULTRTL = 0xff
const UBIDIMAXEXPLICITLEVEL = 125
const UBIDILEVELOVERRIDE = 0x80
const UBIDIKEEPBASECOMBINING = 1
const UBIDIDOMIRRORING = 2
const UBIDIINSERTLRMFORNUMERIC = 4
const UBIDIREMOVEBIDICONTROLS = 8
const UBIDIOUTPUTREVERSE = 16

type UBiDiLevel uint8
type UBiDiDirection c.Int

const (
	UBiDiDirectionUBIDILTR     UBiDiDirection = 0
	UBiDiDirectionUBIDIRTL     UBiDiDirection = 1
	UBiDiDirectionUBIDIMIXED   UBiDiDirection = 2
	UBiDiDirectionUBIDINEUTRAL UBiDiDirection = 3
)
/**
 * Forward declaration of the <code>UBiDi</code> structure for the declaration of
 * the API functions. Its fields are implementation-specific.<p>
 * This structure holds information about a paragraph (or multiple paragraphs)
 * of text with Bidi-algorithm-related details, or about one line of
 * such a paragraph.<p>
 * Reordering can be done on a line, or on one or more paragraphs which are
 * then interpreted each as one single line.
 * @stable ICU 2.0
 */

type UBiDi struct {
	Unused [8]uint8
}
type UBiDiReorderingMode c.Int

const (
	UBiDiReorderingModeUBIDIREORDERDEFAULT                  UBiDiReorderingMode = 0
	UBiDiReorderingModeUBIDIREORDERNUMBERSSPECIAL           UBiDiReorderingMode = 1
	UBiDiReorderingModeUBIDIREORDERGROUPNUMBERSWITHR        UBiDiReorderingMode = 2
	UBiDiReorderingModeUBIDIREORDERRUNSONLY                 UBiDiReorderingMode = 3
	UBiDiReorderingModeUBIDIREORDERINVERSENUMBERSASL        UBiDiReorderingMode = 4
	UBiDiReorderingModeUBIDIREORDERINVERSELIKEDIRECT        UBiDiReorderingMode = 5
	UBiDiReorderingModeUBIDIREORDERINVERSEFORNUMBERSSPECIAL UBiDiReorderingMode = 6
	UBiDiReorderingModeUBIDIREORDERCOUNT                    UBiDiReorderingMode = 7
)

type UBiDiReorderingOption c.Int

const (
	UBiDiReorderingOptionUBIDIOPTIONDEFAULT        UBiDiReorderingOption = 0
	UBiDiReorderingOptionUBIDIOPTIONINSERTMARKS    UBiDiReorderingOption = 1
	UBiDiReorderingOptionUBIDIOPTIONREMOVECONTROLS UBiDiReorderingOption = 2
	UBiDiReorderingOptionUBIDIOPTIONSTREAMING      UBiDiReorderingOption = 4
)
// llgo:type C
type UBiDiClassCallback func(unsafe.Pointer, UChar32) UCharDirection
