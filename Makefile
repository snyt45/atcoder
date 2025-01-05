login:
	oj login https://atcoder.jp/
.PHONY: login

init:
	./shell/init.sh
.PHONY: init

c:
	./shell/contest.sh
.PHONY: c

dl:
	./shell/download.sh
.PHONY: dl

sw:
	./shell/switch.sh $(n)
.PHONY: sw

run:
	./shell/run.sh
.PHONY: run

test:
	./shell/test.sh
.PHONY: test

submit:
	./shell/submit.sh
.PHONY: submit
