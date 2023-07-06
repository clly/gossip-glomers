
MAELSTROM=maelstrom/maelstrom

.PHONY: echo
echo: bin/echo maelstrom
	./maelstrom/maelstrom test -w echo --bin bin/echo --node-count 1 --time-limit 10

.PHONY: maelstrom
maelstrom: $(MAELSTROM)

$(MAELSTROM):
	wget -O- https://github.com/jepsen-io/maelstrom/releases/download/v0.2.3/maelstrom.tar.bz2 | tar -xvj

.PHONY: bin/echo
bin/echo:
	go build -o bin/echo ./echo
