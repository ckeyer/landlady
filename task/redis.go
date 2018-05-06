package task

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gomodule/redigo/redis"
)

var (
	ErrNotExists = errors.New("not found")
)

const (
	rCmdExists = "EXISTS"
	rCmdSet    = "SET"
	rCmdSetnx  = "SETNX"
	rCmdSetex  = "SETEX"
	rCmdGet    = "GET"
	rCmdMget   = "MGET"
	rCmdKeys   = "KEYS"
	rCmdDel    = "DEL"
)

// setKV,
//  exists:
//    nil:   SET
//    true:  SETEX
//    false: SETNX
func (r *TaskController) setKV(key string, v interface{}, exists ...bool) error {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(v); err != nil {
		return err
	}

	cmd := rCmdSet
	if len(exists) == 1 {
		if exists[0] {
			cmd = rCmdSetex
		} else {
			cmd = rCmdSetnx
		}
	}

	n, err := redis.Int(r.Do(cmd, key, buf.String()))
	if err != nil {
		return err
	}
	if n == 0 {
		return fmt.Errorf("insert 0 keys.")
	}

	return nil
}

// getKV
func (r *TaskController) getKV(key string, v interface{}) error {
	data, err := redis.Bytes(r.Do(rCmdGet, key))
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}

func (r *TaskController) existsKey(key string) (bool, error) {
	n, err := redis.Int(r.Do(rCmdExists, key))
	if err != nil {
		return false, err
	}
	if n == 0 {
		return false, nil
	}
	return true, nil
}

func (r *TaskController) deleteKey(key string) error {
	n, err := redis.Int(r.Do(rCmdDel, key))
	if err != nil {
		return err
	}
	if n == 0 {
		return ErrNotExists
	}
	return nil
}

// keys
func (r *TaskController) listKeys(query string) ([]interface{}, error) {
	ret := []interface{}{}
	ks, err := redis.Strings(r.Do(rCmdKeys, query))
	if err != nil {
		return nil, err
	}
	for _, v := range ks {
		ret = append(ret, v)
	}
	return ret, nil
}

func (r *TaskController) listKVs(query string, v interface{}) error {
	ks, err := r.listKeys(query)
	if err != nil {
		return err
	}

	vstrs, err := redis.Strings(r.Do(rCmdMget, ks...))
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.WriteString("[")
	buf.WriteString(strings.Join(vstrs, ","))
	buf.WriteString("]")

	return json.NewDecoder(buf).Decode(v)
}
