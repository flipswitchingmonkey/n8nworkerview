package main

import (
	"fmt"
	"os"
	"time"

	"slices"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

var Ui *tea.Program
var Model model

type model struct {
	ready    bool
	showDetails    bool
	workerNames  []string
	workers  map[string]RedisWorkerResponsePayload
	lastUpdated  map[string]time.Time
	workerTable table.Model
	jobsTable table.Model
	viewport viewport.Model
}

func initialModel() model {
	tWorker:=table.New(
			table.WithColumns([]table.Column{
					{Title: "Worker Id", Width: 26},
					{Title: "FreeMem", Width: 18},
					{Title: "Avg Load", Width: 8},
					{Title: "Jobs", Width: 4},
				}),
		)
	tWorker.SetStyles(TableStyle)
	tWorker.Focus()
	tJobs:=table.New(
			table.WithColumns([]table.Column{
					{Title: "Execution", Width: 10},
					{Title: "Workflow Name", Width: 30},
					{Title: "Workflow Id", Width: 16},
					{Title: "Started At", Width: 20},
					{Title: "Retry of", Width: 10},
			}),
		)
	tJobs.SetStyles(TableStyle)
	tJobs.Blur()
	
	return model{
		ready: false,
		showDetails: false,
		workerNames: []string{},
		workers: map[string]RedisWorkerResponsePayload{},
		lastUpdated: map[string]time.Time{},
		workerTable: tWorker,
		jobsTable: tJobs,
	}
}

func (m model) Init() tea.Cmd {
	SendGetId()
	SendGetStatus()
	return nil
}

func (m model) GetSelectedRowName() (string, error){
	if m.workerTable.SelectedRow() == nil {
		return "", fmt.Errorf("no row selected")
	}
	rowAsStrings := []string(m.workerTable.SelectedRow())
	if len(rowAsStrings) < 1 {
		return "", fmt.Errorf("row has no columns")
	}
	selectedWorkerName := rowAsStrings[0]
	return selectedWorkerName, nil
}

func (m model) GetWorkerDetails(workerName string) (*RedisWorkerResponsePayload, error) {
	details, ok := m.workers[workerName]
	if !ok {
		return nil, fmt.Errorf("no details for worker "+workerName)
	}
	return &details, nil
}

func (m model) GetSelectedWorkerDetails() (*RedisWorkerResponsePayload, error) {
	workerName, err := m.GetSelectedRowName()
	if err  != nil {
		return nil, fmt.Errorf("no worker selected")
	}
	details, ok := m.workers[workerName]
	if !ok {
		return nil, fmt.Errorf("no details for worker "+workerName)
	}
	return &details, nil
}

func (m *model) CheckStaleWorkers() {
	copyNames := slices.Clone(m.workerNames)
	for _, workerName := range copyNames {
		if (m.lastUpdated[workerName].Add(2 * time.Second)).Compare(time.Now()) < 0 {
			delete(m.workers, workerName)
		}
		if (m.lastUpdated[workerName].Add(3 * time.Second)).Compare(time.Now()) < 0 {
			idx := slices.Index(m.workerNames, workerName)
			if idx > -1 {
				if len(m.workerNames) == 1 {
					m.workerNames = []string{}
				} else {
					m.workerNames = append(m.workerNames[:idx], m.workerNames[idx+1:]...)
				}
			}
		}
	}
}

func (m model) GetWorkerNamesAsRows() ([]table.Row) {
	var rows []table.Row
	for _, workerName := range m.workerNames {
		worker, ok := m.workers[workerName]
		if ok {
			rows = append(rows, table.Row{worker.WorkerId, WorkerFreeMem(&worker), WorkerAverageLoad(&worker), fmt.Sprint(len(worker.RunningJobs))})
		} else {
			rows = append(rows, table.Row{workerName, "n/a", "n/a", "n/a"})
		}
	}
	return rows
}

func (m model) GetSelectedWorkerJobsAsRows() ([]table.Row) {
	var rows []table.Row
	details, err := m.GetSelectedWorkerDetails()
	if err != nil {
		return rows
	}
	for _, job := range details.RunningJobs {
			rows = append(rows, table.Row{job})
	}
	return rows
}

func (m model) GetSelectedWorkerJobSummariesAsRows() ([]table.Row) {
	var rows []table.Row
	details, err := m.GetSelectedWorkerDetails()
	if err != nil {
		return rows
	}
	for _, job := range details.RunningJobsSummary {
		startedAt, timeErr := time.Parse(time.RFC3339, job.StartedAt)
		var displayStartedAt string
		if timeErr != nil {
			displayStartedAt = timeErr.Error()
		} else {
			displayStartedAt = startedAt.Format(time.DateTime)
		}
		rows = append(rows, table.Row{job.ExecutionId, job.WorkflowName, job.WorkflowId, displayStartedAt, job.RetryOf})
	}
	return rows
}

func UpdateFromRedis(msg RedisWorkerResponse) {
	Ui.Send(msg)
}

func ShowUi() {
	Model = initialModel()
	Ui = tea.NewProgram(Model, tea.WithAltScreen(), tea.WithMouseCellMotion(),)
	if _, err := Ui.Run(); err != nil { 
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
