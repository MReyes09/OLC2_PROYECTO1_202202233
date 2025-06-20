package funciones

import (
	"strings"
)

type EstandarFunc struct {
	funcionesUsadas map[string]bool
}

func NewEstandarFunc() *EstandarFunc {
	return &EstandarFunc{
		funcionesUsadas: make(map[string]bool),
	}
}

func (l *EstandarFunc) Usar(funcion string) {
	l.funcionesUsadas[funcion] = true
}

func (l *EstandarFunc) ObtenerDefinicionFuncion() string {
	definiciones := []string{}
	for funcion := range l.funcionesUsadas {
		if definicion, ok := DefinicionFunciones[funcion]; ok {
			definiciones = append(definiciones, definicion)
		}
	}
	return strings.Join(definiciones, "\n")
}

var DefinicionFunciones = map[string]string{
	"print_entero": `
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
	`,
	"print_cadena": `
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
	`,
	"print_booleano": `
            .align 2
            print_booleano:
                stp x29, x30, [sp, #-16]!
                cmp x0, #0
                bne print_true
                adr x1, falso
                b print_bool_fin
            .align 2
            print_true:
                adr x1, verdadero
            .align 2
            print_bool_fin:
                mov x0, #1
                mov x2, #5
                mov w8, #64
                svc #0
                ldp x29, x30, [sp], #16
                ret
            .align 2
            verdadero:
                .ascii ""truee""
            .align 2
            falso:
                .ascii ""false""
        `,
	"print_caracter": `
            .align 2
            print_caracter:
                stp x29, x30, [sp, #-16]!
                sub sp, sp, #16
                strb w0, [sp]          
                mov x0, #1             
                mov x1, sp            
                mov x2, #1            
                mov w8, #64           
                svc #0
                add sp, sp, #16
                ldp x29, x30, [sp], #16
                ret
        `,
	"concatenar_cadena": `
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
        `,
	"comparar_cadenas": `
        .align 2
        comparar_cadenas:
            // Guardar registros
            stp x29, x30, [sp, #-16]!
            stp x19, x20, [sp, #-16]!
            stp x21, x22, [sp, #-16]!
            mov x19, x0  
            mov x20, x1
        bucle_comparacion:
            ldrb w21, [x19]
            ldrb w22, [x20]
            cmp w21, w22
            bne cadenas_diferentes
            cbz w21, cadenas_iguales
            add x19, x19, #1
            add x20, x20, #1
            b bucle_comparacion   
        cadenas_diferentes:
            cmp w21, w22
            blt primera_menor  // Si byte1 < byte2
            mov x0, #1
        primera_menor:
            mov x0, #-1
            b fin_comparacion
        cadenas_iguales:
            mov x0, #0
        fin_comparacion:
            // Restaurar registros y retornar
            ldp x21, x22, [sp], #16
            ldp x19, x20, [sp], #16
            ldp x29, x30, [sp], #16
            ret
        `,
	"strconv_atoi": `
        .align 2
        strconv_atoi:
            mov x2, x0      
            mov x3, #0    
            mov x4, #0    
            
            ldrb w5, [x2]
            cmp w5, #45     
            bne check_plus
            mov x4, #1      
            add x2, x2, #1
            b start_conversion
            
        check_plus:
            cmp w5, #43    
            bne start_conversion
            add x2, x2, #1  
        start_conversion:
            mov x3, #0      
        conversion_loop:
            ldrb w5, [x2]   
            cmp w5, #0      
            beq check_sign
            
            sub w5, w5, #48 

            mov x6, #10
            mul x3, x3, x6
            add x3, x3, x5
            
            add x2, x2, #1
            b conversion_loop
            
        check_sign:
            cmp x4, #1
            bne atoi_end
            neg x3, x3
            
        atoi_end:
            mov x0, x3
            ret
        `,
	"print_decimal": `
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
            minus_sign: .ascii ""-""
            dot_char: .ascii "".""
            zero_char: .ascii ""0""
        
        `,
}
