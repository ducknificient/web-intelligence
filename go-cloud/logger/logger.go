package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	configpackage "github.com/ducknificient/web-intelligence/go-cloud/config"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Info(string)
	Error(string)
	Fatal(string)
	CrawlLog(string)
	CrawlError(string)
}

type DefaultLogger struct {
	config         configpackage.Configuration
	logger         *zap.Logger
	crawlLogFile   *os.File
	crawlErrorFile *os.File
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func MyCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(filepath.Base(caller.FullPath()))
}

func NewLogger(config configpackage.Configuration) (logger *DefaultLogger, err error) {

	appconfig := config.GetConfiguration()

	//check log folder is exist
	if _, err := os.Stat(*appconfig.PathLog); os.IsNotExist(err) {
		err := os.MkdirAll(*appconfig.PathLog, 0744)
		if err != nil {
			return logger, err
		}
	}

	//create log file and setting rotate time (24 hours)
	// logFile := _pathlog + _filesep + "app-%Y-%m-%d-%H.log"
	logFile := *appconfig.PathLog + *appconfig.FileSep + "app-%Y-%m-%d.log"
	rotator, err := rotatelogs.New(
		logFile,
		rotatelogs.WithMaxAge(45*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour))
	if err != nil {
		return logger, err
	}

	// initialize the JSON encoding config
	encoderConfig := map[string]string{
		"levelEncoder": "lowercase",
		"levelKey":     "level",
		"timeKey":      "date",
		"timeEncoder":  "iso8601",
		"callerKey":    "caller",
		"messageKey":   "message",
	}
	data, _ := json.Marshal(encoderConfig)

	var encCfg zapcore.EncoderConfig
	if err := json.Unmarshal(data, &encCfg); err != nil {
		return logger, err
	}
	encCfg.EncodeTime = SyslogTimeEncoder
	encCfg.EncodeCaller = MyCaller

	// add the encoder config and rotator to create a new zap logger
	w := zapcore.AddSync(rotator)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encCfg),
		w,
		zap.InfoLevel)
	zap.New(core)

	logger = &DefaultLogger{
		config: config,
		logger: zap.New(core, zap.WithCaller(true), zap.AddStacktrace(zap.ErrorLevel)),
	}

	return logger, nil
}

func (l *DefaultLogger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *DefaultLogger) Error(msg string) {
	l.logger.Error(msg)
}

func (l *DefaultLogger) Fatal(msg string) {
	l.logger.Fatal(msg)
}

func (l *DefaultLogger) CheckEmptyLog() (err error) {

	folderList := make([]string, 0)

	//check log folder is exist
	if _, err := os.Stat(*l.config.GetConfiguration().PathCrawlLog); os.IsNotExist(err) {
		err := os.MkdirAll(*l.config.GetConfiguration().PathCrawlLog, 0744)
		if err != nil {
			return err
		}
	}

	//check log folder is exist
	if _, err := os.Stat(*l.config.GetConfiguration().PathCrawlError); os.IsNotExist(err) {
		err := os.MkdirAll(*l.config.GetConfiguration().PathCrawlError, 0744)
		if err != nil {
			return err
		}
	}

	folderList = append(folderList, *l.config.GetConfiguration().PathCrawlLog)
	folderList = append(folderList, *l.config.GetConfiguration().PathCrawlError)

	for _, folderPath := range folderList {

		// Read the directory contents
		files, err := os.ReadDir(folderPath)
		if err != nil {
			err = errors.New(fmt.Sprintf("Error reading directory: %v\n", err))
			return err
		}

		// Iterate over each file in the directory
		for _, file := range files {
			// Check if it's a regular file
			if file.Type().IsRegular() {
				filePath := filepath.Join(folderPath, file.Name())

				// Read the content of the file
				content, err := os.ReadFile(filePath)
				if err != nil {
					// fmt.Printf("Error reading file %s: %v\n", file.Name(), err)
					// continue

					err = errors.New(fmt.Sprintf("Error reading file %s: %v\n", file.Name(), err))
					return err
				}

				// Check if content is empty
				if len(content) == 0 {
					// Delete the file
					err := os.Remove(filePath)
					if err != nil {
						// fmt.Printf("Error deleting file %s: %v\n", file.Name(), err)
						// continue

						err = errors.New(fmt.Sprintf("Error deleting file %s: %v\n", file.Name(), err))
						return err
					}

					// l.Info(fmt.Sprintf("File %s deleted successfully\n", file.Name()))
				} else {
					// l.Info(fmt.Sprintf("File %s is not empty\n", file.Name()))
				}
			}
		}

	}

	return err
}

func (l *DefaultLogger) SetupCrawlLogFile() (err error) {
	// currentTime := time.Now().Format("2006-01-02_15-04-05")
	currentTime := time.Now().Format("2006-01-02_15")

	pathcrawllog := *l.config.GetConfiguration().PathCrawlLog
	filesep := *l.config.GetConfiguration().FileSep

	// filename := l.PathCrawlLog + *config.Conf.FileSep + currentTime + "_" + "crawl_log"
	filename := pathcrawllog + filesep + currentTime + "_" + "crawl_log"

	// filename := l.PathCrawlLog + *config.Conf.FileSep + currentTime + "_" + "crawl_log"

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	l.crawlLogFile = file

	pathcrawlerror := *l.config.GetConfiguration().PathCrawlError
	error_filename := pathcrawlerror + filesep + currentTime + `_` + `error_log`

	error_file, err := os.OpenFile(error_filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	l.crawlErrorFile = error_file

	return err

}

func (l *DefaultLogger) CrawlLog(msg string) {
	fmt.Fprintf(l.crawlLogFile, "%v", msg)
}

func (l *DefaultLogger) CrawlError(msg string) {
	fmt.Fprintf(l.crawlErrorFile, "%v", msg)
}

func combine() {
	// Define the folder containing the .txt files
	folderPath := "/home/spil/Projects/minicrawler/error-log"

	// Create a new file to store the combined content
	combinedFile, err := os.Create("old-error-log.txt")
	if err != nil {
		fmt.Println("Error creating combined file:", err)
		return
	}
	defer combinedFile.Close()

	// Get a list of .txt files in the folder
	files, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error reading folder:", err)
		return
	}

	// Iterate over each .txt file, read its content, and write it to the combined file
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".txt" {
			continue // Skip directories and non-.txt files
		}

		filePath := filepath.Join(folderPath, file.Name())

		// Read content of the current file
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}

		separator := file.Name() + "\n"

		// Add a newline after each file to separate their content
		if _, err := combinedFile.WriteString(separator); err != nil {
			fmt.Println("Error writing newline to combined file:", err)
			continue
		}

		// Write content to the combined file
		if _, err := combinedFile.Write(content); err != nil {
			fmt.Println("Error writing to combined file:", err)
			continue
		}

		separator = "\n"

		// Add a newline after each file to separate their content
		if _, err := combinedFile.WriteString(separator); err != nil {
			fmt.Println("Error writing newline to combined file:", err)
			continue
		}

		fmt.Println("File", file.Name(), "has been added to the combined file.")
	}

	fmt.Println("All .txt files have been combined into combined.txt")
}
