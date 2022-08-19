# Getting Started with exception
## ThrowException
```go
exception.ThrowRefuseException("error message")
```

## Catch
```go
// Different types of exceptions can be caught at the same time
defer exception.Catch().
    RefuseException(func(exp *exception.RefuseException) {
        if taskGroupDO.Id > 0 {
            log.TaskLogAddService(dto.TaskGroupId, taskGroupDO.JobName, taskGroupDO.Caption, eumLogLevel.Warning, exp.Message)
        }
        exp.ContinueRecover(exp.Message)
    }).
    String(func(exp string) {
        if taskGroupDO.Id > 0 {
            taskGroupDO.Cancel()
            r.repository.Save(taskGroupDO)
            log.TaskLogAddService(taskGroupDO.Id, taskGroupDO.JobName, taskGroupDO.Caption, eumLogLevel.Error, exp)
        }
    })
```