// HACK Assembly Language example
// Draws a rectangle at the top-left corner of the screen.

   @0               // this is an "A" instruction
   D=M              // this is a "C" instruction
   @INFINITE_LOOP
   D;JLE            // this is also a "C" instruction
   @counter         
   M=D
   @SCREEN          // this is also an "A" instruction
   D=A
   @address
   M=D
(LOOP)              // this is a Label
   @address
   A=M
   M=-1
   @address
   D=M
   @32
   D=D+A
   @address
   M=D
   @counter
   MD=M-1
   @LOOP
   D;JGT
(INFINITE_LOOP)
   @INFINITE_LOOP
   0;JMP