GATE_CMD = cd gate && $$(dirname $$(pwd))/bin/gate > gate.log 2>&1 &
GAME_CMD = cd game && $$(dirname $$(pwd))/bin/game > game.log 2>&1 &
DOOR_CMD = cd door && $$(dirname $$(pwd))/bin/door > door.log 2>&1 &

# 目标
run_all:
	@$(GATE_CMD)
	@$(GAME_CMD)
	@$(DOOR_CMD)

.PHONY: pbc
pbc:
	python corev2/message/pbc.py source=pb

ba:
	cd game && go build -o ../bin/game main.go 
	cd gate && go build -o ../bin/gate main.go 
	cd door && go build -o ../bin/door main.go

wscli:
	@node gate/fakecli/wscli.js

tcpcli:
	@node gate/fakecli/tcpcli.js