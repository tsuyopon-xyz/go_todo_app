package store

import (
	"errors"

	"github.com/tsuyopon-xyz/go_todo_app/entity"
)

var (
	Tasks = &TaskStore{Tasks: map[entity.TaskID]*entity.Task{}}

	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	// 動作確認用の仮実装なのであえてexportしている。
	LastID entity.TaskID
	Tasks  map[entity.TaskID]*entity.Task
}

func (ts *TaskStore) Add(t *entity.Task) (entity.TaskID, error) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[t.ID] = t
	return t.ID, nil
}

func (ts *TaskStore) All() entity.Tasks {
	tasks := make([]*entity.Task, len(ts.Tasks))
	// tasks := []*entity.Task{}
	for i, t := range ts.Tasks {
		tasks[i-1] = t
	}

	return tasks
}

func (ts *TaskStore) Get(id entity.TaskID) (*entity.Task, error) {
	if ts, ok := ts.Tasks[id]; ok {
		return ts, nil
	}
	return nil, ErrNotFound
}
