# 定义前端和后端目录
FRONTEND_DIR=frontend
BACKEND_DIR=backend

# front dev 目标
front:
	@echo "Running front-end development..."
	cd $(FRONTEND_DIR) && make dev

# back dev 目标
back:
	@echo "Running back-end development..."
	cd $(BACKEND_DIR) && make dev
