package ngapType

import "free5gc-cli/lib/aper"

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

const (
	CancelAllWarningMessagesPresentTrue aper.Enumerated = 0
)

type CancelAllWarningMessages struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}
