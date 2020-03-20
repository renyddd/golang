## My golang src directory.
Install the golang form: golang.google.cn/dl/

make the **work** directory in your HOME directory, like the following.
```
work
├── bin
│   └── hello
├── pkg
└── src
    └── github.com
        └── renyddd
            ├── hello
            │   └── hello.go
            └── README.md
```
then edit your /etc/profile, using **sudo**.
```bash
export GOPATH=$HOME/work
export PATH=$PATH:$GOPATH/bin
```
