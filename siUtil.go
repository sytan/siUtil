//Package siUtil implements dll operation
package siUtil

import (
	"fmt"
	"strconv"
	"syscall"
	"unsafe"
)

// Error code
const (
	SUCCESS                                = "The operation completed successfully." //operation completed successfully
	ErrorUnableWriteFlash                  = 0x101                                   //Unable to write to Flash.
	ErrorUnableLockFlash                   = 0x102                                   //Unable to lock Flash.
	ErrorTargetNotHalted                   = 0x103                                   //Target device is not halted.
	ErrorInvalidPath                       = 0x104                                   //Invalid path.
	ErrorUnableOpenCOMPort                 = 0x105                                   //Unable to open COM Port.
	ErrorUnableDownlad                     = 0x106                                   //Unable to download program.
	ErrorResetFailed                       = 0x107                                   //Reset failed.
	ErrorEraseFailed                       = 0x108                                   //Erase failed.
	ErrorUnableCloseCOMPort                = 0x109                                   //Unable to close COM Port.
	ErrorDebugAdapterDLLNotFound           = 0x10A                                   //USB Debug Adapter DLL not found.
	ErrorUnableOpenUSBAdapter              = 0x10B                                   //Unable to open USB Debug Adapter.
	ErrorExternalMemoryNotSupported        = 0x10C                                   //External memory is not supported on this device.
	ErrorTargetConnected                   = 0x10D                                   //Target device is connected.
	ErrorUnknownDevice                     = 0x10E                                   //Unknown device.
	ErrorInvalidTargetResponse             = 0x10F                                   //Invalid target response.
	ErrorInvalidArgument                   = 0x110                                   //Invalid arguments.
	ErrorTargetNotRespondMaybeDisconnected = 0x111                                   //Target did not respond and may be disconnected.
	ErrorTargetNotRespondWithEnoughData    = 0x112                                   //Target did not respond with enough data.
	ErrorPortBusy                          = 0x113                                   //The communication port is busy.
	ErrorPortNotOpen                       = 0x114                                   //The communication port is not open.
	ErrorTimerNotAvailable                 = 0x115                                   //Timer resource not available.
	ErrorUnableWriteAllData                = 0x116                                   //Unable to write all the data to the target.
	ErrorNoNewDataReturn                   = 0x117                                   //Serial Interface had no new data to return.
	ErrorUARTNotEnabled                    = 0x118                                   //UART has not been enabled via GPIO configuration.
	ErrorTargetNotInHalt                   = 0x119                                   //Target is not in the HALT state and did not respond to the query.
	ErrorFlashFailure                      = 0x11A                                   //A Flash failure occurred.
	ErrorNotConnectedToTarget              = 0x11B                                   //Not connected to target. (This is a disconnect error, so the target disconnected flagbit will also be set.)
	ErrorTargetNotRespond                  = 0x11C                                   //The target is not responding.
	ErrorCommandFailed                     = 0x11D                                   //The target reported that the command failed.
	ErrorCommunicationPortTimeOut          = 0x11E                                   //Communication port time out.
	ErrorUnableClearError                  = 0x11F                                   //Could not clear communication port error.
	ErrorReceivingByte                     = 0x120                                   //Error receiving byte.
	ErrorReceiveBufferFailure              = 0x121                                   //Receive Buffer Failure.
	ErrorPurgePort                         = 0x122                                   //Error occurred during purge of communication port.
	ErrorFirmwareUpdate                    = 0x123                                   //Error occurred during firmware update.
	ErrorCRCFailure                        = 0x124                                   //A CRC comparison failure has occurred.
	ErrorDeviceReset                       = 0x125                                   //An error occurred during device reset.
	ErrorOperationNotSupported             = 0x126                                   //This operation is not supported on this device.
	ErrorFlashLocked                       = 0x127                                   //Flash is locked.
	ErrorIncorrectResponse                 = 0x128                                   //An incorrect response has been received.
	ErrorAdapterNotAssigned                = 0x129                                   //Valid adapter selection has not been assigned.
	ErrorVoidUSBAdapterHandle              = 0x12A                                   //The USB Debug Adapter handle is void.
	ErrorUnknownFunction                   = 0x12B                                   //USB Debug Adapter unknown function.
	ErrorEraseCancelled                    = 0x12C                                   //Flash erase has been cancelled.
	ErrorPartNotLoaded                     = 0x12D                                   //Part not loaded in XML services.
	ErrorMissingInformation                = 0x12E                                   //The MCU product configuration data is missing information about this device.
	ErrorUnableAccessConfigurationData     = 0x12F                                   //Unable to access MCU product configuration data.
	ErrorDeviceVersionNotSupported         = 0x130                                   //The version of the device you connected to is older than the minimum version supported.
)

