package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shotgun-20/ovs-exporter/ovs"
)

type OvsPromCollector struct {
	ip        string
	port      int
	ovsReader ovs.OvsStatReader
}

var (
	TunFlowPacketsDesc = prometheus.NewDesc(
		"TunflowPackets",
		"The number of packets matched for the given OpenFlow entry.",
		[]string{"match", "action", "table", "priority"},
		nil)

	ExflowPacketsDesc = prometheus.NewDesc(
		"ExflowPackets",
		"The number of packets matched for the given OpenFlow entry.",
		[]string{"match", "action", "table", "priority"},
		nil)

	IntflowPacketsDesc = prometheus.NewDesc(
		"IntflowPackets",
		"The number of packets matched for the given OpenFlow entry.",
		[]string{"match", "action", "table", "priority"},
		nil)

	TunFlowBytesDesc = prometheus.NewDesc(
		"TunflowBytes",
		"The number of bytes matched for the given OpenFlow entry",
		[]string{"match", "action", "table", "priority"},
		nil)

	ExflowBytesDesc = prometheus.NewDesc(
		"ExflowBytes",
		"The number of bytes matched for the given OpenFlow entry",
		[]string{"match", "action", "table", "priority"},
		nil)

	IntflowBytesDesc = prometheus.NewDesc(
		"IntflowBytes",
		"The number of bytes matched for the given OpenFlow entry",
		[]string{"match", "action", "table", "priority"},
		nil)

	TunFlowAgeDesc = prometheus.NewDesc(
		"TunflowAge",
		"The number of seconds have passed since the given OpenFlow entry was created",
		[]string{"match", "action", "table", "priority"},
		nil)

	ExflowAgeDesc = prometheus.NewDesc(
		"ExflowAge",
		"The number of seconds have passed since the given OpenFlow entry was created",
		[]string{"match", "action", "table", "priority"},
		nil)

	IntflowAgeDesc = prometheus.NewDesc(
		"IntflowAge",
		"The number of seconds have passed since the given OpenFlow entry was created",
		[]string{"match", "action", "table", "priority"},
		nil)

	TunFlowIdleTimeDesc = prometheus.NewDesc(
		"TunflowIdleTime",
		"The number of seconds have passed since the last packet has seen for the given OpenFlow entry",
		[]string{"match", "action", "table", "priority"},
		nil)

	ExflowIdleTimeDesc = prometheus.NewDesc(
		"ExflowIdleTime",
		"The number of seconds have passed since the last packet has seen for the given OpenFlow entry",
		[]string{"match", "action", "table", "priority"},
		nil)

	IntflowIdleTimeDesc = prometheus.NewDesc(
		"IntflowIdleTime",
		"The number of seconds have passed since the last packet has seen for the given OpenFlow entry",
		[]string{"match", "action", "table", "priority"},
		nil)

	TunportRxPacketsDesc = prometheus.NewDesc(
		"TunportRxPackets",
		"The number of packet that was recieved by a given port",
		[]string{"port"},
		nil)

	ExportRxPacketsDesc = prometheus.NewDesc(
		"ExportRxPackets",
		"The number of packet that was recieved by a given port",
		[]string{"port"},
		nil)

	IntportRxPacketsDesc = prometheus.NewDesc(
		"IntportRxPackets",
		"The number of packet that was recieved by a given port",
		[]string{"port"},
		nil)

	TunportTxPackets = prometheus.NewDesc(
		"TunportTxPackets",
		"The number of packet that was sent by a given port",
		[]string{"port"},
		nil)

	ExportTxPackets = prometheus.NewDesc(
		"ExportTxPackets",
		"The number of packet that was sent by a given port",
		[]string{"port"},
		nil)

	IntportTxPackets = prometheus.NewDesc(
		"IntportTxPackets",
		"The number of packet that was sent by a given port",
		[]string{"port"},
		nil)

	TunportRxBytesDesc = prometheus.NewDesc(
		"TunportRxBytes",
		"The number of bytes that was recieved by a given port",
		[]string{"port"},
		nil)

	ExportRxBytesDesc = prometheus.NewDesc(
		"ExportRxBytes",
		"The number of bytes that was recieved by a given port",
		[]string{"port"},
		nil)

	IntportRxBytesDesc = prometheus.NewDesc(
		"IntportRxBytes",
		"The number of bytes that was recieved by a given port",
		[]string{"port"},
		nil)

	TunportTxBytes = prometheus.NewDesc(
		"TunportTxBytes",
		"The number of bytes that was sent by a given port",
		[]string{"port"},
		nil)

	ExportTxBytes = prometheus.NewDesc(
		"ExportTxBytes",
		"The number of bytes that was sent by a given port",
		[]string{"port"},
		nil)

	IntportTxBytes = prometheus.NewDesc(
		"IntportTxBytes",
		"The number of bytes that was sent by a given port",
		[]string{"port"},
		nil)

	TunportRxDropsDesc = prometheus.NewDesc(
		"TunportRxDrops",
		"The number of packets that was dropped on receive side by a given port",
		[]string{"port"},
		nil)

	ExportRxDropsDesc = prometheus.NewDesc(
		"ExportRxDrops",
		"The number of packets that was dropped on receive side by a given port",
		[]string{"port"},
		nil)

	IntportRxDropsDesc = prometheus.NewDesc(
		"IntportRxDrops",
		"The number of packets that was dropped on receive side by a given port",
		[]string{"port"},
		nil)

	TunportTxDropsDesc = prometheus.NewDesc(
		"TunportTxDrops",
		"The number of packets that was dropped on sending side by a given port",
		[]string{"port"},
		nil)

	ExportTxDropsDesc = prometheus.NewDesc(
		"ExportTxDrops",
		"The number of packets that was dropped on sending side by a given port",
		[]string{"port"},
		nil)

	IntportTxDropsDesc = prometheus.NewDesc(
		"IntportTxDrops",
		"The number of packets that was dropped on sending side by a given port",
		[]string{"port"},
		nil)

	//	groupPacketsDesc = prometheus.NewDesc(
	//		"groupPackets",
	//		"The number of packet that was sent by a given group",
	//		[]string{"groupId", "groupType"},
	//		nil)
	//
	//	groupBytesDesc = prometheus.NewDesc(
	//		"groupBytes",
	//		"The number of bytes that was sent by a given group",
	//		[]string{"groupId", "groupType"},
	//		nil)
	//
	//	groupDurationDesc = prometheus.NewDesc(
	//		"groupDuration",
	//		"The number of seconds passed since the group entry was added",
	//		[]string{"groupId", "groupType"},
	//		nil)
	//
	//	groupBucketPacketsDesc = prometheus.NewDesc(
	//		"groupBucketPackets",
	//		"The number of packet that was sent by a given group bucket",
	//		[]string{"groupId", "groupType", "bucketActions"},
	//		nil)
	//
	//	groupBucketBytesDesc = prometheus.NewDesc(
	//		"groupBucketBytes",
	//		"The number of bytes that was sent by a given group bucket",
	//		[]string{"groupId", "groupType", "bucketActions"},
	//		nil)
	//
	urlParsingErrorDesc = prometheus.NewDesc(
		"ovs_error",
		"Error scraping target. Correct format is: http://<IP>:<Port>/metrics?target=<targetIP>",
		nil, nil)
)

