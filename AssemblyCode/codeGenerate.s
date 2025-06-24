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
// Función embebida: println
// cadena: "hola"
SUB SP, SP, #8
STR x10, [SP, #0]
// Byte: 104 uso de Heap: 'h'
MOV w0, #104
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
// Byte: 0 uso de Heap: '\x00'
MOV w0, #0
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
LDR x0, [SP], #8
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
// Declaracion no inicializada: miVar1 con tipo int
MOV x0, #0
STR x0, [SP, #-8]!
LDR x0, [SP], #8
MOV x1, #0
SUB x1, x29, x1
STR x0, [x1, #0]
// Float: 12.200000
MOVZ X0, #26214, LSL #0
MOVK X0, #26214, LSL #16
MOVK X0, #26214, LSL #32
MOVK X0, #16424, LSL #48
STR x0, [SP, #-8]!
// Declaracion explicita: miVar2 con tipo float64
LDR x0, [SP], #8
MOV x1, #8
SUB x1, x29, x1
STR x0, [x1, #0]
// Entero: 3
MOV x0, #3
STR x0, [SP, #-8]!
// Declaracion implicita: miVar3
LDR x0, [SP], #8
MOV x1, #16
SUB x1, x29, x1
STR x0, [x1, #0]
// Función embebida: println
MOV x0, #0
SUB x0, x29, x0
LDR x0, [x0, #0]
STR x0, [SP, #-8]!
LDR x0, [SP], #8
MOV X0, x0
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
MOV x0, #16
SUB x0, x29, x0
LDR x0, [x0, #0]
STR x0, [SP, #-8]!
LDR x0, [SP], #8
MOV X0, x0
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
L1:
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
	