package libc

import unsafe "unsafe"

type errmsgstr_t_cgo19_strerror struct {
	str0               [21]int8
	strEILSEQ          [22]int8
	strEDOM            [13]int8
	strERANGE          [25]int8
	strENOTTY          [10]int8
	strEACCES          [18]int8
	strEPERM           [24]int8
	strENOENT          [26]int8
	strESRCH           [16]int8
	strEEXIST          [12]int8
	strEOVERFLOW       [30]int8
	strENOSPC          [24]int8
	strENOMEM          [14]int8
	strEBUSY           [14]int8
	strEINTR           [24]int8
	strEAGAIN          [33]int8
	strESPIPE          [13]int8
	strEXDEV           [18]int8
	strEROFS           [22]int8
	strENOTEMPTY       [20]int8
	strECONNRESET      [25]int8
	strETIMEDOUT       [20]int8
	strECONNREFUSED    [19]int8
	strEHOSTDOWN       [13]int8
	strEHOSTUNREACH    [20]int8
	strEADDRINUSE      [15]int8
	strEPIPE           [12]int8
	strEIO             [10]int8
	strENXIO           [26]int8
	strENOTBLK         [22]int8
	strENODEV          [15]int8
	strENOTDIR         [16]int8
	strEISDIR          [15]int8
	strETXTBSY         [15]int8
	strENOEXEC         [18]int8
	strEINVAL          [17]int8
	strE2BIG           [23]int8
	strELOOP           [19]int8
	strENAMETOOLONG    [18]int8
	strENFILE          [30]int8
	strEMFILE          [30]int8
	strEBADF           [20]int8
	strECHILD          [17]int8
	strEFAULT          [12]int8
	strEFBIG           [15]int8
	strEMLINK          [15]int8
	strENOLCK          [19]int8
	strEDEADLK         [30]int8
	strENOTRECOVERABLE [22]int8
	strEOWNERDEAD      [20]int8
	strECANCELED       [19]int8
	strENOSYS          [25]int8
	strENOMSG          [27]int8
	strEIDRM           [19]int8
	strENOSTR          [20]int8
	strENODATA         [18]int8
	strETIME           [15]int8
	strENOSR           [25]int8
	strENOLINK         [22]int8
	strEPROTO          [15]int8
	strEBADMSG         [12]int8
	strEBADFD          [29]int8
	strENOTSOCK        [13]int8
	strEDESTADDRREQ    [29]int8
	strEMSGSIZE        [18]int8
	strEPROTOTYPE      [31]int8
	strENOPROTOOPT     [23]int8
	strEPROTONOSUPPORT [23]int8
	strESOCKTNOSUPPORT [26]int8
	strENOTSUP         [14]int8
	strEPFNOSUPPORT    [30]int8
	strEAFNOSUPPORT    [41]int8
	strEADDRNOTAVAIL   [22]int8
	strENETDOWN        [16]int8
	strENETUNREACH     [20]int8
	strENETRESET       [28]int8
	strECONNABORTED    [19]int8
	strENOBUFS         [26]int8
	strEISCONN         [20]int8
	strENOTCONN        [21]int8
	strESHUTDOWN       [34]int8
	strEALREADY        [30]int8
	strEINPROGRESS     [22]int8
	strESTALE          [18]int8
	strEREMOTEIO       [17]int8
	strEDQUOT          [15]int8
	strENOMEDIUM       [16]int8
	strEMEDIUMTYPE     [18]int8
	strEMULTIHOP       [19]int8
}

