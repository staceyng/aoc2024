# Path to the skeleton script
SKELETON_SCRIPT := ./skeleton.sh

# Create a new project skeleton
.PHONY: skeleton
skeleton:
	@bash $(SKELETON_SCRIPT) --dir $(DIR)
	@echo "Skeleton created in $(DIR) directory"

# Clean up (remove) a skeleton directory
.PHONY: clean
clean:
	@if [ -d "$(DIR)" ]; then \
		rm -rf "$(DIR)"; \
		echo "Removed $(DIR) directory"; \
	else \
		echo "Directory $(DIR) does not exist"; \
	fi

# Help target to display available commands
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  make skeleton DIR=day1   - Create a new project skeleton (default: day1)"
	@echo "  make clean DIR=day1      - Remove the specified directory"
	@echo "  make help                - Show this help message"