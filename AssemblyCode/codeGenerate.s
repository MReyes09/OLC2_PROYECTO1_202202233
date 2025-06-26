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
SUB SP, SP, #376
// Entero: 4
MOV x1, #4
// Declaracion inferida: x
MOV x0, x1
STR x0, [x29, #-16]
// Entero: 2
MOV x1, #2
// Declaracion inferida: y
MOV x0, x1
STR x0, [x29, #-24]
// Float: 1.500000
MOVZ X0, #0, LSL #0
MOVK X0, #0, LSL #16
MOVK X0, #0, LSL #32
MOVK X0, #16376, LSL #48
FMOV d0, x0
// Declaracion inferida: f
STR d0, [x29, #-32]
// Float: 2.500000
MOVZ X0, #0, LSL #0
MOVK X0, #0, LSL #16
MOVK X0, #0, LSL #32
MOVK X0, #16388, LSL #48
FMOV d0, x0
// Declaracion inferida: g
STR d0, [x29, #-40]
// cadena: "Hola "
MOV x11, x10
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
// Byte: 0 uso de Heap: '\x00'
MOV w0, #0
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Declaracion inferida: s1
MOV x0, x11
STR x0, [x29, #-48]
// cadena: " Mundo!"
MOV x11, x10
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 77 uso de Heap: 'M'
MOV w0, #77
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 117 uso de Heap: 'u'
MOV w0, #117
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
// Declaracion inferida: s2
MOV x0, x11
STR x0, [x29, #-56]
//  ------------- VisitAddSub -------------
// Identificador: x con profundidad: 0
LDR x0, [x29, #-16]
STR x0, [x29, #-64]
//  ------------- VisitAddSub -------------
// Identificador: y con profundidad: 1
LDR x0, [x29, #-24]
STR x0, [x29, #-72]
// Entero: 3
MOV x1, #3
STR x1, [x29, #-80]
// SUMA DE ENTEROS
LDR x2, [x29, #-72]
LDR x1, [x29, #-80]
ADD x0, x2, x1
STR x0, [x29, #-88]
// SUMA DE ENTEROS
LDR x2, [x29, #-64]
LDR x1, [x29, #-88]
ADD x0, x2, x1
STR x0, [x29, #-96]
// Declaracion inferida: res1
LDR x1, [x29, #-96]
MOV x0, x1
STR x0, [x29, #-104]
//  ------------- VisitAddSub -------------
//  ------------- VisitAddSub -------------
// Identificador: x con profundidad: 1
LDR x0, [x29, #-16]
STR x0, [x29, #-112]
// Identificador: y con profundidad: 1
LDR x0, [x29, #-24]
STR x0, [x29, #-120]
// RESTA DE ENTEROS
LDR x2, [x29, #-112]
LDR x1, [x29, #-120]
SUB x0, x2, x1
STR x0, [x29, #-128]
// Entero: 7
MOV x1, #7
STR x1, [x29, #-136]
// SUMA DE ENTEROS
LDR x2, [x29, #-128]
LDR x1, [x29, #-136]
ADD x0, x2, x1
STR x0, [x29, #-144]
// Declaracion inferida: res2
LDR x1, [x29, #-144]
MOV x0, x1
STR x0, [x29, #-152]
//  ------------- VisitAddSub -------------
//  ------------- VisitAddSub -------------
// Identificador: f con profundidad: 1
LDR d0, [x29, #-32]
STR d0, [x29, #-160]
// Identificador: g con profundidad: 1
LDR d0, [x29, #-40]
STR d0, [x29, #-168]
// SUMA DE FLOAT64
LDR d0, [x29, #-160]
LDR d1, [x29, #-168]
FADD d0, d0, d1
STR d0, [x29, #-176]
// Float: 1.000000
MOVZ X0, #0, LSL #0
MOVK X0, #0, LSL #16
MOVK X0, #0, LSL #32
MOVK X0, #16368, LSL #48
FMOV d0, x0
STR d0, [x29, #-184]
// RESTA DE FLOAT64
LDR d0, [x29, #-176]
LDR d1, [x29, #-184]
FSUB d0, d0, d1
STR d0, [x29, #-192]
// Declaracion inferida: res3
LDR d0, [x29, #-192]
STR d0, [x29, #-200]
//  ------------- VisitAddSub -------------
//  ------------- VisitAddSub -------------
// Identificador: x con profundidad: 1
LDR x0, [x29, #-16]
STR x0, [x29, #-208]
// Identificador: f con profundidad: 1
LDR d0, [x29, #-32]
STR d0, [x29, #-216]
// SUMA DE FLOAT64
LDR x0, [x29, #-208]
SCVTF d0, x0
LDR d1, [x29, #-216]
FADD d0, d0, d1
STR d0, [x29, #-224]
//  ------------- VisitAddSub -------------
// Identificador: g con profundidad: 1
LDR d0, [x29, #-40]
STR d0, [x29, #-232]
// Identificador: y con profundidad: 1
LDR x0, [x29, #-24]
STR x0, [x29, #-240]
// RESTA DE FLOAT64
LDR d0, [x29, #-232]
LDR x1, [x29, #-240]
SCVTF d1, x1
FSUB d0, d0, d1
STR d0, [x29, #-248]
// SUMA DE FLOAT64
LDR d0, [x29, #-224]
LDR d1, [x29, #-248]
FADD d0, d0, d1
MOV x9, #-256
ADD x9, x29, x9
STR d0, [x9]
// Declaracion inferida: res4
MOV x9, #-256
ADD x9, x29, x9
LDR d0, [x9]
MOV x9, #-264
ADD x9, x29, x9
STR d0, [x9]
//  ------------- VisitAddSub -------------
// Identificador: x con profundidad: 0
LDR x0, [x29, #-16]
MOV x9, #-272
ADD x9, x29, x9
STR x0, [x9]
//  ------------- VisitAddSub -------------
// Identificador: f con profundidad: 1
LDR d0, [x29, #-32]
MOV x9, #-280
ADD x9, x29, x9
STR d0, [x9]
//  ------------- VisitAddSub -------------
// Identificador: g con profundidad: 2
LDR d0, [x29, #-40]
MOV x9, #-288
ADD x9, x29, x9
STR d0, [x9]
// Float: 1.000000
MOVZ X0, #0, LSL #0
MOVK X0, #0, LSL #16
MOVK X0, #0, LSL #32
MOVK X0, #16368, LSL #48
FMOV d0, x0
MOV x9, #-296
ADD x9, x29, x9
STR d0, [x9]
// RESTA DE FLOAT64
MOV x9, #-288
ADD x9, x29, x9
LDR d0, [x9]
MOV x9, #-296
ADD x9, x29, x9
LDR d1, [x9]
FSUB d0, d0, d1
MOV x9, #-304
ADD x9, x29, x9
STR d0, [x9]
// SUMA DE FLOAT64
MOV x9, #-280
ADD x9, x29, x9
LDR d0, [x9]
MOV x9, #-304
ADD x9, x29, x9
LDR d1, [x9]
FADD d0, d0, d1
MOV x9, #-312
ADD x9, x29, x9
STR d0, [x9]
// SUMA DE FLOAT64
MOV x9, #-272
ADD x9, x29, x9
LDR x0, [x9]
SCVTF d0, x0
MOV x9, #-312
ADD x9, x29, x9
LDR d1, [x9]
FADD d0, d0, d1
MOV x9, #-320
ADD x9, x29, x9
STR d0, [x9]
// Declaracion inferida: res5
MOV x9, #-320
ADD x9, x29, x9
LDR d0, [x9]
MOV x9, #-328
ADD x9, x29, x9
STR d0, [x9]
//  ------------- VisitAddSub -------------
// Identificador: s1 con profundidad: 0
LDR x0, [x29, #-48]
MOV x9, #-336
ADD x9, x29, x9
STR x0, [x9]
// Identificador: s2 con profundidad: 0
LDR x0, [x29, #-56]
MOV x9, #-344
ADD x9, x29, x9
STR x0, [x9]
// CONCATENACION DE CADENAS
MOV x9, #-336
ADD x9, x29, x9
LDR x0, [x9]
MOV x9, #-344
ADD x9, x29, x9
LDR x1, [x9]
ADR x10, buffer
BL concatenar_cadena
MOV x9, #-352
ADD x9, x29, x9
STR x0, [x9]
ADR x10, buffer
// Declaracion inferida: saludo
MOV x9, #-352
ADD x9, x29, x9
LDR x11, [x9]
MOV x0, x11
MOV x9, #-360
ADD x9, x29, x9
STR x0, [x9]
// Función embebida: println
// Identificador: res1
LDR x1, [x29, #-104]
// Preparando entero para imprimir
LDR x1, [x29, #-104]
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
// Identificador: res2
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
// Identificador: res3
LDR d0, [x29, #-200]
FMOV x0, d0
// Preparando float para imprimir
LDR d0, [x29, #-200]
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
// Identificador: res4
MOV x9, #-264
ADD x9, x29, x9
LDR d0, [x9]
FMOV x0, d0
// Preparando float para imprimir
MOV x9, #-264
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
// Identificador: res5
MOV x9, #-328
ADD x9, x29, x9
LDR d0, [x9]
FMOV x0, d0
// Preparando float para imprimir
MOV x9, #-328
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
// Identificador: saludo
MOV x9, #-360
ADD x9, x29, x9
LDR x0, [x9]
MOV x11, x0
// Preparando string para imprimir
MOV x9, #-360
ADD x9, x29, x9
LDR x11, [x9]
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
ADD sp, sp, #376
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
	

        .align 2
        concatenar_cadena:
            stp x29, x30, [sp, #-16]!
            stp x19, x20, [sp, #-16]!
            stp x21, x22, [sp, #-16]!
            mov x19, x0
            mov x20, x1 
            mov x21, x10 
        copiar_primera:
            ldrb w22, [x19]
            cbz w22, primera_terminada
            strb w22, [x10]
            add x19, x19, #1
            add x10, x10, #1
            b copiar_primera
        primera_terminada:
        copiar_segunda:
            ldrb w22, [x20]
            strb w22, [x10]
            cbz w22, concatenacion_completa
            add x20, x20, #1
            add x10, x10, #1
            b copiar_segunda
        concatenacion_completa:
            add x10, x10, #1
            mov x0, x21
            ldp x21, x22, [sp], #16
            ldp x19, x20, [sp], #16
            ldp x29, x30, [sp], #16
            ret
        