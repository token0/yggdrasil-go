// +build linux

package defaults

// Sane defaults for the Linux platform. The "default" options may be
// may be replaced by the running configuration.
func GetDefaults() platformDefaultParameters {
	return platformDefaultParameters{
		// Admin
		DefaultAdminListen: "unix:///opt/var/run/yggdrasil.sock",

		// Configuration (used for yggdrasilctl)
		DefaultConfigFile: "/opt/etc/yggdrasil.conf",

		// Multicast interfaces
		DefaultMulticastInterfaces: []string{
			".*",
		},

		// TUN/TAP
		MaximumIfMTU:     65535,
		DefaultIfMTU:     65535,
		DefaultIfName:    "auto",
		DefaultIfTAPMode: false,
	}
}
