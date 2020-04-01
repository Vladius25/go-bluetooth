// Code generated DO NOT EDIT

package media



import (
   "sync"
   "github.com/muka/go-bluetooth/bluez"
   "github.com/muka/go-bluetooth/util"
   "github.com/muka/go-bluetooth/props"
   "github.com/godbus/dbus"
   "fmt"
)

var MediaControl1Interface = "org.bluez.MediaControl1"


// NewMediaControl1 create a new instance of MediaControl1
//
// Args:
// - objectPath: [variable prefix]/{hci0,hci1,...}/dev_XX_XX_XX_XX_XX_XX
func NewMediaControl1(objectPath dbus.ObjectPath) (*MediaControl1, error) {
	a := new(MediaControl1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: MediaControl1Interface,
			Path:  dbus.ObjectPath(objectPath),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(MediaControl1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}

// NewMediaControl1FromAdapterID create a new instance of MediaControl1
// adapterID: ID of an adapter eg. hci0
func NewMediaControl1FromAdapterID(adapterID string) (*MediaControl1, error) {
	a := new(MediaControl1)
	a.client = bluez.NewClient(
		&bluez.Config{
			Name:  "org.bluez",
			Iface: MediaControl1Interface,
			Path:  dbus.ObjectPath(fmt.Sprintf("/org/bluez/%s", adapterID)),
			Bus:   bluez.SystemBus,
		},
	)
	
	a.Properties = new(MediaControl1Properties)

	_, err := a.GetProperties()
	if err != nil {
		return nil, err
	}
	
	return a, nil
}


/*
MediaControl1 Media Control hierarchy

*/
type MediaControl1 struct {
	client     				*bluez.Client
	propertiesSignal 	chan *dbus.Signal
	objectManagerSignal chan *dbus.Signal
	objectManager       *bluez.ObjectManager
	Properties 				*MediaControl1Properties
	watchPropertiesChannel chan *dbus.Signal
}

// MediaControl1Properties contains the exposed properties of an interface
type MediaControl1Properties struct {
	lock sync.RWMutex `dbus:"ignore"`

	/*
	Connected 
	*/
	Connected bool

	/*
	Player Addressed Player object path.
	*/
	Player dbus.ObjectPath

}

//Lock access to properties
func (p *MediaControl1Properties) Lock() {
	p.lock.Lock()
}

//Unlock access to properties
func (p *MediaControl1Properties) Unlock() {
	p.lock.Unlock()
}






// GetConnected get Connected value
func (a *MediaControl1) GetConnected() (bool, error) {
	v, err := a.GetProperty("Connected")
	if err != nil {
		return false, err
	}
	return v.Value().(bool), nil
}






// GetPlayer get Player value
func (a *MediaControl1) GetPlayer() (dbus.ObjectPath, error) {
	v, err := a.GetProperty("Player")
	if err != nil {
		return dbus.ObjectPath(""), err
	}
	return v.Value().(dbus.ObjectPath), nil
}



// Close the connection
func (a *MediaControl1) Close() {
	
	a.unregisterPropertiesSignal()
	
	a.client.Disconnect()
}

// Path return MediaControl1 object path
func (a *MediaControl1) Path() dbus.ObjectPath {
	return a.client.Config.Path
}

// Client return MediaControl1 dbus client
func (a *MediaControl1) Client() *bluez.Client {
	return a.client
}

// Interface return MediaControl1 interface
func (a *MediaControl1) Interface() string {
	return a.client.Config.Iface
}

// GetObjectManagerSignal return a channel for receiving updates from the ObjectManager
func (a *MediaControl1) GetObjectManagerSignal() (chan *dbus.Signal, func(), error) {

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


// ToMap convert a MediaControl1Properties to map
func (a *MediaControl1Properties) ToMap() (map[string]interface{}, error) {
	return props.ToMap(a), nil
}

// FromMap convert a map to an MediaControl1Properties
func (a *MediaControl1Properties) FromMap(props map[string]interface{}) (*MediaControl1Properties, error) {
	props1 := map[string]dbus.Variant{}
	for k, val := range props {
		props1[k] = dbus.MakeVariant(val)
	}
	return a.FromDBusMap(props1)
}

// FromDBusMap convert a map to an MediaControl1Properties
func (a *MediaControl1Properties) FromDBusMap(props map[string]dbus.Variant) (*MediaControl1Properties, error) {
	s := new(MediaControl1Properties)
	err := util.MapToStruct(s, props)
	return s, err
}

// ToProps return the properties interface
func (a *MediaControl1) ToProps() bluez.Properties {
	return a.Properties
}

// GetWatchPropertiesChannel return the dbus channel to receive properties interface
func (a *MediaControl1) GetWatchPropertiesChannel() chan *dbus.Signal {
	return a.watchPropertiesChannel
}

// SetWatchPropertiesChannel set the dbus channel to receive properties interface
func (a *MediaControl1) SetWatchPropertiesChannel(c chan *dbus.Signal) {
	a.watchPropertiesChannel = c
}

// GetProperties load all available properties
func (a *MediaControl1) GetProperties() (*MediaControl1Properties, error) {
	a.Properties.Lock()
	err := a.client.GetProperties(a.Properties)
	a.Properties.Unlock()
	return a.Properties, err
}

// SetProperty set a property
func (a *MediaControl1) SetProperty(name string, value interface{}) error {
	return a.client.SetProperty(name, value)
}

// GetProperty get a property
func (a *MediaControl1) GetProperty(name string) (dbus.Variant, error) {
	return a.client.GetProperty(name)
}

// GetPropertiesSignal return a channel for receiving udpdates on property changes
func (a *MediaControl1) GetPropertiesSignal() (chan *dbus.Signal, error) {

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
func (a *MediaControl1) unregisterPropertiesSignal() {
	if a.propertiesSignal != nil {
		a.propertiesSignal <- nil
		a.propertiesSignal = nil
	}
}

// WatchProperties updates on property changes
func (a *MediaControl1) WatchProperties() (chan *bluez.PropertyChanged, error) {
	return bluez.WatchProperties(a)
}

func (a *MediaControl1) UnwatchProperties(ch chan *bluez.PropertyChanged) error {
	return bluez.UnwatchProperties(a, ch)
}




/*
Play 
			Resume playback.


*/
func (a *MediaControl1) Play() error {
	
	return a.client.Call("Play", 0, ).Store()
	
}

/*
Pause 
			Pause playback.


*/
func (a *MediaControl1) Pause() error {
	
	return a.client.Call("Pause", 0, ).Store()
	
}

/*
Stop 
			Stop playback.


*/
func (a *MediaControl1) Stop() error {
	
	return a.client.Call("Stop", 0, ).Store()
	
}

/*
Next 
			Next item.


*/
func (a *MediaControl1) Next() error {
	
	return a.client.Call("Next", 0, ).Store()
	
}

/*
Previous 
			Previous item.


*/
func (a *MediaControl1) Previous() error {
	
	return a.client.Call("Previous", 0, ).Store()
	
}

/*
VolumeUp 
			Adjust remote volume one step up


*/
func (a *MediaControl1) VolumeUp() error {
	
	return a.client.Call("VolumeUp", 0, ).Store()
	
}

/*
VolumeDown 
			Adjust remote volume one step down


*/
func (a *MediaControl1) VolumeDown() error {
	
	return a.client.Call("VolumeDown", 0, ).Store()
	
}

/*
FastForward 
			Fast forward playback, this action is only stopped
			when another method in this interface is called.


*/
func (a *MediaControl1) FastForward() error {
	
	return a.client.Call("FastForward", 0, ).Store()
	
}

/*
Rewind 
			Rewind playback, this action is only stopped
			when another method in this interface is called.

Properties

		boolean Connected [readonly]

		object Player [readonly, optional]

			Addressed Player object path.



*/
func (a *MediaControl1) Rewind() error {
	
	return a.client.Call("Rewind", 0, ).Store()
	
}
