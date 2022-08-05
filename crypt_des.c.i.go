package libc

import unsafe "unsafe"

var _cgos_key_shifts_crypt_des [16]uint8 = [16]uint8{uint8(1), uint8(1), uint8(2), uint8(2), uint8(2), uint8(2), uint8(2), uint8(2), uint8(1), uint8(2), uint8(2), uint8(2), uint8(2), uint8(2), uint8(2), uint8(1)}
var _cgos_psbox_crypt_des [8][64]uint32 = [8][64]uint32{[64]uint32{uint32(8421888), uint32(0), uint32(32768), uint32(8421890), uint32(8421378), uint32(33282), uint32(2), uint32(32768), uint32(512), uint32(8421888), uint32(8421890), uint32(512), uint32(8389122), uint32(8421378), uint32(8388608), uint32(2), uint32(514), uint32(8389120), uint32(8389120), uint32(33280), uint32(33280), uint32(8421376), uint32(8421376), uint32(8389122), uint32(32770), uint32(8388610), uint32(8388610), uint32(32770), uint32(0), uint32(514), uint32(33282), uint32(8388608), uint32(32768), uint32(8421890), uint32(2), uint32(8421376), uint32(8421888), uint32(8388608), uint32(8388608), uint32(512), uint32(8421378), uint32(32768), uint32(33280), uint32(8388610), uint32(512), uint32(2), uint32(8389122), uint32(33282), uint32(8421890), uint32(32770), uint32(8421376), uint32(8389122), uint32(8388610), uint32(514), uint32(33282), uint32(8421888), uint32(514), uint32(8389120), uint32(8389120), uint32(0), uint32(32770), uint32(33280), uint32(0), uint32(8421378)}, [64]uint32{uint32(1074282512), uint32(1073758208), uint32(16384), uint32(540688), uint32(524288), uint32(16), uint32(1074266128), uint32(1073758224), uint32(1073741840), uint32(1074282512), uint32(1074282496), uint32(1073741824), uint32(1073758208), uint32(524288), uint32(16), uint32(1074266128), uint32(540672), uint32(524304), uint32(1073758224), uint32(0), uint32(1073741824), uint32(16384), uint32(540688), uint32(1074266112), uint32(524304), uint32(1073741840), uint32(0), uint32(540672), uint32(16400), uint32(1074282496), uint32(1074266112), uint32(16400), uint32(0), uint32(540688), uint32(1074266128), uint32(524288), uint32(1073758224), uint32(1074266112), uint32(1074282496), uint32(16384), uint32(1074266112), uint32(1073758208), uint32(16), uint32(1074282512), uint32(540688), uint32(16), uint32(16384), uint32(1073741824), uint32(16400), uint32(1074282496), uint32(524288), uint32(1073741840), uint32(524304), uint32(1073758224), uint32(1073741840), uint32(524304), uint32(540672), uint32(0), uint32(1073758208), uint32(16400), uint32(1073741824), uint32(1074266128), uint32(1074282512), uint32(540672)}, [64]uint32{uint32(260), uint32(67174656), uint32(0), uint32(67174404), uint32(67109120), uint32(0), uint32(65796), uint32(67109120), uint32(65540), uint32(67108868), uint32(67108868), uint32(65536), uint32(67174660), uint32(65540), uint32(67174400), uint32(260), uint32(67108864), uint32(4), uint32(67174656), uint32(256), uint32(65792), uint32(67174400), uint32(67174404), uint32(65796), uint32(67109124), uint32(65792), uint32(65536), uint32(67109124), uint32(4), uint32(67174660), uint32(256), uint32(67108864), uint32(67174656), uint32(67108864), uint32(65540), uint32(260), uint32(65536), uint32(67174656), uint32(67109120), uint32(0), uint32(256), uint32(65540), uint32(67174660), uint32(67109120), uint32(67108868), uint32(256), uint32(0), uint32(67174404), uint32(67109124), uint32(65536), uint32(67108864), uint32(67174660), uint32(4), uint32(65796), uint32(65792), uint32(67108868), uint32(67174400), uint32(67109124), uint32(260), uint32(67174400), uint32(65796), uint32(4), uint32(67174404), uint32(65792)}, [64]uint32{uint32(2151682048), uint32(2147487808), uint32(2147487808), uint32(64), uint32(4198464), uint32(2151678016), uint32(2151677952), uint32(2147487744), uint32(0), uint32(4198400), uint32(4198400), uint32(2151682112), uint32(2147483712), uint32(0), uint32(4194368), uint32(2151677952), uint32(2147483648), uint32(4096), uint32(4194304), uint32(2151682048), uint32(64), uint32(4194304), uint32(2147487744), uint32(4160), uint32(2151678016), uint32(2147483648), uint32(4160), uint32(4194368), uint32(4096), uint32(4198464), uint32(2151682112), uint32(2147483712), uint32(4194368), uint32(2151677952), uint32(4198400), uint32(2151682112), uint32(2147483712), uint32(0), uint32(0), uint32(4198400), uint32(4160), uint32(4194368), uint32(2151678016), uint32(2147483648), uint32(2151682048), uint32(2147487808), uint32(2147487808), uint32(64), uint32(2151682112), uint32(2147483712), uint32(2147483648), uint32(4096), uint32(2151677952), uint32(2147487744), uint32(4198464), uint32(2151678016), uint32(2147487744), uint32(4160), uint32(4194304), uint32(2151682048), uint32(64), uint32(4194304), uint32(4096), uint32(4198464)}, [64]uint32{uint32(128), uint32(17039488), uint32(17039360), uint32(553648256), uint32(262144), uint32(128), uint32(536870912), uint32(17039360), uint32(537133184), uint32(262144), uint32(16777344), uint32(537133184), uint32(553648256), uint32(553910272), uint32(262272), uint32(536870912), uint32(16777216), uint32(537133056), uint32(537133056), uint32(0), uint32(536871040), uint32(553910400), uint32(553910400), uint32(16777344), uint32(553910272), uint32(536871040), uint32(0), uint32(553648128), uint32(17039488), uint32(16777216), uint32(553648128), uint32(262272), uint32(262144), uint32(553648256), uint32(128), uint32(16777216), uint32(536870912), uint32(17039360), uint32(553648256), uint32(537133184), uint32(16777344), uint32(536870912), uint32(553910272), uint32(17039488), uint32(537133184), uint32(128), uint32(16777216), uint32(553910272), uint32(553910400), uint32(262272), uint32(553648128), uint32(553910400), uint32(17039360), uint32(0), uint32(537133056), uint32(553648128), uint32(262272), uint32(16777344), uint32(536871040), uint32(262144), uint32(0), uint32(537133056), uint32(17039488), uint32(536871040)}, [64]uint32{uint32(268435464), uint32(270532608), uint32(8192), uint32(270540808), uint32(270532608), uint32(8), uint32(270540808), uint32(2097152), uint32(268443648), uint32(2105352), uint32(2097152), uint32(268435464), uint32(2097160), uint32(268443648), uint32(268435456), uint32(8200), uint32(0), uint32(2097160), uint32(268443656), uint32(8192), uint32(2105344), uint32(268443656), uint32(8), uint32(270532616), uint32(270532616), uint32(0), uint32(2105352), uint32(270540800), uint32(8200), uint32(2105344), uint32(270540800), uint32(268435456), uint32(268443648), uint32(8), uint32(270532616), uint32(2105344), uint32(270540808), uint32(2097152), uint32(8200), uint32(268435464), uint32(2097152), uint32(268443648), uint32(268435456), uint32(8200), uint32(268435464), uint32(270540808), uint32(2105344), uint32(270532608), uint32(2105352), uint32(270540800), uint32(0), uint32(270532616), uint32(8), uint32(8192), uint32(270532608), uint32(2105352), uint32(8192), uint32(2097160), uint32(268443656), uint32(0), uint32(270540800), uint32(268435456), uint32(2097160), uint32(268443656)}, [64]uint32{uint32(1048576), uint32(34603009), uint32(33555457), uint32(0), uint32(1024), uint32(33555457), uint32(1049601), uint32(34604032), uint32(34604033), uint32(1048576), uint32(0), uint32(33554433), uint32(1), uint32(33554432), uint32(34603009), uint32(1025), uint32(33555456), uint32(1049601), uint32(1048577), uint32(33555456), uint32(33554433), uint32(34603008), uint32(34604032), uint32(1048577), uint32(34603008), uint32(1024), uint32(1025), uint32(34604033), uint32(1049600), uint32(1), uint32(33554432), uint32(1049600), uint32(33554432), uint32(1049600), uint32(1048576), uint32(33555457), uint32(33555457), uint32(34603009), uint32(34603009), uint32(1), uint32(1048577), uint32(33554432), uint32(33555456), uint32(1048576), uint32(34604032), uint32(1025), uint32(1049601), uint32(34604032), uint32(1025), uint32(33554433), uint32(34604033), uint32(34603008), uint32(1049600), uint32(0), uint32(1), uint32(34604033), uint32(0), uint32(1049601), uint32(34603008), uint32(1024), uint32(33554433), uint32(33555456), uint32(1024), uint32(1048577)}, [64]uint32{uint32(134219808), uint32(2048), uint32(131072), uint32(134350880), uint32(134217728), uint32(134219808), uint32(32), uint32(134217728), uint32(131104), uint32(134348800), uint32(134350880), uint32(133120), uint32(134350848), uint32(133152), uint32(2048), uint32(32), uint32(134348800), uint32(134217760), uint32(134219776), uint32(2080), uint32(133120), uint32(131104), uint32(134348832), uint32(134350848), uint32(2080), uint32(0), uint32(0), uint32(134348832), uint32(134217760), uint32(134219776), uint32(133152), uint32(131072), uint32(133152), uint32(131072), uint32(134350848), uint32(2048), uint32(32), uint32(134348832), uint32(2048), uint32(133152), uint32(134219776), uint32(32), uint32(134217760), uint32(134348800), uint32(134348832), uint32(134217728), uint32(131072), uint32(134219808), uint32(0), uint32(134350880), uint32(131104), uint32(134217760), uint32(134348800), uint32(134219776), uint32(134219808), uint32(0), uint32(134350880), uint32(133120), uint32(133120), uint32(2080), uint32(2080), uint32(131104), uint32(134217728), uint32(134350848)}}
var _cgos_ip_maskl_crypt_des [16][16]uint32 = [16][16]uint32{[16]uint32{uint32(0), uint32(65536), uint32(0), uint32(65536), uint32(16777216), uint32(16842752), uint32(16777216), uint32(16842752), uint32(0), uint32(65536), uint32(0), uint32(65536), uint32(16777216), uint32(16842752), uint32(16777216), uint32(16842752)}, [16]uint32{uint32(0), uint32(1), uint32(0), uint32(1), uint32(256), uint32(257), uint32(256), uint32(257), uint32(0), uint32(1), uint32(0), uint32(1), uint32(256), uint32(257), uint32(256), uint32(257)}, [16]uint32{uint32(0), uint32(131072), uint32(0), uint32(131072), uint32(33554432), uint32(33685504), uint32(33554432), uint32(33685504), uint32(0), uint32(131072), uint32(0), uint32(131072), uint32(33554432), uint32(33685504), uint32(33554432), uint32(33685504)}, [16]uint32{uint32(0), uint32(2), uint32(0), uint32(2), uint32(512), uint32(514), uint32(512), uint32(514), uint32(0), uint32(2), uint32(0), uint32(2), uint32(512), uint32(514), uint32(512), uint32(514)}, [16]uint32{uint32(0), uint32(262144), uint32(0), uint32(262144), uint32(67108864), uint32(67371008), uint32(67108864), uint32(67371008), uint32(0), uint32(262144), uint32(0), uint32(262144), uint32(67108864), uint32(67371008), uint32(67108864), uint32(67371008)}, [16]uint32{uint32(0), uint32(4), uint32(0), uint32(4), uint32(1024), uint32(1028), uint32(1024), uint32(1028), uint32(0), uint32(4), uint32(0), uint32(4), uint32(1024), uint32(1028), uint32(1024), uint32(1028)}, [16]uint32{uint32(0), uint32(524288), uint32(0), uint32(524288), uint32(134217728), uint32(134742016), uint32(134217728), uint32(134742016), uint32(0), uint32(524288), uint32(0), uint32(524288), uint32(134217728), uint32(134742016), uint32(134217728), uint32(134742016)}, [16]uint32{uint32(0), uint32(8), uint32(0), uint32(8), uint32(2048), uint32(2056), uint32(2048), uint32(2056), uint32(0), uint32(8), uint32(0), uint32(8), uint32(2048), uint32(2056), uint32(2048), uint32(2056)}, [16]uint32{uint32(0), uint32(1048576), uint32(0), uint32(1048576), uint32(268435456), uint32(269484032), uint32(268435456), uint32(269484032), uint32(0), uint32(1048576), uint32(0), uint32(1048576), uint32(268435456), uint32(269484032), uint32(268435456), uint32(269484032)}, [16]uint32{uint32(0), uint32(16), uint32(0), uint32(16), uint32(4096), uint32(4112), uint32(4096), uint32(4112), uint32(0), uint32(16), uint32(0), uint32(16), uint32(4096), uint32(4112), uint32(4096), uint32(4112)}, [16]uint32{uint32(0), uint32(2097152), uint32(0), uint32(2097152), uint32(536870912), uint32(538968064), uint32(536870912), uint32(538968064), uint32(0), uint32(2097152), uint32(0), uint32(2097152), uint32(536870912), uint32(538968064), uint32(536870912), uint32(538968064)}, [16]uint32{uint32(0), uint32(32), uint32(0), uint32(32), uint32(8192), uint32(8224), uint32(8192), uint32(8224), uint32(0), uint32(32), uint32(0), uint32(32), uint32(8192), uint32(8224), uint32(8192), uint32(8224)}, [16]uint32{uint32(0), uint32(4194304), uint32(0), uint32(4194304), uint32(1073741824), uint32(1077936128), uint32(1073741824), uint32(1077936128), uint32(0), uint32(4194304), uint32(0), uint32(4194304), uint32(1073741824), uint32(1077936128), uint32(1073741824), uint32(1077936128)}, [16]uint32{uint32(0), uint32(64), uint32(0), uint32(64), uint32(16384), uint32(16448), uint32(16384), uint32(16448), uint32(0), uint32(64), uint32(0), uint32(64), uint32(16384), uint32(16448), uint32(16384), uint32(16448)}, [16]uint32{uint32(0), uint32(8388608), uint32(0), uint32(8388608), uint32(2147483648), uint32(2155872256), uint32(2147483648), uint32(2155872256), uint32(0), uint32(8388608), uint32(0), uint32(8388608), uint32(2147483648), uint32(2155872256), uint32(2147483648), uint32(2155872256)}, [16]uint32{uint32(0), uint32(128), uint32(0), uint32(128), uint32(32768), uint32(32896), uint32(32768), uint32(32896), uint32(0), uint32(128), uint32(0), uint32(128), uint32(32768), uint32(32896), uint32(32768), uint32(32896)}}
var _cgos_ip_maskr_crypt_des [16][16]uint32 = [16][16]uint32{[16]uint32{uint32(0), uint32(0), uint32(65536), uint32(65536), uint32(0), uint32(0), uint32(65536), uint32(65536), uint32(16777216), uint32(16777216), uint32(16842752), uint32(16842752), uint32(16777216), uint32(16777216), uint32(16842752), uint32(16842752)}, [16]uint32{uint32(0), uint32(0), uint32(1), uint32(1), uint32(0), uint32(0), uint32(1), uint32(1), uint32(256), uint32(256), uint32(257), uint32(257), uint32(256), uint32(256), uint32(257), uint32(257)}, [16]uint32{uint32(0), uint32(0), uint32(131072), uint32(131072), uint32(0), uint32(0), uint32(131072), uint32(131072), uint32(33554432), uint32(33554432), uint32(33685504), uint32(33685504), uint32(33554432), uint32(33554432), uint32(33685504), uint32(33685504)}, [16]uint32{uint32(0), uint32(0), uint32(2), uint32(2), uint32(0), uint32(0), uint32(2), uint32(2), uint32(512), uint32(512), uint32(514), uint32(514), uint32(512), uint32(512), uint32(514), uint32(514)}, [16]uint32{uint32(0), uint32(0), uint32(262144), uint32(262144), uint32(0), uint32(0), uint32(262144), uint32(262144), uint32(67108864), uint32(67108864), uint32(67371008), uint32(67371008), uint32(67108864), uint32(67108864), uint32(67371008), uint32(67371008)}, [16]uint32{uint32(0), uint32(0), uint32(4), uint32(4), uint32(0), uint32(0), uint32(4), uint32(4), uint32(1024), uint32(1024), uint32(1028), uint32(1028), uint32(1024), uint32(1024), uint32(1028), uint32(1028)}, [16]uint32{uint32(0), uint32(0), uint32(524288), uint32(524288), uint32(0), uint32(0), uint32(524288), uint32(524288), uint32(134217728), uint32(134217728), uint32(134742016), uint32(134742016), uint32(134217728), uint32(134217728), uint32(134742016), uint32(134742016)}, [16]uint32{uint32(0), uint32(0), uint32(8), uint32(8), uint32(0), uint32(0), uint32(8), uint32(8), uint32(2048), uint32(2048), uint32(2056), uint32(2056), uint32(2048), uint32(2048), uint32(2056), uint32(2056)}, [16]uint32{uint32(0), uint32(0), uint32(1048576), uint32(1048576), uint32(0), uint32(0), uint32(1048576), uint32(1048576), uint32(268435456), uint32(268435456), uint32(269484032), uint32(269484032), uint32(268435456), uint32(268435456), uint32(269484032), uint32(269484032)}, [16]uint32{uint32(0), uint32(0), uint32(16), uint32(16), uint32(0), uint32(0), uint32(16), uint32(16), uint32(4096), uint32(4096), uint32(4112), uint32(4112), uint32(4096), uint32(4096), uint32(4112), uint32(4112)}, [16]uint32{uint32(0), uint32(0), uint32(2097152), uint32(2097152), uint32(0), uint32(0), uint32(2097152), uint32(2097152), uint32(536870912), uint32(536870912), uint32(538968064), uint32(538968064), uint32(536870912), uint32(536870912), uint32(538968064), uint32(538968064)}, [16]uint32{uint32(0), uint32(0), uint32(32), uint32(32), uint32(0), uint32(0), uint32(32), uint32(32), uint32(8192), uint32(8192), uint32(8224), uint32(8224), uint32(8192), uint32(8192), uint32(8224), uint32(8224)}, [16]uint32{uint32(0), uint32(0), uint32(4194304), uint32(4194304), uint32(0), uint32(0), uint32(4194304), uint32(4194304), uint32(1073741824), uint32(1073741824), uint32(1077936128), uint32(1077936128), uint32(1073741824), uint32(1073741824), uint32(1077936128), uint32(1077936128)}, [16]uint32{uint32(0), uint32(0), uint32(64), uint32(64), uint32(0), uint32(0), uint32(64), uint32(64), uint32(16384), uint32(16384), uint32(16448), uint32(16448), uint32(16384), uint32(16384), uint32(16448), uint32(16448)}, [16]uint32{uint32(0), uint32(0), uint32(8388608), uint32(8388608), uint32(0), uint32(0), uint32(8388608), uint32(8388608), uint32(2147483648), uint32(2147483648), uint32(2155872256), uint32(2155872256), uint32(2147483648), uint32(2147483648), uint32(2155872256), uint32(2155872256)}, [16]uint32{uint32(0), uint32(0), uint32(128), uint32(128), uint32(0), uint32(0), uint32(128), uint32(128), uint32(32768), uint32(32768), uint32(32896), uint32(32896), uint32(32768), uint32(32768), uint32(32896), uint32(32896)}}
var _cgos_fp_maskl_crypt_des [8][16]uint32 = [8][16]uint32{[16]uint32{uint32(0), uint32(1073741824), uint32(4194304), uint32(1077936128), uint32(16384), uint32(1073758208), uint32(4210688), uint32(1077952512), uint32(64), uint32(1073741888), uint32(4194368), uint32(1077936192), uint32(16448), uint32(1073758272), uint32(4210752), uint32(1077952576)}, [16]uint32{uint32(0), uint32(268435456), uint32(1048576), uint32(269484032), uint32(4096), uint32(268439552), uint32(1052672), uint32(269488128), uint32(16), uint32(268435472), uint32(1048592), uint32(269484048), uint32(4112), uint32(268439568), uint32(1052688), uint32(269488144)}, [16]uint32{uint32(0), uint32(67108864), uint32(262144), uint32(67371008), uint32(1024), uint32(67109888), uint32(263168), uint32(67372032), uint32(4), uint32(67108868), uint32(262148), uint32(67371012), uint32(1028), uint32(67109892), uint32(263172), uint32(67372036)}, [16]uint32{uint32(0), uint32(16777216), uint32(65536), uint32(16842752), uint32(256), uint32(16777472), uint32(65792), uint32(16843008), uint32(1), uint32(16777217), uint32(65537), uint32(16842753), uint32(257), uint32(16777473), uint32(65793), uint32(16843009)}, [16]uint32{uint32(0), uint32(2147483648), uint32(8388608), uint32(2155872256), uint32(32768), uint32(2147516416), uint32(8421376), uint32(2155905024), uint32(128), uint32(2147483776), uint32(8388736), uint32(2155872384), uint32(32896), uint32(2147516544), uint32(8421504), uint32(2155905152)}, [16]uint32{uint32(0), uint32(536870912), uint32(2097152), uint32(538968064), uint32(8192), uint32(536879104), uint32(2105344), uint32(538976256), uint32(32), uint32(536870944), uint32(2097184), uint32(538968096), uint32(8224), uint32(536879136), uint32(2105376), uint32(538976288)}, [16]uint32{uint32(0), uint32(134217728), uint32(524288), uint32(134742016), uint32(2048), uint32(134219776), uint32(526336), uint32(134744064), uint32(8), uint32(134217736), uint32(524296), uint32(134742024), uint32(2056), uint32(134219784), uint32(526344), uint32(134744072)}, [16]uint32{uint32(0), uint32(33554432), uint32(131072), uint32(33685504), uint32(512), uint32(33554944), uint32(131584), uint32(33686016), uint32(2), uint32(33554434), uint32(131074), uint32(33685506), uint32(514), uint32(33554946), uint32(131586), uint32(33686018)}}
var _cgos_fp_maskr_crypt_des [8][16]uint32 = [8][16]uint32{[16]uint32{uint32(0), uint32(1073741824), uint32(4194304), uint32(1077936128), uint32(16384), uint32(1073758208), uint32(4210688), uint32(1077952512), uint32(64), uint32(1073741888), uint32(4194368), uint32(1077936192), uint32(16448), uint32(1073758272), uint32(4210752), uint32(1077952576)}, [16]uint32{uint32(0), uint32(268435456), uint32(1048576), uint32(269484032), uint32(4096), uint32(268439552), uint32(1052672), uint32(269488128), uint32(16), uint32(268435472), uint32(1048592), uint32(269484048), uint32(4112), uint32(268439568), uint32(1052688), uint32(269488144)}, [16]uint32{uint32(0), uint32(67108864), uint32(262144), uint32(67371008), uint32(1024), uint32(67109888), uint32(263168), uint32(67372032), uint32(4), uint32(67108868), uint32(262148), uint32(67371012), uint32(1028), uint32(67109892), uint32(263172), uint32(67372036)}, [16]uint32{uint32(0), uint32(16777216), uint32(65536), uint32(16842752), uint32(256), uint32(16777472), uint32(65792), uint32(16843008), uint32(1), uint32(16777217), uint32(65537), uint32(16842753), uint32(257), uint32(16777473), uint32(65793), uint32(16843009)}, [16]uint32{uint32(0), uint32(2147483648), uint32(8388608), uint32(2155872256), uint32(32768), uint32(2147516416), uint32(8421376), uint32(2155905024), uint32(128), uint32(2147483776), uint32(8388736), uint32(2155872384), uint32(32896), uint32(2147516544), uint32(8421504), uint32(2155905152)}, [16]uint32{uint32(0), uint32(536870912), uint32(2097152), uint32(538968064), uint32(8192), uint32(536879104), uint32(2105344), uint32(538976256), uint32(32), uint32(536870944), uint32(2097184), uint32(538968096), uint32(8224), uint32(536879136), uint32(2105376), uint32(538976288)}, [16]uint32{uint32(0), uint32(134217728), uint32(524288), uint32(134742016), uint32(2048), uint32(134219776), uint32(526336), uint32(134744064), uint32(8), uint32(134217736), uint32(524296), uint32(134742024), uint32(2056), uint32(134219784), uint32(526344), uint32(134744072)}, [16]uint32{uint32(0), uint32(33554432), uint32(131072), uint32(33685504), uint32(512), uint32(33554944), uint32(131584), uint32(33686016), uint32(2), uint32(33554434), uint32(131074), uint32(33685506), uint32(514), uint32(33554946), uint32(131586), uint32(33686018)}}
var _cgos_key_perm_maskl_crypt_des [8][16]uint32 = [8][16]uint32{[16]uint32{uint32(0), uint32(0), uint32(16), uint32(16), uint32(4096), uint32(4096), uint32(4112), uint32(4112), uint32(1048576), uint32(1048576), uint32(1048592), uint32(1048592), uint32(1052672), uint32(1052672), uint32(1052688), uint32(1052688)}, [16]uint32{uint32(0), uint32(0), uint32(32), uint32(32), uint32(8192), uint32(8192), uint32(8224), uint32(8224), uint32(2097152), uint32(2097152), uint32(2097184), uint32(2097184), uint32(2105344), uint32(2105344), uint32(2105376), uint32(2105376)}, [16]uint32{uint32(0), uint32(0), uint32(64), uint32(64), uint32(16384), uint32(16384), uint32(16448), uint32(16448), uint32(4194304), uint32(4194304), uint32(4194368), uint32(4194368), uint32(4210688), uint32(4210688), uint32(4210752), uint32(4210752)}, [16]uint32{uint32(0), uint32(0), uint32(128), uint32(128), uint32(32768), uint32(32768), uint32(32896), uint32(32896), uint32(8388608), uint32(8388608), uint32(8388736), uint32(8388736), uint32(8421376), uint32(8421376), uint32(8421504), uint32(8421504)}, [16]uint32{uint32(0), uint32(1), uint32(256), uint32(257), uint32(65536), uint32(65537), uint32(65792), uint32(65793), uint32(16777216), uint32(16777217), uint32(16777472), uint32(16777473), uint32(16842752), uint32(16842753), uint32(16843008), uint32(16843009)}, [16]uint32{uint32(0), uint32(2), uint32(512), uint32(514), uint32(131072), uint32(131074), uint32(131584), uint32(131586), uint32(33554432), uint32(33554434), uint32(33554944), uint32(33554946), uint32(33685504), uint32(33685506), uint32(33686016), uint32(33686018)}, [16]uint32{uint32(0), uint32(4), uint32(1024), uint32(1028), uint32(262144), uint32(262148), uint32(263168), uint32(263172), uint32(67108864), uint32(67108868), uint32(67109888), uint32(67109892), uint32(67371008), uint32(67371012), uint32(67372032), uint32(67372036)}, [16]uint32{uint32(0), uint32(8), uint32(2048), uint32(2056), uint32(524288), uint32(524296), uint32(526336), uint32(526344), uint32(134217728), uint32(134217736), uint32(134219776), uint32(134219784), uint32(134742016), uint32(134742024), uint32(134744064), uint32(134744072)}}
var _cgos_key_perm_maskr_crypt_des [12][16]uint32 = [12][16]uint32{[16]uint32{uint32(0), uint32(1), uint32(0), uint32(1), uint32(0), uint32(1), uint32(0), uint32(1), uint32(0), uint32(1), uint32(0), uint32(1), uint32(0), uint32(1), uint32(0), uint32(1)}, [16]uint32{uint32(0), uint32(0), uint32(1048576), uint32(1048576), uint32(4096), uint32(4096), uint32(1052672), uint32(1052672), uint32(16), uint32(16), uint32(1048592), uint32(1048592), uint32(4112), uint32(4112), uint32(1052688), uint32(1052688)}, [16]uint32{uint32(0), uint32(2), uint32(0), uint32(2), uint32(0), uint32(2), uint32(0), uint32(2), uint32(0), uint32(2), uint32(0), uint32(2), uint32(0), uint32(2), uint32(0), uint32(2)}, [16]uint32{uint32(0), uint32(0), uint32(2097152), uint32(2097152), uint32(8192), uint32(8192), uint32(2105344), uint32(2105344), uint32(32), uint32(32), uint32(2097184), uint32(2097184), uint32(8224), uint32(8224), uint32(2105376), uint32(2105376)}, [16]uint32{uint32(0), uint32(4), uint32(0), uint32(4), uint32(0), uint32(4), uint32(0), uint32(4), uint32(0), uint32(4), uint32(0), uint32(4), uint32(0), uint32(4), uint32(0), uint32(4)}, [16]uint32{uint32(0), uint32(0), uint32(4194304), uint32(4194304), uint32(16384), uint32(16384), uint32(4210688), uint32(4210688), uint32(64), uint32(64), uint32(4194368), uint32(4194368), uint32(16448), uint32(16448), uint32(4210752), uint32(4210752)}, [16]uint32{uint32(0), uint32(8), uint32(0), uint32(8), uint32(0), uint32(8), uint32(0), uint32(8), uint32(0), uint32(8), uint32(0), uint32(8), uint32(0), uint32(8), uint32(0), uint32(8)}, [16]uint32{uint32(0), uint32(0), uint32(8388608), uint32(8388608), uint32(32768), uint32(32768), uint32(8421376), uint32(8421376), uint32(128), uint32(128), uint32(8388736), uint32(8388736), uint32(32896), uint32(32896), uint32(8421504), uint32(8421504)}, [16]uint32{uint32(0), uint32(0), uint32(16777216), uint32(16777216), uint32(65536), uint32(65536), uint32(16842752), uint32(16842752), uint32(256), uint32(256), uint32(16777472), uint32(16777472), uint32(65792), uint32(65792), uint32(16843008), uint32(16843008)}, [16]uint32{uint32(0), uint32(0), uint32(33554432), uint32(33554432), uint32(131072), uint32(131072), uint32(33685504), uint32(33685504), uint32(512), uint32(512), uint32(33554944), uint32(33554944), uint32(131584), uint32(131584), uint32(33686016), uint32(33686016)}, [16]uint32{uint32(0), uint32(0), uint32(67108864), uint32(67108864), uint32(262144), uint32(262144), uint32(67371008), uint32(67371008), uint32(1024), uint32(1024), uint32(67109888), uint32(67109888), uint32(263168), uint32(263168), uint32(67372032), uint32(67372032)}, [16]uint32{uint32(0), uint32(0), uint32(134217728), uint32(134217728), uint32(524288), uint32(524288), uint32(134742016), uint32(134742016), uint32(2048), uint32(2048), uint32(134219776), uint32(134219776), uint32(526336), uint32(526336), uint32(134744064), uint32(134744064)}}
var _cgos_comp_maskl0_crypt_des [4][8]uint32 = [4][8]uint32{[8]uint32{uint32(0), uint32(131072), uint32(1), uint32(131073), uint32(524288), uint32(655360), uint32(524289), uint32(655361)}, [8]uint32{uint32(0), uint32(4096), uint32(0), uint32(4096), uint32(64), uint32(4160), uint32(64), uint32(4160)}, [8]uint32{uint32(0), uint32(4194304), uint32(32), uint32(4194336), uint32(32768), uint32(4227072), uint32(32800), uint32(4227104)}, [8]uint32{uint32(0), uint32(1048576), uint32(2048), uint32(1050624), uint32(0), uint32(1048576), uint32(2048), uint32(1050624)}}
var _cgos_comp_maskr0_crypt_des [4][8]uint32 = [4][8]uint32{[8]uint32{uint32(0), uint32(2097152), uint32(131072), uint32(2228224), uint32(2), uint32(2097154), uint32(131074), uint32(2228226)}, [8]uint32{uint32(0), uint32(0), uint32(1048576), uint32(1048576), uint32(4), uint32(4), uint32(1048580), uint32(1048580)}, [8]uint32{uint32(0), uint32(16384), uint32(2048), uint32(18432), uint32(0), uint32(16384), uint32(2048), uint32(18432)}, [8]uint32{uint32(0), uint32(4194304), uint32(32768), uint32(4227072), uint32(8), uint32(4194312), uint32(32776), uint32(4227080)}}
var _cgos_comp_maskl1_crypt_des [4][16]uint32 = [4][16]uint32{[16]uint32{uint32(0), uint32(16), uint32(16384), uint32(16400), uint32(262144), uint32(262160), uint32(278528), uint32(278544), uint32(256), uint32(272), uint32(16640), uint32(16656), uint32(262400), uint32(262416), uint32(278784), uint32(278800)}, [16]uint32{uint32(0), uint32(8388608), uint32(2), uint32(8388610), uint32(512), uint32(8389120), uint32(514), uint32(8389122), uint32(2097152), uint32(10485760), uint32(2097154), uint32(10485762), uint32(2097664), uint32(10486272), uint32(2097666), uint32(10486274)}, [16]uint32{uint32(0), uint32(8192), uint32(4), uint32(8196), uint32(1024), uint32(9216), uint32(1028), uint32(9220), uint32(0), uint32(8192), uint32(4), uint32(8196), uint32(1024), uint32(9216), uint32(1028), uint32(9220)}, [16]uint32{uint32(0), uint32(65536), uint32(8), uint32(65544), uint32(128), uint32(65664), uint32(136), uint32(65672), uint32(0), uint32(65536), uint32(8), uint32(65544), uint32(128), uint32(65664), uint32(136), uint32(65672)}}
var _cgos_comp_maskr1_crypt_des [4][16]uint32 = [4][16]uint32{[16]uint32{uint32(0), uint32(0), uint32(128), uint32(128), uint32(8192), uint32(8192), uint32(8320), uint32(8320), uint32(1), uint32(1), uint32(129), uint32(129), uint32(8193), uint32(8193), uint32(8321), uint32(8321)}, [16]uint32{uint32(0), uint32(16), uint32(8388608), uint32(8388624), uint32(65536), uint32(65552), uint32(8454144), uint32(8454160), uint32(512), uint32(528), uint32(8389120), uint32(8389136), uint32(66048), uint32(66064), uint32(8454656), uint32(8454672)}, [16]uint32{uint32(0), uint32(1024), uint32(4096), uint32(5120), uint32(524288), uint32(525312), uint32(528384), uint32(529408), uint32(32), uint32(1056), uint32(4128), uint32(5152), uint32(524320), uint32(525344), uint32(528416), uint32(529440)}, [16]uint32{uint32(0), uint32(256), uint32(262144), uint32(262400), uint32(0), uint32(256), uint32(262144), uint32(262400), uint32(64), uint32(320), uint32(262208), uint32(262464), uint32(64), uint32(320), uint32(262208), uint32(262464)}}
var _cgos_ascii64_crypt_des [65]uint8 = [65]uint8{'.', '/', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '\x00'}

