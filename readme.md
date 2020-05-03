- Module
    - Types
        - ./x/nameservice/types/types.go
    - Key
        - ./x/nameservice/types/key.go
    - Errors
        - ./x/nameservice/types/errors.go
    - Expected Keepers
        - ./x/nameservice/types/expected_keepers.go
    - Keeper
        - ./x/nameservice/keeper/keeper.go
    - Mesages and Handlers 
        - ./x/nameservice/types/msgs.go
        - ./x/nameservice/handler.go
    - Queries 
        - ./x/nameservice/types/querier.go
        - ./x/nameservice/keeper/querier.go
    - Alias
        - ./x/nameservice/alias.go
    - Codec File
        - ./x/nameservice/types/codec.go
    - Module CLI
        - ./x/nameservice/client/cli/query.go
        - ./x/nameservice/client/cli/tx.go
    - Rest Interface
        - ./x/nameservice/client/rest/rest.go
        - ./x/nameservice/client/rest/query.go
        - ./x/nameservice/client/rest/tx.go
    - AppModule Interface
        - ./x/nameservice/module.go
    - Genesis 
        - ./x/nameservice/genesis.go
        - ./x/nameservice/types/genesis.go
- Application
    - App
        - ./app/app.go
    - Entry Points
        - ./cmd/nsd/main.go
        - ./cmd/nscli/main.go




```bash
scaffold app lvl-1 [user] [repo] [flags]
```

```bash
scaffold module [user] [repo] nameservice
```




```plantuml
@startuml
package Types <<Frame>> {
    class Whois {
        + Value: string
        + Owner: sdk.AccAddress
        + Price: sdk.Coins
        + String()
        
    }

    class Types {
        + NewWhois() Whois
    }
}

@enduml
```

```plantuml
@startuml

package Keeper <<Frame>> {

    class Keeper {
        + CoinKeeper: types.BankKeeper
        + storeKey: sdk.StoreKey
        + cdc: *codec.Codec

        + ResolveName(ctx: sdk.Context, name: string): string
        + SetName(ctx: sdk.Context, name: string, value: string)
        + HasOwner(ctx: sdk.Context, name: string): bool
        + GetOwner(ctx: sdk.Context, name: string): sdk.AccAddress
        + SetOwner(ctx: sdk.Context, name: string, owner: sdk.AccAddress)
        + GetPrice(ctx: sdk.Context, name: string): sdk.Coins
        + SetPrice(ctx: sdk.Context, name: string, price: sdk.Coins)
        + IsNamePresent(ctx: sdk.Context, name: string): bool
        + GetNamesIterator(ctx: sdk.Context): sdk.Iterator
        - SetWhois(ctx: sdk.Context, name: string, whois: types.Whois)
        - GetWhois(ctx: sdk.Context, name: string): types.Whois
        - DeleteWhois(ctx: sdk.Context, name: string)
    }

    class KeeperClass {
        + NewKeeper(coinKeeper: bank.Keeper, storeKey: sdk.StoreKey, cdc: *codec.Codec): Keeper
    }

    Keeper -up[hidden]-KeeperClass
}

@enduml
```



```plantuml
@startuml

package Msg <<Frame>> {

    class MsgSetName {
        + Name: string        
	    + Value: string        
	    + Owner: sdk.AccAddress
        
        + Route(): string
        + Type(): string
        + ValidateBasic(): error
        + GetSignBytes(): []byte
        + GetSigners(): []sdk.AccAddress
    }

    class MsgBuyName {
        + Name: string
        + Bid: sdk.Coins
        + Buyer: sdkAccAddress

        + Route(): string
        + Type(): string
        + ValidateBasic(): error
        + GetSignBytes(): []byte
        + GetSigners(): []sdk.AccAddress
    }
    
    class MsgDeleteName {
        + Name: string
        + Owner: sdk.AccAddress

        + Route(): string
        + Type(): string
        + ValidateBasic(): error
        + GetSignBytes(): []byte
        + GetSigners(): []sdk.AccAddress
    }

    class Msg {
        + NewMsgSetName(name: string, value: string, owner: sdk.AccAddress): MsgSetName
        + NewMsgBuyName(name: string, bid: sdk.Coins, buyer: sdk.AccAddress): MsgBuyName
        + NewMsgDeleteName(name: string, owner: sdk.AccAddress): MsgDeleteName
    }

    class Handler {
        + NewHandler(keeper: Keeper): sdk.Handler
        - handleMsgSetName(ctx: sdk.Context, keeper: Keeper, msg: MsgSetName): (*sdk.Result, error)
        - handleMsgBuyName(ctx: sdk.Context, keeper: Keeper, msg: MsgBuyName): (*sdk.Result, error)
        - handleMsgDeleteName(ctx: sdk.Context, keeper: Keeper, msg: MsgDeleteName): (*sdk.Result, error)
    }

    MsgSetName -up[hidden]- Msg
    MsgBuyName -up[hidden]- Msg
    MsgDeleteName -up[hidden]- Msg

    MsgBuyName -down[hidden]- Handler
    
}
@enduml
```



