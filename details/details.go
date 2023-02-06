package details

import "os"

func GetHostName() (string, error) {
	hostName, err := os.Hostname()

	return hostName, err
}