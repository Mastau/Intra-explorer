OUTPUT_FOLDER := ../data
OUTPUT_FILE := $(OUTPUT_FOLDER)/UserList.json

.PHONY: userList

$(OUTPUT_FOLDER) :
	@mkdir $(OUTPUT_FOLDER)

userList: $(OUTPUT_FOLDER)
	@echo "creating the users list..."
	go run clients_list/main.go $(OUTPUT_FILE)
	@echo "User list saved to $(OUTPUT_FILE)."

.PHONY: clean
clean:
	@echo "Cleaning up generated files..."
	rm -f $(OUTPUT_FILE)
	@echo "Clean complete."
