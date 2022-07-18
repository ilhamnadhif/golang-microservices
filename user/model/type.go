package model

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	TimestampFormat = "20060102150405"
	DateTimeFormat  = "2006-01-02 15:04:05"
	DateFormat      = "2006-01-02"
	TimeFormat      = "15:04:05"
)

type (
	NullString struct {
		sql.NullString
	}
	NullInt struct {
		sql.NullInt64
	}
	NullFloat struct {
		sql.NullFloat64
	}
	NullBool struct {
		sql.NullBool
	}
	NullTime struct {
		sql.NullTime
		Format string
	}
	NullDate struct {
		NullTime
	}
	NullHour struct {
		NullTime
	}
	Time struct {
		time.Time
		Format string
	}
	Date struct {
		Time
	}
	Hour struct {
		Time
	}
	Int struct {
		sql.NullInt64
	}
	IntStr struct {
		Int
	}
	NullIntStr struct {
		NullInt
	}
	String      string
	StringSlice struct {
		Slice []string
	}
	Bool bool
)

func ToNullString(str string) NullString {
	return NullString{sql.NullString{String: str, Valid: true}}
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.NullString.String)
}

func (ns *NullString) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &ns.NullString.String)
	if err == nil {
		ns.NullString.String = strings.TrimSpace(ns.NullString.String)
		ns.NullString.Valid = len(ns.NullString.String) > 0
	}
	return err
}

func (ns *NullString) UnmarshalParam(src string) error {
	if len(src) < 1 {
		return nil
	}
	return ns.UnmarshalJSON([]byte(src))
}

func (sl *StringSlice) UnmarshalParam(src string) error {
	if len(src) < 1 {
		return nil
	}
	var pl []string
	for _, p := range strings.Split(src, ",") {
		p = strings.TrimSpace(p)
		pl = append(pl, p)
	}
	sl.Slice = pl
	return nil
}

func ToNullInt(i int) NullInt {
	return NullInt{sql.NullInt64{Int64: int64(i), Valid: true}}
}

func (ni NullInt) Int() int {
	return int(ni.Int64)
}

func (ni NullInt) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

func (ni *NullInt) UnmarshalJSON(data []byte) error {
	strData := string(data)
	if strData[:1] == "\"" {
		strData = strData[1 : len(strData)-1]
	}
	if strData == "" || strData == "null" {
		return nil
	}
	ii, err := strconv.Atoi(strData)
	if err != nil {
		return err
	}
	ni.NullInt64 = sql.NullInt64{Int64: int64(ii), Valid: true}
	return nil
}

func (ni *NullInt) UnmarshalParam(src string) error {
	if len(src) < 1 {
		return nil
	}
	return ni.UnmarshalJSON([]byte(src))
}

func ToNullFloat(f float64) NullFloat {
	return NullFloat{sql.NullFloat64{Float64: f, Valid: true}}
}

func (t NullFloat) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(t.Float64)
}

func (nf *NullFloat) UnmarshalJSON(data []byte) error {
	strData := string(data)
	if strData[:1] == "\"" {
		strData = strData[1 : len(strData)-1]
	}
	if strData == "" || strData == "null" {
		return nil
	}
	ii, err := strconv.ParseFloat(strData, 64)
	if err != nil {
		return err
	}
	nf.NullFloat64 = sql.NullFloat64{Float64: ii, Valid: true}
	return nil
}

func (nf *NullFloat) UnmarshalParam(src string) error {
	if len(src) < 1 {
		return nil
	}
	return nf.UnmarshalJSON([]byte(src))
}

func ToNullTime(t time.Time) NullTime {
	return NullTime{NullTime: sql.NullTime{Time: t, Valid: !t.IsZero()}}
}

func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil // note: change to nil? trace usage first
	}
	format := nt.Format
	if format == "" {
		format = DateTimeFormat
	}
	return json.Marshal(nt.Time.Format(format))
}

func (nt *NullTime) UnmarshalJSON(data []byte) error {
	strData := string(data)
	if strData[:1] == "\"" {
		strData = strData[1 : len(strData)-1]
	}
	if strData == "" || strData == "null" {
		return nil
	}
	var err error
	format := nt.Format
	if format == "" {
		format = DateTimeFormat
	}
	nt.Time, err = time.Parse(format, strData)
	nt.Valid = err == nil
	return err
}

func (nt *NullTime) UnmarshalParam(src string) error {
	if len(src) < 1 {
		return nil
	}
	return nt.UnmarshalJSON([]byte(src))
}

func (nt *NullTime) Scan(value interface{}) error {
	switch value.(type) {
	default:
		return nt.NullTime.Scan(value)
	case []byte:
		b := value.([]byte)
		s := strings.TrimSpace(string(b))
		t, err := time.Parse(nt.Format, s)
		if err != nil {
			log.Println("NullTime parse error")
			return err
		}
		return nt.NullTime.Scan(t)
	}
}

