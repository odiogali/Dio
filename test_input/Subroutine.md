2024/10/23 14:27
Status: #idea
Tags: #academic 
# What is a Subroutine?
Subroutines, which are routines dedicated to focused or shorter tasks, occur in nearly all embedded code, often to perform initialization routines or to handle algorithms that require handcrafted assembly. They allow code reuse using varying argument values.
___
# About
- There are generally two types:
	- Open
		- Code is inserted inline whereever subroutine is invoked
			- Usualy using a macro preprocessor
		- Arguments are passed in/out using registers
		- Efficient, since overhead of branching and returing is avoided
		- Suitable only for fairly short subroutines
	- Closed
		- There is only one copy of it in memory (machine code for the routine appears only once in RAM)
			- Leads to more compact machine code than with open routines
		- When invoked, control "jumps" to first instruction of the routine
			- PC is loaded with the address of the first instruction
		- When finished, control returns to the next instruction in the calling code
			- PC is loaded with the return address
		- Arguments are placed in registers or on the stack
		- Slower than open routines because of call/return overhead
- Subroutines should not change the state of the machine for the calling code
	- When invoked, a subroutine should save any registers it uses on the stack
	- When it returns, it should restore the original values of the registers
- Arguments to subroutines are considered local variables
	- Subroutine may change their values
# Open (Inline) Subroutines
- Usually implemented using macros
- E.g. Cube function
```ARMASM
define (comment)

comment(cube(1 = input register, 2 = output register))
define(cube, `mul $2, $1, $1 
    mul $2, $1, $2`)
	
	.global main
main:
	stp x29, x30, [sp, -16]!
	...
	mov x19, 8
	cube(x19, x20)
```
- m4 expands this to:
```ARMASM
	.global main
main:
	stp x29, x30, [sp, -16]!
	...
	mov x19, 8
	mul x20, x19, x19 
	mul x20, x20, x19
	...
```
# Closed Subroutine
- General form:
```ARMASM
label:
	stp x29, x30, [sp, alloc]!
	mov x29, sp
	... 
	ldp x29, x30, [sp], -alloc
	ret
```
- `label`: names the subroutine
- `alloc`: the number of bytes (negated) to allocate for the subroutine's stack frame
	- SP must be quadword aligned
	- Minimum of 16 bytes
# Subroutine Linkage
- A subroutine may be invoked using the branch and link instruction: `bl`
	- Form: `bl subroutine_label`
	- Stores the return address into the *link register*: x30
		- Return address is PC + 4, which points to the instruction immediately following `bl`
	- Transfers control to address specified by the label
		- Loads PC register with address of the subroutine's first instruction
- Use the `ret` instruction to return froma subroutine back to the calling code
	- Transfers control to the address stored in the link register (x30)
		- i.e. jumps to the instruction immediately following the original `bl` in calling code
- C example:
```C
int main(){
	...
	func1();
	...
}

void func1(){
	...
	func2();
	...
}

void func2(){
	...
}
```
- ASM example:
```ARMASM
main:
	stp x29, x30, [sp, -16]!
	moiv x29, sp
	...
	bl func1
	...
	ldp x29, x30, [sp], 16
	ret

func1:
	stp x29, x30, [sp, -16]!
	mov x29, sp
	...
	bl func2
	...
	ldp x29, x30, [sp], 16
	ret

func2:                                  // sometimes called the leaf subroutine
	stp x29, x30, [sp, -16]!
	mov x29, sp
	...
	ldp x29, x30, [sp], 16
	ret
```
- The `stp` instructions create a frame record in each function's stack frame
	- Safely stores the LR (x30) in case it is changed by a `bl` in the body of the function
		- Is restored by the `ldp` instruction just before the `ret`
- The FP and stored FP values in the frame records form a linked list
	- E.g: the stack while `func2()` is executing![[Pasted image 20241025142418.png]]
