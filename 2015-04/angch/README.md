This is a fun multiCPU solution.

`go run main.go myinput 6`

A number of goroutines, corresponding to the CPUs available, races to find if they can find an answer,
incrementing by num(goroutines), so we don't need complex coordination among the goroutines.

When an answer is found, other goroutines are signaled with the answer ("bid") and check if they can
find a better answer. If they can't, they quit, if they can, they return the answer to the caller.

The caller figures out the best answer among the goroutines and shows that.

This part may be tricky if one goroutine is faster than the others and found the larger answer when the
correct answer may be from a different goroutine, that's why we needed the second section ("waiting to give up")