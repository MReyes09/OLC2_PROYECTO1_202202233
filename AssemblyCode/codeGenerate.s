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
// cadena: "Mensaje 1"
MOV x11, x10
STR x0, [x29, #-8]
// Byte: 77 uso de Heap: 'M'
MOV w0, #77
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 101 uso de Heap: 'e'
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 110 uso de Heap: 'n'
MOV w0, #110
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
// Byte: 106 uso de Heap: 'j'
MOV w0, #106
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 101 uso de Heap: 'e'
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 49 uso de Heap: '1'
MOV w0, #49
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 0 uso de Heap: '\x00'
MOV w0, #0
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Declaracion implicita: x
MOV x0, x11
STR x0, [x29, #-16]
// Función embebida: println
// cadena: "Saludando"
MOV x11, x10
STR x0, [x29, #-24]
// Byte: 83 uso de Heap: 'S'
MOV w0, #83
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
// Byte: 117 uso de Heap: 'u'
MOV w0, #117
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 100 uso de Heap: 'd'
MOV w0, #100
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 97 uso de Heap: 'a'
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 110 uso de Heap: 'n'
MOV w0, #110
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 100 uso de Heap: 'd'
MOV w0, #100
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 111 uso de Heap: 'o'
MOV w0, #111
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 0 uso de Heap: '\x00'
MOV w0, #0
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
MOV x0, x11
MOV X0, x0
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
// cadena: "Mensaje 2"
MOV x11, x10
STR x0, [x29, #-32]
// Byte: 77 uso de Heap: 'M'
MOV w0, #77
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 101 uso de Heap: 'e'
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 110 uso de Heap: 'n'
MOV w0, #110
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
// Byte: 106 uso de Heap: 'j'
MOV w0, #106
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 101 uso de Heap: 'e'
MOV w0, #101
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
// Declaracion implicita: y
MOV x0, x11
STR x0, [x29, #-40]
// Función embebida: println
LDR x0, [x29, #-16]
MOV x11, x0
MOV x0, x11
MOV X0, x0
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
LDR x0, [x29, #-40]
MOV x11, x0
MOV x0, x11
MOV X0, x0
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
// cadena: "FUNCIONOOOOOOOOOOOOOOOOOO"
MOV x11, x10
STR x0, [x29, #-48]
// Byte: 70 uso de Heap: 'F'
MOV w0, #70
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 85 uso de Heap: 'U'
MOV w0, #85
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 78 uso de Heap: 'N'
MOV w0, #78
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 67 uso de Heap: 'C'
MOV w0, #67
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 73 uso de Heap: 'I'
MOV w0, #73
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 79 uso de Heap: 'O'
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 78 uso de Heap: 'N'
MOV w0, #78
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 79 uso de Heap: 'O'
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 79 uso de Heap: 'O'
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 79 uso de Heap: 'O'
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 79 uso de Heap: 'O'
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 79 uso de Heap: 'O'
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 79 uso de Heap: 'O'
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 79 uso de Heap: 'O'
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 79 uso de Heap: 'O'
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 79 uso de Heap: 'O'
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 79 uso de Heap: 'O'
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 79 uso de Heap: 'O'
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 79 uso de Heap: 'O'
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 79 uso de Heap: 'O'
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 79 uso de Heap: 'O'
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 79 uso de Heap: 'O'
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 79 uso de Heap: 'O'
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 79 uso de Heap: 'O'
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 79 uso de Heap: 'O'
MOV w0, #79
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 0 uso de Heap: '\x00'
MOV w0, #0
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
MOV x0, x11
MOV X0, x0
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
ADD sp, sp, #72
LDP x29, x30, [sp], #16
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
	