# Saving and Restoring Registers
- A called function must save/restore the state of the calling code
	- If it uses any of the registers x19 - x28, it must save their data to the stack at the beginning of the function 
		- Are called "callee-saved registers"
	- The function must restore the data in these registers just before it returns
- E.g:
```ARMASM
x19_size = 8
alloc = -(16 + x19_size) & -16
dealloc = -alloc
x19_save = 16

func2:
	stp x29, x30, [sp, alloc]!
	mov x29, sp
	str x19, [x29, x19_save]    // save x19 in RAM
	...
	mov x19, 13                 // use x19
	...
	ldr x19, [x29, x19_save]    // restore x19
	ldp x29, x30, [sp], dealloc
	ret
``` 
- Note that the callee can also use registers x9 - x15
	- By convention, these registers are not saved/restored by the called function
		- Thus, they are only safe to use in calling code between function calls
	- The calling code can save these registers to the stack, if it is necessary to preserve their value over a function call
		- Are "caller-saved registers"
# Passing Arguments
- 8 or fewer arguments can be passed into a function using registers x0 - x7
	- ints, short ints, and chars use w0 - w7
	- long ints use x0 - x7
- C example:
```C
void sum(int a, int b){
	register int i;
	i = a + b;
	...
}

int main(){
	sum(3,4);
	...
}
```
- ASM example:
```ARMASM
define(i_r, w9) // using temp register so we don't have to worry about restore
sum:
	stp x29, x30, [sp, -16]!
	mov x29, sp
	
	add i_r, w0, w1
	...
	ldp x29, x30, [sp], 16
	ret

main:
	stp x29, x30, [sp, -16]!
	mov x29, sp
	
	mov w0, 3 // set up first arg
	mov w1, 4 // set up second arg
	bl sum
	...
	ldp x29, x30, [sp], 16
	ret
```
- Note that the subroutine is free to overwrite registers x0 - x7 as it executes
	- Register's contents are not preserved over a function call
# Pointer Arguments
- In calling code, the *address* of a variable is passed to the subroutine
	- Implies that the variable must be in RAM, not in a register as registers do not have addresses
- The called subroutine dereferences the address to manipulate the variable being pointed to
	- Usually with a `ldr` or `str` instruction
```C
int main(){
	int a = 5, b =7;
	swap(&a, &b); // Gets the address of a variable - "address of"
	...
}

void swap(int *x, int* y){ // * refers to a pointer 
	register int temp;
	
	temp = *x; // Dereference x and store in temp
	*x = *y;
	*y = temp;
}
```
```ARMASM
	a_size = 4
	b_size = 4
	alloc = -(16 + a_size + b_size) & -16
	dealloc = -alloc
	a_s = 16
	b_s = 20
	...
main:
	stp x29, x30, [sp, alloc]!
	mov x29, sp
	
	mov w19, 5
	str w19, [x29, a_s]
	mov w20, 7
	str w20, [x29, b_s]
	
	add x0, x29, a_s              // first arg is the addr of a
	add x1, x29, b_s              // second arg is the addr of b
	bl swap
	...

define(temp_r, w9)
swap:
	stp x29, x30, [sp, -16]!
	mov x29, sp
	
	ldr temp_r, [x0]
	ldr w10, [x1]
	str w10, [x0]
	str temp_r, [x1]
	
	ldp x29, x30, [sp], 16
	ret
```
# `argc` and `argv`
- The shell passes `argc` in x0
- The address for `argv[]` is passed in x1
# Returning Integers
- A function returns:
	- Long integers in x0
	- Ints, short ints, and chars in w0