func (nt NullTime) String() string {
	return nt.Time.Format(DateTimeFormat)
}

func ToNullDate(t time.Time) NullDate {
	dt := NullDate{NullTime: ToNullTime(t)}
	dt.Format = DateFormat
	return dt
}

func (nd *NullDate) UnmarshalJSON(data []byte) error {
	nd.Format = DateFormat
	return nd.NullTime.UnmarshalJSON(data)
}

func (nd *NullDate) UnmarshalParam(src string) error {
	if len(src) < 1 {
		return nil
	}
	nd.Format = DateFormat
	return nd.NullTime.UnmarshalJSON([]byte(src))
}

func (nd NullDate) String() string {
	return nd.NullTime.Time.Format(DateFormat)
}

func (nd *NullDate) Scan(value interface{}) error {
	nd.Format = DateFormat
	return nd.NullTime.Scan(value)
}

func ToNullHour(t time.Time) NullHour {
	dt := NullHour{NullTime: ToNullTime(t)}
	dt.Format = TimeFormat
	return dt
}

func (nh *NullHour) UnmarshalJSON(data []byte) error {
	nh.Format = TimeFormat
	return nh.NullTime.UnmarshalJSON(data)
}

func (nh *NullHour) UnmarshalParam(src string) error {
	if len(src) < 1 {
		return nil
	}
	nh.Format = TimeFormat
	return nh.NullTime.UnmarshalJSON([]byte(src))
}

func (nh NullHour) String() string {
	return nh.NullTime.Time.Format(TimeFormat)
}

func (nh *NullHour) Scan(value interface{}) error {
	nh.Format = TimeFormat
	return nh.NullTime.Scan(value)
}

func ToTime(t time.Time) Time {
	return Time{Time: t, Format: DateTimeFormat}
}

func (ti Time) MarshalJSON() ([]byte, error) {
	if ti.IsZero() {
		return []byte("null"), nil // note: change to nil? trace usage first
	}
	format := ti.Format
	if format == "" {
		format = DateTimeFormat
	}
	return json.Marshal(ti.Time.Format(format))
}

func (ti *Time) UnmarshalJSON(data []byte) error {
	strData := string(data)
	if strData[:1] == "\"" {
		strData = strData[1 : len(strData)-1]
	}
	if strData == "" || strData == "null" {
		return nil
	}
	var err error
	format := ti.Format
	if format == "" {
		format = DateTimeFormat
	}
	ti.Time, err = time.Parse(format, strData)
	return err
}

func (ti *Time) UnmarshalParam(src string) error {
	if len(src) < 1 {
		return nil
	}
	return ti.UnmarshalJSON([]byte(src))
}

func (ti Time) String() string {
	return ti.Time.String()
}

func (ti Time) Value() (driver.Value, error) {
	return ti.Time, nil
}

func (ti *Time) Scan(value interface{}) error {
	switch value.(type) {
	default:
		t, isT := value.(time.Time)
		if !isT {
			err := errors.New("Time unsupported scan")
			log.Println(err, value)
			return err
		}
		ti.Time = t
		return nil
	case []byte:
		b := value.([]byte)
		s := strings.TrimSpace(string(b))
		t, err := time.Parse(ti.Format, s)
		if err != nil {
			log.Println("Time parse error")
			return err
		}
		ti.Time = t
		return nil
	}
}

func ToDate(t time.Time) Date {
	dt := Date{Time: ToTime(t)}
	dt.Format = DateFormat
	return dt
}

func (dt *Date) UnmarshalJSON(data []byte) error {
	dt.Format = DateFormat
	return dt.Time.UnmarshalJSON(data)
}

func (dt *Date) UnmarshalParam(src string) error {
	if len(src) < 1 {
		return nil
	}
	dt.Format = DateFormat
	return dt.Time.UnmarshalJSON([]byte(src))
}

func (dt Date) String() string {
	return dt.Time.Time.Format(DateFormat)
}

func (dt *Date) Scan(value interface{}) error {
	dt.Format = DateFormat
	return dt.Time.Scan(value)
}

func ToHour(t time.Time) Hour {
	hr := Hour{Time: ToTime(t)}
	hr.Format = TimeFormat
	return hr
}

func (hr *Hour) UnmarshalJSON(data []byte) error {
	hr.Format = TimeFormat
	return hr.Time.UnmarshalJSON(data)
}

func (hr *Hour) UnmarshalParam(src string) error {
	if len(src) < 1 {
		return nil
	}
	hr.Format = TimeFormat
	return hr.Time.UnmarshalJSON([]byte(src))
}

