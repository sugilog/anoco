Anoco
======

Remote Server Agent / Job Manager.


Development
------

Go 1.8+


TODO
------

### Agent for Remote Server.

Functions

- Health Check.
- Job Subscriber. (From Job Manager)
- Job Status-Report Subscriber. (From Job Manager)
- Job Execution Worker.
- Job Reporter. (To Job Manager)

Component

- Agent bin.

### Job Manager

Functions

- Agent Discovery Worker. (To Agent)
- Job Registerer. (On demand, Agent-Handled)
- Job Publisher. (To Agent)
- Job Status Observer. (To Agent)
- Job Status Report Receiver. (From Agent)

Component

- Workers bin.
- Datastore ... ?
