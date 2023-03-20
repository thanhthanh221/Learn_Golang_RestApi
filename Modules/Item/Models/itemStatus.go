package Models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDelete
)

var AllItemStatuses = [3]string{"Doing", "Done", "Deleted"}

func (item *ItemStatus) String() string {
	return AllItemStatuses[*item]
}
func parseSrt2ItemStatus(s string) (ItemStatus, error) {
	for i := range AllItemStatuses {
		if AllItemStatuses[i] == s {
			return ItemStatus(i), nil
		}
	}
	return ItemStatus(0), errors.New("invalid status string")
}
func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New(fmt.Sprintln("fail to scan data from sql: %s", value))
	}
	v, err := parseSrt2ItemStatus(string(bytes))

	if err != nil {
		return errors.New(fmt.Sprintln("fail to scan data from sql: %s", value))
	}
	*item = v
	return nil
}
func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}
func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}
	return item.String(), nil
}
func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")
	itemValue, err := parseSrt2ItemStatus(str)

	if err != nil {
		return err
	}
	*item = itemValue
	return nil
}
