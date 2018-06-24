# compute-mixedreality
Hackathon inspired mixed reality solution for scenario-driven training.

Added golang server -
1. install go 
2. $TARGETDIR is wherever you clone the repository to.
3. to install: 
go install $TARGETDIR/WebApp/src/main.go
4. packages will be installed in $TARGETDIR/WebApp/bin/
5. run server with:
./$TARGETDIR/bin/main bootserv
6. run testclient with:
./$TARGETDIR/bin/main bootclient

*currently not a lot of functionality - but the golang TCP server basics are established.

*Google doc for design in progress: 
