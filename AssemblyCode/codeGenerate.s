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
SUB SP, SP, #168
// Entero: 10
MOV x1, #10
// Declaracion inferida: a
MOV x0, x1
STR x0, [x29, #-16]
// Entero: 3
MOV x1, #3
// Declaracion inferida: b
MOV x0, x1
STR x0, [x29, #-24]
//  ------------- VisitAddSub -------------
// Identificador: a con profundidad: 0
LDR x0, [x29, #-16]
STR x0, [x29, #-32]
// Identificador: b con profundidad: 0
LDR x0, [x29, #-24]
STR x0, [x29, #-40]
// SUMA DE ENTEROS
LDR x2, [x29, #-32]
LDR x1, [x29, #-40]
ADD x0, x2, x1
STR x0, [x29, #-48]
// Declaracion inferida: suma
LDR x1, [x29, #-48]
MOV x0, x1
STR x0, [x29, #-56]
//  ------------- VisitAddSub -------------
// Identificador: a con profundidad: 0
LDR x0, [x29, #-16]
STR x0, [x29, #-64]
// Identificador: b con profundidad: 0
LDR x0, [x29, #-24]
STR x0, [x29, #-72]
// RESTA DE ENTEROS
LDR x2, [x29, #-64]
LDR x1, [x29, #-72]
SUB x0, x2, x1
STR x0, [x29, #-80]
// Declaracion inferida: resta
LDR x1, [x29, #-80]
MOV x0, x1
STR x0, [x29, #-88]
//  ------------- VisitMulDivModulo -------------
// Identificador: a con profundidad: 0
LDR x0, [x29, #-16]
STR x0, [x29, #-96]
// Identificador: b con profundidad: 0
LDR x0, [x29, #-24]
STR x0, [x29, #-104]
// MULTIPLICACION DE ENTEROS
LDR x2, [x29, #-96]
LDR x1, [x29, #-104]
MUL x0, x2, x1
STR x0, [x29, #-112]
// Declaracion inferida: multi
LDR x1, [x29, #-112]
MOV x0, x1
STR x0, [x29, #-120]
//  ------------- VisitMulDivModulo -------------
// Identificador: a con profundidad: 1
LDR x0, [x29, #-16]
STR x0, [x29, #-128]
// Identificador: b con profundidad: 1
LDR x0, [x29, #-24]
STR x0, [x29, #-136]
// MODULO DE ENTEROS
LDR x2, [x29, #-128]
LDR x1, [x29, #-136]
SDIV x3, x2, x1
MSUB x0, x3, x1, x2
STR x0, [x29, #-144]
// Declaracion inferida: mod
LDR x1, [x29, #-144]
MOV x0, x1
STR x0, [x29, #-152]
// Función embebida: println
// Identificador: suma con profundidad: 1
LDR x0, [x29, #-56]
STR x0, [x29, #-160]
// Preparando entero para imprimir
LDR x1, [x29, #-160]
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
// Identificador: resta con profundidad: 1
LDR x0, [x29, #-88]
STR x0, [x29, #-168]
// Preparando entero para imprimir
LDR x1, [x29, #-168]
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
// Identificador: multi con profundidad: 1
LDR x0, [x29, #-120]
STR x0, [x29, #-176]
// Preparando entero para imprimir
LDR x1, [x29, #-176]
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
// Identificador: mod con profundidad: 1
LDR x0, [x29, #-152]
STR x0, [x29, #-184]
// Preparando entero para imprimir
LDR x1, [x29, #-184]
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
L1:
ADD sp, sp, #168
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
	