package recording

/*Recorder is a struct that does the recording process.*/
type Recorder struct {
	Config *Config
}

/*NewRecorder creates a new Recorder with the specified configuration.*/
func NewRecorder(config *Config) *Recorder {
	return &Recorder{config}
}

/*Record starts the recording process.*/
func (recorder *Recorder) Record() bool {
	return true
}
