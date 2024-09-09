
PHONY: run stop build clean

build: build_gate build_door build_game

build_game:
	echo "Building gate..."
	@cd game && go build -v -o ../bin/game
# 构建 gate 服务
build_gate:
	echo "Building gate..."
	@cd gate && go build -v -o ../bin/gate

# 构建 door 服务
build_door:
	echo "Building door..."
	@cd door && go build -v -o ../bin/door

# 定义 bin 目录
BIN_DIR = $(shell pwd)/bin

# 动态获取 bin 目录中的所有可执行文件作为服务列表
SERVICES = $(shell find $(BIN_DIR) -maxdepth 1 -type f -executable -exec basename {} \;)

# 启动所有服务
run:
	@for service in $(SERVICES); do \
	  echo "Starting $$service..."; \
	  #   echo "Command: cd $(shell pwd)/$$service && $(BIN_DIR)/$$service > $(shell pwd)/logs/$$service.log 2>&1 &"; \
	  cd $(shell pwd)/$$service && $(BIN_DIR)/$$service > $(shell pwd)/logs/$$service.log 2>&1 & \
	  sleep 1; \
	  PID=$$(pgrep -f "$(BIN_DIR)/$$service"); \
	  if [ -n "$$PID" ]; then \
	    CMD=$$(ps -p $$PID -o comm=); \
	    echo "Started process $$PID ($$CMD)"; \
	  else \
	    echo "Failed to start $$service"; \
	  fi; \
	done

# 停止所有服务
stop:
	@PIDS=$$(pgrep -f "$$(pwd)/bin/"); \
	if [ -n "$$PIDS" ]; then \
	  for PID in $$PIDS; do \
	    CMD=$$(ps -p $$PID -o comm=); \
	    echo "Killing process $$PID ($$CMD)"; \
	    kill $$PID; \
	  done; \
	else \
	  echo "No processes found for executables in $$(pwd)/bin/"; \
	fi

# 清理日志文件
clean:
	rm -f logs/*.log

wscli:
	@node gate/fakecli/wscli.js

tcpcli:
	@node gate/fakecli/tcpcli.js