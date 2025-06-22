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
// cadena_print: "5 * 2 ="
STR x10, [SP, #-8]!
// Byte: 53 uso de Heap: '5'
MOV w0, #53
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 42 uso de Heap: '*'
MOV w0, #42
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
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 61 uso de Heap: '='
MOV w0, #61
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
// Entero: 5
MOV x0, #5
STR x0, [SP, #-8]!
// Entero: 2
MOV x0, #2
STR x0, [SP, #-8]!
LDR x1, [SP], #8
LDR x0, [SP], #8
MUL x0, x0, x1
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
// cadena_print: "10 / 2 ="
STR x10, [SP, #-8]!
// Byte: 49 uso de Heap: '1'
MOV w0, #49
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 48 uso de Heap: '0'
MOV w0, #48
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 47 uso de Heap: '/'
MOV w0, #47
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
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 61 uso de Heap: '='
MOV w0, #61
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
// Entero: 10
MOV x0, #10
STR x0, [SP, #-8]!
// Entero: 2
MOV x0, #2
STR x0, [SP, #-8]!
LDR x1, [SP], #8
LDR x0, [SP], #8
SDIV x0, x0, x1
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
// cadena_print: "9 % 4 ="
STR x10, [SP, #-8]!
// Byte: 57 uso de Heap: '9'
MOV w0, #57
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 37 uso de Heap: '%'
MOV w0, #37
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 52 uso de Heap: '4'
MOV w0, #52
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 61 uso de Heap: '='
MOV w0, #61
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
// Entero: 9
MOV x0, #9
STR x0, [SP, #-8]!
// Entero: 4
MOV x0, #4
STR x0, [SP, #-8]!
LDR x1, [SP], #8
LDR x0, [SP], #8
SDIV x2, x0, x1
MSUB x0, x2, x1, x0
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
// cadena_print: "5.5 * 2.0 ="
STR x10, [SP, #-8]!
// Byte: 53 uso de Heap: '5'
MOV w0, #53
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 46 uso de Heap: '.'
MOV w0, #46
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 53 uso de Heap: '5'
MOV w0, #53
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 42 uso de Heap: '*'
MOV w0, #42
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
// Byte: 46 uso de Heap: '.'
MOV w0, #46
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 48 uso de Heap: '0'
MOV w0, #48
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 61 uso de Heap: '='
MOV w0, #61
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
// Float: 5.500000
MOVZ X0, #0, LSL #0
MOVK X0, #0, LSL #16
MOVK X0, #0, LSL #32
MOVK X0, #16406, LSL #48
STR x0, [SP, #-8]!
// Float: 2.000000
MOVZ X0, #0, LSL #0
MOVK X0, #0, LSL #16
MOVK X0, #0, LSL #32
MOVK X0, #16384, LSL #48
STR x0, [SP, #-8]!
LDR d1, [SP], #8
LDR d0, [SP], #8
FMUL d0, d0, d1
STR d0, [SP, #-8]!
LDR d0, [SP], #8
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
// cadena_print: "10.0 / 4.0 ="
STR x10, [SP, #-8]!
// Byte: 49 uso de Heap: '1'
MOV w0, #49
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 48 uso de Heap: '0'
MOV w0, #48
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 46 uso de Heap: '.'
MOV w0, #46
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 48 uso de Heap: '0'
MOV w0, #48
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 47 uso de Heap: '/'
MOV w0, #47
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 52 uso de Heap: '4'
MOV w0, #52
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 46 uso de Heap: '.'
MOV w0, #46
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 48 uso de Heap: '0'
MOV w0, #48
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 61 uso de Heap: '='
MOV w0, #61
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
// Float: 10.000000
MOVZ X0, #0, LSL #0
MOVK X0, #0, LSL #16
MOVK X0, #0, LSL #32
MOVK X0, #16420, LSL #48
STR x0, [SP, #-8]!
// Float: 4.000000
MOVZ X0, #0, LSL #0
MOVK X0, #0, LSL #16
MOVK X0, #0, LSL #32
MOVK X0, #16400, LSL #48
STR x0, [SP, #-8]!
LDR d1, [SP], #8
LDR d0, [SP], #8
FDIV d0, d0, d1
STR d0, [SP, #-8]!
LDR d0, [SP], #8
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
// cadena_print: "5 * 2.5 ="
STR x10, [SP, #-8]!
// Byte: 53 uso de Heap: '5'
MOV w0, #53
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 42 uso de Heap: '*'
MOV w0, #42
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
// Byte: 46 uso de Heap: '.'
MOV w0, #46
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 53 uso de Heap: '5'
MOV w0, #53
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 61 uso de Heap: '='
MOV w0, #61
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
// Entero: 5
MOV x0, #5
STR x0, [SP, #-8]!
// Float: 2.500000
MOVZ X0, #0, LSL #0
MOVK X0, #0, LSL #16
MOVK X0, #0, LSL #32
MOVK X0, #16388, LSL #48
STR x0, [SP, #-8]!
LDR d1, [SP], #8
LDR x0, [SP], #8
SCVTF d0, x0
FMUL d0, d0, d1
STR d0, [SP, #-8]!
LDR d0, [SP], #8
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

        