package main
import (
	"github.com/spf13/viper"
	//"fmt"
)

var pathBasedRoutes map[string]string
var siteConfigs []SiteConfig
var SiteIdParam string

type WebHostingConfig struct {
	SiteIdParam string        `json:"siteidparam"`
	Sites       []SiteConfig  `json:"sites"`
}

type SiteConfig struct {
	Name          string         `json:"name"`
	RoutingMethod RoutingMethod  `json:"routingmethod"`
}

type RoutingMethod struct {
	Method  string        `json:"method"`
	UrlPath string        `json:"urlpath"`
	SiteId  string        `json:"siteid"`
}

const UrlPathBasedRouting = "urlpath"

func InitConfigs(serverHome string) {
	readWebappConfigs(serverHome)
	pathBasedRoutes = make(map[string]string)

	for i := range siteConfigs {
		if siteConfigs[i].RoutingMethod.Method == UrlPathBasedRouting {
			pathBasedRoutes[siteConfigs[i].RoutingMethod.SiteId] = serverHome + "/webapps/" + siteConfigs[i].RoutingMethod.UrlPath
		}
	}
}

func AddParamBasedRoute(name string, route string) {
	pathBasedRoutes[name] = route
}

func getParambasedRoute(name string) string {
	if pathBasedRoutes[name] != "" {
		return pathBasedRoutes[name]
	}else {
		return ServerHome + "/webapps/default"
	}
}

func readWebappConfigs(serverHome string) {

	viper.New()
	viper.SetConfigName("config")
	viper.AddConfigPath(serverHome + "/webapps")

//	err := viper.ReadInConfig() // Find and read the config file
//	if err != nil {             // Handle errors reading the config file
//		panic(fmt.Errorf("Fatal error config file: %s \n", err))
//	}
//    xx := viper.Get("sites")
//	print(xx["1"].(string))

}

func getWebappConfigs() [] SiteConfig {
	return siteConfigs
}