```plantuml
@startuml

package Query <<Frame>> {

    class QueryResResolve {
        + Value: string
        + String(): string
    }

    class QueryResNames {
        + String(): string
    }

    class Querier {
        + NewQuerier(keeper: Keeper): sdk.Querier
        - queryResolve(ctx: sdk.Context, path: []string, req: abci.RequestQuery, keeper: Keeper): ([]byte, error)
        - queryWhois(ctx: sdk.Context, path: []string, req: abci.RequestQuery, keeper: Keeper): ([]byte, error)
        - queryNames(ctx: sdk.Context, req: abci.RequestQuery, keeper: Keeper): ([]byte, error)
    }  
}
@enduml
```



```plantuml
@startuml

package ModuleCLI <<Frame>> {
    class Query {
        + GetQueryCmd(queryRoute: string, cdc: *codec.Codec): *cobra.Command
        + GetCmdResolveName(queryRoot: string, cdc: *codec.Codec): *cobra.Command
        + GetCmdWhois(queryRoot: string, cdc: *codec.Codec): *cobra.Command
        + GetCmdNames(queryRoute: string, cdc: *codec.Codec): *cobra.Command
    }

    class Tx {
        + GetTxCmd(queryRoot: string, cdc: *codec.Codec): *cobra.Command
        + GetCmdBuyName(cdc: *codec.Codec): *cobra.Command
        + GetCmdSetName(cdc: *codec.Codec): *cobra.Command
        + GetCmdDeleteName(cdc: *codec.Codec): *cobra.Command
    }
}
@enduml
```



```plantuml
@startuml

package ModuleRest <<Frame>> {

     class setNameReq{
        + BaseReq: rest.BaseReq
        + Name: string
        + Value: string
        + Owner: string
    }

    class buyNameReq {
        + BaseReq: rest.BaseReq
        + Name: string      
        + Amount: string      
        + Buyer: string      
    }

    class Rest {
        + RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, storeName string)
    }

    class Query {
        + resolveNameHandler(cliCtx: context.CLIContext, storeName: string): http.HandlerFunc
        + whoIsHandler(cliCtx: context.CLIContext, storeName: string): http.HandlerFunc
        + namesHandler(cliCtx: context.CLIContext, storeName: string): http.HandlerFunc
    }

    class Tx {
        + buyNameHandler(cliCtx: context.CLIContext): http.HandlerFunc
        + setNameHandler(cliCtx: context.CLIContext): http.HandlerFunc
        + deleteNameHandler(cliCtx: context.CLIContext): http.HandlerFunc
    }

   
}
@enduml
```



```plantuml
@startuml

package Module <<Frame>> {
    
    class AppModule {
        Name(): string
        RegisterInvariants(_: sdk.InvariantRegistry)
        Route(): string
        NewHandler(): sdk.Handler
        QuerierRoute(): string
        NewQuerierHandler(): sdk.Querier
        InitGenesis(ctx: sdk.Context, data: json.RawMessage): []abci.ValidatorUpdate
        ExportGenesis(ctx: sdk.Context): json.RawMessage
        BeginBlock(ctx: sdk.Context, req: abci.RequestBeginBlock)
        EndBlock(_: sdk.Context, _: abci.RequestEndBlock): []abci.ValidatorUpdate
    }

    class AppModuleBasic{
        Name(): string
        RegisterCodec(cdc: *codec.Codec)
        DefaultGenesis(): json.RawMessage
        ValidateGenesis(bz: json.RawMessage): error
        RegisterRESTRoutes(ctx: context.CLIContext, rtr: *mux.Router)
        GetTxCmd(cdc: *codec.Codec): *cobra.Command
        GetQueryCmd(cdc: *codec.Codec): *cobra.Command
    }

    class Module {
        + NewAppModule(k: Keeper, bankKeeper: bank.Keeper): AppModule
    }

}
@enduml
```



