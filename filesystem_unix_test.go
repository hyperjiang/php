//go:build aix || darwin || dragonfly || freebsd || (js && wasm) || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd js,wasm linux netbsd openbsd solaris

package php_test

import (
	"os"

	"github.com/hyperjiang/php"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Filesystem Functions", func() {
	It("Dirname", func() {
		tests := []struct {
			input string
			want  string
		}{
			{"/etc/passwd", "/etc"},
			{"/etc", "/"},
			{".", "."},
		}
		for _, t := range tests {
			Expect(php.Dirname(t.input)).To(Equal(t.want))
		}
	})

	It("DirnameWithLevels", func() {
		type args struct {
			path   string
			levels int
		}
		tests := []struct {
			args args
			want string
		}{
			{args{"/usr/local/lib", 2}, "/usr"},
			{args{"/usr/local/lib", 0}, "/usr/local/lib"},
		}
		for _, t := range tests {
			Expect(php.DirnameWithLevels(t.args.path, t.args.levels)).To(Equal(t.want))
		}
	})

	It("Realpath", func() {
		dir, _ := os.Getwd()
		tests := []struct {
			input string
			want  string
		}{
			{"/etc/", "/etc"},
			{"/etc/..", "/"},
			{".", dir},
		}
		for _, t := range tests {
			Expect(php.Realpath(t.input)).To(Equal(t.want))
		}
	})

	It("Touch and Unlink", func() {
		var filename = "/tmp/hyperjiangphpfilesystemtest.txt"

		err := php.Touch(filename)
		Expect(err).NotTo(HaveOccurred())

		err = php.Unlink(filename)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Mkdir and Rmdir", func() {
		var path = "/tmp/hyperjiangphpfilesystemtestdir"
		var subPath = "/tmp/hyperjiangphpfilesystemtestdir/dir"

		It("no recursive", func() {
			err := php.Mkdir(path, 0666, false)
			Expect(err).NotTo(HaveOccurred())

			// clean up
			err = php.Rmdir(path)
			Expect(err).NotTo(HaveOccurred())
		})

		It("recursive", func() {
			err := php.Mkdir(subPath, 0, true)
			Expect(err).NotTo(HaveOccurred())

			// clean up
			err = php.Rmdir(subPath)
			Expect(err).NotTo(HaveOccurred())
			err = php.Rmdir(path)
			Expect(err).NotTo(HaveOccurred())
		})
	})

	It("IsFile", func() {
		var filename = "/tmp/hyperjiangphpfilesystemtest.txt"

		if php.IsFile(filename) {
			err := php.Unlink(filename)
			Expect(err).NotTo(HaveOccurred())
		}
		Expect(php.IsFile(filename)).To(BeFalse()) // file not exists

		err := php.Touch(filename)
		Expect(err).NotTo(HaveOccurred())
		Expect(php.IsFile(filename)).To(BeTrue())
		err = php.Unlink(filename)
		Expect(err).NotTo(HaveOccurred())
	})

	It("IsDir", func() {
		var pathname = "/tmp/hyperjiangphpfilesystemtestdir"

		if php.IsDir(pathname) {
			err := php.Rmdir(pathname)
			Expect(err).NotTo(HaveOccurred())
		}
		Expect(php.IsDir(pathname)).To(BeFalse()) // dir not exists

		err := php.Mkdir(pathname, 0, false)
		Expect(err).NotTo(HaveOccurred())
		Expect(php.IsDir(pathname)).To(BeTrue())
		err = php.Rmdir(pathname)
		Expect(err).NotTo(HaveOccurred())
	})

	It("IsExecutable", func() {
		var filename = "/tmp/hyperjiangphpfilesystemtest.exe"

		if php.IsExecutable(filename) {
			err := php.Unlink(filename)
			Expect(err).NotTo(HaveOccurred())
		}
		Expect(php.IsExecutable(filename)).To(BeFalse()) // file not exists

		err := php.Touch(filename)
		Expect(err).NotTo(HaveOccurred())
		Expect(php.IsExecutable(filename)).To(BeFalse()) // not executable

		err = php.Chmod(filename, 0777)
		Expect(err).NotTo(HaveOccurred())
		Expect(php.IsExecutable(filename)).To(BeTrue())
		err = php.Unlink(filename)
		Expect(err).NotTo(HaveOccurred())
	})

	It("IsReadable", func() {
		var filename = "/tmp/hyperjiangphpfilesystemtest.txt"

		if php.IsReadable(filename) {
			err := php.Unlink(filename)
			Expect(err).NotTo(HaveOccurred())
		}
		Expect(php.IsReadable(filename)).To(BeFalse()) // file not exists

		err := php.Touch(filename)
		Expect(err).NotTo(HaveOccurred())
		Expect(php.IsReadable(filename)).To(BeTrue())

		err = php.Chmod(filename, 0222)
		Expect(err).NotTo(HaveOccurred())
		Expect(php.IsReadable(filename)).To(BeFalse()) // not readable
		err = php.Unlink(filename)
		Expect(err).NotTo(HaveOccurred())
	})

	It("IsWritable", func() {
		var filename = "/tmp/hyperjiangphpfilesystemtest.txt"

		if php.IsWritable(filename) {
			err := php.Unlink(filename)
			Expect(err).NotTo(HaveOccurred())
		}
		Expect(php.IsWritable(filename)).To(BeFalse()) // file not exists

		err := php.Touch(filename)
		Expect(err).NotTo(HaveOccurred())
		Expect(php.IsWritable(filename)).To(BeTrue())

		err = php.Chmod(filename, 0555)
		Expect(err).NotTo(HaveOccurred())
		Expect(php.IsWritable(filename)).To(BeFalse()) // not writable
		err = php.Unlink(filename)
		Expect(err).NotTo(HaveOccurred())
	})

	It("Symlink and IsLink", func() {
		var filename = "/tmp/hyperjiangphpfilesystemtest.txt"

		err := php.Touch(filename)
		Expect(err).NotTo(HaveOccurred())

		var link = "/tmp/hyperjiangphpfilesystemtest.link"

		if php.IsLink(link) {
			err := php.Unlink(link)
			Expect(err).NotTo(HaveOccurred())
		}
		Expect(php.IsLink(link)).To(BeFalse()) // file not exists

		err = php.Symlink(filename, link)
		Expect(err).NotTo(HaveOccurred())
		err = php.Unlink(filename)
		Expect(err).NotTo(HaveOccurred())

		Expect(php.IsLink(link)).To(BeTrue())
		err = php.Unlink(link)
		Expect(err).NotTo(HaveOccurred())
	})

	It("Basename", func() {
		tests := []struct {
			input string
			want  string
		}{
			{"/etc/sudoers.d", "sudoers.d"},
			{"/etc/", "etc"},
			{".", "."},
			{"/", "/"},
		}
		for _, t := range tests {
			Expect(php.Basename(t.input)).To(Equal(t.want))
		}
	})

	It("Chown", func() {
		var filename = "/tmp/tmpfile"
		err := php.Touch(filename)
		Expect(err).NotTo(HaveOccurred())

		err = php.Chown(filename, os.Getuid(), os.Getegid())
		Expect(err).NotTo(HaveOccurred())

		err = php.Unlink(filename)
		Expect(err).NotTo(HaveOccurred())
	})

	It("Link", func() {
		var filename = "/tmp/hyperjiangphpfilesystemtest.txt"
		err := php.Touch(filename)
		Expect(err).NotTo(HaveOccurred())

		var link = "/tmp/hyperjiangphpfilesystemtest.link"

		if php.IsLink(link) {
			err := php.Unlink(link)
			Expect(err).NotTo(HaveOccurred())
		}

		err = php.Link(filename, link)
		Expect(err).NotTo(HaveOccurred())
		err = php.Unlink(link)
		Expect(err).NotTo(HaveOccurred())

		err = php.Unlink(filename)
		Expect(err).NotTo(HaveOccurred())
	})

	It("Copy", func() {
		var src = "/tmp/hyperjiangphpfilesystemtest.src"
		var dst = "/tmp/hyperjiangphpfilesystemtest.dst"

		if php.FileExists(src) {
			err := php.Unlink(src)
			Expect(err).NotTo(HaveOccurred())
		}

		if php.FileExists(dst) {
			err := php.Unlink(dst)
			Expect(err).NotTo(HaveOccurred())
		}

		err := php.Copy(src, dst)
		Expect(err).To(HaveOccurred(), "Copy should fail because src file does not exist")

		err = php.Touch(src)
		Expect(err).NotTo(HaveOccurred())

		err = php.Copy(src, "/tmp/nodir/hyperjiangphpfilesystemtest.dst")
		Expect(err).To(HaveOccurred(), "Copy should fail because the directory of dst file does not exist")

		err = php.Copy(src, dst)
		Expect(err).NotTo(HaveOccurred())

		err = php.Unlink(dst)
		Expect(err).NotTo(HaveOccurred())

		err = php.Unlink(src)
		Expect(err).NotTo(HaveOccurred())
	})

	It("FileExists", func() {
		var filename = "/tmp/hyperjiangphpfilesystemtest.txt"

		if php.FileExists(filename) {
			err := php.Unlink(filename)
			Expect(err).NotTo(HaveOccurred())
		}
		Expect(php.FileExists(filename)).To(BeFalse()) // file not exists

		err := php.Touch(filename)
		Expect(err).NotTo(HaveOccurred())
		Expect(php.FileExists(filename)).To(BeTrue())
		err = php.Unlink(filename)
		Expect(err).NotTo(HaveOccurred())
	})

	It("FileSize", func() {
		var filename = "/tmp/hyperjiangphpfilesystemtest.txt"

		if php.FileExists(filename) {
			err := php.Unlink(filename)
			Expect(err).NotTo(HaveOccurred())
		}
		_, err := php.FileSize(filename)
		Expect(err).To(HaveOccurred())

		err = php.Touch(filename)
		Expect(err).NotTo(HaveOccurred())
		size, err := php.FileSize(filename)
		Expect(err).NotTo(HaveOccurred())
		Expect(size).To(BeZero())

		err = php.Unlink(filename)
		Expect(err).NotTo(HaveOccurred())
	})

	It("Rename", func() {
		var oldname = "/tmp/hyperjiangphpfilesystemtest.old"
		err := php.Touch(oldname)
		Expect(err).NotTo(HaveOccurred())

		var newname = "/tmp/hyperjiangphpfilesystemtest.new"
		err = php.Rename(oldname, newname)
		Expect(err).NotTo(HaveOccurred())
		Expect(php.FileExists(oldname)).To(BeFalse())
		Expect(php.FileExists(newname)).To(BeTrue())

		err = php.Unlink(newname)
		Expect(err).NotTo(HaveOccurred())
	})

	It("FilePutContents and FileGetContents", func() {
		var filename = "/tmp/hyperjiangphpfilesystemtest.txt"

		if php.FileExists(filename) {
			err := php.Unlink(filename)
			Expect(err).NotTo(HaveOccurred())
		}

		const msg = "Hello world!"
		err := php.FilePutContents(filename, msg)
		Expect(err).NotTo(HaveOccurred())

		str, err := php.FileGetContents(filename)
		Expect(err).NotTo(HaveOccurred())
		Expect(str).To(Equal(msg))

		err = php.Unlink(filename)
		Expect(err).NotTo(HaveOccurred())
	})
})
