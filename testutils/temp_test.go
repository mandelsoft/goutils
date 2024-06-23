package testutils_test

import (
	"os"

	"github.com/mandelsoft/filepath/pkg/filepath"
	"github.com/mandelsoft/goutils/sliceutils"
	. "github.com/mandelsoft/goutils/testutils"
	"github.com/mandelsoft/vfs/pkg/osfs"
	"github.com/mandelsoft/vfs/pkg/vfs"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ttemp dir", func() {
	It("", func() {
		temp := Must(NewTempDir())
		defer Defer(temp.Cleanup)
		path := temp.Path()

		Expect(Must(vfs.DirExists(osfs.OsFs, path))).To(BeTrue())
		Expect(Must(os.ReadDir(path))).To(Equal([]os.DirEntry{}))
		MustBeSuccessful(temp.Cleanup())
		Expect(Must(vfs.FileExists(osfs.OsFs, path))).To(BeFalse())
	})

	Context("with initial content", func() {
		It("handles dir content", func() {
			temp := Must(NewTempDir(WithDirContent("testdata")))
			defer Defer(temp.Cleanup)
			path := temp.Path()

			Expect(sliceutils.Transform(Must(os.ReadDir(path)), DirEntryToName)).To(Equal([]string{"content"}))
			Expect(sliceutils.Transform(Must(os.ReadDir(filepath.Join(path, "content"))), DirEntryToName)).To(Equal([]string{"file"}))
			Expect(Must(os.ReadFile(filepath.Join(path, "content/file")))).To(Equal([]byte("some file content")))
		})

		It("handles file content", func() {
			temp := Must(NewTempDir(WithFileContent("testdata/content/file")))
			defer Defer(temp.Cleanup)
			path := temp.Path()

			Expect(Must(os.ReadFile(filepath.Join(path, "file")))).To(Equal([]byte("some file content")))
		})
	})

	Context("with initial content at dedicated path", func() {
		It("handles dir content", func() {
			temp := Must(NewTempDir(WithDirContent("testdata", "target")))
			defer Defer(temp.Cleanup)
			path := temp.Path()

			Expect(sliceutils.Transform(Must(os.ReadDir(path)), DirEntryToName)).To(Equal([]string{"target"}))
			Expect(sliceutils.Transform(Must(os.ReadDir(filepath.Join(path, "target"))), DirEntryToName)).To(Equal([]string{"content"}))
			Expect(sliceutils.Transform(Must(os.ReadDir(filepath.Join(path, "target/content"))), DirEntryToName)).To(Equal([]string{"file"}))
			Expect(Must(os.ReadFile(filepath.Join(path, "target/content/file")))).To(Equal([]byte("some file content")))
		})

		It("handles dir content (absolute target)", func() {
			temp := Must(NewTempDir(WithDirContent("testdata", "/target")))
			defer Defer(temp.Cleanup)
			path := temp.Path()

			Expect(sliceutils.Transform(Must(os.ReadDir(path)), DirEntryToName)).To(Equal([]string{"target"}))
			Expect(sliceutils.Transform(Must(os.ReadDir(filepath.Join(path, "target"))), DirEntryToName)).To(Equal([]string{"content"}))
			Expect(sliceutils.Transform(Must(os.ReadDir(filepath.Join(path, "target/content"))), DirEntryToName)).To(Equal([]string{"file"}))
			Expect(Must(os.ReadFile(filepath.Join(path, "target/content/file")))).To(Equal([]byte("some file content")))
		})

		It("handles file content", func() {
			temp := Must(NewTempDir(WithFileContent("testdata/content/file", "target/file")))
			defer Defer(temp.Cleanup)
			path := temp.Path()

			Expect(Must(os.ReadFile(filepath.Join(path, "target/file")))).To(Equal([]byte("some file content")))
		})

		It("handles file content (absolute target)", func() {
			temp := Must(NewTempDir(WithFileContent("testdata/content/file", "/target/file")))
			defer Defer(temp.Cleanup)
			path := temp.Path()

			Expect(Must(os.ReadFile(filepath.Join(path, "target/file")))).To(Equal([]byte("some file content")))
		})
	})

	Context("invalid destinations", func() {
		It("handles directory content", func() {
			ExpectError(NewTempDir(WithDirContent("testdata", "../target"))).To(MatchError("destination above root"))
		})

		It("handles file content", func() {
			ExpectError(NewTempDir(WithFileContent("testdata/content/file", "../file"))).To(MatchError("destination above root"))
		})
	})
})

func DirEntryToName(d os.DirEntry) string {
	return d.Name()
}