- E.g. Cube function in C
```C
int cube(int x); // function prototype

int main(){
	register int result;
	
	result = cube(3);
	...
}

int cube(int x){
	return x * x * x;
}
```
```ARMASM
define(result_r, w19)
...
main:
	stp x29, x30, [sp, -16]!
	mov x29, sp
	
	mov w0, 3             
	bl cube
	mov result_r, w0
	...

cube:
	stp x29, x30, [sp, -16]!
	mov x29, sp
	
	mul w9, w0, w0
	mul w0, w9, w0
	
	ldp x29, x30, [sp], 16
	ret
```
# Returning Structures
- In C, a function may return a [[Structures|struct]] by value
```C
struct mystruct {
	long int i;
	long int j;
};

struct mystruct init(){ // return type of 'struct mystruct'
	struct mystruct lvar;
	lvar.i = 0;
	lvar.j = 0;
	return lvar;
}

int main(){
	struct mystruct b;
	
	b = init();
	...
}
```
- In general, a struct is too big to return in x0 or w0
	- Thus another return mechanism is required 
- The calling code provides memory on the stack to store the returned result
	- The address of this memory is put into x8 priorn to the function call
		- x8 is the "indirect result location register"
	- The called subroutine writes to memory at this address using x8 as a pointer to it
# Optimizing Leaf Subroutines
- Leaf subroutines do not call any other subroutines
	- Are leaf nodes on a tree structure diagram
- A frame record is not pushed onto the stack
	- Since the routine does not do a `bl`, LR won't change
	- Since the routine does not call a subroutine, FP won't change
	- Thus we can eliminate the usual `stp`/`ldp` instructions
- If one uses only the registers x0 - x7 and x9 - x15, then a stack frame is not pushed at all
	- No need to save/restore registers
	- No stack variables are used
- E.g. optimized cube function
```ARMASM
cube: 
	mul w9, w0, w0
	mul w0, w9, w0
	ret
```
# Subroutines With 9 or More Arguments
- Arguments beyond the 8th are passed on the stack
	- The calling code allocates memory at the top of the stack and writes the "spilled" argument values there
		- By convention, each argument is allocated 8 bytes
	- The callee reads this memory using the appropriate offset
```C
// Function prototype
int sum (int a1, int a2, int a3, int a4, int a5, int a6, int a7, int a8, int           int a9, int a10);

int main(){
	register int result;
	
	result = sum(10, 20, 30, 40, 50, 60, 70, 80, 90, 100);
}

int sum (int a1, int a2, int a3, int a4, int a5, int a6, int a7, int a8, int           int a9, int a10){
	return a1 + a2 + a3 + a4 + a5 + a6 + a7 + a8 + a9 + a10;
}
```
```ARMASM
define(result_r, w19)

	spilled_mem_size = 16
	alloc = -spilled_mem_size & -16
	dealloc = -alloc
	
	.global main
	
main:
	stp x29, x30, [sp, -16]!
	mov x29, sp
	
	// Set up first 8 args
	mov w0, 10
	mov w1, 20
	mov w2, 30
	mov w3, 40
	mov w4, 50
	mov w5, 60
	mov w6, 70
	mov w7, 80
	
	// Allocate memory for args 9 and 10
	add sp, sp, alloc
	
	// Write spilled arguments to top of stack
	mov w9, 90
	str w9, [sp, 0] // set up arg 9
	mov w9, 100
	str w9, [sp, 8] // set up arg 10
	
	bl sum // call sum function
	mov result_r, w0 
	
	// Deallocate memory for spilled arguments
	add sp, sp, dealloc
	...

	arg9_s = 16
	arg10_s = 24
	
sum:
	stp x29, x30, [sp, -16]!
	mov x29, sp
	
	// add first 8 arguments
	add w0, w0, w1
	add w0, w0, w2
	add w0, w0, w3
	add w0, w0, w4
	add w0, w0, w5
	add w0, w0, w6
	add w0, w0, w7
	
	// add 9th and 10th args
	ldr w9, [x29, arg9_s]
	add w0, w0, w9
	ldr w9, [x29, arg10_s]
	add w0, w0, w9
	
	ldp x29, x30, [sp], 16
	ret
```
- When in `sum()`, the stack appears as:![[Pasted image 20241101171639.png]]
# Questions