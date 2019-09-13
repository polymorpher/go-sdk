package cmd

import (
	"fmt"

	color "github.com/fatih/color"
)

const (
	hmyDocsDir      = "hmycli-docs"
	defaultNodeAddr = "http://localhost:9500"
)

var (
	g           = color.New(color.FgGreen).SprintFunc()
	cookbookDoc = fmt.Sprintf(`
Cookbook of usage, note that every subcommand recognizes a '--help' flag

%s
hmy --node="https://api.s1.b.hmny.io/" balance <SOME_ONE_ADDRESS>

%s
hmy --node="https://api.s1.b.hmny.io" blockchain transaction-by-hash <SOME_TRANSACTION_HASH>

%s
hmy keys list

%s
./hmy --node="https://api.s0.b.hmny.io/" transfer \
    --from one1yc06ghr2p8xnl2380kpfayweguuhxdtupkhqzw \
    --to one1q6gkzcap0uruuu8r6sldxuu47pd4ww9w9t7tg6 \
    --from-shard 0 --to-shard 1 --amount 200
`,
		g("1. Check Balances"),
		g("2. Check completed transaction"),
		g("3. List local keys"),
		g("4. Sending a transaction"),
	)
)