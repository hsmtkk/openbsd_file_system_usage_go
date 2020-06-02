package fsusage

import(
  "os/exec"
  "fmt"
  "strconv"
  "strings"
)

type FsUsage struct {
  FileSystem string
  Blocks512 int64
  Used int64
  Avail int64
  Capacity int
  MountedOn string
}

func GetFsUsages()([]FsUsage, error){
  out, err := ExecDfCommand()
  if err != nil {
    return nil, fmt.Errorf("failed to exec df command; %w", err)
  }
  usages, err := ParseDfOutput(out)
  if err != nil {
    return nil, fmt.Errorf("failed to parse df command output; %w", err)
  }
  return usages, nil
}

func ExecDfCommand()(string, error){
  out, err := exec.Command("df").Output()
  if err != nil {
    return "", err
  }
  return string(out), nil
}

func ParseDfOutput(dfOutput string)([]FsUsage, error){
  usages := []FsUsage{}
  lines := strings.Split(dfOutput, "\n")
  for i, line := range lines {
    // skip header
    if i == 0 {
      continue
    }
    // skip empty line
    if strings.TrimSpace(line) == "" {
      continue
    }
    usage, err := ParseDfOutputLine(line)
    if err != nil {
      return nil, err
    }
    usages = append(usages, usage)
  }
  return usages, nil
}

func ParseDfOutputLine(dfOutputLine string)(FsUsage, error){
  empty := FsUsage{}
  columns := strings.Fields(dfOutputLine)
  if len(columns) != 6 {
    return empty, fmt.Errorf("invalid column length: %d", len(columns))
  }
  blocks512, err := atoi64(columns[1])
  if err != nil {
    return empty, fmt.Errorf("failed to parse 512-blocks column: %s", columns[1])
  }
  used, err := atoi64(columns[2])
  if err != nil {
    return empty, fmt.Errorf("failed to parse used column: %s", columns[2])
  }
  avail, err := atoi64(columns[3])
  if err != nil {
    return empty, fmt.Errorf("failed to parse avail column: %s", columns[3])
  }
  if !strings.HasSuffix(columns[4], "%") {
    return empty, fmt.Errorf("failed to parse capacity column: %s", columns[4])
  }
  capacity, err := strconv.Atoi(strings.TrimSuffix(columns[4], "%"))
  if err != nil {
    return empty, fmt.Errorf("failed to parse capacity column: %s", columns[4])
  }

  return FsUsage{
    FileSystem: columns[0],
    Blocks512: blocks512,
    Used: used,
    Avail: avail,
    Capacity: capacity,
    MountedOn: columns[5],
  }, nil
}

func atoi64(s string)(int64, error){
  return strconv.ParseInt(s, 10, 64)
}
