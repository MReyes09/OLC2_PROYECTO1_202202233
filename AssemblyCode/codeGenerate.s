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
// cadena_print: "Boolean:"
STR x10, [SP, #-8]!
// Byte: 66 uso de Heap: 'B'
MOV w0, #66
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 111 uso de Heap: 'o'
MOV w0, #111
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
// Byte: 101 uso de Heap: 'e'
MOV w0, #101
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
LDR x0, [SP], #8
MOV X0, x0
BL print_cadena
MOV X0, #1
ADR X1, espacio_str
MOV X2, #1
MOV W8, #64
SVC #0
// Booleano: 1
MOV x0, #1
STR x0, [SP, #-8]!
LDR x0, [SP], #8
MOV X0, x0
BL print_booleano
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
// cadena_print: "Boolean2:"
STR x10, [SP, #-8]!
// Byte: 66 uso de Heap: 'B'
MOV w0, #66
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 111 uso de Heap: 'o'
MOV w0, #111
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
// Byte: 101 uso de Heap: 'e'
MOV w0, #101
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
// Byte: 50 uso de Heap: '2'
MOV w0, #50
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
LDR x0, [SP], #8
MOV X0, x0
BL print_cadena
MOV X0, #1
ADR X1, espacio_str
MOV X2, #1
MOV W8, #64
SVC #0
// Booleano: 0
MOV x0, #0
STR x0, [SP, #-8]!
LDR x0, [SP], #8
MOV X0, x0
BL print_booleano
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
// cadena_print: "Prueba Texto"
STR x10, [SP, #-8]!
// Byte: 80 uso de Heap: 'P'
MOV w0, #80
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 114 uso de Heap: 'r'
MOV w0, #114
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 117 uso de Heap: 'u'
MOV w0, #117
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 101 uso de Heap: 'e'
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 98 uso de Heap: 'b'
MOV w0, #98
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
// Byte: 84 uso de Heap: 'T'
MOV w0, #84
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 101 uso de Heap: 'e'
MOV w0, #101
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 120 uso de Heap: 'x'
MOV w0, #120
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 116 uso de Heap: 't'
MOV w0, #116
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
// Función embebida: println
// cadena_print: "Y numero"
STR x10, [SP, #-8]!
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
// Función embebida: println
// Entero: 15
MOV x0, #15
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
// cadena_print: "string y"
STR x10, [SP, #-8]!
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
// Byte: 114 uso de Heap: 'r'
MOV w0, #114
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 105 uso de Heap: 'i'
MOV w0, #105
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 110 uso de Heap: 'n'
MOV w0, #110
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 103 uso de Heap: 'g'
MOV w0, #103
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 32 uso de Heap: ' '
MOV w0, #32
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 121 uso de Heap: 'y'
MOV w0, #121
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
// Entero: 199
MOV x0, #199
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
// cadena_print: "double:"
STR x10, [SP, #-8]!
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
// Byte: 117 uso de Heap: 'u'
MOV w0, #117
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 98 uso de Heap: 'b'
MOV w0, #98
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 108 uso de Heap: 'l'
MOV w0, #108
STRB w0, [x10]
MOV x0, #1
ADD x10, x10, x0
// Byte: 101 uso de Heap: 'e'
MOV w0, #101
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
LDR x0, [SP], #8
MOV X0, x0
BL print_cadena
MOV X0, #1
ADR X1, espacio_str
MOV X2, #1
MOV W8, #64
SVC #0
// Float: 1.250000
MOVZ X0, #0, LSL #0
MOVK X0, #0, LSL #16
MOVK X0, #0, LSL #32
MOVK X0, #16372, LSL #48
STR x0, [SP, #-8]!
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
            stp x29, x30, [sp, #-16]!    
            stp x19, x20, [sp, #-16]!
            stp x21, x22, [sp, #-16]!
            stp x23, x24, [sp, #-16]!
            
            // Check if number is negative
            fmov x19, d0
            tst x19, #(1 << 63)       // Comprueba el bit de signo
            beq skip_minus

            // Print minus sign
            mov x0, #1
            adr x1, minus_sign
            mov x2, #1
            mov x8, #64
            svc #0

            // Make value positive
            fneg d0, d0

        skip_minus:
            // Convert integer part
            fcvtzs x0, d0             // x0 = int(d0)
            bl print_entero

            // Print dot '.'
            mov x0, #1
            adr x1, dot_char
            mov x2, #1
            mov x8, #64
            svc #0

            // Get fractional part: frac = d0 - float(int(d0))
            frintm d4, d0             // d4 = floor(d0)
            fsub d2, d0, d4           // d2 = d0 - floor(d0) (exact fraction)

            // Para 2.5, d2 debe ser exactamente 0.5

            // Multiplicar por 1_000_000 (6 decimales)
            movz x1, #0x000F, lsl #16
            movk x1, #0x4240, lsl #0   // x1 = 1000000
            scvtf d3, x1              // d3 = 1000000.0
            fmul d2, d2, d3           // d2 = frac * 1_000_000
            
            // Redondear al entero más cercano para evitar errores de precisión
            frintn d2, d2             // d2 = round(d2)
            fcvtzs x0, d2             // x0 = int(d2)

            // print ceros a la izquierda si es necesario
            mov x20, x0               // x20 = fracción entera
            movz x21, #0x0001, lsl #16
            movk x21, #0x86A0, lsl #0  // x21 = 100000
            mov x22, #0               // inicializar contador de ceros
            mov x23, #10              // constante para división

        leading_zero_loop:
            udiv x24, x20, x21        // x24 = x20 / x21
            cbnz x24, done_leading_zeros  // Si hay un dígito no cero, salir del bucle

            // print '0'
            mov x0, #1
            adr x1, zero_char
            mov x2, #1
            mov x8, #64
            svc #0

            udiv x21, x21, x23        // x21 /= 10
            add x22, x22, #1          // incrementar contador de ceros
            cmp x21, #0               // verificar si llegamos al final
            beq print_remaining       // si divisor es 0, saltar a print el resto
            b leading_zero_loop

        done_leading_zeros:
            // Print the remaining fractional part
            mov x0, x20
            bl print_entero
            b exit_function

        print_remaining:
            // Caso especial cuando la parte fraccionaria es 0 después de print ceros
            cmp x20, #0
            bne exit_function
            
            // Ya imprimimos todos los ceros necesarios
            // No hace falta print nada más

        exit_function:
            // Restore context
            ldp x23, x24, [sp], #16
            ldp x21, x22, [sp], #16
            ldp x19, x20, [sp], #16
            ldp x29, x30, [sp], #16
            ret
            minus_sign: .ascii "-"
            dot_char: .ascii "."
            zero_char: .ascii "0"
        
        

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
        