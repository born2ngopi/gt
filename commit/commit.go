package commit

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path"

	"gt/common"
	"gt/workspace"
)

func Commit() {
	// get folder localtion
	root_path, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}

	git_path := path.Join(root_path, ".git")
	db_path := path.Join(git_path, "objects")
	_ = db_path

	// get list folder inside root path
	wrkSpc := workspace.Init(root_path).ListDir()
	if wrkSpc == nil {
		log.Fatal("cannot find list files")
	}

	for _, file := range wrkSpc.Files {
		f, err := wrkSpc.ReadFile(file)
		if err != nil {
			log.Printf("cannot read file : %v", err)
			return
		}

		buf := bytes.NewBuffer(nil)
		io.Copy(buf, f)

		content := fmt.Sprintf("%s %d %s", wrkSpc.GetFileType(f), len(buf.Bytes()), buf.String())

		oid := GetOid(content)

		if err := Store(oid, content, db_path); err != nil {
			log.Println(err)
		}

		// close file
		f.Close()
	}
}

func GetOid(content string) string {
	h := sha1.New()
	h.Write([]byte(content))
	bs := h.Sum(nil)

	return hex.EncodeToString(bs)
}

func Store(oid, content, db_path string) error {

	object_path := path.Join(db_path, oid[0:2])
	if _, err := os.Stat(object_path); os.IsNotExist(err) {
		if err := os.Mkdir(object_path, 0755); err != nil {
			return err
		}
	}

	temp_path := path.Join(object_path, common.GetRandString())
	if _, err := os.Stat(temp_path); os.IsNotExist(err) {

		f, err := os.Create(temp_path)
		if err != nil {
			return err
		}
		defer f.Close()

		w := bufio.NewWriter(f)
		if _, err := w.WriteString(content); err != nil {
			return err
		}

		if err := w.Flush(); err != nil {
			return err
		}

	} else {

		f, err := os.Open(temp_path)
		if err != nil {
			return err
		}
		defer f.Close()

		w := bufio.NewWriter(f)
		if _, err := w.WriteString(content); err != nil {
			return err
		}

		if err := w.Flush(); err != nil {
			return err
		}

	}

	return nil
}
