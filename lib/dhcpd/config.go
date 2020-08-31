package dhcpd

import (
	"errors"
	"github.com/apparentlymart/go-cidr/cidr"
	"hcc/harp/action/grpc/client"
	pb "hcc/harp/action/grpc/pb/rpcharp"
	"hcc/harp/dao"
	"hcc/harp/lib/config"
	"hcc/harp/lib/fileutil"
	"hcc/harp/lib/ifconfig"
	"hcc/harp/lib/iputil"
	"hcc/harp/lib/logger"
	"hcc/harp/lib/servicecontrol"
	"hcc/harp/model"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type nodeEntries struct {
	PXEMACAddress string
	IP            string
	NodeName      string
}

func getPXEFilename(serverUUID string) (string, error) {
	if len(serverUUID) == 0 {
		return "", errors.New("failed to get server UUID")
	}
	return model.DefaultPXEdir + "/" + serverUUID + "/pxelinux.0", nil
}

// NodeData : Data structure of list_node
type NodeData struct {
	Data struct {
		Node model.Node `json:"node"`
	} `json:"data"`
}

func getNodePXEMACAddress(nodeUUID string) (string, error) {
	node, err := client.RC.GetNode(nodeUUID)
	if err != nil {
		return "", err
	}

	return node.PXEMacAddr, nil
}

func writeConfigFile(input string, name string) error {
	err := fileutil.CreateDirIfNotExist(config.DHCPD.ConfigFileLocation)
	if err != nil {
		return err
	}

	dhcpdConfLocation := config.DHCPD.ConfigFileLocation + "/" + name + ".conf"
	err = fileutil.WriteFile(dhcpdConfLocation, input)
	if err != nil {
		return err
	}

	return nil
}

// CheckNodeUUIDs : Check UUIDs of nodes
func CheckNodeUUIDs(subnet net.IPNet, nodeUUIDs []string, leaderNodeUUID string) error {
	if len(nodeUUIDs) == 0 {
		return errors.New("provided nodeUUIDs[] is empty")
	}
	count := int(cidr.AddressCount(&subnet) - 2)
	if len(nodeUUIDs) > count {
		return errors.New("node count is bigger than available IP addresses count")
	}

	var leaderNodeUUIDfound = false
	for _, uuid := range nodeUUIDs {
		if uuid == leaderNodeUUID {
			leaderNodeUUIDfound = true
			break
		}
	}
	if leaderNodeUUIDfound == false {
		return errors.New("leaderNodeUUID is not found from provided nodeUUIDs[]")
	}

	return nil
}

func doWriteConfig(subnet *pb.Subnet, firstIP net.IP, lastIP net.IP, pxeFileName string, nodeUUIDs []string,
	useSamePXEFileForCompute bool) error {
	confContent := subnetConfBase
	confContent = strings.Replace(confContent, "HARP_DHCPD_SUBNET", subnet.NetworkIP, -1)
	confContent = strings.Replace(confContent, "HARP_DHCPD_NETMASK", subnet.Netmask, -1)
	confContent = strings.Replace(confContent, "HARP_DHCPD_START_IP", firstIP.String(), -1)
	confContent = strings.Replace(confContent, "HARP_DHCPD_LAST_IP", lastIP.String(), -1)

	confContent = strings.Replace(confContent, "HARP_DHCPD_NEXT_SERVER", subnet.NextServer, -1)
	if useSamePXEFileForCompute {
		confContent = strings.Replace(confContent, "HARP_DHCPD_PXE_FILENAME", pxeFileName, -1)
	} else {
		confContent = strings.Replace(confContent, "    filename \"HARP_DHCPD_PXE_FILENAME\";", "", -1)
	}

	confContent = strings.Replace(confContent, "HARP_DHCPD_DOMAIN_NAME_SERVER", subnet.NameServer, -1)
	confContent = strings.Replace(confContent, "HARP_DHCPD_DOMAIN_NAME", subnet.DomainName, -1)
	confContent = strings.Replace(confContent, "HARP_DHCPD_GATEWAY", subnet.Gateway, -1)

	confContent = strings.Replace(confContent, "HARP_DHCPD_MIN_LEASE_TIME", strconv.Itoa(int(config.DHCPD.MinLeaseTime)), -1)
	confContent = strings.Replace(confContent, "HARP_DHCPD_DEFAULT_LEASE_TIME", strconv.Itoa(int(config.DHCPD.DefaultLeaseTime)), -1)
	confContent = strings.Replace(confContent, "HARP_DHCPD_MAX_LEASE_TIME", strconv.Itoa(int(config.DHCPD.MaxLeaseTime)), -1)

	nextIP := cidr.Inc(firstIP)

	var nodeEntryConfPart = ""
	for i, uuid := range nodeUUIDs {
		pxeMacAddr, err := getNodePXEMACAddress(uuid)
		if err != nil {
			return err
		}
		pxeMacAddr = strings.Replace(pxeMacAddr, "-", ":", -1)

		var node = new(nodeEntries)
		node.NodeName = "node" + strconv.Itoa(i) + "." + subnet.SubnetName
		node.PXEMACAddress = pxeMacAddr
		if uuid == subnet.LeaderNodeUUID {
			node.IP = firstIP.String()
		} else {
			node.IP = nextIP.String()
			nextIP = cidr.Inc(nextIP)
		}

		var nodeConfPart = nodeEntry
		nodeConfPart = strings.Replace(nodeConfPart, "HARP_DHCPD_NODE_NAME", node.NodeName, -1)
		nodeConfPart = strings.Replace(nodeConfPart, "HARP_DHCPD_NODE_PXE_MAC", node.PXEMACAddress, -1)
		nodeConfPart = strings.Replace(nodeConfPart, "HARP_DHCPD_NODE_IP", node.IP, -1)
		if useSamePXEFileForCompute {
			nodeConfPart = strings.Replace(nodeConfPart, "        filename \"HARP_DHCPD_PXE_FILENAME\";", "", -1)
		} else {
			var otherPXEFileName string
			if uuid == subnet.LeaderNodeUUID {
				otherPXEFileName = strings.Replace(pxeFileName, subnet.ServerUUID, subnet.ServerUUID+"/Leader", -1)
			} else {
				otherPXEFileName = strings.Replace(pxeFileName, subnet.ServerUUID, subnet.ServerUUID+"/Compute", -1)
			}
			nodeConfPart = strings.Replace(nodeConfPart, "HARP_DHCPD_PXE_FILENAME", otherPXEFileName, -1)
		}

		nodeEntryConfPart += nodeConfPart
	}

	confContent = strings.Replace(confContent, "HARP_DHCPD_NODES_ENTRIES", nodeEntryConfPart, -1)

	err := writeConfigFile(confContent, subnet.ServerUUID)
	if err != nil {
		return err
	}

	return nil
}

// CreateConfig : Get needed parameters for make dhcpd config file then generate config file for each subnet
func CreateConfig(subnetUUID string, nodeUUIDs []string) error {
	if len(subnetUUID) == 0 {
		return errors.New("subnetUUID is needed for make dhcpd config file")
	}

	subnet, err := dao.ReadSubnet(subnetUUID)
	if err != nil {
		return err
	}

	if len(subnet.SubnetName) == 0 {
		return errors.New("name is needed for make dhcpd config file")
	}

	netIPnetworkIP := iputil.CheckValidIP(subnet.NetworkIP)
	if netIPnetworkIP == nil {
		return errors.New("wrong network IP address")
	}

	mask, err := iputil.CheckNetmask(subnet.Netmask)
	if err != nil {
		return err
	}

	ipNet := net.IPNet{
		IP:   netIPnetworkIP,
		Mask: mask,
	}

	err = iputil.CheckIPisInSubnet(ipNet, subnet.Gateway)
	if err != nil {
		return err
	}

	netIPnextServer := iputil.CheckValidIP(subnet.NextServer)
	if netIPnextServer == nil {
		return errors.New("wrong next server IP")
	}

	netIPnameServer := iputil.CheckValidIP(subnet.NameServer)
	if netIPnameServer == nil {
		return errors.New("wrong name server IP")
	}

	err = CheckNodeUUIDs(ipNet, nodeUUIDs, subnet.LeaderNodeUUID)
	if err != nil {
		return err
	}

	pxeFileName, err := getPXEFilename(subnet.ServerUUID)
	if err != nil {
		return err
	}

	firstIP, _ := cidr.AddressRange(&ipNet)
	firstIP = cidr.Inc(firstIP)
	lastIP := firstIP

	for i := 0; i < len(nodeUUIDs)-1; i++ {
		lastIP = cidr.Inc(lastIP)
	}

	err = doWriteConfig(subnet, firstIP, lastIP, pxeFileName, nodeUUIDs, false)
	if err != nil {
		return err
	}

	// Allocate gateway IP address to internal interface
	err = ifconfig.CreateAndLoadIfconfigScriptInternal(config.AdaptiveIP.InternalIfaceName, subnet.Gateway, subnet.Netmask)
	if err != nil {
		return err
	}

	return nil
}

// CheckLocalDHCPDConfig : Check if harp dhcpd config file is included in local dhcpd server config file
func CheckLocalDHCPDConfig() error {
	include := includeStr
	include = strings.Replace(include, "HARP_DHCPD_CONF_LOCATION",
		config.DHCPD.ConfigFileLocation+"/harp_dhcpd.conf", -1)

	confData, err := ioutil.ReadFile(config.DHCPD.LocalConfigFileLocation)
	if err != nil {
		return errors.New("failed to reading data from local dhcpd config file location")
	}

	isHarpDHCPDIncluded := strings.Contains(string(confData), include)
	if !isHarpDHCPDIncluded {
		logger.Logger.Println("Please add this line to dhcpd config file!\n" + include)
		return errors.New("cannot find harp dhcp config include line from local dhcpd config file")
	}

	return nil
}

func getSubnetConfFiles() ([]string, error) {
	var files []string

	folder := config.DHCPD.ConfigFileLocation
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}

// UpdateHarpDHCPDConfig : Update harp dhcpd main config file. Write subnet config files include lines to 'harp_dhcpd.conf'
func UpdateHarpDHCPDConfig() (int, error) {
	configFiles, err := getSubnetConfFiles()
	if err != nil {
		return 0, err
	}

	harpDHCPDConf := harpDHCPDConfBase
	var allIncludeLines = ""
	var files = 0

	for _, filename := range configFiles {
		if strings.Contains(filename, "harp_dhcpd.conf") ||
			filename == config.DHCPD.ConfigFileLocation {
			continue
		}

		include := "    " + includeStr
		include = strings.Replace(include, "HARP_DHCPD_CONF_LOCATION", filename, -1)
		allIncludeLines += include + "\n"

		files++
	}

	if files == 0 {
		harpDHCPDConf = ""
	} else {
		harpDHCPDConf = strings.Replace(harpDHCPDConf, "HARP_DHCPD_INCLUDE_STRINGS", allIncludeLines, -1)
	}

	err = fileutil.CreateDirIfNotExist(config.DHCPD.ConfigFileLocation)
	if err != nil {
		return 0, err
	}

	err = fileutil.WriteFile(config.DHCPD.ConfigFileLocation+"/harp_dhcpd.conf", harpDHCPDConf)
	if err != nil {
		return 0, err
	}

	return files, nil
}

// CreateDHCPDConfig : Do dhcpd config file creation works
func CreateDHCPDConfig(in *pb.ReqCreateDHPCDConf) (string, error) {
	subnetUUID := in.GetSubnetUUID()
	subnetUUIDOk := len(subnetUUID) != 0

	if !subnetUUIDOk {
		return "", errors.New("need a SubnetUUID argument")
	}

	subnet, err := dao.ReadSubnet(subnetUUID)
	if err != nil {
		return "", err
	}

	nodes, err := client.RC.GetNodeList(subnet.ServerUUID)
	if err != nil {
		logger.Logger.Println(errors.New("Failed to get nodes by server UUID: " +
			subnet.ServerUUID + " (" + err.Error() + ")"))
		return "", err
	}

	var nodeUUIDs []string

	for i := range nodes {
		nodeUUIDs = append(nodeUUIDs, nodes[i].UUID)
	}

	err = CreateConfig(subnetUUID, nodeUUIDs)
	if err != nil {
		return "", err
	}

	_, err = UpdateHarpDHCPDConfig()
	if err != nil {
		return "", err
	}

	err = servicecontrol.RestartDHCPDServer()
	if err != nil {
		logger.Logger.Println("Failed to restart dhcpd server (" + config.DHCPD.LocalDHCPDServiceName + ")")
		return "", err
	}

	return "CreateDHCPDConfig: succeed", nil
}

// DeleteDHCPDConfigFile : Do dhcpd config file deletion works
func DeleteDHCPDConfigFile(in *pb.ReqDeleteDHPCDConf) (string, error) {
	subnetUUID := in.GetSubnetUUID()
	subnetUUIDOk := len(subnetUUID) != 0

	if !subnetUUIDOk {
		return "", errors.New("need a SubnetUUID argument")
	}

	subnet, err := dao.ReadSubnet(subnetUUID)
	if err != nil {
		return "", err
	}

	dhcpdConfLocation := config.DHCPD.ConfigFileLocation + "/" + subnet.ServerUUID + ".conf"
	err = fileutil.DeleteFile(dhcpdConfLocation)
	if err != nil {
		return "", err
	}

	err = servicecontrol.RestartDHCPDServer()
	if err != nil {
		logger.Logger.Println("Failed to restart dhcpd server (" + config.DHCPD.LocalDHCPDServiceName + ")")
		return "", err
	}

	return "DeleteDHCPDConfigFile: succeed", nil
}

// CheckDatabaseAndGenerateDHCPDConfigs : Check database and generate dhcpd configs
func CheckDatabaseAndGenerateDHCPDConfigs() error {
	serverUUIDs, err := client.RC.AllServerUUID()
	if err != nil {
		return err
	}

	for i := range serverUUIDs {
		nodes, err := client.RC.GetNodeList(serverUUIDs[i])
		if err != nil {
			logger.Logger.Println(errors.New("Failed to get nodes by server UUID: " +
				serverUUIDs[i] + " (" + err.Error() + ")"))
			continue
		}

		var nodeUUIDs []string

		for i := range nodes {
			nodeUUIDs = append(nodeUUIDs, nodes[i].UUID)
		}

		subnet, err := dao.ReadSubnetByServer(serverUUIDs[i])
		if err != nil {
			logger.Logger.Println("Failed to get subnet by server UUID: " +
				serverUUIDs[i] + " (" + err.Error() + ")")
			continue
		}

		subnetUUID := subnet.UUID
		err = CreateConfig(subnetUUID, nodeUUIDs)
		if err != nil {
			logger.Logger.Println("Failed to create dhcpd config of subnetUUID=" +
				subnetUUID + " (" + err.Error() + ")")
			continue
		}

		logger.Logger.Println("Created dhcpd config of subnetUUID=" + subnetUUID)
	}

	files, err := UpdateHarpDHCPDConfig()
	if err != nil {
		return err
	}

	if files != 0 {
		err = servicecontrol.RestartDHCPDServer()
		if err != nil {
			logger.Logger.Println("Failed to restart dhcpd server (" + config.DHCPD.LocalDHCPDServiceName + ")")
			return err
		}
	}

	return nil
}