var errmsgstr_cgo20_strerror errmsgstr_t_cgo19_strerror = errmsgstr_t_cgo19_strerror{[21]int8{'N', 'o', ' ', 'e', 'r', 'r', 'o', 'r', ' ', 'i', 'n', 'f', 'o', 'r', 'm', 'a', 't', 'i', 'o', 'n', '\x00'}, [22]int8{'I', 'l', 'l', 'e', 'g', 'a', 'l', ' ', 'b', 'y', 't', 'e', ' ', 's', 'e', 'q', 'u', 'e', 'n', 'c', 'e', '\x00'}, [13]int8{'D', 'o', 'm', 'a', 'i', 'n', ' ', 'e', 'r', 'r', 'o', 'r', '\x00'}, [25]int8{'R', 'e', 's', 'u', 'l', 't', ' ', 'n', 'o', 't', ' ', 'r', 'e', 'p', 'r', 'e', 's', 'e', 'n', 't', 'a', 'b', 'l', 'e', '\x00'}, [10]int8{'N', 'o', 't', ' ', 'a', ' ', 't', 't', 'y', '\x00'}, [18]int8{'P', 'e', 'r', 'm', 'i', 's', 's', 'i', 'o', 'n', ' ', 'd', 'e', 'n', 'i', 'e', 'd', '\x00'}, [24]int8{'O', 'p', 'e', 'r', 'a', 't', 'i', 'o', 'n', ' ', 'n', 'o', 't', ' ', 'p', 'e', 'r', 'm', 'i', 't', 't', 'e', 'd', '\x00'}, [26]int8{'N', 'o', ' ', 's', 'u', 'c', 'h', ' ', 'f', 'i', 'l', 'e', ' ', 'o', 'r', ' ', 'd', 'i', 'r', 'e', 'c', 't', 'o', 'r', 'y', '\x00'}, [16]int8{'N', 'o', ' ', 's', 'u', 'c', 'h', ' ', 'p', 'r', 'o', 'c', 'e', 's', 's', '\x00'}, [12]int8{'F', 'i', 'l', 'e', ' ', 'e', 'x', 'i', 's', 't', 's', '\x00'}, [30]int8{'V', 'a', 'l', 'u', 'e', ' ', 't', 'o', 'o', ' ', 'l', 'a', 'r', 'g', 'e', ' ', 'f', 'o', 'r', ' ', 'd', 'a', 't', 'a', ' ', 't', 'y', 'p', 'e', '\x00'}, [24]int8{'N', 'o', ' ', 's', 'p', 'a', 'c', 'e', ' ', 'l', 'e', 'f', 't', ' ', 'o', 'n', ' ', 'd', 'e', 'v', 'i', 'c', 'e', '\x00'}, [14]int8{'O', 'u', 't', ' ', 'o', 'f', ' ', 'm', 'e', 'm', 'o', 'r', 'y', '\x00'}, [14]int8{'R', 'e', 's', 'o', 'u', 'r', 'c', 'e', ' ', 'b', 'u', 's', 'y', '\x00'}, [24]int8{'I', 'n', 't', 'e', 'r', 'r', 'u', 'p', 't', 'e', 'd', ' ', 's', 'y', 's', 't', 'e', 'm', ' ', 'c', 'a', 'l', 'l', '\x00'}, [33]int8{'R', 'e', 's', 'o', 'u', 'r', 'c', 'e', ' ', 't', 'e', 'm', 'p', 'o', 'r', 'a', 'r', 'i', 'l', 'y', ' ', 'u', 'n', 'a', 'v', 'a', 'i', 'l', 'a', 'b', 'l', 'e', '\x00'}, [13]int8{'I', 'n', 'v', 'a', 'l', 'i', 'd', ' ', 's', 'e', 'e', 'k', '\x00'}, [18]int8{'C', 'r', 'o', 's', 's', '-', 'd', 'e', 'v', 'i', 'c', 'e', ' ', 'l', 'i', 'n', 'k', '\x00'}, [22]int8{'R', 'e', 'a', 'd', '-', 'o', 'n', 'l', 'y', ' ', 'f', 'i', 'l', 'e', ' ', 's', 'y', 's', 't', 'e', 'm', '\x00'}, [20]int8{'D', 'i', 'r', 'e', 'c', 't', 'o', 'r', 'y', ' ', 'n', 'o', 't', ' ', 'e', 'm', 'p', 't', 'y', '\x00'}, [25]int8{'C', 'o', 'n', 'n', 'e', 'c', 't', 'i', 'o', 'n', ' ', 'r', 'e', 's', 'e', 't', ' ', 'b', 'y', ' ', 'p', 'e', 'e', 'r', '\x00'}, [20]int8{'O', 'p', 'e', 'r', 'a', 't', 'i', 'o', 'n', ' ', 't', 'i', 'm', 'e', 'd', ' ', 'o', 'u', 't', '\x00'}, [19]int8{'C', 'o', 'n', 'n', 'e', 'c', 't', 'i', 'o', 'n', ' ', 'r', 'e', 'f', 'u', 's', 'e', 'd', '\x00'}, [13]int8{'H', 'o', 's', 't', ' ', 'i', 's', ' ', 'd', 'o', 'w', 'n', '\x00'}, [20]int8{'H', 'o', 's', 't', ' ', 'i', 's', ' ', 'u', 'n', 'r', 'e', 'a', 'c', 'h', 'a', 'b', 'l', 'e', '\x00'}, [15]int8{'A', 'd', 'd', 'r', 'e', 's', 's', ' ', 'i', 'n', ' ', 'u', 's', 'e', '\x00'}, [12]int8{'B', 'r', 'o', 'k', 'e', 'n', ' ', 'p', 'i', 'p', 'e', '\x00'}, [10]int8{'I', '/', 'O', ' ', 'e', 'r', 'r', 'o', 'r', '\x00'}, [26]int8{'N', 'o', ' ', 's', 'u', 'c', 'h', ' ', 'd', 'e', 'v', 'i', 'c', 'e', ' ', 'o', 'r', ' ', 'a', 'd', 'd', 'r', 'e', 's', 's', '\x00'}, [22]int8{'B', 'l', 'o', 'c', 'k', ' ', 'd', 'e', 'v', 'i', 'c', 'e', ' ', 'r', 'e', 'q', 'u', 'i', 'r', 'e', 'd', '\x00'}, [15]int8{'N', 'o', ' ', 's', 'u', 'c', 'h', ' ', 'd', 'e', 'v', 'i', 'c', 'e', '\x00'}, [16]int8{'N', 'o', 't', ' ', 'a', ' ', 'd', 'i', 'r', 'e', 'c', 't', 'o', 'r', 'y', '\x00'}, [15]int8{'I', 's', ' ', 'a', ' ', 'd', 'i', 'r', 'e', 'c', 't', 'o', 'r', 'y', '\x00'}, [15]int8{'T', 'e', 'x', 't', ' ', 'f', 'i', 'l', 'e', ' ', 'b', 'u', 's', 'y', '\x00'}, [18]int8{'E', 'x', 'e', 'c', ' ', 'f', 'o', 'r', 'm', 'a', 't', ' ', 'e', 'r', 'r', 'o', 'r', '\x00'}, [17]int8{'I', 'n', 'v', 'a', 'l', 'i', 'd', ' ', 'a', 'r', 'g', 'u', 'm', 'e', 'n', 't', '\x00'}, [23]int8{'A', 'r', 'g', 'u', 'm', 'e', 'n', 't', ' ', 'l', 'i', 's', 't', ' ', 't', 'o', 'o', ' ', 'l', 'o', 'n', 'g', '\x00'}, [19]int8{'S', 'y', 'm', 'b', 'o', 'l', 'i', 'c', ' ', 'l', 'i', 'n', 'k', ' ', 'l', 'o', 'o', 'p', '\x00'}, [18]int8{'F', 'i', 'l', 'e', 'n', 'a', 'm', 'e', ' ', 't', 'o', 'o', ' ', 'l', 'o', 'n', 'g', '\x00'}, [30]int8{'T', 'o', 'o', ' ', 'm', 'a', 'n', 'y', ' ', 'o', 'p', 'e', 'n', ' ', 'f', 'i', 'l', 'e', 's', ' ', 'i', 'n', ' ', 's', 'y', 's', 't', 'e', 'm', '\x00'}, [30]int8{'N', 'o', ' ', 'f', 'i', 'l', 'e', ' ', 'd', 'e', 's', 'c', 'r', 'i', 'p', 't', 'o', 'r', 's', ' ', 'a', 'v', 'a', 'i', 'l', 'a', 'b', 'l', 'e', '\x00'}, [20]int8{'B', 'a', 'd', ' ', 'f', 'i', 'l', 'e', ' ', 'd', 'e', 's', 'c', 'r', 'i', 'p', 't', 'o', 'r', '\x00'}, [17]int8{'N', 'o', ' ', 'c', 'h', 'i', 'l', 'd', ' ', 'p', 'r', 'o', 'c', 'e', 's', 's', '\x00'}, [12]int8{'B', 'a', 'd', ' ', 'a', 'd', 'd', 'r', 'e', 's', 's', '\x00'}, [15]int8{'F', 'i', 'l', 'e', ' ', 't', 'o', 'o', ' ', 'l', 'a', 'r', 'g', 'e', '\x00'}, [15]int8{'T', 'o', 'o', ' ', 'm', 'a', 'n', 'y', ' ', 'l', 'i', 'n', 'k', 's', '\x00'}, [19]int8{'N', 'o', ' ', 'l', 'o', 'c', 'k', 's', ' ', 'a', 'v', 'a', 'i', 'l', 'a', 'b', 'l', 'e', '\x00'}, [30]int8{'R', 'e', 's', 'o', 'u', 'r', 'c', 'e', ' ', 'd', 'e', 'a', 'd', 'l', 'o', 'c', 'k', ' ', 'w', 'o', 'u', 'l', 'd', ' ', 'o', 'c', 'c', 'u', 'r', '\x00'}, [22]int8{'S', 't', 'a', 't', 'e', ' ', 'n', 'o', 't', ' ', 'r', 'e', 'c', 'o', 'v', 'e', 'r', 'a', 'b', 'l', 'e', '\x00'}, [20]int8{'P', 'r', 'e', 'v', 'i', 'o', 'u', 's', ' ', 'o', 'w', 'n', 'e', 'r', ' ', 'd', 'i', 'e', 'd', '\x00'}, [19]int8{'O', 'p', 'e', 'r', 'a', 't', 'i', 'o', 'n', ' ', 'c', 'a', 'n', 'c', 'e', 'l', 'e', 'd', '\x00'}, [25]int8{'F', 'u', 'n', 'c', 't', 'i', 'o', 'n', ' ', 'n', 'o', 't', ' ', 'i', 'm', 'p', 'l', 'e', 'm', 'e', 'n', 't', 'e', 'd', '\x00'}, [27]int8{'N', 'o', ' ', 'm', 'e', 's', 's', 'a', 'g', 'e', ' ', 'o', 'f', ' ', 'd', 'e', 's', 'i', 'r', 'e', 'd', ' ', 't', 'y', 'p', 'e', '\x00'}, [19]int8{'I', 'd', 'e', 'n', 't', 'i', 'f', 'i', 'e', 'r', ' ', 'r', 'e', 'm', 'o', 'v', 'e', 'd', '\x00'}, [20]int8{'D', 'e', 'v', 'i', 'c', 'e', ' ', 'n', 'o', 't', ' ', 'a', ' ', 's', 't', 'r', 'e', 'a', 'm', '\x00'}, [18]int8{'N', 'o', ' ', 'd', 'a', 't', 'a', ' ', 'a', 'v', 'a', 'i', 'l', 'a', 'b', 'l', 'e', '\x00'}, [15]int8{'D', 'e', 'v', 'i', 'c', 'e', ' ', 't', 'i', 'm', 'e', 'o', 'u', 't', '\x00'}, [25]int8{'O', 'u', 't', ' ', 'o', 'f', ' ', 's', 't', 'r', 'e', 'a', 'm', 's', ' ', 'r', 'e', 's', 'o', 'u', 'r', 'c', 'e', 's', '\x00'}, [22]int8{'L', 'i', 'n', 'k', ' ', 'h', 'a', 's', ' ', 'b', 'e', 'e', 'n', ' ', 's', 'e', 'v', 'e', 'r', 'e', 'd', '\x00'}, [15]int8{'P', 'r', 'o', 't', 'o', 'c', 'o', 'l', ' ', 'e', 'r', 'r', 'o', 'r', '\x00'}, [12]int8{'B', 'a', 'd', ' ', 'm', 'e', 's', 's', 'a', 'g', 'e', '\x00'}, [29]int8{'F', 'i', 'l', 'e', ' ', 'd', 'e', 's', 'c', 'r', 'i', 'p', 't', 'o', 'r', ' ', 'i', 'n', ' ', 'b', 'a', 'd', ' ', 's', 't', 'a', 't', 'e', '\x00'}, [13]int8{'N', 'o', 't', ' ', 'a', ' ', 's', 'o', 'c', 'k', 'e', 't', '\x00'}, [29]int8{'D', 'e', 's', 't', 'i', 'n', 'a', 't', 'i', 'o', 'n', ' ', 'a', 'd', 'd', 'r', 'e', 's', 's', ' ', 'r', 'e', 'q', 'u', 'i', 'r', 'e', 'd', '\x00'}, [18]int8{'M', 'e', 's', 's', 'a', 'g', 'e', ' ', 't', 'o', 'o', ' ', 'l', 'a', 'r', 'g', 'e', '\x00'}, [31]int8{'P', 'r', 'o', 't', 'o', 'c', 'o', 'l', ' ', 'w', 'r', 'o', 'n', 'g', ' ', 't', 'y', 'p', 'e', ' ', 'f', 'o', 'r', ' ', 's', 'o', 'c', 'k', 'e', 't', '\x00'}, [23]int8{'P', 'r', 'o', 't', 'o', 'c', 'o', 'l', ' ', 'n', 'o', 't', ' ', 'a', 'v', 'a', 'i', 'l', 'a', 'b', 'l', 'e', '\x00'}, [23]int8{'P', 'r', 'o', 't', 'o', 'c', 'o', 'l', ' ', 'n', 'o', 't', ' ', 's', 'u', 'p', 'p', 'o', 'r', 't', 'e', 'd', '\x00'}, [26]int8{'S', 'o', 'c', 'k', 'e', 't', ' ', 't', 'y', 'p', 'e', ' ', 'n', 'o', 't', ' ', 's', 'u', 'p', 'p', 'o', 'r', 't', 'e', 'd', '\x00'}, [14]int8{'N', 'o', 't', ' ', 's', 'u', 'p', 'p', 'o', 'r', 't', 'e', 'd', '\x00'}, [30]int8{'P', 'r', 'o', 't', 'o', 'c', 'o', 'l', ' ', 'f', 'a', 'm', 'i', 'l', 'y', ' ', 'n', 'o', 't', ' ', 's', 'u', 'p', 'p', 'o', 'r', 't', 'e', 'd', '\x00'}, [41]int8{'A', 'd', 'd', 'r', 'e', 's', 's', ' ', 'f', 'a', 'm', 'i', 'l', 'y', ' ', 'n', 'o', 't', ' ', 's', 'u', 'p', 'p', 'o', 'r', 't', 'e', 'd', ' ', 'b', 'y', ' ', 'p', 'r', 'o', 't', 'o', 'c', 'o', 'l', '\x00'}, [22]int8{'A', 'd', 'd', 'r', 'e', 's', 's', ' ', 'n', 'o', 't', ' ', 'a', 'v', 'a', 'i', 'l', 'a', 'b', 'l', 'e', '\x00'}, [16]int8{'N', 'e', 't', 'w', 'o', 'r', 'k', ' ', 'i', 's', ' ', 'd', 'o', 'w', 'n', '\x00'}, [20]int8{'N', 'e', 't', 'w', 'o', 'r', 'k', ' ', 'u', 'n', 'r', 'e', 'a', 'c', 'h', 'a', 'b', 'l', 'e', '\x00'}, [28]int8{'C', 'o', 'n', 'n', 'e', 'c', 't', 'i', 'o', 'n', ' ', 'r', 'e', 's', 'e', 't', ' ', 'b', 'y', ' ', 'n', 'e', 't', 'w', 'o', 'r', 'k', '\x00'}, [19]int8{'C', 'o', 'n', 'n', 'e', 'c', 't', 'i', 'o', 'n', ' ', 'a', 'b', 'o', 'r', 't', 'e', 'd', '\x00'}, [26]int8{'N', 'o', ' ', 'b', 'u', 'f', 'f', 'e', 'r', ' ', 's', 'p', 'a', 'c', 'e', ' ', 'a', 'v', 'a', 'i', 'l', 'a', 'b', 'l', 'e', '\x00'}, [20]int8{'S', 'o', 'c', 'k', 'e', 't', ' ', 'i', 's', ' ', 'c', 'o', 'n', 'n', 'e', 'c', 't', 'e', 'd', '\x00'}, [21]int8{'S', 'o', 'c', 'k', 'e', 't', ' ', 'n', 'o', 't', ' ', 'c', 'o', 'n', 'n', 'e', 'c', 't', 'e', 'd', '\x00'}, [34]int8{'C', 'a', 'n', 'n', 'o', 't', ' ', 's', 'e', 'n', 'd', ' ', 'a', 'f', 't', 'e', 'r', ' ', 's', 'o', 'c', 'k', 'e', 't', ' ', 's', 'h', 'u', 't', 'd', 'o', 'w', 'n', '\x00'}, [30]int8{'O', 'p', 'e', 'r', 'a', 't', 'i', 'o', 'n', ' ', 'a', 'l', 'r', 'e', 'a', 'd', 'y', ' ', 'i', 'n', ' ', 'p', 'r', 'o', 'g', 'r', 'e', 's', 's', '\x00'}, [22]int8{'O', 'p', 'e', 'r', 'a', 't', 'i', 'o', 'n', ' ', 'i', 'n', ' ', 'p', 'r', 'o', 'g', 'r', 'e', 's', 's', '\x00'}, [18]int8{'S', 't', 'a', 'l', 'e', ' ', 'f', 'i', 'l', 'e', ' ', 'h', 'a', 'n', 'd', 'l', 'e', '\x00'}, [17]int8{'R', 'e', 'm', 'o', 't', 'e', ' ', 'I', '/', 'O', ' ', 'e', 'r', 'r', 'o', 'r', '\x00'}, [15]int8{'Q', 'u', 'o', 't', 'a', ' ', 'e', 'x', 'c', 'e', 'e', 'd', 'e', 'd', '\x00'}, [16]int8{'N', 'o', ' ', 'm', 'e', 'd', 'i', 'u', 'm', ' ', 'f', 'o', 'u', 'n', 'd', '\x00'}, [18]int8{'W', 'r', 'o', 'n', 'g', ' ', 'm', 'e', 'd', 'i', 'u', 'm', ' ', 't', 'y', 'p', 'e', '\x00'}, [19]int8{'M', 'u', 'l', 't', 'i', 'h', 'o', 'p', ' ', 'a', 't', 't', 'e', 'm', 'p', 't', 'e', 'd', '\x00'}}
var errmsgidx_cgo21_strerror [132]uint16 = [132]uint16{uint16(0), uint16(109), uint16(133), uint16(159), uint16(269), uint16(523), uint16(533), uint16(677), uint16(642), uint16(797), uint16(817), uint16(293), uint16(241), uint16(91), uint16(834), uint16(559), uint16(255), uint16(175), uint16(339), uint16(581), uint16(596), uint16(612), uint16(660), uint16(737), uint16(767), uint16(81), uint16(627), uint16(846), uint16(217), uint16(326), uint16(357), uint16(861), uint16(511), uint16(43), uint16(56), uint16(895), uint16(719), uint16(876), uint16(986), uint16(379), uint16(700), 0, uint16(1011), uint16(1038), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, uint16(1057), uint16(1077), uint16(1095), uint16(1110), 0, 0, 0, uint16(1135), 0, 0, 0, uint16(1157), uint16(1803), 0, uint16(1172), uint16(187), 0, uint16(1184), 0, 0, 0, 0, 0, 0, uint16(21), 0, 0, 0, uint16(1213), uint16(1226), uint16(1255), uint16(1273), uint16(1304), uint16(1327), uint16(1350), uint16(1376), uint16(1390), uint16(1420), uint16(496), uint16(1461), uint16(1483), uint16(1499), uint16(1519), uint16(1547), uint16(399), uint16(1566), uint16(1592), uint16(1612), uint16(1633), 0, uint16(424), uint16(444), uint16(463), uint16(476), uint16(1667), uint16(1697), uint16(1719), 0, 0, 0, 0, uint16(1737), uint16(1754), uint16(1769), uint16(1785), uint16(967), 0, 0, 0, 0, uint16(947), uint16(925)}

func __strerror_l(e int32, loc *struct___locale_struct) *int8 {
	var s *int8
	if uint64(e) >= 132 {
		e = int32(0)
	}
	s = (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&errmsgstr_cgo20_strerror)))) + uintptr(int32(*(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint16)(unsafe.Pointer(&errmsgidx_cgo21_strerror)))) + uintptr(e)*2))))))
	return (*int8)(unsafe.Pointer(__lctrans(s, *(**struct___locale_map)(unsafe.Pointer(uintptr(unsafe.Pointer((**struct___locale_map)(unsafe.Pointer(&loc.cat)))) + uintptr(int32(5))*8)))))
}
func Strerror(e int32) *int8 {
	return __strerror_l(e, __pthread_self().locale)
}
