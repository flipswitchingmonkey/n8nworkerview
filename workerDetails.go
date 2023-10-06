package main

import (
	"fmt"
	"strings"
)

func newLine(b *strings.Builder) {
	b.WriteString("\n")
}

func addDetail(b *strings.Builder, label string, content string) {
	b.WriteString(H2Style.Render(label))
	b.WriteString(NormalTextStyle.Render(content))
	newLine(b)
}

func RenderWorkerDetails(m model) string {
	details, err := m.GetSelectedWorkerDetails()
	if err != nil {
		return ""
	}
	var b strings.Builder
	b.WriteString(H1Style.Render(details.WorkerId) + "\n")
	// b.WriteString(H2Style.Render("Hostname"))
	// b.WriteString(NormalTextStyle.Render(details.Hostname))
	// newLine(&b)
	addDetail(&b, "Hostname", details.Hostname)
	addDetail(&b, "Architecture", details.Arch)
	addDetail(&b, "Platform", details.Platform)
	// addDetail(&b, "Jobs", strings.Join(details.RunningJobsSummary, "\n     "))
	loads := ""
	for i, load := range details.LoadAvg {
		loads += fmt.Sprintf("%f", load)
		if i < len(details.LoadAvg)-1 {
			loads += ", " 
		}
	}
	addDetail(&b, "Load", loads)
	addDetail(&b, "Free Memory", fmt.Sprintf("%.2f / %.2f GB", details.FreeMem/1024/1024/1024, details.TotalMem/1024/1024/1024))
	addDetail(&b, "CPUs", details.Cpus)
	addDetail(&b, "Network", strings.Join(details.Net, "\n         "))
	return b.String()
}

func WorkerFreeMem(worker *RedisWorkerResponsePayload) string {
	return fmt.Sprintf("%.2f / %.2f GB", worker.FreeMem/1024/1024/1024, worker.TotalMem/1024/1024/1024)
}

func WorkerAverageLoad(worker *RedisWorkerResponsePayload) string {
	if len(worker.LoadAvg) == 0 {
		return ""
	}
	var loads float32 = 0
	for _, load := range worker.LoadAvg {
		loads += load
	}
	return fmt.Sprintf("%.2f", loads / float32(len(worker.LoadAvg)))
}
