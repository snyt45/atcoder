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

dlu:
	./shell/downloadurl.sh
.PHONY: dlu

sw:
	./shell/switch.sh $(n)
.PHONY: sw

run:
	./shell/run.sh
.PHONY: run

watch:
	./shell/watch.sh
.PHONY: watch

test:
	./shell/test.sh
.PHONY: test

submit:
	./shell/submit.sh
.PHONY: submit

move:
	./shell/movefile.sh
.PHONY: move