var (
	err    error
	siUtil syscall.Handle
)

func abort(dllname string, err error) {
	panic(fmt.Sprintf("Failed to load library %s: %v", dllname, err))
}
func init() {
	dllname := "SiUtil.dll"
	siUtil, err = syscall.LoadLibrary(dllname)
	if err != nil {
		abort(dllname, err)
	}
}

// FreeLibrary implements release loaded labrary
func FreeLibrary() {
	syscall.FreeLibrary(siUtil)
}

// Connect implements Connects to a target C8051Fxxx device using a Serial Adapter.
func Connect(nComPort, nDisableDialogBoxes, nECprotocol, nBaudRateIndex int) int32 {
	connect, _ := syscall.GetProcAddress(siUtil, "Connect")
	num := 4
	r, _, _ := syscall.Syscall6(
		uintptr(connect),
		uintptr(num),
		uintptr(nComPort),
		uintptr(nDisableDialogBoxes),
		uintptr(nECprotocol),
		uintptr(nBaudRateIndex),
		0,
		0)
	return int32(r)
}

// Disconnect implements Disconnects from a target C8051Fxxx device using a Serial Adapter.
func Disconnect(nComPort int) int32 {
	disconnect, _ := syscall.GetProcAddress(siUtil, "Disconnect")
	num := 1
	r, _, _ := syscall.Syscall(uintptr(disconnect), uintptr(num), uintptr(nComPort), 0, 0)
	return int32(r)
}

// ConnectUSB implements Connects to a target C8051Fxxx device using a USB Debug Adapter.
func ConnectUSB(sSerialNumber string, nECprotocol int, nPowerTarget int, nDisableDialogBoxes int) int32 {
	connectUSB, _ := syscall.GetProcAddress(siUtil, "ConnectUSB")
	num := 4
	r, _, _ := syscall.Syscall6(
		uintptr(connectUSB),
		uintptr(num),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(sSerialNumber))),
		uintptr(nECprotocol),
		uintptr(nPowerTarget),
		uintptr(nDisableDialogBoxes),
		0,
		0)
	return int32(r)
}

// DisconnectUSB implements Disconnects from a target C8051Fxxx device using a USB Debug Adapter.
func DisconnectUSB() int32 {
	disconnectUSB, _ := syscall.GetProcAddress(siUtil, "DisconnectUSB")
	num := 0
	r, _, _ := syscall.Syscall(uintptr(disconnectUSB), uintptr(num), 0, 0, 0)
	return int32(r)
}

// Connected implements Returns the connection state of the target C8051Fxxx device.
func Connected() int {
	connected, _ := syscall.GetProcAddress(siUtil, "Connected")
	num := 0
	r, _, _ := syscall.Syscall(uintptr(connected), uintptr(num), 0, 0, 0)
	return int(r)
}

// Download implements Downloads a hex file to the target C8051Fxxx device.
func Download(sDownloadFile string, nDeviceErase int, nDisableDialogBoxed int, nDownloadScratchPadSFLF int, nBankSelect int, nLockFlash int, bPersistFlash int) int32 {
	download, _ := syscall.GetProcAddress(siUtil, "Download")
	num := 7
	f := []byte(sDownloadFile)
	r, _, _ := syscall.Syscall9(
		uintptr(download),
		uintptr(num),
		uintptr(unsafe.Pointer(&f[0])), //sDownloadFile—A character pointer to the beginning of a character array (string) containing the full path and filename of the file to be downloaded.
		uintptr(nDeviceErase),
		uintptr(nDisableDialogBoxed),
		uintptr(nDownloadScratchPadSFLF),
		uintptr(nBankSelect),
		uintptr(nLockFlash),
		uintptr(bPersistFlash),
		0,
		0)
	return int32(r)
}

// ISupportBanking implements Checks to see if banking is supported.
func ISupportBanking() (int, int32) {
	iSupportBanking, _ := syscall.GetProcAddress(siUtil, "ISupportBanking")
	num := 1
	var nSupportedBanks int
	r, _, _ := syscall.Syscall(uintptr(iSupportBanking), uintptr(num), uintptr(unsafe.Pointer(&nSupportedBanks)), 0, 0)
	return nSupportedBanks, int32(r)
}

//GetSAFirmwareVersion implements Returns the Serial Adapter firmware version.
func GetSAFirmwareVersion() string {
	getSAFirmwareVersion, _ := syscall.GetProcAddress(siUtil, "GetSAFirmwareVersion")
	num := 0
	r, _, _ := syscall.Syscall(uintptr(getSAFirmwareVersion), uintptr(num), 0, 0, 0)
	return strconv.Itoa(int(r))
}

// GetUSBFirmwareVersion implements Returns the USB Debug Adapter firmware version.
func GetUSBFirmwareVersion() string {
	getUSBFirmwareVersion, _ := syscall.GetProcAddress(siUtil, "GetUSBFirmwareVersion")
	num := 0
	r, _, _ := syscall.Syscall(uintptr(getUSBFirmwareVersion), uintptr(num), 0, 0, 0)
	return strconv.Itoa(int(r))
}

// GetDLLVersion implements get version of dll
func GetDLLVersion() string {
	getDLLVersion, _ := syscall.GetProcAddress(siUtil, "GetDLLVersion")
	num := 0
	r, _, _ := syscall.Syscall(uintptr(getDLLVersion), uintptr(num), 0, 0, 0)
	bStr := []byte{}
	for i := 0; i < 256; i++ {
		b := (*byte)(unsafe.Pointer(r + uintptr(i)))
		if *b == 0 || *b == 0xFF {
			break
		}
		bStr = append(bStr, *b)
	}
	return string(bStr)
}

// GetDeviceName implements Returns the name of the target C8051Fxxx device.
func GetDeviceName() (string, int32) {
	sDeviceName := make([]byte, 9) //long enough to hold devicename
	getDeviceName, _ := syscall.GetProcAddress(siUtil, "GetDeviceName")
	num := 1
	r, _, _ := syscall.Syscall(uintptr(getDeviceName), uintptr(num), uintptr(unsafe.Pointer(&sDeviceName)), 0, 0)
	return string(sDeviceName), int32(r)
}

// GetRAMMemory implements Read RAM memory from a specified address.
func GetRAMMemory() {

}

// GetXRAMMemory implements Read XRAM memory from a specified address.
func GetXRAMMemory() {

}

// GetCodeMemory uintptr implements Read Code memory from a specified address.
func GetCodeMemory(nStartAddress, nLength uint32) ([]byte, int32) {
	buf := make([]byte, nLength)
	getCodeMemory, _ := syscall.GetProcAddress(siUtil, "GetCodeMemory")
	num := 3
	r, _, _ := syscall.Syscall(uintptr(getCodeMemory), uintptr(num), uintptr(unsafe.Pointer(&buf[0])), uintptr(nStartAddress), uintptr(nLength))
	return buf, int32(r)
}

// SetRAMMemory implements Writes value to a specified address in RAM memory.
func SetRAMMemory() {

}

// SetXRAMMemory implements Writes value to a specified address in XRAM memory.
func SetXRAMMemory() {

}

//SetCodeMemory implements Writes value to a specified address in code memory.
func SetCodeMemory(buf []byte, nStartAddress uint32, nLength uint32, nDisableDialogs int) int32 {
	setCodeMemory, _ := syscall.GetProcAddress(siUtil, "SetCodeMemory")
	num := 4
	r, _, _ := syscall.Syscall6(uintptr(setCodeMemory), uintptr(num), uintptr(unsafe.Pointer(&buf[0])), uintptr(nStartAddress), uintptr(nLength), uintptr(nDisableDialogs), 0, 0)
	return int32(r)
}

// SetTargetGo implements Puts the target C8051Fxxx device in a Run state.
func SetTargetGo() int32 {
	setTargetGo, _ := syscall.GetProcAddress(siUtil, "SetTargetGo")
	num := 0
	r, _, _ := syscall.Syscall(uintptr(setTargetGo), uintptr(num), 0, 0, 0)
	return int32(r)
}

