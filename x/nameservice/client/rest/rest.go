package rest

import (
	"fmt"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

const (
	restName = "name"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string) {
	r.HandleFunc(fmt.Sprintf("/%s/names", storeName), namesHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/names", storeName), buyNameHandler(cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("/%s/names", storeName), setNameHandler(cliCtx)).Methods("PUT")
	r.HandleFunc(fmt.Sprintf("/%s/names/{%s}", storeName, restName), resolveNameHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/names/{%s}/whois", storeName, restName), whoIsHandler(cliCtx, storeName)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/%s/names", storeName), deleteNameHandler(cliCtx)).Methods("DELETE")
}

// // RegisterRoutes registers nameservice-related REST handlers to a router
// func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
// 	registerQueryRoutes(cliCtx, r)
// 	registerTxRoutes(cliCtx, r)
// }
