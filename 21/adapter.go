package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.


type Monitor struct {}

func (m *Monitor) PlugIn (cabel MonitorIn) {
	cabel.PlugInVGA()
	fmt.Println("Cabel inserted")
}
//  interface for connecting a VGA cabel
type MonitorIn interface {
	PlugInVGA()
}
// СableVGA cabel
type СableVGA struct {}

// PlugInHDMI - connects monitor and cabel VGA
func (vga *СableVGA) PlugInVGA(){
	fmt.Println("Cabel plugged vga")
}
// СableHDMI cabel
type СableHDMI struct {}

// PlugInHDMI - connects monitor and cabel HDMI
func (hmdi *СableHDMI) PlugInHDMI(){
	fmt.Println("Cabel plugged hdmi")
}

// AdapterHDMIToVGA adapter afford to connect HDMI cabel to VGA interface
type AdapterHDMIToVGA struct {
	hdmi *СableHDMI
}

func (adapter *AdapterHDMIToVGA) PlugInVGA(){
	fmt.Println("Cabel plugged vga")
	adapter.hdmi.PlugInHDMI()
}


func main() {

	vga := СableVGA{}
	hmdi := СableHDMI{}

	monitor := Monitor{}

	adapter := AdapterHDMIToVGA{
		hdmi : &hmdi,
	}

	monitor.PlugIn(&vga)
	monitor.PlugIn(&adapter)

}
