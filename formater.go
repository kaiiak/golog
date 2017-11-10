package golog

// Formater All leval output function
type Formater interface {
	Printer
	Debuger
	Infoer
	Warner
	Errorer
	Fataler
	Panicer
}

// Printer Normal output function
type Printer interface {
	Print(args ...interface{})
	Println(args ...interface{})
	Printf(format string, args ...interface{})
}

// Debuger Output debug message
// if IsDebug() return false, function will not output
type Debuger interface {
	IsDebug() bool
	Debug(args ...interface{})
	Debugln(args ...interface{})
	Debugf(format string, args ...interface{})
}

// Infoer Output information message
type Infoer interface {
	Info(args ...interface{})
	Infoln(args ...interface{})
	Infof(format string, args ...interface{})
}

// Warner Output warning message
type Warner interface {
	Warn(args ...interface{})
	Warnln(args ...interface{})
	Warnf(format string, args ...interface{})
}

// Errorer Output error message
type Errorer interface {
	Error(args ...interface{})
	Errorln(args ...interface{})
	Errorf(format string, args ...interface{})
}

// Panicer Output panic message
type Panicer interface {
	Panic(args ...interface{})
	Panicln(args ...interface{})
	Panicf(format string, args ...interface{})
}

// Fataler Output Fataler message
type Fataler interface {
	Fatal(args ...interface{})
	Fatalln(args ...interface{})
	Fatalf(format string, args ...interface{})
}
