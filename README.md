# twins

A simple master/slave architecture based on GRPC implementation

## Add your task

1. Create `your_own_task.go` file under the directory `task`
2. Implement all interfaces.
3. Register your task when `RegisterTask()` is called

## Note

This item is only applicable to you only have two servers, and you need to ask for high availability :)