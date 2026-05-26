## Task Tracker

Task Tracker is an application that allows you to manage and track your tasks. All tasks will be stored in a json file in the root of the project.

![alt text](carbon.png)

You can:
- Add tasks
- Update tasks
- Delete tasks
- Mark tasks as done or in progress
- View all tasks
- View tasks filtered by status (in progress, done)

## Running the Application

```bash
go build -o task-cli ./cmd
```

You can replace the executable file name `task-cli` with any name you prefer.

## Available Commands

Add a new task:

```bash
./task-cli add "Ask ChatGPT how to center a div for the 47th time"
```

Output:

```text
Task added successfully (ID: 1)
```

Update a task:

```bash
./task-cli update 1 "Ask ChatGPT how to center a div for the 48th time"
```

Output:

```text
Task 1 updated successfully
```

Delete a task:

```bash
./task-cli delete 1
```

Output:

```text
Task 1 deleted successfully
```

Mark a task as done:

```bash
./task-cli mark-done 1
```

Output:

```text
Task 1 marked as done
```

Mark a task as in progress:

```bash
./task-cli mark-in-progress 1
```

Output:

```text
Task 1 marked as in progress
```

View all tasks:

```bash
./task-cli list
```

Output:

```text
ID: 1, Description: Ask ChatGPT how to center a div for the 48th time, Status: in-progress, CreatedAt: 2026-05-22 11:44:00, UpdatedAt: 2026-05-22 11:48:00

ID: 2, Description: Survive another AI-generated meeting, Status: todo, CreatedAt: 2026-05-22 11:48:53, UpdatedAt: 2026-05-22 11:48:53

ID: 3, Description: Convince myself this bug is a feature, Status: todo, CreatedAt: 2026-05-22 11:48:59, UpdatedAt: 2026-05-22 11:48:59

ID: 4, Description: Stop debugging. It was a typo again., Status: todo, CreatedAt: 2026-05-22 11:49:06, UpdatedAt: 2026-05-22 11:49:06
```

Commands for viewing tasks by status:

```bash
./task-cli list done
./task-cli list todo
./task-cli list in-progress
```