package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const USHOWCPLUSPLUSAPI = 0
const USHOWCPLUSPLUSHEADERAPI = 0
const UICUDATATYPELETTER = "l"
const UUSEUSRDATA = 0

type UDate float64
type UErrorCode c.Int

const (
	UErrorCodeUUSINGFALLBACKWARNING             UErrorCode = -128
	UErrorCodeUERRORWARNINGSTART                UErrorCode = -128
	UErrorCodeUUSINGDEFAULTWARNING              UErrorCode = -127
	UErrorCodeUSAFECLONEALLOCATEDWARNING        UErrorCode = -126
	UErrorCodeUSTATEOLDWARNING                  UErrorCode = -125
	UErrorCodeUSTRINGNOTTERMINATEDWARNING       UErrorCode = -124
	UErrorCodeUSORTKEYTOOSHORTWARNING           UErrorCode = -123
	UErrorCodeUAMBIGUOUSALIASWARNING            UErrorCode = -122
	UErrorCodeUDIFFERENTUCAVERSION              UErrorCode = -121
	UErrorCodeUPLUGINCHANGEDLEVELWARNING        UErrorCode = -120
	UErrorCodeUERRORWARNINGLIMIT                UErrorCode = -119
	UErrorCodeUZEROERROR                        UErrorCode = 0
	UErrorCodeUILLEGALARGUMENTERROR             UErrorCode = 1
	UErrorCodeUMISSINGRESOURCEERROR             UErrorCode = 2
	UErrorCodeUINVALIDFORMATERROR               UErrorCode = 3
	UErrorCodeUFILEACCESSERROR                  UErrorCode = 4
	UErrorCodeUINTERNALPROGRAMERROR             UErrorCode = 5
	UErrorCodeUMESSAGEPARSEERROR                UErrorCode = 6
	UErrorCodeUMEMORYALLOCATIONERROR            UErrorCode = 7
	UErrorCodeUINDEXOUTOFBOUNDSERROR            UErrorCode = 8
	UErrorCodeUPARSEERROR                       UErrorCode = 9
	UErrorCodeUINVALIDCHARFOUND                 UErrorCode = 10
	UErrorCodeUTRUNCATEDCHARFOUND               UErrorCode = 11
	UErrorCodeUILLEGALCHARFOUND                 UErrorCode = 12
	UErrorCodeUINVALIDTABLEFORMAT               UErrorCode = 13
	UErrorCodeUINVALIDTABLEFILE                 UErrorCode = 14
	UErrorCodeUBUFFEROVERFLOWERROR              UErrorCode = 15
	UErrorCodeUUNSUPPORTEDERROR                 UErrorCode = 16
	UErrorCodeURESOURCETYPEMISMATCH             UErrorCode = 17
	UErrorCodeUILLEGALESCAPESEQUENCE            UErrorCode = 18
	UErrorCodeUUNSUPPORTEDESCAPESEQUENCE        UErrorCode = 19
	UErrorCodeUNOSPACEAVAILABLE                 UErrorCode = 20
	UErrorCodeUCENOTFOUNDERROR                  UErrorCode = 21
	UErrorCodeUPRIMARYTOOLONGERROR              UErrorCode = 22
	UErrorCodeUSTATETOOOLDERROR                 UErrorCode = 23
	UErrorCodeUTOOMANYALIASESERROR              UErrorCode = 24
	UErrorCodeUENUMOUTOFSYNCERROR               UErrorCode = 25
	UErrorCodeUINVARIANTCONVERSIONERROR         UErrorCode = 26
	UErrorCodeUINVALIDSTATEERROR                UErrorCode = 27
	UErrorCodeUCOLLATORVERSIONMISMATCH          UErrorCode = 28
	UErrorCodeUUSELESSCOLLATORERROR             UErrorCode = 29
	UErrorCodeUNOWRITEPERMISSION                UErrorCode = 30
	UErrorCodeUINPUTTOOLONGERROR                UErrorCode = 31
	UErrorCodeUSTANDARDERRORLIMIT               UErrorCode = 32
	UErrorCodeUBADVARIABLEDEFINITION            UErrorCode = 65536
	UErrorCodeUPARSEERRORSTART                  UErrorCode = 65536
	UErrorCodeUMALFORMEDRULE                    UErrorCode = 65537
	UErrorCodeUMALFORMEDSET                     UErrorCode = 65538
	UErrorCodeUMALFORMEDSYMBOLREFERENCE         UErrorCode = 65539
	UErrorCodeUMALFORMEDUNICODEESCAPE           UErrorCode = 65540
	UErrorCodeUMALFORMEDVARIABLEDEFINITION      UErrorCode = 65541
	UErrorCodeUMALFORMEDVARIABLEREFERENCE       UErrorCode = 65542
	UErrorCodeUMISMATCHEDSEGMENTDELIMITERS      UErrorCode = 65543
	UErrorCodeUMISPLACEDANCHORSTART             UErrorCode = 65544
	UErrorCodeUMISPLACEDCURSOROFFSET            UErrorCode = 65545
	UErrorCodeUMISPLACEDQUANTIFIER              UErrorCode = 65546
	UErrorCodeUMISSINGOPERATOR                  UErrorCode = 65547
	UErrorCodeUMISSINGSEGMENTCLOSE              UErrorCode = 65548
	UErrorCodeUMULTIPLEANTECONTEXTS             UErrorCode = 65549
	UErrorCodeUMULTIPLECURSORS                  UErrorCode = 65550
	UErrorCodeUMULTIPLEPOSTCONTEXTS             UErrorCode = 65551
	UErrorCodeUTRAILINGBACKSLASH                UErrorCode = 65552
	UErrorCodeUUNDEFINEDSEGMENTREFERENCE        UErrorCode = 65553
	UErrorCodeUUNDEFINEDVARIABLE                UErrorCode = 65554
	UErrorCodeUUNQUOTEDSPECIAL                  UErrorCode = 65555
	UErrorCodeUUNTERMINATEDQUOTE                UErrorCode = 65556
	UErrorCodeURULEMASKERROR                    UErrorCode = 65557
	UErrorCodeUMISPLACEDCOMPOUNDFILTER          UErrorCode = 65558
	UErrorCodeUMULTIPLECOMPOUNDFILTERS          UErrorCode = 65559
	UErrorCodeUINVALIDRBTSYNTAX                 UErrorCode = 65560
	UErrorCodeUINVALIDPROPERTYPATTERN           UErrorCode = 65561
	UErrorCodeUMALFORMEDPRAGMA                  UErrorCode = 65562
	UErrorCodeUUNCLOSEDSEGMENT                  UErrorCode = 65563
	UErrorCodeUILLEGALCHARINSEGMENT             UErrorCode = 65564
	UErrorCodeUVARIABLERANGEEXHAUSTED           UErrorCode = 65565
	UErrorCodeUVARIABLERANGEOVERLAP             UErrorCode = 65566
	UErrorCodeUILLEGALCHARACTER                 UErrorCode = 65567
	UErrorCodeUINTERNALTRANSLITERATORERROR      UErrorCode = 65568
	UErrorCodeUINVALIDID                        UErrorCode = 65569
	UErrorCodeUINVALIDFUNCTION                  UErrorCode = 65570
	UErrorCodeUPARSEERRORLIMIT                  UErrorCode = 65571
	UErrorCodeUUNEXPECTEDTOKEN                  UErrorCode = 65792
	UErrorCodeUFMTPARSEERRORSTART               UErrorCode = 65792
	UErrorCodeUMULTIPLEDECIMALSEPARATORS        UErrorCode = 65793
	UErrorCodeUMULTIPLEDECIMALSEPERATORS        UErrorCode = 65793
	UErrorCodeUMULTIPLEEXPONENTIALSYMBOLS       UErrorCode = 65794
	UErrorCodeUMALFORMEDEXPONENTIALPATTERN      UErrorCode = 65795
	UErrorCodeUMULTIPLEPERCENTSYMBOLS           UErrorCode = 65796
	UErrorCodeUMULTIPLEPERMILLSYMBOLS           UErrorCode = 65797
	UErrorCodeUMULTIPLEPADSPECIFIERS            UErrorCode = 65798
	UErrorCodeUPATTERNSYNTAXERROR               UErrorCode = 65799
	UErrorCodeUILLEGALPADPOSITION               UErrorCode = 65800
	UErrorCodeUUNMATCHEDBRACES                  UErrorCode = 65801
	UErrorCodeUUNSUPPORTEDPROPERTY              UErrorCode = 65802
	UErrorCodeUUNSUPPORTEDATTRIBUTE             UErrorCode = 65803
	UErrorCodeUARGUMENTTYPEMISMATCH             UErrorCode = 65804
	UErrorCodeUDUPLICATEKEYWORD                 UErrorCode = 65805
	UErrorCodeUUNDEFINEDKEYWORD                 UErrorCode = 65806
	UErrorCodeUDEFAULTKEYWORDMISSING            UErrorCode = 65807
	UErrorCodeUDECIMALNUMBERSYNTAXERROR         UErrorCode = 65808
	UErrorCodeUFORMATINEXACTERROR               UErrorCode = 65809
	UErrorCodeUNUMBERARGOUTOFBOUNDSERROR        UErrorCode = 65810
	UErrorCodeUNUMBERSKELETONSYNTAXERROR        UErrorCode = 65811
	UErrorCodeUMFUNRESOLVEDVARIABLEERROR        UErrorCode = 65812
	UErrorCodeUMFSYNTAXERROR                    UErrorCode = 65813
	UErrorCodeUMFUNKNOWNFUNCTIONERROR           UErrorCode = 65814
	UErrorCodeUMFVARIANTKEYMISMATCHERROR        UErrorCode = 65815
	UErrorCodeUMFFORMATTINGERROR                UErrorCode = 65816
	UErrorCodeUMFNONEXHAUSTIVEPATTERNERROR      UErrorCode = 65817
	UErrorCodeUMFDUPLICATEOPTIONNAMEERROR       UErrorCode = 65818
	UErrorCodeUMFSELECTORERROR                  UErrorCode = 65819
	UErrorCodeUMFMISSINGSELECTORANNOTATIONERROR UErrorCode = 65820
	UErrorCodeUMFDUPLICATEDECLARATIONERROR      UErrorCode = 65821
	UErrorCodeUMFOPERANDMISMATCHERROR           UErrorCode = 65822
	UErrorCodeUMFDUPLICATEVARIANTERROR          UErrorCode = 65823
	UErrorCodeUFMTPARSEERRORLIMIT               UErrorCode = 65824
	UErrorCodeUBRKINTERNALERROR                 UErrorCode = 66048
	UErrorCodeUBRKERRORSTART                    UErrorCode = 66048
	UErrorCodeUBRKHEXDIGITSEXPECTED             UErrorCode = 66049
	UErrorCodeUBRKSEMICOLONEXPECTED             UErrorCode = 66050
	UErrorCodeUBRKRULESYNTAX                    UErrorCode = 66051
	UErrorCodeUBRKUNCLOSEDSET                   UErrorCode = 66052
	UErrorCodeUBRKASSIGNERROR                   UErrorCode = 66053
	UErrorCodeUBRKVARIABLEREDFINITION           UErrorCode = 66054
	UErrorCodeUBRKMISMATCHEDPAREN               UErrorCode = 66055
	UErrorCodeUBRKNEWLINEINQUOTEDSTRING         UErrorCode = 66056
	UErrorCodeUBRKUNDEFINEDVARIABLE             UErrorCode = 66057
	UErrorCodeUBRKINITERROR                     UErrorCode = 66058
	UErrorCodeUBRKRULEEMPTYSET                  UErrorCode = 66059
	UErrorCodeUBRKUNRECOGNIZEDOPTION            UErrorCode = 66060
	UErrorCodeUBRKMALFORMEDRULETAG              UErrorCode = 66061
	UErrorCodeUBRKERRORLIMIT                    UErrorCode = 66062
	UErrorCodeUREGEXINTERNALERROR               UErrorCode = 66304
	UErrorCodeUREGEXERRORSTART                  UErrorCode = 66304
	UErrorCodeUREGEXRULESYNTAX                  UErrorCode = 66305
	UErrorCodeUREGEXINVALIDSTATE                UErrorCode = 66306
	UErrorCodeUREGEXBADESCAPESEQUENCE           UErrorCode = 66307
	UErrorCodeUREGEXPROPERTYSYNTAX              UErrorCode = 66308
	UErrorCodeUREGEXUNIMPLEMENTED               UErrorCode = 66309
	UErrorCodeUREGEXMISMATCHEDPAREN             UErrorCode = 66310
	UErrorCodeUREGEXNUMBERTOOBIG                UErrorCode = 66311
	UErrorCodeUREGEXBADINTERVAL                 UErrorCode = 66312
	UErrorCodeUREGEXMAXLTMIN                    UErrorCode = 66313
	UErrorCodeUREGEXINVALIDBACKREF              UErrorCode = 66314
	UErrorCodeUREGEXINVALIDFLAG                 UErrorCode = 66315
	UErrorCodeUREGEXLOOKBEHINDLIMIT             UErrorCode = 66316
	UErrorCodeUREGEXSETCONTAINSSTRING           UErrorCode = 66317
	UErrorCodeUREGEXOCTALTOOBIG                 UErrorCode = 66318
	UErrorCodeUREGEXMISSINGCLOSEBRACKET         UErrorCode = 66319
	UErrorCodeUREGEXINVALIDRANGE                UErrorCode = 66320
	UErrorCodeUREGEXSTACKOVERFLOW               UErrorCode = 66321
	UErrorCodeUREGEXTIMEOUT                     UErrorCode = 66322
	UErrorCodeUREGEXSTOPPEDBYCALLER             UErrorCode = 66323
	UErrorCodeUREGEXPATTERNTOOBIG               UErrorCode = 66324
	UErrorCodeUREGEXINVALIDCAPTUREGROUPNAME     UErrorCode = 66325
	UErrorCodeUREGEXERRORLIMIT                  UErrorCode = 66326
	UErrorCodeUIDNAPROHIBITEDERROR              UErrorCode = 66560
	UErrorCodeUIDNAERRORSTART                   UErrorCode = 66560
	UErrorCodeUIDNAUNASSIGNEDERROR              UErrorCode = 66561
	UErrorCodeUIDNACHECKBIDIERROR               UErrorCode = 66562
	UErrorCodeUIDNASTD3ASCIIRULESERROR          UErrorCode = 66563
	UErrorCodeUIDNAACEPREFIXERROR               UErrorCode = 66564
	UErrorCodeUIDNAVERIFICATIONERROR            UErrorCode = 66565
	UErrorCodeUIDNALABELTOOLONGERROR            UErrorCode = 66566
	UErrorCodeUIDNAZEROLENGTHLABELERROR         UErrorCode = 66567
	UErrorCodeUIDNADOMAINNAMETOOLONGERROR       UErrorCode = 66568
	UErrorCodeUIDNAERRORLIMIT                   UErrorCode = 66569
	UErrorCodeUSTRINGPREPPROHIBITEDERROR        UErrorCode = 66560
	UErrorCodeUSTRINGPREPUNASSIGNEDERROR        UErrorCode = 66561
	UErrorCodeUSTRINGPREPCHECKBIDIERROR         UErrorCode = 66562
	UErrorCodeUPLUGINERRORSTART                 UErrorCode = 66816
	UErrorCodeUPLUGINTOOHIGH                    UErrorCode = 66816
	UErrorCodeUPLUGINDIDNTSETLEVEL              UErrorCode = 66817
	UErrorCodeUPLUGINERRORLIMIT                 UErrorCode = 66818
	UErrorCodeUERRORLIMIT                       UErrorCode = 66818
)
