package logconfig

import (
	"bytes"
	"html/template"
	"os"
	"path"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	yaml "gopkg.in/yaml.v2"
)

func format(templateString string, data map[string]interface{}) []byte {
	t := template.Must(template.New("logconfig").Parse(templateString))
	buf := &bytes.Buffer{}

	if err := t.Execute(buf, data); err != nil {
		panic(err)
	}

	return []byte(buf.String())
}

//NewZapLogger returns configured zap logger.
func NewZapLogger(loggingConfig []byte) (*zap.Logger, error) {
	var config zap.Config

	if err := yaml.Unmarshal(loggingConfig, &config); err != nil {
		panic(err)
	}
	var paths []string

	paths = append(paths, config.OutputPaths...)
	paths = append(paths, config.ErrorOutputPaths...)

	for _, file := range paths {
		// base := path.Base(file)
		dir := path.Dir(file)
		// ignore if not file path (stdout / stderr)
		if dir != "." {
			//create dirs and file if not already existed.
			if _, err := os.Stat(dir); os.IsNotExist(err) {
				os.MkdirAll(dir, os.ModePerm)
			}

			if _, err := os.Stat(file); os.IsNotExist(err) {
				os.OpenFile(file, os.O_CREATE|os.O_APPEND, 0755)
			}
		}
	}

	// config.OutputPaths
	// config.ErrorOutputPaths

	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	return config.Build()
}
