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
SUB SP, SP, #104
// Función embebida: println
// cadena: "Valor numero:"
MOV x11, x10
STR x0, [x29, #-16]
// Byte: 86 uso de Heap: 'V'
MOV w0, #86
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
// Byte: 111 uso de Heap: 'o'
MOV w0, #111
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 114 uso de Heap: 'r'
MOV w0, #114
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 110 uso de Heap: 'n'
MOV w0, #110
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 117 uso de Heap: 'u'
MOV w0, #117
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 109 uso de Heap: 'm'
MOV w0, #109
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 101 uso de Heap: 'e'
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 114 uso de Heap: 'r'
MOV w0, #114
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 111 uso de Heap: 'o'
MOV w0, #111
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 58 uso de Heap: ':'
MOV w0, #58
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
// Entero: 10
MOV x1, #10
MOV X0, x1
BL print_entero
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
// cadena: "Declaracion x := 123"
MOV x11, x10
STR x0, [x29, #-24]
// Byte: 68 uso de Heap: 'D'
MOV w0, #68
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 101 uso de Heap: 'e'
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 99 uso de Heap: 'c'
MOV w0, #99
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
// Byte: 114 uso de Heap: 'r'
MOV w0, #114
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 97 uso de Heap: 'a'
MOV w0, #97
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 99 uso de Heap: 'c'
MOV w0, #99
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
// Byte: 110 uso de Heap: 'n'
MOV w0, #110
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 120 uso de Heap: 'x'
MOV w0, #120
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 58 uso de Heap: ':'
MOV w0, #58
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 61 uso de Heap: '='
MOV w0, #61
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
// Byte: 50 uso de Heap: '2'
MOV w0, #50
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 51 uso de Heap: '3'
MOV w0, #51
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
// cadena: "Saludando!"
MOV x11, x10
STR x0, [x29, #-32]
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
// Byte: 33 uso de Heap: '!'
MOV w0, #33
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 0 uso de Heap: '\x00'
MOV w0, #0
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Declaracion inferida: y
MOV x0, x11
STR x0, [x29, #-40]
// Entero: 123
MOV x1, #123
// Declaracion inferida: x
MOV x0, x1
STR x0, [x29, #-48]
// Función embebida: println
// cadena: "Y ->"
MOV x11, x10
STR x0, [x29, #-56]
// Byte: 89 uso de Heap: 'Y'
MOV w0, #89
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 45 uso de Heap: '-'
MOV w0, #45
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 62 uso de Heap: '>'
MOV w0, #62
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
// cadena: "X ->"
MOV x11, x10
STR x0, [x29, #-64]
// Byte: 88 uso de Heap: 'X'
MOV w0, #88
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 45 uso de Heap: '-'
MOV w0, #45
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 62 uso de Heap: '>'
MOV w0, #62
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
LDR x1, [x29, #-48]
MOV X0, x1
BL print_entero
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
// cadena: "al reves"
MOV x11, x10
STR x0, [x29, #-72]
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
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 114 uso de Heap: 'r'
MOV w0, #114
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 101 uso de Heap: 'e'
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 118 uso de Heap: 'v'
MOV w0, #118
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
// Función embebida: println
// cadena: "X ->"
MOV x11, x10
STR x0, [x29, #-80]
// Byte: 88 uso de Heap: 'X'
MOV w0, #88
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 45 uso de Heap: '-'
MOV w0, #45
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 62 uso de Heap: '>'
MOV w0, #62
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
LDR x1, [x29, #-48]
MOV X0, x1
BL print_entero
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
// cadena: "Y ->"
MOV x11, x10
STR x0, [x29, #-88]
// Byte: 89 uso de Heap: 'Y'
MOV w0, #89
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 45 uso de Heap: '-'
MOV w0, #45
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 62 uso de Heap: '>'
MOV w0, #62
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
L1:
ADD sp, sp, #104
LDP x29, x30, [sp], #16
RET

//Funciones De Impresion:

		.align 2
		print_entero:
			stp x29, x30, [sp, #-16]!
			stp x19, x20, [sp, #-16]!
			stp x21, x22, [sp, #-16]!
			stp x23, x24, [sp, #-16]!
			stp x25, x26, [sp, #-16]!
			stp x27, x28, [sp, #-16]!
			mov x19, x0
			cmp x19, #0 
			bge numero_positivo   
			mov x0, #1                 
			adr x1, signo_menos1        
			mov x2, #1                 
			mov w8, #64                
			svc #0
			neg x19, x19               
		.align 2
		numero_positivo:
			sub sp, sp, #32            
			mov x22, sp                
			mov x23, #0            
			cmp x19, #0
			bne bucle_conversion
			mov w24, #48           
			strb w24, [x22, x23]
			add x23, x23, #1 
			b print_resultado
		.align 2
		bucle_conversion:
			mov x24, #10
			udiv x25, x19, x24
			msub x26, x25, x24, x19
			add x26, x26, #48         
			strb w26, [x22, x23]     
			add x23, x23, #1          
			mov x19, x25               
			cbnz x19, bucle_conversion 
			mov x27, #0      
		.align 2
		bucle_inversion:
			sub x28, x23, x27        
			sub x28, x28, #1     
			cmp x27, x28              
			bge print_resultado  
			ldrb w24, [x22, x27]      
			ldrb w25, [x22, x28]       
			strb w25, [x22, x27]       
			strb w24, [x22, x28]      
			add x27, x27, #1           
			b bucle_inversion         
		.align 2
		print_resultado:
			mov x0, #1                 
			mov x1, x22               
			mov x2, x23               
			mov w8, #64              
			svc #0
			// Restaurar registros
			add sp, sp, #32            
			ldp x27, x28, [sp], #16   
			ldp x25, x26, [sp], #16
			ldp x23, x24, [sp], #16
			ldp x21, x22, [sp], #16
			ldp x19, x20, [sp], #16
			ldp x29, x30, [sp], #16    
			ret                       
		.align 2
		signo_menos1:
			.ascii "-"
	

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
	