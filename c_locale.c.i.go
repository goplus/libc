package libc

import unsafe "unsafe"

var _cgos_empty_mo__c_locale [5]uint32 = [5]uint32{uint32(2500072158), uint32(0), uint32(4294967295), uint32(4294967295), uint32(4294967295)}
var __c_dot_utf8 struct___locale_map = struct___locale_map{unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_empty_mo__c_locale))), 20, [24]int8{'C', '.', 'U', 'T', 'F', '-', '8', '\x00'}, nil}
var __c_locale struct___locale_struct = struct___locale_struct{[6]*struct___locale_map{nil}}
var __c_dot_utf8_locale struct___locale_struct = struct___locale_struct{[6]*struct___locale_map{&__c_dot_utf8}}
