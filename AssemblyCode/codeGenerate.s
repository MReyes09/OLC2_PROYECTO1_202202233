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
SUB SP, SP, #72
//  ------------- VisitAddSub -------------
// Entero: 6
MOV x1, #6
STR x1, [x29, #-16]
//  ------------- VisitAddSub -------------
// Entero: 9
MOV x1, #9
STR x1, [x29, #-24]
// Entero: 8
MOV x1, #8
STR x1, [x29, #-32]
// SUMA DE ENTEROS
LDR x2, [x29, #-24]
LDR x1, [x29, #-32]
ADD x0, x2, x1
STR x0, [x29, #-40]
// SUMA DE ENTEROS
LDR x2, [x29, #-16]
LDR x1, [x29, #-40]
ADD x0, x2, x1
STR x0, [x29, #-48]
// Declaracion inferida: numero
LDR x1, [x29, #-48]
MOV x0, x1
STR x0, [x29, #-56]
L1:
ADD sp, sp, #72
LDP x29, x30, [sp], #16
RET

//Funciones De Impresion:
