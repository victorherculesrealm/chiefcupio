package main
import ("context";"fmt";"os";"os/signal";"time")
func worker(ctx context.Context,id string){fmt.Printf("[%s] Worker started\n",id);select{case <-ctx.Done():fmt.Printf("[%s] Worker stopped\n",id);case <-time.After(2*time.Second):fmt.Printf("[%s] Worker completed\n",id)}}
func main(){appID:="alert-daemon-89c7d1";fmt.Printf("Starting %s...\n",appID);ctx,cancel:=signal.NotifyContext(context.Background(),os.Interrupt);defer cancel();worker(ctx,appID);fmt.Println("Done.")}
