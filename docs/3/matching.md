**What happens if you remove the go-command from the Seek call in the main function?**

Then Anna always sends to Bob and Cody always to Dave whilst noone reads Evas message. The reason is the the parallel execution with the go command makes it random which thread reads the message so that it's different every time. Now they are all executed in a determined order.

**What happens if you switch the declaration wg := new(sync.WaitGroup) to var wg sync.WaitGroup and the parameter wg sync.WaitGroup to wg sync.WaitGroup?**

wg becomes a copy/instance instead of pointer so the wg sent into the message with not have the same values as the one in the main routine. Therefore wg will never be done so it'll continue to wait until all threads sleep and a deadlock has arosen.

**What happens if you remove the buffer on the channel match?**

It becomes an unbuffered channel. The last message creates a deadlock, as there is no buffer the message is expected to be read instantly which won't happen as the main routine reads the message after the waitgroup is finished.

**What happens if you remove the default-case from the case-statement in the main function?**

It makes no difference as the last message always will be recieved. If there is an even amount of people tho then there isn't an extra message waiting. This will lead to a deadlock as all messages have been sent but the main method will no longer go on.