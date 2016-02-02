/**
 * Created with IntelliJ IDEA.
 * User: medcl
 * Date: 13-11-8
 * Time: 下午6:32
 * To change this template use File | Settings | File Templates.
 */
package http

import (
	"net/http"
	"github.com/pantsing/gograce/ghttp"
	log "github.com/cihub/seelog"
	. "github.com/medcl/gopa/src/config"
)

func index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("gopa!"))
	w.Write([]byte("\nversion: "+config.Version))
	w.Write([]byte("\ncluster: "+config.ClusterConfig.Name))
}

var config RuntimeConfig
func Start(runtimeConfig *RuntimeConfig) {
	config=*runtimeConfig
	http.HandleFunc("/", index)
	ghttp.ListenAndServe(":8001", nil)
	log.Info("http server is up,http://localhost:8001/")

}
