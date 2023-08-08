#!/bin/bash

# Output the data in JSON format
echo "{"
echo "\"openwrt_device\": {"

# Get the MAC address of the OpenWrt device
mac_address=$(ifconfig eth0 | awk '/ether/ {print $2}')

# Get the hostname of the OpenWrt device
hostname=$(hostname)

# Output the OpenWrt device details in JSON format
echo "\"mac_address\": \"$mac_address\","
echo "\"hostname\": \"$hostname\""

echo "}"
echo "}"