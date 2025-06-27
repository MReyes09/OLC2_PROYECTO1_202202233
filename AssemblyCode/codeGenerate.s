.data
buffer: .space 256
salto_linea_str: .ascii "\n"
espacio_str: .ascii " "
corchete_apertura: .ascii "["
corchete_cierre: .ascii "]"
coma_espacio: .ascii ", "
heap: .space 4096
.text
.global _start
_start:
 mov x29, sp
 adr x10, heap
 bl main
MOV x0, #0
MOV x8, #93
SVC #0

//Funciones Foraneas:
main:
stp x29, x30, [sp, #-16]!
mov x29, sp
SUB SP, SP, #264
// Entero: 8
MOV x1, #8
// Declaracion inferida: a
MOV x0, x1
STR x0, [x29, #-16]
// Entero: 5
MOV x1, #5
// Declaracion inferida: b
MOV x0, x1
STR x0, [x29, #-24]
// Entero: 8
MOV x1, #8
// Declaracion inferida: c
MOV x0, x1
STR x0, [x29, #-32]
// Float: 2.500000
MOVZ X0, #0, LSL #0
MOVK X0, #0, LSL #16
MOVK X0, #0, LSL #32
MOVK X0, #16388, LSL #48
FMOV d0, x0
// Declaracion inferida: x
STR d0, [x29, #-40]
// Float: 3.500000
MOVZ X0, #0, LSL #0
MOVK X0, #0, LSL #16
MOVK X0, #0, LSL #32
MOVK X0, #16396, LSL #48
FMOV d0, x0
// Declaracion inferida: y
STR d0, [x29, #-48]
// Float: 2.500000
MOVZ X0, #0, LSL #0
MOVK X0, #0, LSL #16
MOVK X0, #0, LSL #32
MOVK X0, #16388, LSL #48
FMOV d0, x0
// Declaracion inferida: z
STR d0, [x29, #-56]
//  ------------- Operación Relacional -------------
// Identificador: a con profundidad: 0
LDR x0, [x29, #-16]
STR x0, [x29, #-64]
// Identificador: b con profundidad: 0
LDR x0, [x29, #-24]
STR x0, [x29, #-72]
// Comparación de Enteros
LDR x2, [x29, #-64]
LDR x1, [x29, #-72]
CMP x2, x1
B.GT Ltrue0
MOV x0, #0
B Lfalse0
Ltrue0:
MOV x0, #1
Lfalse0:
STR x0, [x29, #-80]
// Declaracion inferida: r1
LDR x1, [x29, #-80]
MOV x0, x1
STR x0, [x29, #-88]
//  ------------- Operación Relacional -------------
// Identificador: b con profundidad: 0
LDR x0, [x29, #-24]
STR x0, [x29, #-96]
// Identificador: a con profundidad: 0
LDR x0, [x29, #-16]
STR x0, [x29, #-104]
// Comparación de Enteros
LDR x2, [x29, #-96]
LDR x1, [x29, #-104]
CMP x2, x1
B.GE Ltrue1
MOV x0, #0
B Lfalse1
Ltrue1:
MOV x0, #1
Lfalse1:
STR x0, [x29, #-112]
// Declaracion inferida: r2
LDR x1, [x29, #-112]
MOV x0, x1
STR x0, [x29, #-120]
//  ------------- Operación Relacional -------------
// Identificador: a con profundidad: 0
LDR x0, [x29, #-16]
STR x0, [x29, #-128]
// Identificador: c con profundidad: 0
LDR x0, [x29, #-32]
STR x0, [x29, #-136]
// Comparación de Enteros
LDR x2, [x29, #-128]
LDR x1, [x29, #-136]
CMP x2, x1
B.LE Ltrue2
MOV x0, #0
B Lfalse2
Ltrue2:
MOV x0, #1
Lfalse2:
STR x0, [x29, #-144]
// Declaracion inferida: r3
LDR x1, [x29, #-144]
MOV x0, x1
STR x0, [x29, #-152]
//  ------------- Operación Relacional -------------
// Identificador: x con profundidad: 0
LDR d0, [x29, #-40]
STR d0, [x29, #-160]
// Identificador: y con profundidad: 0
LDR d0, [x29, #-48]
STR d0, [x29, #-168]
// Comparación de Float64
LDR d0, [x29, #-160]
LDR d1, [x29, #-168]
FCMP d0, d1
B.LT Ltrue3
MOV x0, #0
B Lfalse3
Ltrue3:
MOV x0, #1
Lfalse3:
STR x0, [x29, #-176]
// Declaracion inferida: r4
LDR x1, [x29, #-176]
MOV x0, x1
STR x0, [x29, #-184]
//  ------------- Operación Relacional -------------
// Identificador: y con profundidad: 0
LDR d0, [x29, #-48]
STR d0, [x29, #-192]
// Identificador: z con profundidad: 0
LDR d0, [x29, #-56]
STR d0, [x29, #-200]
// Comparación de Float64
LDR d0, [x29, #-192]
LDR d1, [x29, #-200]
FCMP d0, d1
B.LE Ltrue4
MOV x0, #0
B Lfalse4
Ltrue4:
MOV x0, #1
Lfalse4:
STR x0, [x29, #-208]
// Declaracion inferida: r5
LDR x1, [x29, #-208]
MOV x0, x1
STR x0, [x29, #-216]
//  ------------- Operación Relacional -------------
// Identificador: x con profundidad: 0
LDR d0, [x29, #-40]
STR d0, [x29, #-224]
// Identificador: z con profundidad: 0
LDR d0, [x29, #-56]
STR d0, [x29, #-232]
// Comparación de Float64
LDR d0, [x29, #-224]
LDR d1, [x29, #-232]
FCMP d0, d1
B.GE Ltrue5
MOV x0, #0
B Lfalse5
Ltrue5:
MOV x0, #1
Lfalse5:
STR x0, [x29, #-240]
// Declaracion inferida: r6
LDR x1, [x29, #-240]
MOV x0, x1
STR x0, [x29, #-248]
// Función embebida: println
// Identificador: r1
LDR x1, [x29, #-88]
LDR x1, [x29, #-88]
MOV X0, x1
BL print_booleano
MOV X0, x1
BL print_cadena
MOV X0, #1
ADR X1, espacio_str
MOV X2, #1
MOV W8, #64
SVC #0
MOV X0, #1
ADR X1, salto_linea_str
MOV X2, #1
MOV W8, #64
SVC #0
// Función embebida: println
// Identificador: r2
LDR x1, [x29, #-120]
LDR x1, [x29, #-120]
MOV X0, x1
BL print_booleano
MOV X0, x1
BL print_cadena
MOV X0, #1
ADR X1, espacio_str
MOV X2, #1
MOV W8, #64
SVC #0
MOV X0, #1
ADR X1, salto_linea_str
MOV X2, #1
MOV W8, #64
SVC #0
// Función embebida: println
// Identificador: r3
LDR x1, [x29, #-152]
LDR x1, [x29, #-152]
MOV X0, x1
BL print_booleano
MOV X0, x1
BL print_cadena
MOV X0, #1
ADR X1, espacio_str
MOV X2, #1
MOV W8, #64
SVC #0
MOV X0, #1
ADR X1, salto_linea_str
MOV X2, #1
MOV W8, #64
SVC #0
// Función embebida: println
// Identificador: r4
LDR x1, [x29, #-184]
LDR x1, [x29, #-184]
MOV X0, x1
BL print_booleano
MOV X0, x1
BL print_cadena
MOV X0, #1
ADR X1, espacio_str
MOV X2, #1
MOV W8, #64
SVC #0
MOV X0, #1
ADR X1, salto_linea_str
MOV X2, #1
MOV W8, #64
SVC #0
// Función embebida: println
// Identificador: r5
LDR x1, [x29, #-216]
LDR x1, [x29, #-216]
MOV X0, x1
BL print_booleano
MOV X0, x1
BL print_cadena
MOV X0, #1
ADR X1, espacio_str
MOV X2, #1
MOV W8, #64
SVC #0
MOV X0, #1
ADR X1, salto_linea_str
MOV X2, #1
MOV W8, #64
SVC #0
// Función embebida: println
// Identificador: r6
LDR x1, [x29, #-248]
LDR x1, [x29, #-248]
MOV X0, x1
BL print_booleano
MOV X0, x1
BL print_cadena
MOV X0, #1
ADR X1, espacio_str
MOV X2, #1
MOV W8, #64
SVC #0
MOV X0, #1
ADR X1, salto_linea_str
MOV X2, #1
MOV W8, #64
SVC #0
L1:
ADD sp, sp, #264
LDP x29, x30, [sp], #16
RET

//Funciones De Impresion:

            .align 2
            print_booleano:
                stp x29, x30, [sp, #-16]!
                cmp x0, #0
                bne print_true
                adr x0, falso
                b print_bool_call
            .align 2
            print_true:
                adr x0, verdadero
            .align 2
            print_bool_call:
                bl print_cadena    // Usa tu función existente
                ldp x29, x30, [sp], #16
                ret
            .align 2
            verdadero:
                .asciz "true"      // .asciz añade automáticamente '\0'
            .align 2
            falso:
                .asciz "false"
        

		.align 2
		print_cadena:
			stp     x29, x30, [sp, #-16]!
			stp     x19, x20, [sp, #-16]!
			mov     x19, x0
		.align 2
		print_loop:
			ldrb    w20, [x19]
			cbz     w20, print_done
			mov     x0, #1      
			mov     x1, x19
			mov     x2, #1
			mov     x8, #64             
			svc     #0      
			add     x19, x19, #1
			b       print_loop
		.align 2    
		print_done:
			ldp     x19, x20, [sp], #16
			ldp     x29, x30, [sp], #16
			ret
	