// SetTargetHalt implements Puts the target C8051Fxxx device in a Halt state.
func SetTargetHalt() int32 {
	setTargetHalt, _ := syscall.GetProcAddress(siUtil, "SetTargetHalt")
	num := 0
	r, _, _ := syscall.Syscall(uintptr(setTargetHalt), uintptr(num), 0, 0, 0)
	return int32(r)
}

// USBDebugDevices implements Determines how many USB Debug Adapters are present.
func USBDebugDevices() (int, int32) {
	var nDevices int
	usbDebugDevices, _ := syscall.GetProcAddress(siUtil, "USBDebugDevices")
	num := 1
	r, _, _ := syscall.Syscall(uintptr(usbDebugDevices), uintptr(num), uintptr(unsafe.Pointer(&nDevices)), 0, 0)
	return nDevices, int32(r)
}

// GetUSBDeviceSN implements Obtains the serial number of the enumerated USB Debug Adapters.
func GetUSBDeviceSN(nDeviceNum int) (string, int32) {
	sSerialNum := make([]byte, 11) //lenth of sSerialNum should be as same as usbDeviceSN
	getUSBDeviceSN, _ := syscall.GetProcAddress(siUtil, "GetUSBDeviceSN")
	num := 2
	r, _, _ := syscall.Syscall(uintptr(getUSBDeviceSN), uintptr(num), uintptr(nDeviceNum), uintptr(unsafe.Pointer(&sSerialNum)), 0)
	return string(sSerialNum), int32(r)
}

// GetUSBDLLVersion implements Returns the version of the USB Debug Adapter driver file.
func GetUSBDLLVersion() (string, int32) {
	sVersion := make([]byte, 7)
	getUSBDLLVersion, _ := syscall.GetProcAddress(siUtil, "GetUSBDLLVersion")
	num := 1
	r, _, _ := syscall.Syscall(uintptr(getUSBDLLVersion), uintptr(num), uintptr(unsafe.Pointer(&sVersion)), 0, 0)
	return string(sVersion), int32(r)
}

// FLASHErase implements Erase the Flash program memory using a Serial Adapter.
func FLASHErase(nComPort, nDisableDialogBoxes, nECprotocol int) int32 {
	flashErase, _ := syscall.GetProcAddress(siUtil, "FLASHErase")
	num := 3
	r, _, _ := syscall.Syscall(
		uintptr(flashErase),
		uintptr(num),
		uintptr(nComPort),
		uintptr(nDisableDialogBoxes),
		uintptr(nECprotocol))
	return int32(r)
}

// FLASHEraseUSB implements Erase the Flash program memory using a USB Debug Adapter.
func FLASHEraseUSB(sSerialNumber string, nDisableDialogBoxes, nECprotocol int) int32 {
	flashEraseUSB, _ := syscall.GetProcAddress(siUtil, "FLASHEraseUSB")
	num := 3
	r, _, _ := syscall.Syscall(
		uintptr(flashEraseUSB),
		uintptr(num),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(sSerialNumber))),
		uintptr(nDisableDialogBoxes),
		uintptr(nECprotocol))
	return int32(r)
}

// SetJTAGDeviceAndConnect implements Configure a connection to a C8051Fxxx device on a JTAG chain using a Serial Adapter.
func SetJTAGDeviceAndConnect() {

}

// SetJTAGDeviceAndConnectUSB implements Configure a connection to a C8051Fxxx device on a JTAG chain using a USB Debug Adapter.
func SetJTAGDeviceAndConnectUSB() {

}

// GetErrorMsg implements returns a string describing an error which has been generated by the DLL.
func GetErrorMsg(errorCode int64) string {
	getErrorMsg, _ := syscall.GetProcAddress(siUtil, "GetErrorMsg")
	num := 1
	r, _, _ := syscall.Syscall(uintptr(getErrorMsg), uintptr(num), uintptr(errorCode), 0, 0)
	bStr := []byte{}
	for i := 0; i < 256; i++ {
		b := (*byte)(unsafe.Pointer(r + uintptr(i)))
		if *b == 0x0 || *b == 0xFF {
			break
		}
		bStr = append(bStr, *b)
	}
	fmt.Println(r, bStr)
	return string(bStr)
}
