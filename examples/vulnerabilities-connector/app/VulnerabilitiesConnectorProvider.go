package app

type VulnerabilitiesConnectorProvider string

const (
	CROUD_STRIKE VulnerabilitiesConnectorProvider = "CroudStrike"
	NUCLEUS      VulnerabilitiesConnectorProvider = "Nucleus"
	QUALYS       VulnerabilitiesConnectorProvider = "Qualys"
	RAPID7       VulnerabilitiesConnectorProvider = "Rapid7"
	TANIUM       VulnerabilitiesConnectorProvider = "Tanium"
	TENABLE      VulnerabilitiesConnectorProvider = "Tenable"
)
