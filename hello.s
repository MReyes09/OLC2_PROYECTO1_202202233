// hello.s — ARM64 "Hello, World!" (con syscall write y exit)
.global _start

.section .data
msg:    .asciz "Hola, mundo desde ARM64!\n"

.section .text
_start:
    mov     x0, 1              // file descriptor 1 = stdout
    ldr     x1, =msg           // dirección del mensaje
    mov     x2, 25             // longitud del mensaje
    mov     x8, 64             // syscall write (64 en ARM64)
    svc     0                  // llamada al sistema

    mov     x0, 0              // código de salida
    mov     x8, 93             // syscall exit (93 en ARM64)
    svc     0
