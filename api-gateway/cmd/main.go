package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type Route struct {
	Name    string `mapstructure:"name"`
	Context string `mapstructure:"context"`
	Target  string `mapstructure:"target"`
}

func main() {
	log.SetOutput(os.Stdout)

	viper.AddConfigPath("./config") //Viper looks here for the files.
	viper.SetConfigType("yaml")     //Sets the format of the config file.
	viper.SetConfigName("default")  // So that Viper loads default.yml.
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Warning could not load configuration: %v", err)
	}
	viper.AutomaticEnv() // Merges any overrides set through env vars.

	gatewayConfig := &GatewayConfig{}

	err = viper.UnmarshalKey("gateway", gatewayConfig)
	if err != nil {
		panic(err)
	}

	log.Println("Initializing routes...")

	r := mux.NewRouter()

	for _, route := range gatewayConfig.Routes {
		// Returns a proxy for the target url.
		proxy, err := NewProxy(route.Target)
		if err != nil {
			panic(err)
		}
		// Just logging the mapping.
		log.Printf("Mapping '%v' | %v ---> %v", route.Name, route.Context, route.Target)
		// Maps the HandlerFunc fn returned by NewHandler() fn
		// that delegates the requests to the proxy.
		r.HandleFunc(route.Context+"/{targetPath:.*}", NewHandler(proxy))
	}

	log.Printf("Started server on %v", gatewayConfig.ListenAddr)
	log.Fatal(http.ListenAndServe(gatewayConfig.ListenAddr, r))
}

type GatewayConfig struct {
	ListenAddr string  `mapstructure:"listenAddr"`
	Routes     []Route `mapstructure:"routes"`
}

func NewProxy(targetUrl string) (*httputil.ReverseProxy, error) {
	target, err := url.Parse(targetUrl)
	if err != nil {
		return nil, err
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ModifyResponse = func(response *http.Response) error {
		dumpResponse, err := httputil.DumpResponse(response, false)
		if err != nil {
			return err
		}
		log.Println("Response: \r\n", string(dumpResponse))
		return nil
	}
	return proxy, nil
}

func NewHandler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = mux.Vars(r)["targetPath"]
		log.Println("Request URL: ", r.URL.String())
		p.ServeHTTP(w, r)
	}
}

// func JwtVerify(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		authHeader := r.Header.Get("Authorization")
// 		if authHeader == "" {
// 			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
// 			return
// 		}

// 		tokenString := strings.Split(authHeader, " ")[1]
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 			}
// 			return []byte("secret"), nil
// 		})

// 		if err != nil || !token.Valid {
// 			http.Error(w, "Invalid token", http.StatusUnauthorized)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

// func main() {

// 	// config.LoadConfig()
// 	// logger.InitLogger()
// 	viper.AddConfigPath("./config") //Viper looks here for the files.
// 	viper.SetConfigType("yaml")     //Sets the format of the config file.
// 	viper.SetConfigName("default")  // So that Viper loads default.yml.
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		log.Println("Warning could not load configuration", err.Error())
// 	}
// 	viper.AutomaticEnv()                         // Merges any overrides set through env vars.
// 	gatewayConfig := &GatewayConfig{}            // declare and get address.
// 	viper.UnmarshalKey("gateway", gatewayConfig) //Pass the address
// 	r := handlers.SetupRouter()
// 	// r := mux.NewRouter()
// 	log.Println("Starting API Gateway on port 8080...")
// 	if err := http.ListenAndServe(":8080", r); err != nil {
// 		log.Fatalf("Could not start server: %s\n", err)
// 	}
// }

// func NewProxy(targetUrl string) (*httputil.ReverseProxy, error) {
// 	target, err := url.Parse(targetUrl)
// 	if err != nil {
// 		return nil, err
// 	}
// 	proxy := httputil.NewSingleHostReverseProxy(target)
// 	proxy.ModifyResponse = func(response *http.Response) error {
// 		dumpedResponse, err := httputil.DumpResponse(response, false)
// 		if err != nil {
// 			return err
// 		}
// 		log.Println("Response: \r\n", string(dumpedResponse))
// 		return nil
// 	}
// 	return proxy, nil
// }
// func NewHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		r.URL.Path = mux.Vars(r)["targetPath"]
// 		log.Println("Request URL: ", r.URL.String())
// 		proxy.ServeHTTP(w, r)
// 	}
// }
