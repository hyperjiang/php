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
		php.Unlink(filename)
		Expect(php.IsFile(filename)).To(BeFalse()) // file not exists

		php.Touch(filename)
		Expect(php.IsFile(filename)).To(BeTrue())
		php.Unlink(filename)
	})

	It("IsDir", func() {
		var pathname = "/tmp/hyperjiangphpfilesystemtestdir"
		php.Rmdir(pathname)
		Expect(php.IsDir(pathname)).To(BeFalse()) // dir not exists

		php.Mkdir(pathname, 0, false)
		Expect(php.IsDir(pathname)).To(BeTrue())
		php.Rmdir(pathname)
	})

	It("IsExecutable", func() {
		var filename = "/tmp/hyperjiangphpfilesystemtest.exe"
		php.Unlink(filename)
		Expect(php.IsExecutable(filename)).To(BeFalse()) // file not exists

		php.Touch(filename)
		Expect(php.IsExecutable(filename)).To(BeFalse()) // not executable

		php.Chmod(filename, 0777)
		Expect(php.IsExecutable(filename)).To(BeTrue())
		php.Unlink(filename)
	})

	It("IsReadable", func() {
		var filename = "/tmp/hyperjiangphpfilesystemtest.txt"
		php.Unlink(filename)
		Expect(php.IsReadable(filename)).To(BeFalse()) // file not exists

		php.Touch(filename)
		Expect(php.IsReadable(filename)).To(BeTrue())

		php.Chmod(filename, 0222)
		Expect(php.IsReadable(filename)).To(BeFalse()) // not readable
		php.Unlink(filename)
	})

	It("IsWritable", func() {
		var filename = "/tmp/hyperjiangphpfilesystemtest.txt"
		php.Unlink(filename)
		Expect(php.IsWritable(filename)).To(BeFalse()) // file not exists

		php.Touch(filename)
		Expect(php.IsWritable(filename)).To(BeTrue())

		php.Chmod(filename, 0555)
		Expect(php.IsWritable(filename)).To(BeFalse()) // not writable
		php.Unlink(filename)
	})

	It("Symlink and IsLink", func() {
		var filename = "/tmp/hyperjiangphpfilesystemtest.txt"
		php.Touch(filename)
		defer php.Unlink(filename)

		var link = "/tmp/hyperjiangphpfilesystemtest.link"
		php.Unlink(link)
		Expect(php.IsLink(link)).To(BeFalse()) // file not exists

		err := php.Symlink(filename, link)
		Expect(err).NotTo(HaveOccurred())

		Expect(php.IsLink(link)).To(BeTrue())
		php.Unlink(link)
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
		php.Touch(filename)
		defer php.Unlink(filename)

		err := php.Chown(filename, os.Getuid(), os.Getegid())
		Expect(err).NotTo(HaveOccurred())
	})

	It("Link", func() {
		var filename = "/tmp/hyperjiangphpfilesystemtest.txt"
		php.Touch(filename)
		defer php.Unlink(filename)

		var link = "/tmp/hyperjiangphpfilesystemtest.link"
		php.Unlink(link)

		err := php.Link(filename, link)
		Expect(err).NotTo(HaveOccurred())
		php.Unlink(link)
	})

	It("Copy", func() {
		var src = "/tmp/hyperjiangphpfilesystemtest.src"
		var dst = "/tmp/hyperjiangphpfilesystemtest.dst"
		php.Unlink(src)
		php.Unlink(dst)

		err := php.Copy(src, dst)
		Expect(err).To(HaveOccurred(), "Copy should fail because src file does not exist")

		php.Touch(src)
		defer php.Unlink(src)

		err = php.Copy(src, "/tmp/nodir/hyperjiangphpfilesystemtest.dst")
		Expect(err).To(HaveOccurred(), "Copy should fail because the directory of dst file does not exist")

		err = php.Copy(src, dst)
		Expect(err).NotTo(HaveOccurred())

		php.Unlink(dst)
	})

	It("FileExists", func() {
		var filename = "/tmp/hyperjiangphpfilesystemtest.txt"
		php.Unlink(filename)
		Expect(php.FileExists(filename)).To(BeFalse()) // file not exists

		php.Touch(filename)
		Expect(php.FileExists(filename)).To(BeTrue())
		php.Unlink(filename)
	})

	It("FileSize", func() {
		var filename = "/tmp/hyperjiangphpfilesystemtest.txt"
		php.Unlink(filename)
		_, err := php.FileSize(filename)
		Expect(err).To(HaveOccurred())

		php.Touch(filename)
		size, err := php.FileSize(filename)
		Expect(err).NotTo(HaveOccurred())
		Expect(size).To(BeZero())

		php.Unlink(filename)
	})

	It("Rename", func() {
		var oldname = "/tmp/hyperjiangphpfilesystemtest.old"
		php.Touch(oldname)

		var newname = "/tmp/hyperjiangphpfilesystemtest.new"
		php.Rename(oldname, newname)

		Expect(php.FileExists(oldname)).To(BeFalse())
		Expect(php.FileExists(newname)).To(BeTrue())

		php.Unlink(newname)
	})

	It("FilePutContents and FileGetContents", func() {
		var filename = "/tmp/hyperjiangphpfilesystemtest.txt"
		php.Unlink(filename)

		const msg = "Hello world!"
		err := php.FilePutContents(filename, msg)
		Expect(err).NotTo(HaveOccurred())

		str, err := php.FileGetContents(filename)
		Expect(err).NotTo(HaveOccurred())
		Expect(str).To(Equal(msg))

		php.Unlink(filename)
	})
})
