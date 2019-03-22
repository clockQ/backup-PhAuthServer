package main

import (
	"os"
	"log"
	"net/http"
	"github.com/manyminds/api2go"
	"github.com/julienschmidt/httprouter"

	"github.com/alfredyang1986/BmServiceDef/BmConfig"
	"github.com/alfredyang1986/BmServiceDef/BmPodsDefine"
	"github.com/alfredyang1986/BmServiceDef/BmApiResolver"

	"github.com/PharbersDeveloper/PhAuthServer/PhFactory"
)

func main() {
	const (
		version  = "v0"
		confHome = "PH_AUTH_HOME"
	)
	// 本机测试，添加上
	//os.Setenv(confHome, "resources")
	log.Println("Pharbers Auth Server begins, version =", version)

	fac := PhFactory.PhTable{}
	var pod = BmPodsDefine.Pod{Name: "Pharbers Auth", Factory: fac}
	prodEnv := os.Getenv(confHome)
	pod.RegisterSerFromYAML(prodEnv + "/resource/service-def.yaml")

	var phRouter BmConfig.BmRouterConfig
	phRouter.GenerateConfig(confHome)

	addr := phRouter.Host + ":" + phRouter.Port
	log.Println("Pharbers Auth Server Listening on", addr)
	api := api2go.NewAPIWithResolver(version, &BmApiResolver.RequestURL{Addr: addr})
	pod.RegisterAllResource(api)

	pod.RegisterAllFunctions(version, api)
	pod.RegisterAllMiddleware(api)

	handler := api.Handler().(*httprouter.Router)
	pod.RegisterPanicHandler(handler)
	http.ListenAndServe(":"+phRouter.Port, handler)

	log.Println("Pharbers Auth Server begins, version =", version)
}
