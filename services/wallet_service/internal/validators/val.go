package val

import (
	"fmt"
	"regexp"
)

func ValidateWalletAddress(walletAddress string) error {
	re := regexp.MustCompile(`^0x[a-zA-Z0-9]{40}$`)

	if !re.MatchString(walletAddress) {
		return fmt.Errorf("Invalid wallet address")
	}
	return nil
}
