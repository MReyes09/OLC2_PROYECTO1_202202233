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
SUB SP, SP, #424
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
// Entero: 2
MOV x1, #2
// Declaracion inferida: c
MOV x0, x1
STR x0, [x29, #-32]
// Float: 3.200000
MOVZ X0, #39322, LSL #0
MOVK X0, #39321, LSL #16
MOVK X0, #39321, LSL #32
MOVK X0, #16393, LSL #48
FMOV d0, x0
// Declaracion inferida: x
STR d0, [x29, #-40]
// Float: 1.800000
MOVZ X0, #52429, LSL #0
MOVK X0, #52428, LSL #16
MOVK X0, #52428, LSL #32
MOVK X0, #16380, LSL #48
FMOV d0, x0
// Declaracion inferida: y
STR d0, [x29, #-48]
// Float: 6.000000
MOVZ X0, #0, LSL #0
MOVK X0, #0, LSL #16
MOVK X0, #0, LSL #32
MOVK X0, #16408, LSL #48
FMOV d0, x0
// Declaracion inferida: z
STR d0, [x29, #-56]
//  ------------- VisitAddSub -------------
// Identificador: a con profundidad: 0
LDR x0, [x29, #-16]
STR x0, [x29, #-64]
// Identificador: c con profundidad: 0
LDR x0, [x29, #-32]
STR x0, [x29, #-72]
// SUMA DE ENTEROS
LDR x2, [x29, #-64]
LDR x1, [x29, #-72]
ADD x0, x2, x1
STR x0, [x29, #-80]
// Declaracion inferida: suma
LDR x1, [x29, #-80]
MOV x0, x1
STR x0, [x29, #-88]
//  ------------- VisitAddSub -------------
// Identificador: b con profundidad: 0
LDR x0, [x29, #-24]
STR x0, [x29, #-96]
// Identificador: c con profundidad: 0
LDR x0, [x29, #-32]
STR x0, [x29, #-104]
// RESTA DE ENTEROS
LDR x2, [x29, #-96]
LDR x1, [x29, #-104]
SUB x0, x2, x1
STR x0, [x29, #-112]
// Declaracion inferida: resta
LDR x1, [x29, #-112]
MOV x0, x1
STR x0, [x29, #-120]
//  ------------- VisitMulDivModulo -------------
// Identificador: a con profundidad: 0
LDR x0, [x29, #-16]
STR x0, [x29, #-128]
// Identificador: b con profundidad: 0
LDR x0, [x29, #-24]
STR x0, [x29, #-136]
// MULTIPLICACION DE ENTEROS
LDR x2, [x29, #-128]
LDR x1, [x29, #-136]
MUL x0, x2, x1
STR x0, [x29, #-144]
// Declaracion inferida: multi
LDR x1, [x29, #-144]
MOV x0, x1
STR x0, [x29, #-152]
//  ------------- VisitMulDivModulo -------------
// Identificador: a con profundidad: 0
LDR x0, [x29, #-16]
STR x0, [x29, #-160]
// Identificador: c con profundidad: 0
LDR x0, [x29, #-32]
STR x0, [x29, #-168]
// DIVISION DE ENTEROS
LDR x2, [x29, #-160]
LDR x1, [x29, #-168]
SDIV x0, x2, x1
STR x0, [x29, #-176]
// Declaracion inferida: div
LDR x1, [x29, #-176]
MOV x0, x1
STR x0, [x29, #-184]
//  ------------- VisitMulDivModulo -------------
// Identificador: b con profundidad: 0
LDR x0, [x29, #-24]
STR x0, [x29, #-192]
// Identificador: c con profundidad: 0
LDR x0, [x29, #-32]
STR x0, [x29, #-200]
// MODULO DE ENTEROS
LDR x2, [x29, #-192]
LDR x1, [x29, #-200]
SDIV x3, x2, x1
MSUB x0, x3, x1, x2
STR x0, [x29, #-208]
// Declaracion inferida: mod
LDR x1, [x29, #-208]
MOV x0, x1
STR x0, [x29, #-216]
//  ------------- VisitAddSub -------------
// Identificador: x con profundidad: 0
LDR d0, [x29, #-40]
STR d0, [x29, #-224]
// Identificador: z con profundidad: 0
LDR d0, [x29, #-56]
STR d0, [x29, #-232]
// SUMA DE FLOAT64
LDR d0, [x29, #-224]
LDR d1, [x29, #-232]
FADD d0, d0, d1
STR d0, [x29, #-240]
// Declaracion inferida: sumaF
LDR d0, [x29, #-240]
STR d0, [x29, #-248]
//  ------------- VisitAddSub -------------
// Identificador: z con profundidad: 0
LDR d0, [x29, #-56]
MOV x9, #-256
ADD x9, x29, x9
STR d0, [x9]
// Identificador: y con profundidad: 0
LDR d0, [x29, #-48]
MOV x9, #-264
ADD x9, x29, x9
STR d0, [x9]
// RESTA DE FLOAT64
MOV x9, #-256
ADD x9, x29, x9
LDR d0, [x9]
MOV x9, #-264
ADD x9, x29, x9
LDR d1, [x9]
FSUB d0, d0, d1
MOV x9, #-272
ADD x9, x29, x9
STR d0, [x9]
// Declaracion inferida: restaF
MOV x9, #-272
ADD x9, x29, x9
LDR d0, [x9]
MOV x9, #-280
ADD x9, x29, x9
STR d0, [x9]
//  ------------- VisitMulDivModulo -------------
// Identificador: y con profundidad: 0
LDR d0, [x29, #-48]
MOV x9, #-288
ADD x9, x29, x9
STR d0, [x9]
// Identificador: x con profundidad: 0
LDR d0, [x29, #-40]
MOV x9, #-296
ADD x9, x29, x9
STR d0, [x9]
// MULTIPLICACION DE FLOAT64
MOV x9, #-288
ADD x9, x29, x9
LDR d0, [x9]
MOV x9, #-296
ADD x9, x29, x9
LDR d1, [x9]
FMUL d0, d0, d1
MOV x9, #-304
ADD x9, x29, x9
STR d0, [x9]
// Declaracion inferida: multiF
MOV x9, #-304
ADD x9, x29, x9
LDR d0, [x9]
MOV x9, #-312
ADD x9, x29, x9
STR d0, [x9]
//  ------------- VisitMulDivModulo -------------
// Identificador: z con profundidad: 0
LDR d0, [x29, #-56]
MOV x9, #-320
ADD x9, x29, x9
STR d0, [x9]
// Identificador: x con profundidad: 0
LDR d0, [x29, #-40]
MOV x9, #-328
ADD x9, x29, x9
STR d0, [x9]
// DIVISION DE FLOAT64
MOV x9, #-320
ADD x9, x29, x9
LDR d0, [x9]
MOV x9, #-328
ADD x9, x29, x9
LDR d1, [x9]
FDIV d0, d0, d1
MOV x9, #-336
ADD x9, x29, x9
STR d0, [x9]
// Declaracion inferida: divF
MOV x9, #-336
ADD x9, x29, x9
LDR d0, [x9]
MOV x9, #-344
ADD x9, x29, x9
STR d0, [x9]
//  ------------- VisitAddSub -------------
//  ------------- VisitAddSub -------------
// Identificador: a con profundidad: 1
LDR x0, [x29, #-16]
MOV x9, #-352
ADD x9, x29, x9
STR x0, [x9]
//  ------------- VisitMulDivModulo -------------
// Identificador: x con profundidad: 2
LDR d0, [x29, #-40]
MOV x9, #-360
ADD x9, x29, x9
STR d0, [x9]
// Identificador: b con profundidad: 2
LDR x0, [x29, #-24]
MOV x9, #-368
ADD x9, x29, x9
STR x0, [x9]
// MULTIPLICACION DE FLOAT64
MOV x9, #-360
ADD x9, x29, x9
LDR d0, [x9]
MOV x9, #-368
ADD x9, x29, x9
LDR x1, [x9]
SCVTF d1, x1
FMUL d0, d0, d1
MOV x9, #-376
ADD x9, x29, x9
STR d0, [x9]
// SUMA DE FLOAT64
MOV x9, #-352
ADD x9, x29, x9
LDR x0, [x9]
SCVTF d0, x0
MOV x9, #-376
ADD x9, x29, x9
LDR d1, [x9]
FADD d0, d0, d1
MOV x9, #-384
ADD x9, x29, x9
STR d0, [x9]
// Identificador: y con profundidad: 0
LDR d0, [x29, #-48]
MOV x9, #-392
ADD x9, x29, x9
STR d0, [x9]
// RESTA DE FLOAT64
MOV x9, #-384
ADD x9, x29, x9
LDR d0, [x9]
MOV x9, #-392
ADD x9, x29, x9
LDR d1, [x9]
FSUB d0, d0, d1
MOV x9, #-400
ADD x9, x29, x9
STR d0, [x9]
// Declaracion inferida: mezcla
MOV x9, #-400
ADD x9, x29, x9
LDR d0, [x9]
MOV x9, #-408
ADD x9, x29, x9
STR d0, [x9]
// Función embebida: println
// Identificador: suma
LDR x1, [x29, #-88]
// Preparando entero para imprimir
LDR x1, [x29, #-88]
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
// Identificador: resta
LDR x1, [x29, #-120]
// Preparando entero para imprimir
LDR x1, [x29, #-120]
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
// Identificador: multi
LDR x1, [x29, #-152]
// Preparando entero para imprimir
LDR x1, [x29, #-152]
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
// Identificador: div
LDR x1, [x29, #-184]
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
// Función embebida: println
// Identificador: mod
LDR x1, [x29, #-216]
// Preparando entero para imprimir
LDR x1, [x29, #-216]
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
// Identificador: sumaF
LDR d0, [x29, #-248]
FMOV x0, d0
// Preparando float para imprimir
LDR d0, [x29, #-248]
BL print_decimal
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
// Identificador: restaF
MOV x9, #-280
ADD x9, x29, x9
LDR d0, [x9]
FMOV x0, d0
// Preparando float para imprimir
MOV x9, #-280
ADD x9, x29, x9
LDR d0, [x9]
BL print_decimal
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
// Identificador: multiF
MOV x9, #-312
ADD x9, x29, x9
LDR d0, [x9]
FMOV x0, d0
// Preparando float para imprimir
MOV x9, #-312
ADD x9, x29, x9
LDR d0, [x9]
BL print_decimal
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
// Identificador: divF
MOV x9, #-344
ADD x9, x29, x9
LDR d0, [x9]
FMOV x0, d0
// Preparando float para imprimir
MOV x9, #-344
ADD x9, x29, x9
LDR d0, [x9]
BL print_decimal
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
// Identificador: mezcla
MOV x9, #-408
ADD x9, x29, x9
LDR d0, [x9]
FMOV x0, d0
// Preparando float para imprimir
MOV x9, #-408
ADD x9, x29, x9
LDR d0, [x9]
BL print_decimal
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
ADD sp, sp, #424
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
        print_decimal:
            // Guardar contexto
            stp x29, x30, [sp, #-16]!
            stp x19, x20, [sp, #-16]!
            stp x21, x22, [sp, #-16]!
            stp x23, x24, [sp, #-16]!

            // Verificar si es negativo
            fcmp d0, #0.0
            bge skip_minus

            // Imprimir el signo '-'
            mov x0, #1
            adr x1, minus_sign
            mov x2, #1
            mov x8, #64
            svc #0

            // Convertir a valor positivo
            fneg d0, d0

        skip_minus:
            // Parte entera
            fcvtzs x0, d0         // x0 = int(d0)
            bl print_entero

            // Imprimir punto decimal '.'
            mov x0, #1
            adr x1, dot_char
            mov x2, #1
            mov x8, #64
            svc #0

            // Parte fraccionaria: frac = d0 - floor(d0)
            frintm d4, d0         // d4 = floor(d0)
            fsub d2, d0, d4       // d2 = d0 - d4 (solo parte decimal)

            // Multiplicar por 1_000_000 (para 6 decimales)
            movz x1, #0x000F, lsl #16
            movk x1, #0x4240, lsl #0   // x1 = 1000000
            scvtf d3, x1               // d3 = 1000000.0
            fmul d2, d2, d3            // d2 = frac * 1_000_000
            frintn d2, d2              // redondear
            fcvtzs x0, d2              // x0 = int(frac * 1_000_000)

            // Preparar impresión con ceros a la izquierda
            mov x20, x0                // x20 = valor entero de parte decimal
            movz x21, #0x0001, lsl #16
            movk x21, #0x86A0, lsl #0  // x21 = 100000
            mov x23, #10              // constante para dividir

        leading_zero_loop:
            udiv x24, x20, x21        // x24 = x20 / x21
            cbnz x24, done_leading_zeros

            // Imprimir '0'
            mov x0, #1
            adr x1, zero_char
            mov x2, #1
            mov x8, #64
            svc #0

            udiv x21, x21, x23        // x21 /= 10
            cmp x21, #0
            beq print_remaining
            b leading_zero_loop

        done_leading_zeros:
            mov x0, x20
            bl print_entero
            b exit_function

        print_remaining:
            cmp x20, #0
            bne exit_function
            // Si parte decimal es cero, no imprimir más

        exit_function:
            // Restaurar contexto
            ldp x23, x24, [sp], #16
            ldp x21, x22, [sp], #16
            ldp x19, x20, [sp], #16
            ldp x29, x30, [sp], #16
            ret

        // Constantes utilizadas
        minus_sign: .ascii "-"
        dot_char:   .ascii "."
        zero_char:  .ascii "0"

        