TEXT ·rotate16(SB), 0, $16-8
    MOVD arg+0(FP), R0
    ADD R0@>16, ZR, R0
    MOVD R0, ret+8(FP)
    RET
