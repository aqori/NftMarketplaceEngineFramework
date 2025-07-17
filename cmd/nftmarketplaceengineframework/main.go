// cmd/nftmarketplaceengineframework/main.go
package main

import (
"flag"
"log"
"os"

"nftmarketplaceengineframework/internal/nftmarketplaceengineframework"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nftmarketplaceengineframework.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