// Describe implements Prometheus.Collector.
func (c OvsPromCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- prometheus.NewDesc("dummy", "dummy", nil, nil)
}

// Collect implements Prometheus.Collector.
func (c OvsPromCollector) Collect(ch chan<- prometheus.Metric) {
	if c.ip == "" {
		ch <- prometheus.NewInvalidMetric(urlParsingErrorDesc, nil)
		return
	}

	//Creating Prometheus compatible output for:
	//	- number of packets as "flowPackets" type Counter
	//	- number of bytes as "flowBytes" type Counter
	//	- age of the flow as "flowAge" type Gauge
	//	- idle time as "flowIdleTime" type Gauge

	TunFlowEntries, err := c.ovsReader.TunFlows(c.ip, c.port)

	if err != nil {
		ch <- prometheus.NewInvalidMetric(prometheus.NewDesc("ovs_error", "Error parsing flow dump", nil, nil), err)
		return
	}

	for _, entry := range TunFlowEntries {

		ch <- prometheus.MustNewConstMetric(
			TunFlowPacketsDesc,
			prometheus.CounterValue,
			float64(entry.Packets),
			entry.Match,
			entry.Action,
			entry.Table,
			entry.Priority)

		ch <- prometheus.MustNewConstMetric(
			TunFlowBytesDesc,
			prometheus.CounterValue,
			float64(entry.Bytes),
			entry.Match,
			entry.Action,
			entry.Table,
			entry.Priority)

		ch <- prometheus.MustNewConstMetric(
			TunFlowAgeDesc,
			prometheus.GaugeValue,
			float64(entry.Duration),
			entry.Match,
			entry.Action,
			entry.Table,
			entry.Priority)

		ch <- prometheus.MustNewConstMetric(
			TunFlowIdleTimeDesc,
			prometheus.GaugeValue,
			float64(entry.IdleAge),
			entry.Match,
			entry.Action,
			entry.Table,
			entry.Priority)
	}

	ExflowEntries, err := c.ovsReader.ExFlows(c.ip, c.port)

	if err != nil {
		ch <- prometheus.NewInvalidMetric(prometheus.NewDesc("ovs_error", "Error parsing flow dump", nil, nil), err)
		return
	}

	for _, entry := range ExflowEntries {

		ch <- prometheus.MustNewConstMetric(
			ExflowPacketsDesc,
			prometheus.CounterValue,
			float64(entry.Packets),
			entry.Match,
			entry.Action,
			entry.Table,
			entry.Priority)

		ch <- prometheus.MustNewConstMetric(
			ExflowBytesDesc,
			prometheus.CounterValue,
			float64(entry.Bytes),
			entry.Match,
			entry.Action,
			entry.Table,
			entry.Priority)

		ch <- prometheus.MustNewConstMetric(
			ExflowAgeDesc,
			prometheus.GaugeValue,
			float64(entry.Duration),
			entry.Match,
			entry.Action,
			entry.Table,
			entry.Priority)

		ch <- prometheus.MustNewConstMetric(
			ExflowIdleTimeDesc,
			prometheus.GaugeValue,
			float64(entry.IdleAge),
			entry.Match,
			entry.Action,
			entry.Table,
			entry.Priority)
	}

	IntflowEntries, err := c.ovsReader.IntFlows(c.ip, c.port)

	if err != nil {
		ch <- prometheus.NewInvalidMetric(prometheus.NewDesc("ovs_error", "Error parsing flow dump", nil, nil), err)
		return
	}

	for _, entry := range IntflowEntries {

		ch <- prometheus.MustNewConstMetric(
			IntflowPacketsDesc,
			prometheus.CounterValue,
			float64(entry.Packets),
			entry.Match,
			entry.Action,
			entry.Table,
			entry.Priority)

		ch <- prometheus.MustNewConstMetric(
			IntflowBytesDesc,
			prometheus.CounterValue,
			float64(entry.Bytes),
			entry.Match,
			entry.Action,
			entry.Table,
			entry.Priority)

		ch <- prometheus.MustNewConstMetric(
			IntflowAgeDesc,
			prometheus.GaugeValue,
			float64(entry.Duration),
			entry.Match,
			entry.Action,
			entry.Table,
			entry.Priority)

		ch <- prometheus.MustNewConstMetric(
			IntflowIdleTimeDesc,
			prometheus.GaugeValue,
			float64(entry.IdleAge),
			entry.Match,
			entry.Action,
			entry.Table,
			entry.Priority)
	}

	TunPortEntries, err := c.ovsReader.TunPorts(c.ip, c.port)
	//if error was occured we return
	if err != nil {
		ch <- prometheus.NewInvalidMetric(prometheus.NewDesc("ovs_error", "Error parsing port dump", nil, nil), err)
		return
	}

	for _, entry := range TunPortEntries {

		ch <- prometheus.MustNewConstMetric(
			TunportRxPacketsDesc,
			prometheus.CounterValue,
			float64(entry.RxPackets),
			entry.PortNumber)

		ch <- prometheus.MustNewConstMetric(
			TunportTxPackets,
			prometheus.CounterValue,
			float64(entry.TxPackets),
			entry.PortNumber)

		ch <- prometheus.MustNewConstMetric(
			TunportRxBytesDesc,
			prometheus.CounterValue,
			float64(entry.RxBytes),
			entry.PortNumber)

		ch <- prometheus.MustNewConstMetric(
			TunportTxBytes,
			prometheus.CounterValue,
			float64(entry.TxBytes),
			entry.PortNumber)

		ch <- prometheus.MustNewConstMetric(
			TunportRxDropsDesc,
			prometheus.CounterValue,
			float64(entry.RxDrops),
			entry.PortNumber)

		ch <- prometheus.MustNewConstMetric(
			TunportTxDropsDesc,
			prometheus.CounterValue,
			float64(entry.TxDrops),
			entry.PortNumber)
	}

	ExportEntries, err := c.ovsReader.ExPorts(c.ip, c.port)
	//if error was occured we return
	if err != nil {
		ch <- prometheus.NewInvalidMetric(prometheus.NewDesc("ovs_error", "Error parsing port dump", nil, nil), err)
		return
	}

	for _, entry := range ExportEntries {

		ch <- prometheus.MustNewConstMetric(
			ExportRxPacketsDesc,
			prometheus.CounterValue,
			float64(entry.RxPackets),
			entry.PortNumber)

		ch <- prometheus.MustNewConstMetric(
			ExportTxPackets,
			prometheus.CounterValue,
			float64(entry.TxPackets),
			entry.PortNumber)

		ch <- prometheus.MustNewConstMetric(
			ExportRxBytesDesc,
			prometheus.CounterValue,
			float64(entry.RxBytes),
			entry.PortNumber)

		ch <- prometheus.MustNewConstMetric(
			ExportTxBytes,
			prometheus.CounterValue,
			float64(entry.TxBytes),
			entry.PortNumber)

		ch <- prometheus.MustNewConstMetric(
			ExportRxDropsDesc,
			prometheus.CounterValue,
			float64(entry.RxDrops),
			entry.PortNumber)

		ch <- prometheus.MustNewConstMetric(
			ExportTxDropsDesc,
			prometheus.CounterValue,
			float64(entry.TxDrops),
			entry.PortNumber)
	}

	IntportEntries, err := c.ovsReader.IntPorts(c.ip, c.port)
	//if error was occured we return
	if err != nil {
		ch <- prometheus.NewInvalidMetric(prometheus.NewDesc("ovs_error", "Error parsing port dump", nil, nil), err)
		return
	}

	for _, entry := range IntportEntries {

		ch <- prometheus.MustNewConstMetric(
			IntportRxPacketsDesc,
			prometheus.CounterValue,
			float64(entry.RxPackets),
			entry.PortNumber)

		ch <- prometheus.MustNewConstMetric(
			IntportTxPackets,
			prometheus.CounterValue,
			float64(entry.TxPackets),
			entry.PortNumber)

		ch <- prometheus.MustNewConstMetric(
			IntportRxBytesDesc,
			prometheus.CounterValue,
			float64(entry.RxBytes),
			entry.PortNumber)

		ch <- prometheus.MustNewConstMetric(
			IntportTxBytes,
			prometheus.CounterValue,
			float64(entry.TxBytes),
			entry.PortNumber)

		ch <- prometheus.MustNewConstMetric(
			IntportRxDropsDesc,
			prometheus.CounterValue,
			float64(entry.RxDrops),
			entry.PortNumber)

		ch <- prometheus.MustNewConstMetric(
			IntportTxDropsDesc,
			prometheus.CounterValue,
			float64(entry.TxDrops),
			entry.PortNumber)
	}
	//	groupEntries, err := c.ovsReader.Groups(c.ip, c.port)
	//	//if error was occured we return
	//	if err != nil {
	//		ch <- prometheus.NewInvalidMetric(prometheus.NewDesc("ovs_error", "Error parsing group dump", nil, nil), err)
	//		return
	//	}
	//
	//	//Creating Prometheus compatible output for every group stat with groupId label:
	//	//	- number of packets that was forwarded by a group rule as "groupPackets" type Counter
	//	//	- number of bytes that was forwarded by a group rule as "groupBytes" type Counter
	//	//	- number of second that passed since a group rule was added as "groupPackets" type Gauge
	//	//	- number of packets that was forwarded by a bucket in a group rule as "groupBucketPackets" type Counter
	//	//	- number of bytes that was forwarded by a bucket in a group rule as "groupBucketBytes" type Counter
	//
	//	for _, entry := range groupEntries {
	//
	//		ch <- prometheus.MustNewConstMetric(
	//			groupPacketsDesc,
	//			prometheus.CounterValue,
	//			float64(entry.Packets),
	//			entry.GroupId,
	//			entry.GroupType)
	//
	//		ch <- prometheus.MustNewConstMetric(
	//			groupBytesDesc,
	//			prometheus.CounterValue,
	//			float64(entry.Bytes),
	//			entry.GroupId,
	//			entry.GroupType)
	//
	//		ch <- prometheus.MustNewConstMetric(
	//			groupDurationDesc,
	//			prometheus.CounterValue,
	//			float64(entry.Duration),
	//			entry.GroupId,
	//			entry.GroupType)
	//
	//		for _, bucket := range entry.Buckets {
	//			ch <- prometheus.MustNewConstMetric(
	//				groupBucketPacketsDesc,
	//				prometheus.CounterValue,
	//				float64(bucket.Packets),
	//				entry.GroupId,
	//				entry.GroupType,
	//				bucket.Actions)
	//
	//			ch <- prometheus.MustNewConstMetric(
	//				groupBucketBytesDesc,
	//				prometheus.CounterValue,
	//				float64(bucket.Bytes),
	//				entry.GroupId,
	//				entry.GroupType,
	//				bucket.Actions)
	//		}
	//	}
}
