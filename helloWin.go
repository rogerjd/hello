package hello

import (
	"net/http"
	//C:\Users\roger\AppData\Local\Google\Cloud SDK\google-cloud-sdk\platform\google_appengine\google\appengine\datastore
	_ "golang.org/x/net/context" //as of 1.7 it is std lib, but gae may not be using latest(1.7)

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	//"appengine/datastore"
)

func init() {

	http.HandleFunc("/", handler)

}

type (
	Entity struct {
		Value string
	}
)

func handler(w http.ResponseWriter, r *http.Request) {
	//watch: file loc
	//fmt.Fprint(w, "Hello, world Windows!!!")

	ctx := appengine.NewContext(r)
	k := datastore.NewKey(ctx, "Entity", "", 0, nil)
	e := new(Entity)
	e.Value = "gae test"
	datastore.Put(ctx, k, e)
}