func (hr Hour) String() string {
	return hr.Time.Time.Format(TimeFormat)
}

func (hr *Hour) Scan(value interface{}) error {
	hr.Format = TimeFormat
	return hr.Time.Scan(value)
}

func ToNullBool(b bool) NullBool {
	return NullBool{sql.NullBool{Bool: b, Valid: true}}
}

func (nb NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nb.Bool)
}

func (nb *NullBool) UnmarshalJSON(data []byte) error {
	strData := string(data)
	if strData[:1] == "\"" {
		strData = strData[1 : len(strData)-1]
	}
	if strData == "" || strData == "null" {
		return nil
	}
	b, err := strconv.ParseBool(strData)
	if err != nil {
		return err
	}
	nb.NullBool = sql.NullBool{Bool: b, Valid: true}
	return nil
}

func ToInt(i int) Int {
	return Int{sql.NullInt64{Int64: int64(i), Valid: true}}
}

func (i Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Int64)
}

func (i *Int) UnmarshalJSON(data []byte) error {
	strData := string(data)
	if strData[:1] == "\"" {
		strData = strData[1 : len(strData)-1]
	}
	if strData == "" || strData == "null" {
		i.NullInt64 = sql.NullInt64{Int64: 0, Valid: true}
		return nil
	}
	ii, err := strconv.Atoi(strData)
	if err != nil {
		return err
	}
	i.NullInt64 = sql.NullInt64{Int64: int64(ii), Valid: true}
	return nil
}

func (i *Int) UnmarshalParam(src string) error {
	if len(src) < 1 {
		return nil
	}
	return i.UnmarshalJSON([]byte(src))
}

func (i *Int) Scan(value interface{}) error {
	err := i.NullInt64.Scan(value)
	if err != nil {
		return err
	}
	i.NullInt64.Valid = true
	return nil
}

func (i Int) Value() (driver.Value, error) {
	i.Valid = true
	return i.NullInt64.Value()
}

func (i Int) Int() int {
	return int(i.Int64)
}

func ToIntStr(i int) IntStr {
	return IntStr{Int: ToInt(i)}
}

func (is IntStr) MarshalJSON() ([]byte, error) {
	return json.Marshal(is.Int64)
}

func (is *IntStr) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if value == nil || !ok {
		log.Println("IntStr unsupported scan", value)
		return nil
	}
	s := strings.TrimSpace(string(b))
	if len(s) == 0 {
		is.Int.Valid = true
		return nil
	}
	return is.Int.Scan(s)
}

func (is IntStr) Value() (driver.Value, error) {
	is.Valid = true
	return is.String(), nil
}

func (is IntStr) String() string {
	return strconv.Itoa(int(is.Int64))
}

func ToNullIntStr(i int) NullIntStr {
	return NullIntStr{NullInt: ToNullInt(i)}
}

func (nis NullIntStr) MarshalJSON() ([]byte, error) {
	return json.Marshal(nis.Int64)
}

func (nis *NullIntStr) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if value == nil || !ok {
		log.Println("NullIntStr unsupported scan", value)
		return nil
	}
	s := strings.TrimSpace(string(b))
	if len(s) == 0 {
		nis.NullInt.Valid = true
		return nil
	}
	return nis.NullInt.Scan(s)
}

func (nis NullIntStr) Value() (driver.Value, error) {
	if !nis.Valid {
		return nis.NullInt64.Value()
	}
	return nis.String(), nil
}

func (nis NullIntStr) String() string {
	return strconv.Itoa(int(nis.Int64))
}

func (b Bool) MarshalJSON() ([]byte, error) {
	return json.Marshal(b)
}

func (b *Bool) UnmarshalJSON(data []byte) error {
	strData := string(data)

	if strData == "" {
		*b = false
		return nil
	}
	if strData[0] == '"' {
		strData = strData[1 : len(strData)-1]
	}
	strData = strings.ToLower(strData)
	switch strData {
	case "0":
		*b = false
		return nil
	case "1":
		*b = true
		return nil
	default:
		val, err := strconv.ParseBool(strData)
		*b = Bool(val)
		return err
	}
}

func (b Bool) Bool() bool {
	return bool(b)
}

func ValidateCustomType(field reflect.Value) interface{} {
	switch fld := field.Interface().(type) {
	case Int:
		return fld.Int()
	case NullInt:
		return fld.Int()
	case Time:
		return fld.Time
	case NullTime:
		return fld.Time
	case Date:
		return fld.Time
	case NullDate:
		return fld.Time
	case Hour:
		return fld.Time
	case NullHour:
		return fld.Time
	case Bool:
		return fld.Bool()
	case NullFloat:
		return fld.Float64
	default:
		panic(fmt.Sprintf("trying to validate unhandled custom type: %v", fld))
	}
}
