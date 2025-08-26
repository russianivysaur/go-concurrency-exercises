//////////////////////////////////////////////////////////////////////
//
// Given is a mock process which runs indefinitely and blocks the
// program. Right now the only way to stop the program is to send a
// SIGINT (Ctrl-C). Killing a process like that is not graceful, so we
// want to try to gracefully stop the process first.
//
// Change the program to do the following:
//   1. On SIGINT try to gracefully stop the process using
//          `proc.Stop()`
//   2. If SIGINT is called again, just kill the program (last resort)
//

package main
import "os"
import "os/signal"
func main() {
	// Create a process
	proc := MockProcess{}
    signalChannel := make(chan os.Signal)
    signal.Notify(signalChannel,os.Interrupt)
    go func(){
      count := 0
      for {
        select {
          case v,_ := <-signalChannel:
            if v==os.Interrupt{
              count++
            }
            if count > 1{
              os.Exit(1)
              return
            }
            go proc.Stop()
        }
      }
    }()
	// Run the process (blocking)
	proc.Run()
}
