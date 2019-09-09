package gocec

/*
#include <stdlib.h>
#include <libcec/cecc.h>

extern void invokeLogCallback(int, cec_log_message *);
static inline void logCallback(void *cbParam, const cec_log_message *message)
{
	invokeLogCallback(*(int*)cbParam, (cec_log_message*)message);
}

static inline void createCallbacks(libcec_configuration *config, int cbParam)
{
	config->callbacks = malloc(sizeof(struct ICECCallbacks));
	config->callbacks->logMessage = logCallback;
	config->callbacks->keyPress = NULL;
	config->callbacks->commandReceived = NULL;
	config->callbacks->configurationChanged = NULL;
	config->callbacks->alert = NULL;    
	config->callbacks->menuStateChanged = NULL;
	config->callbacks->sourceActivated = NULL;
	
	config->callbackParam = malloc(sizeof(int));
	*(int*)config->callbackParam = cbParam;
}
*/
import "C"

import (
	"sync"
	"time"
)

var mutex sync.Mutex
var index int
var registeredCallbacks = make(map[int]*callbacks)

type LogLevel int

const (
	LogLevelError LogLevel = 1
	LogLevelWarning LogLevel = 2
	LogLevelNotice LogLevel = 4
	LogLevelTraffic LogLevel = 8
	LogLevelDebug LogLevel = 16
	LogLevelAll LogLevel = 31
)

type LogMessage struct {
	Message string
	Level LogLevel
	Time time.Time
}

type LogCallback func (*LogMessage)

type callbacks struct {
	logCallback LogCallback
}

func initialiseCallbacks(config *Configuration) {
	config.callbacks = &callbacks{}
	
	mutex.Lock()
	defer mutex.Unlock()
	
	index++
	for registeredCallbacks[index] != nil {
		index++
	}
	registeredCallbacks[index] = config.callbacks
	
	C.createCallbacks(&config.configuration, C.int(index))
}

func (config *Configuration) SetLogCallback(callback LogCallback) {
	config.callbacks.logCallback = callback
}

//export invokeLogCallback
func invokeLogCallback(cbParam C.int, cMessage *C.cec_log_message) {
	callbacks := lookupCallbacks(int(cbParam))

	if callbacks.logCallback == nil {
		return
	}


	message := &LogMessage{
		Message: C.GoString(cMessage.message),
		Level: LogLevel(cMessage.level),
		Time: time.Unix(int64(cMessage.time), 0),
	}

	callbacks.logCallback(message)
}

func lookupCallbacks(i int) *callbacks {
	mutex.Lock()
	defer mutex.Unlock()
	
	return registeredCallbacks[i]
}
