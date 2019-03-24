# PWRP Coding standards

- Comments

Documentation comments should be added in the form of `/**/`, like this:

```go
/*this is a documentation comment*/
func function() {}
```

`//` comments should be used for debugging purposes only, like this:

```go
/*this is a documentation comment*/
func function() {
    // this is a debugging comemnt
    b := 3 + 5
}
```

- Logging

Functions that start other functions(or processes) should log the start & end events themselves, and delegate the rest of the logging to the called function.

For example, this is some sample code of starting the recording process:

```go
log.Debug("root cmd: ", "running recorder.Record()")
if success := recorder.Record(); success == false {
    log.Fatal("recorder reported cannot continue")
}
log.Info("recorder reported can continue")
```

```go
/*Record starts the recording process.*/
func (recorder *Recorder) Record() bool {
    log.Debug("recorder.Record(): ", "starting iteration over repository list")
    for _, repositoryURL := range recorder.Config.Repositories {
        repository, err := gitops.Clone(repositoryURL, recorder.Config.StoragePath)
        if err != nil {
            log.Fatal("recording: ", "fatal error encountered when cloning - ", err)
            return false
        }
        log.Info("recording: ", "repository "+repositoryURL+"cloned successfully")

        _, err = repository.Branch("_pwrp")
    }
    return true
}
```