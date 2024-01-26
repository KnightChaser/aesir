package datastructure

import (
	"fmt"
	"log"

	"github.com/tidwall/gjson"
)

func EventDataStructureJsonify(id int64, stat string) map[string]interface{} {

	eventDataStructure := make(map[string]interface{})

	switch id {
	case 1:
		eventDataStructure["CommandLine"] = gjson.Get(stat, "Event.EventData.CommandLine").String()
		eventDataStructure["Company"] = gjson.Get(stat, "Event.EventData.Company").String()
		eventDataStructure["CurrentDirectory"] = gjson.Get(stat, "Event.EventData.CurrentDirectory").String()
		eventDataStructure["Description"] = gjson.Get(stat, "Event.EventData.Description").String()
		eventDataStructure["FileVersion"] = gjson.Get(stat, "Event.EventData.FileVersion").String()
		eventDataStructure["Hashes"] = gjson.Get(stat, "Event.EventData.Hashes").String()
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["IntegrityLevel"] = gjson.Get(stat, "Event.EventData.IntegrityLevel").Int()
		eventDataStructure["LogonGuid"] = gjson.Get(stat, "Event.EventData.LogonGuid").String()
		eventDataStructure["LogonId"] = gjson.Get(stat, "Event.EventData.LogonId").String()
		eventDataStructure["OriginalFileName"] = gjson.Get(stat, "Event.EventData.OriginalFileName").String()
		eventDataStructure["ParentCommandLine"] = gjson.Get(stat, "Event.EventData.ParentCommandLine").String()
		eventDataStructure["ParentImage"] = gjson.Get(stat, "Event.EventData.ParentImage").String()
		eventDataStructure["ParentProcessGuid"] = gjson.Get(stat, "Event.EventData.ParentProcessGuid").String()
		eventDataStructure["ParentProcessId"] = gjson.Get(stat, "Event.EventData.ParentProcessId").Int()
		eventDataStructure["ParentUser"] = gjson.Get(stat, "Event.EventData.ParentUser").String()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").String()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["Product"] = gjson.Get(stat, "Event.EventData.Product").String()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["TerminalSessionId"] = gjson.Get(stat, "Event.EventData.TerminalSessionId").String()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 2:
		eventDataStructure["CreationUtcTime"] = gjson.Get(stat, "Event.EventData.CreationUtcTime").Time()
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["PreviousCreationUtcTime"] = gjson.Get(stat, "Event.EventData.PreviousCreationUtcTime").Time()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").String()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["TargetFilename"] = gjson.Get(stat, "Event.EventData.TargetFilename").String()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 3:
		eventDataStructure["DestinationHostname"] = gjson.Get(stat, "Event.EventData.DestinationHostname").String()
		eventDataStructure["DestinationIp"] = gjson.Get(stat, "Event.EventData.DestinationIp").String()
		eventDataStructure["DestinationIsIpv6"] = gjson.Get(stat, "Event.EventData.DestinationIsIpv6").Bool()
		eventDataStructure["DestinationPort"] = gjson.Get(stat, "Event.EventData.DestinationPort").Int()
		eventDataStructure["DestinationPortName"] = gjson.Get(stat, "Event.EventData.DestinationPortName").String()
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["Initiated"] = gjson.Get(stat, "Event.EventData.Initiated").Bool()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").String()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["Protocol"] = gjson.Get(stat, "Event.EventData.Protocol").String()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["SourceHostname"] = gjson.Get(stat, "Event.EventData.SourceHostname").String()
		eventDataStructure["SourceIp"] = gjson.Get(stat, "Event.EventData.SourceIp").String()
		eventDataStructure["SourceIsIpv6"] = gjson.Get(stat, "Event.EventData.SourceIsIpv6").Bool()
		eventDataStructure["SourcePort"] = gjson.Get(stat, "Event.EventData.SourcePort").Int()
		eventDataStructure["SourcePortName"] = gjson.Get(stat, "Event.EventData.SourcePortName").String()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 4:
		eventDataStructure["SchemaVersion"] = gjson.Get(stat, "Event.EventData.SchemaVersion").String()
		eventDataStructure["State"] = gjson.Get(stat, "Event.EventData.State").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		eventDataStructure["Version"] = gjson.Get(stat, "Event.EventData.Version").String()
		break

	case 5:
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").String()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 6:
		eventDataStructure["Hashes"] = gjson.Get(stat, "Event.EventData.Hashes").String()
		eventDataStructure["ImageLoaded"] = gjson.Get(stat, "Event.EventData.ImageLoaded").String()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["Signature"] = gjson.Get(stat, "Event.EventData.Signature").String()
		eventDataStructure["SignatureStatus"] = gjson.Get(stat, "Event.EventData.SignatureStatus").String()
		eventDataStructure["Signed"] = gjson.Get(stat, "Event.EventData.Signed").Bool()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 7:
		eventDataStructure["Company"] = gjson.Get(stat, "Event.EventData.Company").String()
		eventDataStructure["Description"] = gjson.Get(stat, "Event.EventData.Description").String()
		eventDataStructure["FileVersion"] = gjson.Get(stat, "Event.EventData.FileVersion").String()
		eventDataStructure["Hashes"] = gjson.Get(stat, "Event.EventData.Hashes").String()
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["ImageLoaded"] = gjson.Get(stat, "Event.EventData.ImageLoaded").String()
		eventDataStructure["OriginalFileName"] = gjson.Get(stat, "Event.EventData.OriginalFileName").String()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").String()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["Product"] = gjson.Get(stat, "Event.EventData.Product").String()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["Signature"] = gjson.Get(stat, "Event.EventData.Signature").String()
		eventDataStructure["SignatureStatus"] = gjson.Get(stat, "Event.EventData.SignatureStatus").String()
		eventDataStructure["Signed"] = gjson.Get(stat, "Event.EventData.Signed").Bool()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 8:
		eventDataStructure["NewThreadId"] = gjson.Get(stat, "Event.EventData.NewThreadId").Int()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["SourceImage"] = gjson.Get(stat, "Event.EventData.SourceImage").String()
		eventDataStructure["SourceProcessGuid"] = gjson.Get(stat, "Event.EventData.SourceProcessGuid").String()
		eventDataStructure["SourceProcessId"] = gjson.Get(stat, "Event.EventData.SourceProcessId").Int()
		eventDataStructure["SourceUser"] = gjson.Get(stat, "Event.EventData.SourceUser").String()
		eventDataStructure["StartAddress"] = gjson.Get(stat, "Event.EventData.StartAddress").String()
		eventDataStructure["StartFunction"] = gjson.Get(stat, "Event.EventData.StartFunction").String()
		eventDataStructure["StartModule"] = gjson.Get(stat, "Event.EventData.StartModule").String()
		eventDataStructure["TargetImage"] = gjson.Get(stat, "Event.EventData.TargetImage").String()
		eventDataStructure["TargetProcessGuid"] = gjson.Get(stat, "Event.EventData.TargetProcessGuid").String()
		eventDataStructure["TargetProcessId"] = gjson.Get(stat, "Event.EventData.TargetProcessId").Int()
		eventDataStructure["TargetUser"] = gjson.Get(stat, "Event.EventData.TargetUser").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 9:
		eventDataStructure["Device"] = gjson.Get(stat, "Event.EventData.Device").String()
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").String()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 10:
		eventDataStructure["CallTrace"] = gjson.Get(stat, "Event.EventData.CallTrace").String()
		eventDataStructure["GrantedAccess"] = gjson.Get(stat, "Event.EventData.GrantedAccess").String()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["SourceImage"] = gjson.Get(stat, "Event.EventData.SourceImage").String()
		eventDataStructure["SourceProcessGUID"] = gjson.Get(stat, "Event.EventData.SourceProcessGUID").String()
		eventDataStructure["SourceProcessId"] = gjson.Get(stat, "Event.EventData.SourceProcessId").Int()
		eventDataStructure["SourceThreadId"] = gjson.Get(stat, "Event.EventData.SourceThreadId").Int()
		eventDataStructure["SourceUser"] = gjson.Get(stat, "Event.EventData.SourceUser").String()
		eventDataStructure["TargetImage"] = gjson.Get(stat, "Event.EventData.TargetImage").String()
		eventDataStructure["TargetProcessGUID"] = gjson.Get(stat, "Event.EventData.TargetProcessGUID").String()
		eventDataStructure["TargetProcessId"] = gjson.Get(stat, "Event.EventData.TargetProcessId").Int()
		eventDataStructure["TargetUser"] = gjson.Get(stat, "Event.EventData.TargetUser").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 11:
		eventDataStructure["CreationUtcTime "] = gjson.Get(stat, "Event.EventData.CreationUtcTime").Time()
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").Int()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["TargetFilename"] = gjson.Get(stat, "Event.EventData.TargetFilename").String()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 12:
		eventDataStructure["EventType"] = gjson.Get(stat, "Event.EventData.EventType").String()
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").String()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["TargetObject"] = gjson.Get(stat, "Event.EventData.TargetObject").String()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 13:
		eventDataStructure["Details"] = gjson.Get(stat, "Event.EventData.Details").String()
		eventDataStructure["EventType"] = gjson.Get(stat, "Event.EventData.EventType").String()
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").String()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["TargetObject"] = gjson.Get(stat, "Event.EventData.TargetObject").String()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 14:
		eventDataStructure["EventType"] = gjson.Get(stat, "Event.EventData.EventType").String()
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["NewName"] = gjson.Get(stat, "Event.EventData.NewName").String()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").String()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["TargetObject"] = gjson.Get(stat, "Event.EventData.TargetObject").String()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 15:
		eventDataStructure["Contents"] = gjson.Get(stat, "Event.EventData.Contents").String()
		eventDataStructure["CreationUtcTime"] = gjson.Get(stat, "Event.EventData.CreationUtcTime").Time()
		eventDataStructure["Hash"] = gjson.Get(stat, "Event.EventData.Hash").String()
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").String()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["TargetFilename"] = gjson.Get(stat, "Event.EventData.TargetFilename").String()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 16:
		eventDataStructure["Configuration"] = gjson.Get(stat, "Event.EventData.Configuration").String()
		eventDataStructure["ConfigurationFileHash"] = gjson.Get(stat, "Event.EventData.ConfigurationFileHash").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 17:
		eventDataStructure["EventType"] = gjson.Get(stat, "Event.EventData.EventType").String()
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["PipeName"] = gjson.Get(stat, "Event.EventData.PipeName").String()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").String()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 18:
		eventDataStructure["EventType"] = gjson.Get(stat, "Event.EventData.EventType").String()
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["PipeName"] = gjson.Get(stat, "Event.EventData.PipeName").String()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").String()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 19:
		eventDataStructure["EventNamespace"] = gjson.Get(stat, "Event.EventData.EventNamespace").String()
		eventDataStructure["EventType"] = gjson.Get(stat, "Event.EventData.EventType").String()
		eventDataStructure["Name"] = gjson.Get(stat, "Event.EventData.Name").String()
		eventDataStructure["Operation"] = gjson.Get(stat, "Event.EventData.Operation").String()
		eventDataStructure["Query"] = gjson.Get(stat, "Event.EventData.Query").String()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 20:
		eventDataStructure["Destination"] = gjson.Get(stat, "Event.EventData.Destination").String()
		eventDataStructure["EventType"] = gjson.Get(stat, "Event.EventData.EventType").String()
		eventDataStructure["Name"] = gjson.Get(stat, "Event.EventData.Name").String()
		eventDataStructure["Operation"] = gjson.Get(stat, "Event.EventData.Operation").String()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["Type"] = gjson.Get(stat, "Event.EventData.Type").String()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 21:
		eventDataStructure["Consumer"] = gjson.Get(stat, "Event.EventData.Consumer").String()
		eventDataStructure["EventType"] = gjson.Get(stat, "Event.EventData.EventType").String()
		eventDataStructure["Filter"] = gjson.Get(stat, "Event.EventData.Filter").String()
		eventDataStructure["Operation"] = gjson.Get(stat, "Event.EventData.Operation").String()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 22:
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").String()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["QueryName"] = gjson.Get(stat, "Event.EventData.QueryName").String()
		eventDataStructure["QueryResults"] = gjson.Get(stat, "Event.EventData.QueryResults").String()
		eventDataStructure["QueryStatus"] = gjson.Get(stat, "Event.EventData.QueryStatus").String()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 23:
		eventDataStructure["Archived"] = gjson.Get(stat, "Event.EventData.Archived").String()
		eventDataStructure["Hashes"] = gjson.Get(stat, "Event.EventData.Hashes").String()
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["IsExecutable"] = gjson.Get(stat, "Event.EventData.IsExecutable").String()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").String()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["TargetFilename"] = gjson.Get(stat, "Event.EventData.TargetFilename").String()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 24:
		eventDataStructure["Archived"] = gjson.Get(stat, "Event.EventData.Archived").Bool()
		eventDataStructure["ClientInfo"] = gjson.Get(stat, "Event.EventData.ClientInfo").String()
		eventDataStructure["Hashes"] = gjson.Get(stat, "Event.EventData.Hashes").String()
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").String()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["Session"] = gjson.Get(stat, "Event.EventData.Session").Int()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 25:
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").String()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["Type"] = gjson.Get(stat, "Event.EventData.Type").String()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		break

	case 26:
		eventDataStructure["Hashes"] = gjson.Get(stat, "Event.EventData.Hashes").String()
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["IsExecutable"] = gjson.Get(stat, "Event.EventData.IsExecutable").String()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").String()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["TargetFilename"] = gjson.Get(stat, "Event.EventData.TargetFilename").String()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		break

	case 27:
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").String()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["TargetFilename"] = gjson.Get(stat, "Event.EventData.TargetFilename").String()
		eventDataStructure["Hashes"] = gjson.Get(stat, "Event.EventData.Hashes").String()
		break

	case 28:
		eventDataStructure["RuleName"] = gjson.Get(stat, "Event.EventData.RuleName").String()
		eventDataStructure["UtcTime"] = gjson.Get(stat, "Event.EventData.UtcTime").Time()
		eventDataStructure["ProcessGuid"] = gjson.Get(stat, "Event.EventData.ProcessGuid").String()
		eventDataStructure["ProcessId"] = gjson.Get(stat, "Event.EventData.ProcessId").Int()
		eventDataStructure["User"] = gjson.Get(stat, "Event.EventData.User").String()
		eventDataStructure["Image"] = gjson.Get(stat, "Event.EventData.Image").String()
		eventDataStructure["TargetFilename"] = gjson.Get(stat, "Event.EventData.TargetFilename").String()
		eventDataStructure["Hashes"] = gjson.Get(stat, "Event.EventData.Hashes").String()
		eventDataStructure["IsExecutable"] = gjson.Get(stat, "Event.EventData.IsExecutable").String()
		break

	default:
		log.Panic(fmt.Sprintf("ID %d received; seems to be invalid Sysmon EVTX file.\n", id))
	}

	return eventDataStructure
}
