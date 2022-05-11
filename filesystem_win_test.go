//go:build windows
// +build windows

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
			{"C:\\etc\\passwd", "C:\\etc"},
			{"C:\\etc", "C:\\"},
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
			{args{"C:\\usr\\local\\lib", 2}, "C:\\usr"},
			{args{"C:\\usr\\local\\lib", 0}, "C:\\usr\\local\\lib"},
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
			{"C:\\etc\\", "C:\\etc"},
			{"C:\\etc\\..", "C:\\"},
			{".", dir},
		}
		for _, t := range tests {
			Expect(php.Realpath(t.input)).To(Equal(t.want))
		}
	})

	It("Touch and Unlink", func() {
		var filename = "C:\\touch-and-unlink-test.txt"

		err := php.Touch(filename)
		Expect(err).NotTo(HaveOccurred())

		err = php.Unlink(filename)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Mkdir and Rmdir", func() {
		var path = "C:\\mkdir-and-rmdir-test"
		var subPath = "C:\\mkdir-and-rmdir-test\\dir"

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
		var filename = "C:\\is-file-test.txt"
		Expect(php.IsFile(filename)).To(BeFalse()) // file not exists

		err := php.Touch(filename)
		Expect(err).NotTo(HaveOccurred())
		Expect(php.IsFile(filename)).To(BeTrue())
		err = php.Unlink(filename)
		Expect(err).NotTo(HaveOccurred())
	})

	It("IsDir", func() {
		var pathname = "C:\\is-dir-test"
		Expect(php.IsDir(pathname)).To(BeFalse()) // dir not exists

		err := php.Mkdir(pathname, 0, false)
		Expect(err).NotTo(HaveOccurred())
		Expect(php.IsDir(pathname)).To(BeTrue())
		err = php.Rmdir(pathname)
		Expect(err).NotTo(HaveOccurred())
	})

	It("IsExecutable", func() {
		var filename = "C:\\is-executable-test.exe"
		Expect(php.IsReadable(filename)).To(BeFalse()) // file not exists

		err := php.Touch(filename)
		Expect(err).NotTo(HaveOccurred())
		Expect(php.IsExecutable(filename)).To(BeTrue())

		err = php.Unlink(filename)
		Expect(err).NotTo(HaveOccurred())
	})

	It("IsReadable", func() {
		var filename = "C:\\is-readable-test.txt"
		Expect(php.IsReadable(filename)).To(BeFalse()) // file not exists

		err := php.Touch(filename)
		Expect(err).NotTo(HaveOccurred())
		Expect(php.IsReadable(filename)).To(BeTrue())

		err = php.Unlink(filename)
		Expect(err).NotTo(HaveOccurred())
	})

	It("IsWritable", func() {
		var filename = "C:\\is-writable-test.txt"
		Expect(php.IsWritable(filename)).To(BeFalse()) // file not exists

		err := php.Touch(filename)
		Expect(err).NotTo(HaveOccurred())
		Expect(php.IsWritable(filename)).To(BeTrue())

		err = php.Chmod(filename, 0555)
		Expect(err).NotTo(HaveOccurred())
		Expect(php.IsWritable(filename)).To(BeFalse()) // not writable
		err = php.Unlink(filename)
		Expect(err).To(HaveOccurred())
	})

	It("Symlink and IsLink", func() {
		var filename = "C:\\symlink-and-islink-test.txt"
		php.Touch(filename)

		var link = "C:\\hyperjiangphpfilesystemtest.link"

		if php.IsLink(link) {
			err := php.Unlink(link)
			Expect(err).NotTo(HaveOccurred())
		}
		Expect(php.IsLink(link)).To(BeFalse()) // file not exists

		err := php.Symlink(filename, link)
		Expect(err).NotTo(HaveOccurred())

		Expect(php.IsLink(link)).To(BeTrue())
		err = php.Unlink(link)
		Expect(err).NotTo(HaveOccurred())

		err = php.Unlink(filename)
		Expect(err).NotTo(HaveOccurred())
	})

	It("Basename", func() {
		tests := []struct {
			input string
			want  string
		}{
			{"C:\\etc\\sudoers.d", "sudoers.d"},
			{"C:\\etc\\", "etc"},
			{".", "."},
			{"C:\\", "\\"},
		}
		for _, t := range tests {
			Expect(php.Basename(t.input)).To(Equal(t.want))
		}
	})

	It("Link", func() {
		var filename = "C:\\link-test.txt"
		err := php.Touch(filename)
		Expect(err).NotTo(HaveOccurred())

		var link = "C:\\link-test.link"
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
		var src = "C:\\copy-test.src"
		var dst = "C:\\copy-test.dst"

		if php.IsLink(src) {
			err := php.Unlink(src)
			Expect(err).NotTo(HaveOccurred())
		}

		if php.IsLink(dst) {
			err := php.Unlink(dst)
			Expect(err).NotTo(HaveOccurred())
		}

		err := php.Copy(src, dst)
		Expect(err).To(HaveOccurred(), "Copy should fail because src file does not exist")

		err = php.Touch(src)
		Expect(err).NotTo(HaveOccurred())

		err = php.Copy(src, "C:\\nodir\\copy-test.dst")
		Expect(err).To(HaveOccurred(), "Copy should fail because the directory of dst file does not exist")

		err = php.Copy(src, dst)
		Expect(err).NotTo(HaveOccurred())

		err = php.Unlink(dst)
		Expect(err).NotTo(HaveOccurred())

		err = php.Unlink(src)
		Expect(err).NotTo(HaveOccurred())
	})

	It("FileExists", func() {
		var filename = "C:\\file-exists-test.txt"
		Expect(php.FileExists(filename)).To(BeFalse()) // file not exists

		err := php.Touch(filename)
		Expect(err).NotTo(HaveOccurred())
		Expect(php.FileExists(filename)).To(BeTrue())
		err = php.Unlink(filename)
		Expect(err).NotTo(HaveOccurred())
	})

	It("FileSize", func() {
		var filename = "C:\\file-size-test.txt"
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
		var oldname = "C:\\rename-test.old"
		err := php.Touch(oldname)
		Expect(err).NotTo(HaveOccurred())

		var newname = "C:\\rename-test.new"
		err = php.Rename(oldname, newname)

		Expect(err).NotTo(HaveOccurred())
		Expect(php.FileExists(oldname)).To(BeFalse())
		Expect(php.FileExists(newname)).To(BeTrue())

		err = php.Unlink(newname)
		Expect(err).NotTo(HaveOccurred())
	})

	It("FilePutContents and FileGetContents", func() {
		var filename = "C:\\file-put-and-get-contents-test.txt"
		if php.IsLink(filename) {
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
