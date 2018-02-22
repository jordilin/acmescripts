# Installation notes

## Plan9

git clone https://github.com/9fans/plan9port

I prefer euro.9.font as default, so remove euro.8.font and add:

src/cmd/acme/acme.c:    "/lib/font/bit/lucm/euro.9.font"

./INSTALL

## Tools

* go get golang.org/x/tools/cmd/goimports
* go get github.com/jordilin/go/acme/acmego
* go get github.com/jordilin/Watch

* cp plumbing to $HOME/lib
* cp $HOME/opensource/plan9port/plumb/fileaddr $HOME/lib

