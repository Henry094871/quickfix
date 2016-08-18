package config

//NOTE: Additions to this file should be made to both config/doc.go and http://www.quickfixgo.org/docs/

//Const configuration settings
const (
	BeginString              string = "BeginString"
	SenderCompID             string = "SenderCompID"
	TargetCompID             string = "TargetCompID"
	SessionQualifier         string = "SessionQualifier"
	SocketAcceptHost         string = "SocketAcceptHost"
	SocketAcceptPort         string = "SocketAcceptPort"
	SocketConnectHost        string = "SocketConnectHost"
	SocketConnectPort        string = "SocketConnectPort"
	SocketPrivateKeyFile     string = "SocketPrivateKeyFile"
	SocketCertificateFile    string = "SocketCertificateFile"
	SocketCAFile             string = "SocketCAFile"
	DefaultApplVerID         string = "DefaultApplVerID"
	StartTime                string = "StartTime"
	EndTime                  string = "EndTime"
	StartDay                 string = "StartDay"
	EndDay                   string = "EndDay"
	TimeZone                 string = "TimeZone"
	DataDictionary           string = "DataDictionary"
	TransportDataDictionary  string = "TransportDataDictionary"
	AppDataDictionary        string = "AppDataDictionary"
	ResetOnLogon             string = "ResetOnLogon"
	RefreshOnLogon           string = "RefreshOnLogon"
	ResetOnLogout            string = "ResetOnLogout"
	ReconnectInterval        string = "ReconnectInterval"
	HeartBtInt               string = "HeartBtInt"
	FileLogPath              string = "FileLogPath"
	FileStorePath            string = "FileStorePath"
	SQLDriver                string = "SQLDriver"
	SQLDataSourceName        string = "SQLDataSourceName"
	SQLConnMaxLifetime       string = "SQLConnMaxLifetime"
	ValidateFieldsOutOfOrder string = "ValidateFieldsOutOfOrder"
)
