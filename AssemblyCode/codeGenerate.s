.data
salto_linea_str: .ascii "\n"
espacio_str: .ascii " "
corchete_apertura: .ascii "["
corchete_cierre: .ascii "]"
coma_espacio: .ascii ", "
heap: .space 4096
.text
.global _start
_start:
 adr x10, heap
 bl main
MOV x0, #0
MOV x8, #93
SVC #0

//Funciones Foraneas:
main:
// Función embebida: println
// cadena_print: "Hola papu"
STR x10, [SP, #-8]!
// Byte: 72 uso de Heap: 'H'
MOV w0, #72
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 111 uso de Heap: 'o'
MOV w0, #111
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 108 uso de Heap: 'l'
MOV w0, #108
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 97 uso de Heap: 'a'
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 112 uso de Heap: 'p'
MOV w0, #112
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 97 uso de Heap: 'a'
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 112 uso de Heap: 'p'
MOV w0, #112
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 117 uso de Heap: 'u'
MOV w0, #117
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 0 uso de Heap: '\x00'
MOV w0, #0
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
LDR x0, [SP], #8
MOV X0, x0
BL print_cadena
MOV X0, #1
ADR X1, salto_linea_str
MOV X2, #1
MOV W8, #64
SVC #0
// Función embebida: println
// cadena_print: "Como estas?"
STR x10, [SP, #-8]!
// Byte: 67 uso de Heap: 'C'
MOV w0, #67
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 111 uso de Heap: 'o'
MOV w0, #111
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 109 uso de Heap: 'm'
MOV w0, #109
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 111 uso de Heap: 'o'
MOV w0, #111
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 101 uso de Heap: 'e'
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 115 uso de Heap: 's'
MOV w0, #115
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 116 uso de Heap: 't'
MOV w0, #116
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 97 uso de Heap: 'a'
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 115 uso de Heap: 's'
MOV w0, #115
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 63 uso de Heap: '?'
MOV w0, #63
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 0 uso de Heap: '\x00'
MOV w0, #0
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
LDR x0, [SP], #8
MOV X0, x0
BL print_cadena
MOV X0, #1
ADR X1, salto_linea_str
MOV X2, #1
MOV W8, #64
SVC #0
// Función embebida: println
// cadena_print: "Ya salio compi 2"
STR x10, [SP, #-8]!
// Byte: 89 uso de Heap: 'Y'
MOV w0, #89
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 97 uso de Heap: 'a'
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 115 uso de Heap: 's'
MOV w0, #115
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 97 uso de Heap: 'a'
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 108 uso de Heap: 'l'
MOV w0, #108
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 105 uso de Heap: 'i'
MOV w0, #105
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 111 uso de Heap: 'o'
MOV w0, #111
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 99 uso de Heap: 'c'
MOV w0, #99
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 111 uso de Heap: 'o'
MOV w0, #111
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 109 uso de Heap: 'm'
MOV w0, #109
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 112 uso de Heap: 'p'
MOV w0, #112
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 105 uso de Heap: 'i'
MOV w0, #105
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 50 uso de Heap: '2'
MOV w0, #50
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 0 uso de Heap: '\x00'
MOV w0, #0
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
LDR x0, [SP], #8
MOV X0, x0
BL print_cadena
MOV X0, #1
ADR X1, salto_linea_str
MOV X2, #1
MOV W8, #64
SVC #0
L1:
ADD x0, x29, xzr
LDR x30, [x0, #0]
RET

//Funciones De Impresion:

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
	