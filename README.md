# Goroutines

Goroutines: Goroutines are lightweight, independently scheduled functions that can be executed concurrently. They are similar to threads but are managed by Go’s runtime, making them more efficient. Goroutines enable concurrent execution in Go.

WaitGroup: Block execution until all goroutines are done.
It sets a counter to the number of goroutines to wait for and decrements it by one when a goroutine is done.

Channels: Channels are a communication mechanism in Go that allows goroutines to communicate and synchronize. They facilitate coordination between concurrent tasks, which is crucial for many concurrent programs.

Go also supports parallelism through multiple ways, including the use of multiple CPU cores. Go’s runtime scheduler can automatically distribute goroutines across available CPU cores to achieve parallel execution when possible.

Use waitgroups when you don't want goroutine to return some data.
Use channels when you do want goroutine to return some data.
