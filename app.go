package main
import(
	"fmt"
	"github.com/golang/protobuf/proto"
	"crypto/tls"
	"net/http"
	"encoding/base64"
	example "github.com/abhatnag/main/protospb/example_ID"
)


func GetProtos(brandIDs ...string){
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	for _,brand :=range brandIDs{
		message := &example.Message{
			ID: brand,
		}
		data, err := proto.Marshal(message)
		if err != nil {
				panic(err)
				return
			}
	urlEncoding := base64.RawURLEncoding.EncodeToString(data)
	fmt.Printf("For brandID: %s the proto is, %s " ,brand,urlEncoding)
	}
	
}

func main(){
	GetProtos("c40d3c81-ca1a-4800-a26f-b58d0314159f","zzzzz","e5afa9dc-ba60-4f2c-a75d-8ee81cea1b91", "")
}

