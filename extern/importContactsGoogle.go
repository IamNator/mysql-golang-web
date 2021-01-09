package extern

//GET https://people.googleapis.com/v1/%5BRESOURCENAME%5D/connections?key=[YOUR_API_KEY] HTTP/1.1
//
//Authorization: Bearer [YOUR_ACCESS_TOKEN]
//Accept: application/json
import (
	"context"
	"fmt"
	"github.com/IamNator/mysql-golang-web/models"
	"io"
	"log"
	"net/http"
	"time"
)

type db models.DBData



func (db *db) getContact(){
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(5 * time.Second))
	defer cancelFunc()

	respBody, err := fetchContact(ctx)
	if err != nil {
		log.Println("request err : ", err.Error())
	}

	fmt.Println(respBody)
}


func fetchContact(ctx context.Context) (* io.ReadCloser, error) {

	rwChan := make(chan io.ReadCloser)
	errContextTimedOut := fmt.Errorf("context timed out")

	go func() {
		resp, err := http.Get("https://people.googleapis.com/v1/%5BRESOURCENAME%5D/connections?key=[YOUR_API_KEY] HTTP/1.1")
		//req, err := http.NewRequestWithContext()
		if err != nil {
			log.Println("request err : ", err.Error())
		}
		rwChan <- resp.Body
	}()

	for {
		select {
			case rw := <- rwChan :
				return &rw, nil
			case <- ctx.Done() :
				return nil, errContextTimedOut
		}
	}
}