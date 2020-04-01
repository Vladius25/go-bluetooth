// Code generated DO NOT EDIT

package input



import (
   "sync"
   "github.com/muka/go-bluetooth/bluez"
   "github.com/muka/go-bluetooth/util"
   "github.com/muka/go-bluetooth/props"
   "github.com/godbus/dbus"
)

var Input1Interface = "org.bluez.Input1"


// NewInput1 create a new instance of Input1
//
// Args:
// - objectPath: [variable prefix]/{hci0,hci1,...}/dev_XX_XX_XX_XX_XX_XX
func NewInput1(objectPath dbus.ObjectPath) (*Input1, error) {
	a := new(Input1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: Input1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(Input1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


/*
Input1 Input hierarchy

*/
type Input1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	objectManagerSignal chan *dbus.Signal
	objectManager       *bluez.ObjectManager
	Properties 				*Input1Properties
	watchPropertiesChannel chan *dbus.Signal
}

// Input1Properties contains the exposed properties of an interface
type Input1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

	/*
	ReconnectMode Determines the Connectability mode of the HID device as
			defined by the HID Profile specification, Section 5.4.2.

			This mode is based in the two properties
			HIDReconnectInitiate (see Section 5.3.4.6) and
			HIDNormallyConnectable (see Section 5.3.4.14) which
			define the following four possible values:

			"none"		Device and host are not required to
					automatically restore the connection.

			"host"		Bluetooth HID host restores connection.

			"device"	Bluetooth HID device restores
					connection.

			"any"		Bluetooth HID device shall attempt to
					restore the lost connection, but
					Bluetooth HID Host may also restore the
					connection.
	*/
	ReconnectMode string

}

//Lock access to properties
func (p *Input1Properties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *Input1Properties) Unlock() {
	p.lock.Unlock()
}






// GetReconnectMode get ReconnectMode value
func (a *Input1) GetReconnectMode() (string, error) {
	v, err := a.GetProperty("ReconnectMode")
	if err != nil {
		return "", err
	}
	return v.Value().(string), nil
}



// Close the connection
func (a *Input1) Close() {
	
	a.unregisterPropertiesSignal()
	
	a.client.Disconnect()
}

// Path return Input1 object path
func (a *Input1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return Input1 dbus client
func (a *Input1) Client() *bluez.Client {
	return a.client
}

// Interface return Input1 interface
func (a *Input1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *Input1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

	if a.objectManagerSignal == nil {
		if a.objectManager == nil {
			om, err := bluez.GetObjectManager()
			if err != nil {
				return nil, nil, err
			}
			a.objectManager = om
		}

		s, err := a.objectManager.Register()
		if err != nil {
			return nil, nil, err
		}
		a.objectManagerSignal = s
	}

	cancel := func() {
		if a.objectManagerSignal == nil {
			return
		}
		a.objectManagerSignal <- nil
		a.objectManager.Unregister(a.objectManagerSignal)
		a.objectManagerSignal = nil
	}

	return a.objectManagerSignal, cancel, nil
}


// ToMap convert a Input1Properties to map
func (a *Input1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an Input1Properties
func (a *Input1Properties) FromMap(props map[string]interface{}) (*Input1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an Input1Properties
func (a *Input1Properties) FromDBusMap(props map[string]dbus.Variant) (*Input1Properties, error) {
	s := new(Input1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *Input1) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *Input1) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *Input1) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *Input1) GetProperties() (*Input1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *Input1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *Input1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *Input1) GetPropertiesSignal() (chan *dbus.Signal, error) {

	if a.propertiesSignal == nil {
		s, err := a.client.Register(a.client.Config.Path, bluez.PropertiesInterface)
		if err != nil {
			return nil, err
		}
		a.propertiesSignal = s
	}

	return a.propertiesSignal, nil
}

// Unregister for changes signalling
func (a *Input1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *Input1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *Input1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}



