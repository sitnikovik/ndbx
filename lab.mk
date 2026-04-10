LAB_DIR = autograder/cmd

.PHONY: autograder
autograder: lab1 lab2 lab3 lab4 lab5

.PHONY: lab1
lab1:
	@cd $(LAB_DIR)/lab1 && set -a; source $(ENV_FILE); set +a; go run .

.PHONY: lab2
lab2:
	@cd $(LAB_DIR)/lab2 && set -a; source $(ENV_FILE); set +a; go run .

.PHONY: lab3
lab3:
	@cd $(LAB_DIR)/lab3 && set -a; source $(ENV_FILE); set +a; go run .

.PHONY: lab4
lab4:
	@cd $(LAB_DIR)/lab4 && set -a; source $(ENV_FILE); set +a; go run .

.PHONY: lab5
lab5:
	@cd $(LAB_DIR)/lab5 && set -a; source $(ENV_FILE); set +a; go run .