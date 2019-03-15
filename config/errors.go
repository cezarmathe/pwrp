package config

/*ErrCreateStorageDir describes the creation of a certain storage directory*/
type ErrCreateStorageDir struct {
	Dirname string
}

/*NewErrCreateStorageDir returns a new ErrCreateStorageDir*/
func NewErrCreateStorageDir(dirname string) *ErrCreateStorageDir {
	return &ErrCreateStorageDir{dirname}
}

func (err *ErrCreateStorageDir) Error() string {
	return "encountered an error when creating the storage directory at \"" + err.Dirname + "\""
}

/*ErrNoPermissions describes a missing branch*/
type ErrNoPermissions struct {
	Path string
}

/*NewErrNoPermissions returns a new ErrNoPermissions*/
func NewErrNoPermissions(path string) *ErrNoPermissions {
	return &ErrNoPermissions{path}
}

func (err *ErrNoPermissions) Error() string {
	return "no permissions for writing data at " + err.Path
}
