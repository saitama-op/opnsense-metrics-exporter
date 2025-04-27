package config

type ProductCheck struct {
	APIVersion          string        `json:"api_version"`
	Connection          string        `json:"connection"`
	DowngradePackages   []interface{} `json:"downgrade_packages"`
	DownloadSize        string        `json:"download_size"`
	LastCheck           string        `json:"last_check"`
	NeedsReboot         string        `json:"needs_reboot"`
	NewPackages         []interface{} `json:"new_packages"`
	OSVersion           string        `json:"os_version"`
	ProductID           string        `json:"product_id"`
	ProductTarget       string        `json:"product_target"`
	ProductVersion      string        `json:"product_version"`
	ProductABI          string        `json:"product_abi"`
	ReinstallPackages   []interface{} `json:"reinstall_packages"`
	RemovePackages      []interface{} `json:"remove_packages"`
	Repository          string        `json:"repository"`
	UpgradeMajorMessage string        `json:"upgrade_major_message"`
	UpgradeMajorVersion string        `json:"upgrade_major_version"`
	UpgradeNeedsReboot  string        `json:"upgrade_needs_reboot"`
	UpgradePackages     []interface{} `json:"upgrade_packages"`
	UpgradeSets         []interface{} `json:"upgrade_sets"`
}

type Product struct {
	ProductABI            string        `json:"product_abi"`
	ProductArch           string        `json:"product_arch"`
	ProductCheck          ProductCheck  `json:"product_check"`
	ProductConflicts      string        `json:"product_conflicts"`
	ProductCopyrightOwner string        `json:"product_copyright_owner"`
	ProductCopyrightURL   string        `json:"product_copyright_url"`
	ProductCopyrightYears string        `json:"product_copyright_years"`
	ProductEmail          string        `json:"product_email"`
	ProductHash           string        `json:"product_hash"`
	ProductID             string        `json:"product_id"`
	ProductLatest         string        `json:"product_latest"`
	ProductLicense        []interface{} `json:"product_license"`
	ProductLog            int           `json:"product_log"`
	ProductMirror         string        `json:"product_mirror"`
	ProductName           string        `json:"product_name"`
	ProductNickname       string        `json:"product_nickname"`
	ProductRepos          string        `json:"product_repos"`
	ProductSeries         string        `json:"product_series"`
	ProductTier           string        `json:"product_tier"`
	ProductTime           string        `json:"product_time"`
	ProductVersion        string        `json:"product_version"`
	ProductWebsite        string        `json:"product_website"`
}

type OPNsenseStatus struct {
	APIVersion          string        `json:"api_version"`
	Connection          string        `json:"connection"`
	DowngradePackages   []interface{} `json:"downgrade_packages"`
	DownloadSize        string        `json:"download_size"`
	LastCheck           string        `json:"last_check"`
	NeedsReboot         string        `json:"needs_reboot"`
	NewPackages         []interface{} `json:"new_packages"`
	OSVersion           string        `json:"os_version"`
	ProductID           string        `json:"product_id"`
	ProductTarget       string        `json:"product_target"`
	ProductVersion      string        `json:"product_version"`
	ProductABI          string        `json:"product_abi"`
	ReinstallPackages   []interface{} `json:"reinstall_packages"`
	RemovePackages      []interface{} `json:"remove_packages"`
	Repository          string        `json:"repository"`
	UpgradeMajorMessage string        `json:"upgrade_major_message"`
	UpgradeMajorVersion string        `json:"upgrade_major_version"`
	UpgradeNeedsReboot  string        `json:"upgrade_needs_reboot"`
	UpgradePackages     []interface{} `json:"upgrade_packages"`
	UpgradeSets         []interface{} `json:"upgrade_sets"`
	Product             Product       `json:"product"`
	AllPackages         []interface{} `json:"all_packages"`
	AllSets             []interface{} `json:"all_sets"`
	StatusMsg           string        `json:"status_msg"`
	Status              string        `json:"status"`
}
