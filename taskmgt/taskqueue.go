package taskmgt

import (
	"container/list"
	"sync"
)



/*
* TaskQueue
*/

type TaskQueue struct {
	TaskList list.List
	TaskTable map[int32](*list.Element)

	Rwlock sync.RWMutex
}

func NewTaskQueue()  *TaskQueue {
	return &TaskQueue{
		TaskList:  list.List{},
		TaskTable: make(map[int32](*list.Element)),
	}
}

func (tq *TaskQueue) AddTask(task *TaskEntity) {
	tq.Rwlock.Lock()
	e := tq.TaskList.PushBack(task)
	tq.TaskTable[task.TaskId] = e
	tq.Rwlock.Unlock()
}

func (tq *TaskQueue) RemoveTask(TaskId int32) {
	if e, ok := tq.TaskTable[TaskId]; ok {
		tq.Rwlock.Lock()
		delete(tq.TaskTable, TaskId)
		tq.TaskList.Remove(e)
		tq.Rwlock.Unlock()
	}
}

func (tq *TaskQueue) FindTask(TaskId int32) *TaskEntity {
	if e, ok := tq.TaskTable[TaskId]; ok {
		return e.Value.(*TaskEntity)
	}
	return nil
}

func (tq *TaskQueue) GettaskNum() int {
	return tq.TaskList.Len()
}


//从json文件读入任务列表
/*func ReadTaskList(listpath string) {
	filePtr, err := os.Open(listpath)
  if err != nil {
    fmt.Println("Open file failed [Err:%s]", err.Error())
    return
  }
  defer filePtr.Close()
  log.Printf("opened %s", filePtr)

  var task []TaskEntity
  // 创建json解码器
  decoder := json.NewDecoder(filePtr)
  err = decoder.Decode(&task)
  if err != nil {
    fmt.Println("Decoder failed", err.Error())
  } else {
    fmt.Println("Decoder success")
    fmt.Println(task)
  }
  fmt.Println("---taskloading end---")
}*/