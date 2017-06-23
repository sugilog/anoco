Anoco
======

Remote Server Agent / Job Manager.


Development
------

Go 1.8+

[make2help](https://github.com/Songmu/make2help)


TODO
------

### Agent for Remote Server.

Functions

- [x] Health Check.
- [x] Job Subscriber. (From Job Manager)
- [x] Job Status-Report Subscriber. (From Job Manager)
- [x] Job Execution Worker.
- [ ] Job Reporter. (To Job Manager)

Component

- Agent bin.

### Job Manager

Functions

- [ ] Agent Discovery Worker. (To Agent)
- [ ] Job Registerer. (On demand, Agent-Handled)
- [ ] Job Publisher. (To Agent, Retry support)
- [ ] Job Status Observer. (To Agent)
- [ ] Job Status Report Receiver. (From Agent)

Component

- Workers bin ... ?
- Datastore ... ?
