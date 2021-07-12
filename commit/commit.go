package commit

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path"

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

		h := sha1.New()
		h.Write([]byte(content))
		bs := h.Sum(nil)

		oid := hex.EncodeToString(bs)

		if err := Store(oid, content, db_path); err != nil {
			log.Println(err)
		}

		// close file
		f.Close()
	}
}

func Store(oid, content, db_path string) error {

	object_path := path.Join(db_path, oid[0:2])

	if _, err := os.Stat(object_path); os.IsNotExist(err) {
		if err := os.Mkdir(object_path, 0755); err != nil {
			return err
		}
	}

	// TODO

	return nil
}
