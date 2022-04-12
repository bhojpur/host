package condition

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"reflect"
	"time"

	"github.com/bhojpur/host/pkg/common/generic"
	"github.com/sirupsen/logrus"
)

type Cond string

func (c Cond) GetStatus(obj interface{}) string {
	return getStatus(obj, string(c))
}

func (c Cond) SetError(obj interface{}, reason string, err error) {
	if err == nil || err == generic.ErrSkip {
		c.True(obj)
		c.Message(obj, "")
		c.Reason(obj, reason)
		return
	}
	if reason == "" {
		reason = "Error"
	}
	c.False(obj)
	c.Message(obj, err.Error())
	c.Reason(obj, reason)
}

func (c Cond) MatchesError(obj interface{}, reason string, err error) bool {
	if err == nil {
		return c.IsTrue(obj) &&
			c.GetMessage(obj) == "" &&
			c.GetReason(obj) == reason
	}
	if reason == "" {
		reason = "Error"
	}
	return c.IsFalse(obj) &&
		c.GetMessage(obj) == err.Error() &&
		c.GetReason(obj) == reason
}

func (c Cond) SetStatus(obj interface{}, status string) {
	setStatus(obj, string(c), status)
}

func (c Cond) SetStatusBool(obj interface{}, val bool) {
	if val {
		setStatus(obj, string(c), "True")
	} else {
		setStatus(obj, string(c), "False")
	}
}

func (c Cond) True(obj interface{}) {
	setStatus(obj, string(c), "True")
}

func (c Cond) IsTrue(obj interface{}) bool {
	return getStatus(obj, string(c)) == "True"
}

func (c Cond) False(obj interface{}) {
	setStatus(obj, string(c), "False")
}

func (c Cond) IsFalse(obj interface{}) bool {
	return getStatus(obj, string(c)) == "False"
}

func (c Cond) Unknown(obj interface{}) {
	setStatus(obj, string(c), "Unknown")
}

func (c Cond) IsUnknown(obj interface{}) bool {
	return getStatus(obj, string(c)) == "Unknown"
}

func (c Cond) LastUpdated(obj interface{}, ts string) {
	setTS(obj, string(c), ts)
}

func (c Cond) GetLastUpdated(obj interface{}) string {
	return getTS(obj, string(c))
}

func (c Cond) CreateUnknownIfNotExists(obj interface{}) {
	condSlice := getValue(obj, "Status", "Conditions")
	cond := findCond(obj, condSlice, string(c))
	if cond == nil {
		c.Unknown(obj)
	}
}

func (c Cond) Reason(obj interface{}, reason string) {
	cond := findOrCreateCond(obj, string(c))
	getFieldValue(cond, "Reason").SetString(reason)
}

func (c Cond) GetReason(obj interface{}) string {
	cond := findOrNotCreateCond(obj, string(c))
	if cond == nil {
		return ""
	}
	return getFieldValue(*cond, "Reason").String()
}

func (c Cond) SetMessageIfBlank(obj interface{}, message string) {
	if c.GetMessage(obj) == "" {
		c.Message(obj, message)
	}
}

func (c Cond) Message(obj interface{}, message string) {
	cond := findOrCreateCond(obj, string(c))
	setValue(cond, "Message", message)
}

func (c Cond) GetMessage(obj interface{}) string {
	cond := findOrNotCreateCond(obj, string(c))
	if cond == nil {
		return ""
	}
	return getFieldValue(*cond, "Message").String()
}

func touchTS(value reflect.Value) {
	now := time.Now().UTC().Format(time.RFC3339)
	getFieldValue(value, "LastUpdateTime").SetString(now)
}

func getStatus(obj interface{}, condName string) string {
	cond := findOrNotCreateCond(obj, condName)
	if cond == nil {
		return ""
	}
	return getFieldValue(*cond, "Status").String()
}

func setTS(obj interface{}, condName, ts string) {
	cond := findOrCreateCond(obj, condName)
	getFieldValue(cond, "LastUpdateTime").SetString(ts)
}

func getTS(obj interface{}, condName string) string {
	cond := findOrNotCreateCond(obj, condName)
	if cond == nil {
		return ""
	}
	return getFieldValue(*cond, "LastUpdateTime").String()
}

func setStatus(obj interface{}, condName, status string) {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		panic("obj passed must be a pointer")
	}
	cond := findOrCreateCond(obj, condName)
	setValue(cond, "Status", status)
}

func setValue(cond reflect.Value, fieldName, newValue string) {
	value := getFieldValue(cond, fieldName)
	if value.String() != newValue {
		value.SetString(newValue)
		touchTS(cond)
	}
}

func findOrNotCreateCond(obj interface{}, condName string) *reflect.Value {
	condSlice := getValue(obj, "Status", "Conditions")
	if !condSlice.IsValid() {
		condSlice = getValue(obj, "Conditions")
	}
	return findCond(obj, condSlice, condName)
}

func findOrCreateCond(obj interface{}, condName string) reflect.Value {
	condSlice := getValue(obj, "Status", "Conditions")
	if !condSlice.IsValid() {
		condSlice = getValue(obj, "Conditions")
	}
	cond := findCond(obj, condSlice, condName)
	if cond != nil {
		return *cond
	}

	newCond := reflect.New(condSlice.Type().Elem()).Elem()
	newCond.FieldByName("Type").SetString(condName)
	newCond.FieldByName("Status").SetString("Unknown")
	condSlice.Set(reflect.Append(condSlice, newCond))
	return *findCond(obj, condSlice, condName)
}

func findCond(obj interface{}, val reflect.Value, name string) *reflect.Value {
	defer func() {
		if recover() != nil {
			logrus.Fatalf("failed to find .Status.Conditions field on %v", reflect.TypeOf(obj))
		}
	}()

	for i := 0; i < val.Len(); i++ {
		cond := val.Index(i)
		typeVal := getFieldValue(cond, "Type")
		if typeVal.String() == name {
			return &cond
		}
	}

	return nil
}

func getValue(obj interface{}, name ...string) reflect.Value {
	if obj == nil {
		return reflect.Value{}
	}
	v := reflect.ValueOf(obj)
	t := v.Type()
	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = v.Type()
	}

	field := v.FieldByName(name[0])
	if len(name) == 1 {
		return field
	}
	return getFieldValue(field, name[1:]...)
}

func getFieldValue(v reflect.Value, name ...string) reflect.Value {
	if !v.IsValid() {
		return v
	}
	field := v.FieldByName(name[0])
	if len(name) == 1 {
		return field
	}
	return getFieldValue(field, name[1:]...)
}

func Error(reason string, err error) error {
	return &conditionError{
		reason:  reason,
		message: err.Error(),
	}
}

type conditionError struct {
	reason  string
	message string
}

func (e *conditionError) Error() string {
	return e.message
}
