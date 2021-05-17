package ipLink

import (
	"hcc/harp/lib/config"
	"hcc/harp/lib/iputil"
	"hcc/harp/lib/vnstat"
	"strconv"
	"strings"
)

func getIfaceVNUM(ip string) (vnum int) {
	var ifaceVNUM = 0

	ipSplit := strings.Split(ip, ".")
	for _, ipSplited := range ipSplit {
		ipSplitedInt, _ := strconv.Atoi(ipSplited)
		ifaceVNUM += ipSplitedInt
	}

	return ifaceVNUM
}

func isHarpInternalDeviceExist(ip string) bool {
	err := runIP("link show " +
		HarpInternalPrefix + strconv.Itoa(getIfaceVNUM(ip)))
	if err != nil {
		return false
	}

	return true
}

func addHarpInternalDevice(ip string) error {
	err := runIP("link add link " + config.AdaptiveIP.InternalIfaceName + " " +
		HarpInternalPrefix + strconv.Itoa(getIfaceVNUM(ip)) +
		" address " + generateMACAddress(ip) +
		" type macvlan")
	if err != nil {
		return err
	}

	return nil
}

func deleteHarpInternalDevice(ip string) error {
	err := runIP("link delete " +
		HarpInternalPrefix + strconv.Itoa(getIfaceVNUM(ip)))
	if err != nil {
		return err
	}

	return nil
}

func upHarpInternalDevice(ip string) error {
	err := runIP("link set dev " +
		HarpInternalPrefix + strconv.Itoa(getIfaceVNUM(ip)) + " up")
	if err != nil {
		return err
	}

	return nil
}

func downHarpInternalDevice(ip string) error {
	err := runIP("link set dev " +
		HarpInternalPrefix + strconv.Itoa(getIfaceVNUM(ip)) + " down")
	if err != nil {
		return err
	}

	return nil
}

func setIPtoHarpInternalDevice(ip string, cidr int) error {
	err := runIP("address add " +
		ip + "/" + strconv.Itoa(cidr) +
		" dev " +
		HarpInternalPrefix + strconv.Itoa(getIfaceVNUM(ip)))
	if err != nil {
		return err
	}

	return nil
}

// SetHarpInternalDevice : Setting up harp internal gateway device
func SetHarpInternalDevice(ip string, netmask string) error {
	if isHarpInternalDeviceExist(ip) {
		err := deleteHarpInternalDevice(ip)
		if err != nil {
			return err
		}
	}

	err := addHarpInternalDevice(ip)
	if err != nil {
		return err
	}

	mask, err := iputil.CheckNetmask(netmask)
	if err != nil {
		return err
	}
	maskLen, _ := mask.Size()
	err = setIPtoHarpInternalDevice(ip, maskLen)
	if err != nil {
		return err
	}

	err = upHarpInternalDevice(ip)
	if err != nil {
		return err
	}

	vnstat.ScheduleUpdateVnStat(HarpInternalPrefix+strconv.Itoa(getIfaceVNUM(ip)), true)

	return nil
}

// UnsetHarpInternalDevice : Delete harp internal gateway device
func UnsetHarpInternalDevice(ip string) error {
	if !isHarpInternalDeviceExist(ip) {
		return nil
	}

	err := downHarpInternalDevice(ip)
	if err != nil {
		return err
	}

	err = deleteHarpInternalDevice(ip)
	if err != nil {
		return err
	}

	vnstat.RemoveUpdateVnStat(HarpInternalPrefix + strconv.Itoa(getIfaceVNUM(ip)))

	return nil
}

// AddOrDeleteIPToHarpExternalDevice : Add/Delete AdaptiveIP address to/from external interface
func AddOrDeleteIPToHarpExternalDevice(ip string, netmask string, isAdd bool) error {
	mask, err := iputil.CheckNetmask(netmask)
	if err != nil {
		return err
	}
	maskLen, _ := mask.Size()

	var addArg string
	if isAdd {
		addArg = "add"
	} else {
		addArg = "del"
	}
	err = runIP("address " + addArg + " " +
		ip + "/" + strconv.Itoa(maskLen) +
		" dev " +
		config.AdaptiveIP.ExternalIfaceName)
	if err != nil {
		if strings.Contains(err.Error(), "RTNETLINK") &&
			strings.Contains(err.Error(), "exist") {
			return nil
		}
		return err
	}

	return nil
}
