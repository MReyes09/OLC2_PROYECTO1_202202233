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
SUB SP, SP, #56
//  ------------- Operación NOT -------------
//  ------------- Operación NOT -------------
// Booleano: 1
MOV x1, #1
STR x1, [x29, #-16]
LDR x0, [x29, #-16]
EOR x0, x0, #1
STR x0, [x29, #-24]
LDR x0, [x29, #-24]
EOR x0, x0, #1
STR x0, [x29, #-32]
// Declaracion inferida: var1
LDR x1, [x29, #-32]
MOV x0, x1
STR x0, [x29, #-40]
// Función embebida: println
// Identificador: var1
LDR x1, [x29, #-40]
LDR x1, [x29, #-40]
MOV X0, x1
BL print_booleano
MOV X0, x1
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
ADD sp, sp, #56
LDP x29, x30, [sp], #16
RET

//Funciones De Impresion:

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
	