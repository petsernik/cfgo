package io

import (
	"bufio"
	"os"
	"reflect"
	"strconv"
	"unicode"
)

// IO wraps buffered reader and writer
type IO struct {
	reader *bufio.Reader
	writer *bufio.Writer
}

// NewIO creates a new IO instance
func NewIO() *IO {
	return &IO{
		reader: bufio.NewReaderSize(os.Stdin, 1<<20),  // 1MB buffer
		writer: bufio.NewWriterSize(os.Stdout, 1<<20), // 1MB buffer
	}
}

// Read reads the next token and writes it into pointer v
func (io *IO) Read(a ...any) {
	for _, v := range a {
		rv := reflect.ValueOf(v)
		if rv.Kind() != reflect.Pointer {
			panic("Read: argument must be a pointer")
		}
		rv = rv.Elem()

		switch rv.Kind() {
		case reflect.Int:
			x, err := strconv.Atoi(io.nextToken())
			if err != nil {
				panic(err)
			}
			rv.SetInt(int64(x))
		case reflect.Int64:
			x, err := strconv.ParseInt(io.nextToken(), 10, 64)
			if err != nil {
				panic(err)
			}
			rv.SetInt(x)
		case reflect.Float64:
			x, err := strconv.ParseFloat(io.nextToken(), 64)
			if err != nil {
				panic(err)
			}
			rv.SetFloat(x)
		case reflect.String:
			rv.SetString(io.nextToken())
		case reflect.Bool:
			x, err := strconv.ParseBool(io.nextToken())
			if err != nil {
				panic(err)
			}
			rv.SetBool(x)
		case reflect.Slice, reflect.Array:
			for i := range rv.Len() {
				elemPtr := reflect.New(rv.Type().Elem())
				io.Read(elemPtr.Interface())
				rv.Index(i).Set(elemPtr.Elem())
			}
		default:
			panic("Read: unsupported type " + rv.Kind().String())
		}
	}
}

// Print writes one or more values without newline
func (io *IO) Print(a ...any) {
	length := len(a)
	for i, v := range a {
		io.writeValue(v)
		if i < length-1 {
			io.writer.WriteByte(' ')
		}
	}
}

// Println writes one or more values followed by newline
func (io *IO) Println(a ...any) {
	io.Print(a...)
	io.writer.WriteByte('\n')
}

// Flush flushes the buffered writer
func (io *IO) Flush() {
	io.writer.Flush()
}

// Helpers
func (io *IO) nextToken() string {
	var token []byte
LOOP:
	for {
		b, err := io.reader.ReadByte()
		if err != nil {
			break LOOP
		}
		switch {
		case unicode.IsSpace(rune(b)):
			if len(token) > 0 {
				break LOOP
			}
		default:
			token = append(token, b)
		}
	}
	return string(token)
}

func (io *IO) writeValue(v any) {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		io.writer.WriteString(strconv.FormatInt(rv.Int(), 10))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		io.writer.WriteString(strconv.FormatUint(rv.Uint(), 10))
	case reflect.Float32, reflect.Float64:
		io.writer.WriteString(strconv.FormatFloat(rv.Float(), 'f', -1, 64))
	case reflect.Bool: // specific for competitive programming: YES or NO
		if rv.Bool() {
			io.writer.WriteString("YES")
		} else {
			io.writer.WriteString("NO")
		}
	case reflect.String:
		io.writer.WriteString(rv.String())
	case reflect.Slice, reflect.Array:
		for i := range rv.Len() {
			io.writeValue(rv.Index(i).Interface())
			switch rv.Index(i).Kind() {
			case reflect.Slice, reflect.Array:
				io.writer.WriteString("\n")
			default:
				if i != rv.Len()-1 {
					io.writer.WriteString(" ")
				}
			}
		}
	default:
		io.writer.WriteString("<this value unsupported for print>")
	}
}
