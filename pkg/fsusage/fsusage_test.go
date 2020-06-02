package fsusage_test

import(
  "testing"
  "github.com/hsmtkk/openbsd_file_system_usage_go/pkg/fsusage"
  "github.com/stretchr/testify/assert"
)

func TestExecDfCommand(t *testing.T){
  out, err := fsusage.ExecDfCommand()
  assert.Nil(t, err, "should be nil")
  usages, err := fsusage.ParseDfOutput(out)
  assert.Nil(t, err, "should be nil")
  assert.Greater(t, len(usages), 0, "should be greater than zero")
}

func TestParseDfOutput(t *testing.T){
  dfOutput := `Filesystem  512-blocks      Used     Avail Capacity  Mounted on
/dev/wd0a      2057756    156472   1798400     8%    /
/dev/wd0k     23104892   7501628  14448020    34%    /home`
  want := []fsusage.FsUsage{}
  want = append(want, fsusage.FsUsage{
    FileSystem : "/dev/wd0a",
    Blocks512:2057756,
    Used: 156472,
    Avail: 1798400,
    Capacity: 8,
    MountedOn:"/",
  })
  want = append(want, fsusage.FsUsage{
    FileSystem : "/dev/wd0k",
    Blocks512:23104892,
    Used: 7501628,
    Avail: 14448020,
    Capacity: 34,
    MountedOn:"/home",
  })
  got, err := fsusage.ParseDfOutput(dfOutput)
  assert.Nil(t, err, "should be nil")
  assert.Equal(t, want, got, "should be equal")
}

func TestParseDfOutputLine(t *testing.T){
  dfOutputLine := "/dev/wd0a      2057756    156472   1798400     8%    /"
  want := fsusage.FsUsage{
    FileSystem : "/dev/wd0a",
    Blocks512:2057756,
    Used: 156472,
    Avail: 1798400,
    Capacity: 8,
    MountedOn:"/",
  }
  got, err := fsusage.ParseDfOutputLine(dfOutputLine)
  assert.Nil(t, err, "should be nil")
  assert.Equal(t, want, got, "should be equal")
}
