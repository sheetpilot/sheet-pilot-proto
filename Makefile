
PLATFORM := ${sh -c uname}

.PHONY: gen-hooks
gen-hooks:
		chmod +x gen.sh
		touch .git/hooks/pre-commit
		chmod +x .git/hooks/pre-commit
		echo "$$PWD/gen.sh" | tee .git/hooks/pre-commit
