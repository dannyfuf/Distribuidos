package main
import (
	"fmt"
	"math/rand"
	"time"

	"src/common"
	
)

func main() {
	

	s1 := rand.NewSource(time.Now().UnixNano())
	
	mensaje := "Informante"
	i := 0

	for i<10 {

		if mensaje == "Informante"{
			
			r1 := rand.New(s1)
			rand := r1.Intn(3)
			
			var port string = common.Get_env_var("FULCRUM_PORT")
			var ip string	
			
			if rand == 0{
				ip = common.Get_env_var("IP_SERVER_18")
				//ip = "172.17.0.3"
			} else if rand == 1{

				ip = common.Get_env_var("IP_SERVER_19")
				//ip = "172.17.0.4"
			} else if rand == 2{

				ip = common.Get_env_var("IP_SERVER_20")
				//ip = "172.17.0.5"
			}

			fmt.Printf("%s:%s\n",ip,port)
		}
		i++
	}
}