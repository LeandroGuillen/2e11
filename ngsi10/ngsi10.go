package ngsi10

import (
	"github.com/codegangsta/martini"
	"net/http"
	"encoding/json"
	"log"
// 	"fmt"
// 	"io/ioutil"
)


func GetIndividualContextEntity(params martini.Params) string {
	return "GET " + params["EntityID"]
}

func PutIndividualContextEntity(params martini.Params) string {
	return "PUT " + params["EntityID"]
}

func PostIndividualContextEntity(params martini.Params) string {
	return "POST " + params["EntityID"]
}

func PostUpdateContextRequest(req *http.Request) string {
	decoder := json.NewDecoder(req.Body)
	var t UpdateContextRequest
	err := decoder.Decode(&t)
	if(err != nil) {
		panic(err)
	}
	log.Println(t)
	return t.ToString()
}



// {
//     "contextElements": [
//         {
//             "type": "Room",
//             "isPattern": "false",
//             "id": "Room1",
//             "attributes": [
//             {
//                 "name": "temperature",
//                 "type": "centigrade",
//                 "value": "23"
//             },
//             {
//                 "name": "pressure",
//                 "type": "mmHg",
//                 "value": "720"
//             }
//             ]
//         }
//     ],
//     "updateAction": "APPEND"
// }