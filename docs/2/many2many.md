**What happens if you switch the order of the statements wgp.Wait() and close(ch) in the end of the main function?**

You get the error: panic: send on closed channel. The channel is closed by the main thread before waiting for all producers, so they don't get the time required to send their messages before the channel is closed.

**What happens if you move the close(ch) from the main function and instead close the channel in the end of the function Produce?**

Same error as above, the first Produce closes the channel so it's closed before all the other Produce routines finish sending strings

**What happens if you remove the statement close(ch) completely?**

It does not make a different, the code still works as exspected. The recievers continue to listen for more data but the main thread stops waiting (due to waitgroup) and therefore terminates the program

**What happens if you increase the number of consumers from 2 to 4?**

The messages are split up on 4 consumers instead of 2 so the code executes approximately at half the time as each routine recieves half the amount of strings

**Can you be sure that all strings are printed before the program stops?**

No the waitgroup waits for the producers to finish but not the consumers so it is possible that the main thread finishes before all messages have been recieved and therefore terminates the program