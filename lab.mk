LAB_DIR = autograder/cmd

.PHONY: lab5
lab5:
	@cd $(LAB_DIR)/lab5 && set -a; source $(ENV_FILE); set +a; go run .