#include "textflag.h"

GLOBL ·TestInt(SB),NOPTR,$8

DATA ·TestInt+0(SB)/1,$0x37
DATA ·TestInt+1(SB)/1,$0x25
DATA ·TestInt+2(SB)/1,$0x00
DATA ·TestInt+3(SB)/1,$0x00
DATA ·TestInt+4(SB)/1,$0x00
DATA ·TestInt+5(SB)/1,$0x00
DATA ·TestInt+6(SB)/1,$0x00
DATA ·TestInt+7(SB)/1,$0x00

GLOBL ·TestStr(SB),NOPTR,$40

DATA ·TestStr+0(SB)/8,$·TestStr+16(SB)
DATA ·TestStr+8(SB)/8,$17
DATA ·TestStr+16(SB)/24,$"Hello go assembly"
