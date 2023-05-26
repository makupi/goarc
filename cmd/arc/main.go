package main

import (
	"fmt"
	"github.com/makupi/goarc/pkg/decoder"
	"github.com/makupi/goarc/pkg/encoder"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "encode",
				Action: func(cCtx *cli.Context) error {
					enc, err := encoder.ReserveAddressFromCID(cCtx.Args().Get(0))
					if err != nil {
						return err
					}
					fmt.Println(enc)
					return nil
				},
			},
			{
				Name: "decode",
				Action: func(cCtx *cli.Context) error {
					dec, err := decoder.ParseASAUrl(cCtx.Args().Get(0), cCtx.Args().Get(1))
					if err != nil {
						return err
					}
					fmt.Println(dec)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
