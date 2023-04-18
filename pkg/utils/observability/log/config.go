package log

type Config struct {
	// Writer is used to specify where to send the logs.
	// default is os.Stdout, which is the standard output.
	// values can be: os.File, os.Stdout, os.Stderr.
	Writer string `yaml:"writer"`
	// Level is the minimum level to log.
	Level string `yaml:"level"`
	// Path is the path to the log file. which is used when Writer is os.File.  if the given path is a directory (the given path ends with a slash),
	// the logger will create a file named with the datetime to write log.
	Path string `yaml:"path"`
	// Format is the log format. value is json or text.
	Format string `yaml:"format"`
}
