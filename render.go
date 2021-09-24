package main

import (
        "fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type NameserversStruct struct {
    Addresses []string
    Search []string
}

type InterfaceStruct struct {
    Dhcp4 bool `yaml:"dhcp4"`
    Dhcp6 bool `yaml:"dhcp6"`
    Addresses []string
    Gateway4 string `yaml:"gateway4"`
    Nameservers NameserversStruct `yaml:"nameservers"`
}

type EthernetsStruct struct {
    Interface InterfaceStruct `yaml:"ens160"`
}

type NetworkStruct struct {
    Version int8 `yaml:"version"`
    Renderer string `yaml:"renderer"`
    Ethernets EthernetsStruct `yaml:"ethernets"`
}

type Netplan struct {
    Network NetworkStruct `yaml:"network"`
}


func Render(ip, prefix, gateway, dns1, dns2, suffix string) {

	full_ip := fmt.Sprintf("%s/%s", ip, prefix)

	netplan := Netplan{
          Network: NetworkStruct{
            Version: 2,
	    Renderer: "networkd",
	    Ethernets: EthernetsStruct{
	      Interface: InterfaceStruct{
		Dhcp4: false,
		Dhcp6: false,
		Addresses: []string{full_ip},
		Gateway4: gateway,
		Nameservers: NameserversStruct{
                  Addresses: []string{dns1, dns2, "8.8.8.8"},
		  Search: []string{suffix},
	        },
              },
            },
          },
	}

	yamlData, err := yaml.Marshal(&netplan)

        if err != nil {
          fmt.Printf("Error while Marshaling. %v", err)
        }

        fmt.Println("[INFO] Rendered Netplan configuration ...")
	fmt.Println(string(yamlData))

        fileName := "/etc/netplan/vcloud-inet.yaml"
        err = ioutil.WriteFile(fileName, yamlData, 0644)
        if err != nil {
            panic("Unable to write data into the file")
        }
}
