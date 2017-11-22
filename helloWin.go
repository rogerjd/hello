package hello

import (
	"fmt"
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
	//Entity a struct to r/w to/from datastore
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
	k2, err := datastore.Put(ctx, k, e)
	if err != nil {
		fmt.Fprint(w, "error occurred")
	}
	fmt.Fprint(w, "no error")

	e2 := new(Entity)
	datastore.Get(ctx, k2, e2)
	fmt.Fprintf(w, "%#v", e2)
}