func _cgos_ascii_to_bin_crypt_des(ch int32) uint32 {
	var sch int32 = func() int32 {
		if ch < int32(128) {
			return ch
		} else {
			return -(int32(256) - ch)
		}
	}()
	var retval int32
	retval = sch - '.'
	if sch >= 'A' {
		retval = sch - 53
		if sch >= 'a' {
			retval = sch - 59
		}
	}
	retval &= int32(63)
	return uint32(retval)
}
func _cgos_ascii_is_unsafe_crypt_des(ch uint8) int32 {
	return func() int32 {
		if !(ch != 0) || int32(ch) == '\n' || int32(ch) == ':' {
			return 1
		} else {
			return 0
		}
	}()
}
func _cgos_setup_salt_crypt_des(salt uint32) uint32 {
	var obit uint32
	var saltbit uint32
	var saltbits uint32
	var i uint32
	saltbits = uint32(0)
	saltbit = uint32(1)
	obit = uint32(8388608)
	for i = uint32(0); i < uint32(24); i++ {
		if salt&saltbit != 0 {
			saltbits |= obit
		}
		saltbit <<= int32(1)
		obit >>= int32(1)
	}
	return saltbits
}
func __des_setkey(key *uint8, ekey *struct_expanded_key) {
	var k0 uint32
	var k1 uint32
	var rawkey0 uint32
	var rawkey1 uint32
	var shifts uint32
	var round uint32
	var i uint32
	var ibit uint32
	rawkey0 = uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(key)) + uintptr(int32(3))))) | uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(key)) + uintptr(int32(2)))))<<int32(8) | uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(key)) + uintptr(int32(1)))))<<int32(16) | uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(key)) + uintptr(int32(0)))))<<int32(24)
	rawkey1 = uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(key)) + uintptr(int32(7))))) | uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(key)) + uintptr(int32(6)))))<<int32(8) | uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(key)) + uintptr(int32(5)))))<<int32(16) | uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(key)) + uintptr(int32(4)))))<<int32(24)
	k0 = func() (_cgo_ret uint32) {
		_cgo_addr := &k1
		*_cgo_addr = uint32(0)
		return *_cgo_addr
	}()
	for func() uint32 {
		i = uint32(0)
		return func() (_cgo_ret uint32) {
			_cgo_addr := &ibit
			*_cgo_addr = uint32(28)
			return *_cgo_addr
		}()
	}(); i < uint32(4); func() uint32 {
		i++
		return func() (_cgo_ret uint32) {
			_cgo_addr := &ibit
			*_cgo_addr -= uint32(4)
			return *_cgo_addr
		}()
	}() {
		var j uint32 = i << int32(1)
		k0 |= *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[16]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[16]uint32)(unsafe.Pointer(&_cgos_key_perm_maskl_crypt_des)))) + uintptr(i)*64)))))) + uintptr(rawkey0>>ibit&uint32(15))*4)) | *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[16]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[16]uint32)(unsafe.Pointer(&_cgos_key_perm_maskl_crypt_des)))) + uintptr(i+uint32(4))*64)))))) + uintptr(rawkey1>>ibit&uint32(15))*4))
		k1 |= *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[16]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[16]uint32)(unsafe.Pointer(&_cgos_key_perm_maskr_crypt_des)))) + uintptr(j)*64)))))) + uintptr(rawkey0>>ibit&uint32(15))*4))
		ibit -= uint32(4)
		k1 |= *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[16]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[16]uint32)(unsafe.Pointer(&_cgos_key_perm_maskr_crypt_des)))) + uintptr(j+uint32(1))*64)))))) + uintptr(rawkey0>>ibit&uint32(15))*4)) | *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[16]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[16]uint32)(unsafe.Pointer(&_cgos_key_perm_maskr_crypt_des)))) + uintptr(i+uint32(8))*64)))))) + uintptr(rawkey1>>ibit&uint32(15))*4))
	}
	shifts = uint32(0)
	for round = uint32(0); round < uint32(16); round++ {
		var t0 uint32
		var t1 uint32
		var kl uint32
		var kr uint32
		shifts += uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_key_shifts_crypt_des)))) + uintptr(round))))
		t0 = k0<<shifts | k0>>(uint32(28)-shifts)
		t1 = k1<<shifts | k1>>(uint32(28)-shifts)
		kl = func() (_cgo_ret uint32) {
			_cgo_addr := &kr
			*_cgo_addr = uint32(0)
			return *_cgo_addr
		}()
		ibit = uint32(25)
		for i = uint32(0); i < uint32(4); i++ {
			kl |= *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[8]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[8]uint32)(unsafe.Pointer(&_cgos_comp_maskl0_crypt_des)))) + uintptr(i)*32)))))) + uintptr(t0>>ibit&uint32(7))*4))
			kr |= *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[8]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[8]uint32)(unsafe.Pointer(&_cgos_comp_maskr0_crypt_des)))) + uintptr(i)*32)))))) + uintptr(t1>>ibit&uint32(7))*4))
			ibit -= uint32(4)
			kl |= *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[16]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[16]uint32)(unsafe.Pointer(&_cgos_comp_maskl1_crypt_des)))) + uintptr(i)*64)))))) + uintptr(t0>>ibit&uint32(15))*4))
			kr |= *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[16]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[16]uint32)(unsafe.Pointer(&_cgos_comp_maskr1_crypt_des)))) + uintptr(i)*64)))))) + uintptr(t1>>ibit&uint32(15))*4))
			ibit -= uint32(3)
		}
		*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&ekey.l)))) + uintptr(round)*4)) = kl
		*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&ekey.r)))) + uintptr(round)*4)) = kr
	}
}
func __do_des(l_in uint32, r_in uint32, l_out *uint32, r_out *uint32, count uint32, saltbits uint32, ekey *struct_expanded_key) {
	var l uint32
	var r uint32
	l = func() (_cgo_ret uint32) {
		_cgo_addr := &r
		*_cgo_addr = uint32(0)
		return *_cgo_addr
	}()
	if l_in|r_in != 0 {
		var i uint32
		var ibit uint32
		for func() uint32 {
			i = uint32(0)
			return func() (_cgo_ret uint32) {
				_cgo_addr := &ibit
				*_cgo_addr = uint32(28)
				return *_cgo_addr
			}()
		}(); i < uint32(8); func() uint32 {
			i++
			return func() (_cgo_ret uint32) {
				_cgo_addr := &ibit
				*_cgo_addr -= uint32(4)
				return *_cgo_addr
			}()
		}() {
			l |= *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[16]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[16]uint32)(unsafe.Pointer(&_cgos_ip_maskl_crypt_des)))) + uintptr(i)*64)))))) + uintptr(l_in>>ibit&uint32(15))*4)) | *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[16]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[16]uint32)(unsafe.Pointer(&_cgos_ip_maskl_crypt_des)))) + uintptr(i+uint32(8))*64)))))) + uintptr(r_in>>ibit&uint32(15))*4))
			r |= *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[16]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[16]uint32)(unsafe.Pointer(&_cgos_ip_maskr_crypt_des)))) + uintptr(i)*64)))))) + uintptr(l_in>>ibit&uint32(15))*4)) | *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[16]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[16]uint32)(unsafe.Pointer(&_cgos_ip_maskr_crypt_des)))) + uintptr(i+uint32(8))*64)))))) + uintptr(r_in>>ibit&uint32(15))*4))
		}
	}
	for func() (_cgo_ret uint32) {
		_cgo_addr := &count
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		var round uint32 = uint32(16)
		var kl *uint32 = (*uint32)(unsafe.Pointer(&ekey.l))
		var kr *uint32 = (*uint32)(unsafe.Pointer(&ekey.r))
		var f uint32
		for func() (_cgo_ret uint32) {
			_cgo_addr := &round
			_cgo_ret = *_cgo_addr
			*_cgo_addr--
			return
		}() != 0 {
			var r48l uint32
			var r48r uint32
			r48l = r&uint32(1)<<int32(23) | r&uint32(4160749568)>>int32(9) | r&uint32(528482304)>>int32(11) | r&uint32(33030144)>>int32(13) | r&uint32(2064384)>>int32(15)
			r48r = r&uint32(129024)<<int32(7) | r&uint32(8064)<<int32(5) | r&uint32(504)<<int32(3) | r&uint32(31)<<int32(1) | r&uint32(2147483648)>>int32(31)
			f = (r48l ^ r48r) & saltbits
			r48l ^= f ^ *func() (_cgo_ret *uint32) {
				_cgo_addr := &kl
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}()
			r48r ^= f ^ *func() (_cgo_ret *uint32) {
				_cgo_addr := &kr
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}()
			f = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[64]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[64]uint32)(unsafe.Pointer(&_cgos_psbox_crypt_des)))) + uintptr(int32(0))*256)))))) + uintptr(r48l>>int32(18))*4)) | *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[64]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[64]uint32)(unsafe.Pointer(&_cgos_psbox_crypt_des)))) + uintptr(int32(1))*256)))))) + uintptr(r48l>>int32(12)&uint32(63))*4)) | *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[64]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[64]uint32)(unsafe.Pointer(&_cgos_psbox_crypt_des)))) + uintptr(int32(2))*256)))))) + uintptr(r48l>>int32(6)&uint32(63))*4)) | *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[64]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[64]uint32)(unsafe.Pointer(&_cgos_psbox_crypt_des)))) + uintptr(int32(3))*256)))))) + uintptr(r48l&uint32(63))*4)) | *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[64]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[64]uint32)(unsafe.Pointer(&_cgos_psbox_crypt_des)))) + uintptr(int32(4))*256)))))) + uintptr(r48r>>int32(18))*4)) | *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[64]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[64]uint32)(unsafe.Pointer(&_cgos_psbox_crypt_des)))) + uintptr(int32(5))*256)))))) + uintptr(r48r>>int32(12)&uint32(63))*4)) | *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[64]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[64]uint32)(unsafe.Pointer(&_cgos_psbox_crypt_des)))) + uintptr(int32(6))*256)))))) + uintptr(r48r>>int32(6)&uint32(63))*4)) | *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[64]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[64]uint32)(unsafe.Pointer(&_cgos_psbox_crypt_des)))) + uintptr(int32(7))*256)))))) + uintptr(r48r&uint32(63))*4))
			f ^= l
			l = r
			r = f
		}
		r = l
		l = f
	}
	{
		var i uint32
		var ibit uint32
		var lo uint32
		var ro uint32
		lo = func() (_cgo_ret uint32) {
			_cgo_addr := &ro
			*_cgo_addr = uint32(0)
			return *_cgo_addr
		}()
		for func() uint32 {
			i = uint32(0)
			return func() (_cgo_ret uint32) {
				_cgo_addr := &ibit
				*_cgo_addr = uint32(28)
				return *_cgo_addr
			}()
		}(); i < uint32(4); func() uint32 {
			i++
			return func() (_cgo_ret uint32) {
				_cgo_addr := &ibit
				*_cgo_addr -= uint32(4)
				return *_cgo_addr
			}()
		}() {
			ro |= *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[16]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[16]uint32)(unsafe.Pointer(&_cgos_fp_maskr_crypt_des)))) + uintptr(i)*64)))))) + uintptr(l>>ibit&uint32(15))*4)) | *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[16]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[16]uint32)(unsafe.Pointer(&_cgos_fp_maskr_crypt_des)))) + uintptr(i+uint32(4))*64)))))) + uintptr(r>>ibit&uint32(15))*4))
			ibit -= uint32(4)
			lo |= *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[16]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[16]uint32)(unsafe.Pointer(&_cgos_fp_maskl_crypt_des)))) + uintptr(i)*64)))))) + uintptr(l>>ibit&uint32(15))*4)) | *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&*(*[16]uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[16]uint32)(unsafe.Pointer(&_cgos_fp_maskl_crypt_des)))) + uintptr(i+uint32(4))*64)))))) + uintptr(r>>ibit&uint32(15))*4))
		}
		*l_out = lo
		*r_out = ro
	}
}
func _cgos_des_cipher_crypt_des(in *uint8, out *uint8, count uint32, saltbits uint32, ekey *struct_expanded_key) {
	var l_out uint32
	var r_out uint32
	var rawl uint32
	var rawr uint32
	rawl = uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(in)) + uintptr(int32(3))))) | uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(in)) + uintptr(int32(2)))))<<int32(8) | uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(in)) + uintptr(int32(1)))))<<int32(16) | uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(in)) + uintptr(int32(0)))))<<int32(24)
	rawr = uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(in)) + uintptr(int32(7))))) | uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(in)) + uintptr(int32(6)))))<<int32(8) | uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(in)) + uintptr(int32(5)))))<<int32(16) | uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(in)) + uintptr(int32(4)))))<<int32(24)
	__do_des(rawl, rawr, &l_out, &r_out, count, saltbits, ekey)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out)) + uintptr(int32(0)))) = uint8(l_out >> int32(24))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out)) + uintptr(int32(1)))) = uint8(l_out >> int32(16))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out)) + uintptr(int32(2)))) = uint8(l_out >> int32(8))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out)) + uintptr(int32(3)))) = uint8(l_out)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out)) + uintptr(int32(4)))) = uint8(r_out >> int32(24))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out)) + uintptr(int32(5)))) = uint8(r_out >> int32(16))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out)) + uintptr(int32(6)))) = uint8(r_out >> int32(8))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out)) + uintptr(int32(7)))) = uint8(r_out)
}
func _cgos__crypt_extended_r_uut_crypt_des(_key *int8, _setting *int8, output *int8) *int8 {
	var key *uint8 = (*uint8)(unsafe.Pointer(_key))
	var setting *uint8 = (*uint8)(unsafe.Pointer(_setting))
	var ekey struct_expanded_key
	var keybuf [8]uint8
	var p *uint8
	var q *uint8
	var count uint32
	var salt uint32
	var l uint32
	var r0 uint32
	var r1 uint32
	var i uint32
	q = (*uint8)(unsafe.Pointer(&keybuf))
	for uintptr(unsafe.Pointer(q)) <= uintptr(unsafe.Pointer(&*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&keybuf)))) + uintptr(7))))) {
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &q
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = uint8(int32(*key) << int32(1))
		if *key != 0 {
			*(*uintptr)(unsafe.Pointer(&key))++
		}
	}
	__des_setkey((*uint8)(unsafe.Pointer(&keybuf)), &ekey)
	if int32(*setting) == '_' {
		for func() uint32 {
			i = uint32(1)
			return func() (_cgo_ret uint32) {
				_cgo_addr := &count
				*_cgo_addr = uint32(0)
				return *_cgo_addr
			}()
		}(); i < uint32(5); i++ {
			var value uint32 = _cgos_ascii_to_bin_crypt_des(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(setting)) + uintptr(i)))))
			if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_ascii64_crypt_des)))) + uintptr(value)))) != int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(setting)) + uintptr(i)))) {
				return (*int8)(nil)
			}
			count |= value << ((i - uint32(1)) * uint32(6))
		}
		if !(count != 0) {
			return (*int8)(nil)
		}
		for func() uint32 {
			i = uint32(5)
			return func() (_cgo_ret uint32) {
				_cgo_addr := &salt
				*_cgo_addr = uint32(0)
				return *_cgo_addr
			}()
		}(); i < uint32(9); i++ {
			var value uint32 = _cgos_ascii_to_bin_crypt_des(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(setting)) + uintptr(i)))))
			if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_ascii64_crypt_des)))) + uintptr(value)))) != int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(setting)) + uintptr(i)))) {
				return (*int8)(nil)
			}
			salt |= value << ((i - uint32(5)) * uint32(6))
		}
		for *key != 0 {
			_cgos_des_cipher_crypt_des((*uint8)(unsafe.Pointer(&keybuf)), (*uint8)(unsafe.Pointer(&keybuf)), uint32(1), uint32(0), &ekey)
			q = (*uint8)(unsafe.Pointer(&keybuf))
			for uintptr(unsafe.Pointer(q)) <= uintptr(unsafe.Pointer(&*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&keybuf)))) + uintptr(7))))) && int32(*key) != 0 {
				*func() (_cgo_ret *uint8) {
					_cgo_addr := &q
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}() ^= uint8(int32(*func() (_cgo_ret *uint8) {
					_cgo_addr := &key
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}()) << int32(1))
			}
			__des_setkey((*uint8)(unsafe.Pointer(&keybuf)), &ekey)
		}
		Memcpy(unsafe.Pointer(output), unsafe.Pointer(setting), uint64(9))
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(output)) + uintptr(int32(9)))) = int8('\x00')
		p = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(output)))) + uintptr(int32(9))))
	} else {
		count = uint32(25)
		if _cgos_ascii_is_unsafe_crypt_des(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(setting)) + uintptr(int32(0))))) != 0 || _cgos_ascii_is_unsafe_crypt_des(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(setting)) + uintptr(int32(1))))) != 0 {
			return (*int8)(nil)
		}
		salt = _cgos_ascii_to_bin_crypt_des(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(setting)) + uintptr(int32(1))))))<<int32(6) | _cgos_ascii_to_bin_crypt_des(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(setting)) + uintptr(int32(0))))))
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(output)) + uintptr(int32(0)))) = int8(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(setting)) + uintptr(int32(0)))))
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(output)) + uintptr(int32(1)))) = int8(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(setting)) + uintptr(int32(1)))))
		p = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(output)))) + uintptr(int32(2))))
	}
	__do_des(uint32(0), uint32(0), &r0, &r1, count, _cgos_setup_salt_crypt_des(salt), &ekey)
	l = r0 >> int32(8)
	*func() (_cgo_ret *uint8) {
		_cgo_addr := &p
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}() = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_ascii64_crypt_des)))) + uintptr(l>>int32(18)&uint32(63))))
	*func() (_cgo_ret *uint8) {
		_cgo_addr := &p
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}() = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_ascii64_crypt_des)))) + uintptr(l>>int32(12)&uint32(63))))
	*func() (_cgo_ret *uint8) {
		_cgo_addr := &p
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}() = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_ascii64_crypt_des)))) + uintptr(l>>int32(6)&uint32(63))))
	*func() (_cgo_ret *uint8) {
		_cgo_addr := &p
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}() = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_ascii64_crypt_des)))) + uintptr(l&uint32(63))))
	l = r0<<int32(16) | r1>>int32(16)&uint32(65535)
	*func() (_cgo_ret *uint8) {
		_cgo_addr := &p
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}() = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_ascii64_crypt_des)))) + uintptr(l>>int32(18)&uint32(63))))
	*func() (_cgo_ret *uint8) {
		_cgo_addr := &p
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}() = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_ascii64_crypt_des)))) + uintptr(l>>int32(12)&uint32(63))))
	*func() (_cgo_ret *uint8) {
		_cgo_addr := &p
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}() = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_ascii64_crypt_des)))) + uintptr(l>>int32(6)&uint32(63))))
	*func() (_cgo_ret *uint8) {
		_cgo_addr := &p
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}() = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_ascii64_crypt_des)))) + uintptr(l&uint32(63))))
	l = r1 << int32(2)
	*func() (_cgo_ret *uint8) {
		_cgo_addr := &p
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}() = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_ascii64_crypt_des)))) + uintptr(l>>int32(12)&uint32(63))))
	*func() (_cgo_ret *uint8) {
		_cgo_addr := &p
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}() = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_ascii64_crypt_des)))) + uintptr(l>>int32(6)&uint32(63))))
	*func() (_cgo_ret *uint8) {
		_cgo_addr := &p
		_cgo_ret = *_cgo_addr
		*(*uintptr)(unsafe.Pointer(_cgo_addr))++
		return
	}() = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&_cgos_ascii64_crypt_des)))) + uintptr(l&uint32(63))))
	*p = uint8(0)
	return output
}
func __crypt_des(key *int8, setting *int8, output *int8) *int8 {
	var test_key *int8 = (*int8)(unsafe.Pointer(&[21]int8{-128, -1, -128, '\x01', ' ', '\u007f', -127, -128, -128, '\r', '\n', -1, '\u007f', ' ', -127, ' ', 't', 'e', 's', 't', '\x00'}))
	var test_setting *int8 = (*int8)(unsafe.Pointer(&[10]int8{'_', '0', '.', '.', '.', '/', '9', 'Z', 'z', '\x00'}))
	var test_hash *int8 = (*int8)(unsafe.Pointer(&[21]int8{'_', '0', '.', '.', '.', '/', '9', 'Z', 'z', 'X', '7', 'i', 'S', 'J', 'N', 'd', '2', '1', 's', 'U', '\x00'}))
	var test_buf [21]int8
	var retval *int8
	var p *int8
	if int32(*setting) != '_' {
		test_setting = (*int8)(unsafe.Pointer(&[3]int8{-128, 'x', '\x00'}))
		test_hash = (*int8)(unsafe.Pointer(&[14]int8{-128, 'x', '2', '2', '/', 'w', 'K', '5', '2', 'Z', 'K', 'G', 'A', '\x00'}))
	}
	retval = _cgos__crypt_extended_r_uut_crypt_des(key, setting, output)
	p = _cgos__crypt_extended_r_uut_crypt_des(test_key, test_setting, (*int8)(unsafe.Pointer(&test_buf)))
	if p != nil && !(Strcmp(p, test_hash) != 0) && retval != nil {
		return retval
	}
	return func() *int8 {
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(setting)) + uintptr(int32(0))))) == '*' {
			return (*int8)(unsafe.Pointer(&[2]int8{'x', '\x00'}))
		} else {
			return (*int8)(unsafe.Pointer(&[2]int8{'*', '\x00'}))
		}
	}()
}
