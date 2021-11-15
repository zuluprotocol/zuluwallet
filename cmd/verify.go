package cmd

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"code.vegaprotocol.io/vegawallet/cmd/cli"
	"code.vegaprotocol.io/vegawallet/cmd/flags"
	"code.vegaprotocol.io/vegawallet/cmd/printer"
	"code.vegaprotocol.io/vegawallet/crypto"
	"github.com/spf13/cobra"
)

var (
	verifyMessageLong = cli.LongDesc(`
		Verify a message against a signature.

		The signature has to be generated by Ed25519 cryptographic algorithm.
	`)

	verifyMessageExample = cli.Examples(`
		# Verify the signature of a message
		vegawallet verify --message MESSAGE --signature SIGNATURE --pubkey PUBKEY
	`)
)

type VerifyMessageHandler func(*crypto.VerifyMessageRequest) (bool, error)

func NewCmdVerifyMessage(w io.Writer, rf *RootFlags) *cobra.Command {
	return BuildCmdVerifyMessage(w, crypto.VerifyMessage, rf)
}

func BuildCmdVerifyMessage(w io.Writer, handler VerifyMessageHandler, rf *RootFlags) *cobra.Command {
	f := &VerifyMessageFlags{}

	cmd := &cobra.Command{
		Use:     "verify",
		Short:   "Verify a message against a signature",
		Long:    verifyMessageLong,
		Example: verifyMessageExample,
		RunE: func(_ *cobra.Command, _ []string) error {
			req, err := f.Validate()
			if err != nil {
				return err
			}

			isValid, err := handler(req)
			if err != nil {
				return err
			}

			switch rf.Output {
			case flags.InteractiveOutput:
				PrintVerifyMessageResponse(w, isValid)
			case flags.JSONOutput:
				return printer.FprintJSON(w, struct {
					IsValid bool `json:"isValid"`
				}{
					IsValid: isValid,
				})
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&f.PubKey,
		"pubkey", "k",
		"",
		"Public key associated to the signature (hex-encoded)",
	)
	cmd.Flags().StringVarP(&f.Message,
		"message", "m",
		"",
		"Message to be verified (base64-encoded)",
	)
	cmd.Flags().StringVarP(&f.Signature,
		"signature", "s",
		"",
		"Signature of the message (base64-encoded)",
	)

	return cmd
}

type VerifyMessageFlags struct {
	Signature string
	Message   string
	PubKey    string
}

func (f *VerifyMessageFlags) Validate() (*crypto.VerifyMessageRequest, error) {
	req := &crypto.VerifyMessageRequest{}

	if len(f.PubKey) == 0 {
		return nil, flags.FlagMustBeSpecifiedError("pubkey")
	}
	req.PubKey = f.PubKey

	if len(f.Signature) == 0 {
		return nil, flags.FlagMustBeSpecifiedError("signature")
	}
	decodedSignature, err := base64.StdEncoding.DecodeString(f.Signature)
	if err != nil {
		return nil, flags.MustBase64EncodedError("signature")
	}
	req.Signature = decodedSignature

	if len(f.Message) == 0 {
		return nil, flags.FlagMustBeSpecifiedError("message")
	}
	decodedMessage, err := base64.StdEncoding.DecodeString(f.Message)
	if err != nil {
		return nil, flags.MustBase64EncodedError("message")
	}
	req.Message = decodedMessage

	return req, nil
}

func PrintVerifyMessageResponse(w io.Writer, isValid bool) {
	p := printer.NewInteractivePrinter(w)

	if isValid {
		p.CheckMark().SuccessText("Valid signature").NextSection()
	} else {
		p.CrossMark().DangerText("Invalid signature").NextSection()
	}

	p.BlueArrow().InfoText("Sign a message").NextLine()
	p.Text("To sign a message, see the following command:").NextSection()
	p.Code(fmt.Sprintf("%s sign --help", os.Args[0])).NextSection()
}
